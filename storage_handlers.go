package main

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"github.com/spaolacci/murmur3"
	"time"
	"encoding/hex"
	"hash"
	"errors"
)

// KeyError is a standard error for when a key is not found in the storage engine
type KeyError struct{}

func (e KeyError) Error() string {
	return "Key not found"
}

// StorageHandler is a standard interface to a storage backend,
// used by AuthorisationManager to read and write key values to the backend
type StorageHandler interface {
	GetKey(string) (string, error) // Returned string is expected to be a JSON object (SessionState)
	SetKey(string, string, int64)  // Second input string is expected to be a JSON object (SessionState)
	GetExp(string) (int64, error)  // Returns expiry of a key
	GetKeys(string) []string
	DeleteKey(string) bool
	DeleteRawKey(string) bool
	Connect() bool
	GetKeysAndValues() map[string]string
	GetKeysAndValuesWithFilter(string) map[string]string
	DeleteKeys([]string) bool
	Decrement(string)
	IncrememntWithExpire(string, int64) int64
	SetRollingWindow(string, int64, int64) int
}

// InMemoryStorageManager implements the StorageHandler interface,
// it uses an in-memory map to store sessions, should only be used
// for testing purposes
type InMemoryStorageManager struct {
	Sessions map[string]string
}

// Decrement is a dummy function
func (s *InMemoryStorageManager) Decrement(n string) {
	log.Warning("Not implemented!")
}

func (s *InMemoryStorageManager) SetRollingWindow(keyName string, per int64, expire int64) int {
	log.Warning("Not Implemented!")
	return 0
}

func (s *InMemoryStorageManager) IncrememntWithExpire(n string, i int64) int64 {
	log.Warning("Not implemented!")
	return 0
}

func (s *InMemoryStorageManager) Connect() bool {
	return true
}

// GetKey retrieves the key from the in-memory map
func (s InMemoryStorageManager) GetKey(keyName string) (string, error) {
	value, ok := s.Sessions[keyName]
	if !ok {
		return "", KeyError{}
	}

	return value, nil

}

// SetKey updates the in-memory key
func (s InMemoryStorageManager) SetKey(keyName string, sessionState string, timeout int64) {
	s.Sessions[keyName] = sessionState
}

func (s InMemoryStorageManager) GetExp(keyName string) (int64, error) {
	return 0, nil
}

// GetKeys will retreive multiple keys based on a filter (prefix, e.g. tyk.keys)
func (s InMemoryStorageManager) GetKeys(filter string) []string {
	sessions := make([]string, 0, len(s.Sessions))
	for key := range s.Sessions {
		if strings.Contains(key, filter) {
			sessions = append(sessions, key)
		}
	}

	return sessions
}

// GetKeysAndValues returns all keys and their data, very expensive call.
func (s InMemoryStorageManager) GetKeysAndValues() map[string]string {
	return s.Sessions
}

// GetKeysAndValuesWithFilter does nothing here
func (s InMemoryStorageManager) GetKeysAndValuesWithFilter(filter string) map[string]string {
	log.Warning("NOT IMPLEMENTED")
	return s.Sessions
}

// DeleteKey will remove a key from the storage engine
func (s InMemoryStorageManager) DeleteKey(keyName string) bool {
	delete(s.Sessions, keyName)
	return true
}

// DeleteRawKey will remove a key from the storage engine
func (s InMemoryStorageManager) DeleteRawKey(keyName string) bool {
	delete(s.Sessions, keyName)
	return true
}

// DeleteKeys remove keys from sessions DB
func (s InMemoryStorageManager) DeleteKeys(keys []string) bool {

	for _, keyName := range keys {
		delete(s.Sessions, keyName)
	}

	return true
}

// ------------------- REDIS STORAGE MANAGER -------------------------------

// RedisStorageManager is a storage manager that uses the redis database.
type RedisStorageManager struct {
	pool      *redis.Pool
	KeyPrefix string
	HashKeys bool
}

func (r *RedisStorageManager) newPool(server, password string, database int) *redis.Pool {
	maxIdle := 100
	if config.Storage.MaxIdle > 0 {
		maxIdle = config.Storage.MaxIdle
	}
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if database > 0 {
				if _, err := c.Do("SELECT", database); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Connect will establish a connection to the DB
func (r *RedisStorageManager) Connect() bool {

	fullPath := config.Storage.Host + ":" + strconv.Itoa(config.Storage.Port)
	log.Info("Connecting to redis on: ", fullPath)
	r.pool = r.newPool(fullPath, config.Storage.Password, config.Storage.Database)

	return true
}

func doHash(in string) string {
	var h hash.Hash32 = murmur3.New32()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}

//Public function for use in classes that bypass elements of the storage manager
func publicHash(in string) string {
	if !config.HashKeys {
		// Not hashing? Return the raw key
		return in
	}
	
	return doHash(in)
}

func (r *RedisStorageManager) hashKey(in string) string {
	if !r.HashKeys {
		// Not hashing? Return the raw key
		return in
	}
	return doHash(in)
}

func (r *RedisStorageManager) fixKey(keyName string) string {
	setKeyName := r.KeyPrefix + r.hashKey(keyName)
	
	log.Debug("Input key was: ", setKeyName)
	
	
	return setKeyName
}

func (r *RedisStorageManager) cleanKey(keyName string) string {
	setKeyName := strings.Replace(keyName, r.KeyPrefix, "", 1)
	return setKeyName
}

// GetKey will retreive a key from the database
func (r *RedisStorageManager) GetKey(keyName string) (string, error) {
	db := r.pool.Get()
	defer db.Close()
	
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.GetKey(keyName)
	}

	log.Debug("[STORE] Getting: ", r.fixKey(keyName))
	value, err := redis.String(db.Do("GET", r.fixKey(keyName)))
	if err != nil {
		log.Debug("Error trying to get value:", err)
		return "", KeyError{}
	}

	return value, nil
}

func (r *RedisStorageManager) GetExp(keyName string) (int64, error) {
	db := r.pool.Get()
	defer db.Close()
	log.Debug("Getting exp for key: ", r.fixKey(keyName))
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.GetExp(keyName)
	}

	value, err := redis.Int64(db.Do("TTL", r.fixKey(keyName)))
	if err != nil {
		log.Error("Error trying to get TTL: ", err)
	} else {
		return value, nil
	}

	return 0, KeyError{}
}

// SetKey will create (or update) a key value in the store
func (r *RedisStorageManager) SetKey(keyName string, sessionState string, timeout int64) {
	db := r.pool.Get()
	defer db.Close()
	log.Debug("[STORE] Raw key is: ", keyName)
	log.Debug("[STORE] Setting key: ", r.fixKey(keyName))
	
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		r.SetKey(keyName, sessionState, timeout)
	} else {
		_, err := db.Do("SET", r.fixKey(keyName), sessionState)
		if timeout > 0 {
			_, expErr := db.Do("EXPIRE", r.fixKey(keyName), timeout)
			if expErr != nil {
				log.Error("Could not EXPIRE key")
				log.Error(expErr)
			}
		}
		if err != nil {
			log.Error("Error trying to set value:")
			log.Error(err)
		}
	}
}

// Decrement will decrement a key in redis
func (r *RedisStorageManager) Decrement(keyName string) {
	db := r.pool.Get()
	defer db.Close()

	keyName = r.fixKey(keyName)
	log.Debug("Decrementing key: ", keyName)
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		r.Decrement(keyName)
	} else {
		err := db.Send("DECR", keyName)

		if err != nil {
			log.Error("Error trying to decrement value:", err)
		}
	}
}

// IncrementWithExpire will increment a key in redis
func (r *RedisStorageManager) IncrememntWithExpire(keyName string, expire int64) int64 {
	db := r.pool.Get()
	defer db.Close()

	log.Debug("Incrementing raw key: ", keyName)
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		r.IncrememntWithExpire(keyName, expire)
	} else {
		// This function uses a raw key, so we shouldn't call fixKey
		fixedKey := keyName
		val, err := redis.Int64(db.Do("INCR", fixedKey))
		log.Debug("Incremented key: ", fixedKey, ", val is: ", val)
		if val == 1 {
			log.Debug("--> Setting Expire")
			db.Send("EXPIRE", fixedKey, expire)
		}
		if err != nil {
			log.Error("Error trying to increment value:", err)
		}
		return val
	}
	return 0
}

// GetKeys will return all keys according to the filter (filter is a prefix - e.g. tyk.keys.*)
func (r *RedisStorageManager) GetKeys(filter string) []string {
	db := r.pool.Get()
	defer db.Close()
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.GetKeys(filter)
	}

	searchStr := r.KeyPrefix + r.hashKey(filter) + "*"
	sessionsInterface, err := db.Do("KEYS", searchStr)
	if err != nil {
		log.Error("Error trying to get all keys:")
		log.Error(err)

	} else {
		sessions, _ := redis.Strings(sessionsInterface, err)
		for i, v := range sessions {
			sessions[i] = r.cleanKey(v)
		}

		return sessions
	}

	return []string{}
}

// GetKeysAndValuesWithFilter will return all keys and their values with a filter
func (r *RedisStorageManager) GetKeysAndValuesWithFilter(filter string) map[string]string {
	db := r.pool.Get()
	defer db.Close()
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.GetKeysAndValuesWithFilter(filter)
	}

	searchStr := r.KeyPrefix + r.hashKey(filter) + "*"
	log.Debug("[STORE] Getting list by: ", searchStr)
	sessionsInterface, err := db.Do("KEYS", searchStr)
	if err != nil {
		log.Error("Error trying to get filtered client keys:")
		log.Error(err)

	} else {
		keys, _ := redis.Strings(sessionsInterface, err)
		valueObj, err := db.Do("MGET", sessionsInterface.([]interface{})...)
		values, err := redis.Strings(valueObj, err)

		returnValues := make(map[string]string)
		for i, v := range keys {
			returnValues[r.cleanKey(v)] = values[i]
		}

		return returnValues
	}

	return map[string]string{}
}

// GetKeysAndValues will return all keys and their values - not to be used lightly
func (r *RedisStorageManager) GetKeysAndValues() map[string]string {
	db := r.pool.Get()
	defer db.Close()
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.GetKeysAndValues()
	}

	searchStr := r.KeyPrefix + "*"
	sessionsInterface, err := db.Do("KEYS", searchStr)
	if err != nil {
		log.Error("Error trying to get all keys:")
		log.Error(err)

	} else {
		keys, _ := redis.Strings(sessionsInterface, err)
		valueObj, err := db.Do("MGET", sessionsInterface.([]interface{})...)
		values, err := redis.Strings(valueObj, err)

		returnValues := make(map[string]string)
		for i, v := range keys {
			returnValues[r.cleanKey(v)] = values[i]
		}

		return returnValues
	}

	return map[string]string{}
}

// DeleteKey will remove a key from the database
func (r *RedisStorageManager) DeleteKey(keyName string) bool {
	db := r.pool.Get()
	defer db.Close()
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.DeleteKey(keyName)
	}

	_, err := db.Do("DEL", r.fixKey(keyName))
	if err != nil {
		log.Error("Error trying to delete key:")
		log.Error(err)
	}

	return true
}

// DeleteKey will remove a key from the database without prefixing, assumes user knows what they are doing
func (r *RedisStorageManager) DeleteRawKey(keyName string) bool {
	db := r.pool.Get()
	defer db.Close()
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.DeleteRawKey(keyName)
	}

	_, err := db.Do("DEL", keyName)
	if err != nil {
		log.Error("Error trying to delete key:")
		log.Error(err)
	}

	return true
}

// DeleteKeys will remove a group of keys in bulk
func (r *RedisStorageManager) DeleteKeys(keys []string) bool {
	db := r.pool.Get()
	defer db.Close()
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.DeleteKeys(keys)
	}

	if len(keys) > 0 {
		asInterface := make([]interface{}, len(keys))
		for i, v := range keys {
			asInterface[i] = interface{}(r.fixKey(v))
		}
		
		log.Debug("Deleting: ", asInterface)
		_, err := db.Do("DEL", asInterface...)
		if err != nil {
			log.Error("Error trying to delete keys:")
			log.Error(err)
		}
	} else {
		log.Debug("RedisStorageManager called DEL - Nothing to delete")
	}

	return true
}

// DeleteKeys will remove a group of keys in bulk without a prefix handler
func (r *RedisStorageManager) DeleteRawKeys(keys []string, prefix string) bool {
	db := r.pool.Get()
	defer db.Close()
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		return r.DeleteKeys(keys)
	}

	if len(keys) > 0 {
		asInterface := make([]interface{}, len(keys))
		for i, v := range keys {
			asInterface[i] = interface{}(prefix + v)
		}
		
		log.Debug("Deleting: ", asInterface)
		_, err := db.Do("DEL", asInterface...)
		if err != nil {
			log.Error("Error trying to delete keys:")
			log.Error(err)
		}
	} else {
		log.Debug("RedisStorageManager called DEL - Nothing to delete")
	}

	return true
}

// StartPubSubHandler will listen for a signal and run the callback with the message
func (r *RedisStorageManager) StartPubSubHandler(channel string, callback func(redis.Message)) error {
	psc := redis.PubSubConn{r.pool.Get()}
	psc.Subscribe(channel)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			callback(v)

		case redis.Subscription:
			log.Info("Subscription started: ", v.Channel)

		case error:
			log.Error("Redis disconnected or error received, attempting to reconnect: ", v)
			
			return v
		}
	}
	return errors.New("Connection closed.")
}

func (r *RedisStorageManager) Publish(channel string, message string) error {
	db := r.pool.Get()
	defer db.Close()
	if r.pool == nil {
		log.Info("Connection dropped, Connecting..")
		r.Connect()
		r.Publish(channel, message)
	} else {
		_, err := db.Do("PUBLISH", channel, message)
		if err != nil {
			log.Error("Error trying to set value:")
			log.Error(err)
			return err
		}
	}
	return nil
}

// IncrementWithExpire will increment a key in redis
func (r *RedisStorageManager) SetRollingWindow(keyName string, per int64, expire int64) int {
	db := r.pool.Get()
	defer db.Close()

	log.Debug("Incrementing raw key: ", keyName)
	if db == nil {
		log.Info("Connection dropped, connecting..")
		r.Connect()
		r.SetRollingWindow(keyName, per, expire)
	} else {
		log.Debug("keyName is: ", keyName)
		now := time.Now()
		log.Debug("Now is:", now)
		onePeriodAgo := now.Add(time.Duration(-1 * per) * time.Second)
		log.Debug("Then is: ", onePeriodAgo)
		
		db.Send("MULTI")
		// Drop the last period so we get current bucket
		db.Send("ZREMRANGEBYSCORE", keyName, "-inf", onePeriodAgo.UnixNano())
		// Get the set
		db.Send("ZRANGE", keyName, 0, -1)
		// Add this request to the pile
		db.Send("ZADD", keyName, now.UnixNano(), strconv.Itoa(int(now.UnixNano())))
		// REset the TTL so the key lives as long as the requests pile in
		db.Send("EXPIRE", keyName, per)
		r, err := redis.Values(db.Do("EXEC"))
		
		intVal := len(r[1].([]interface{}))
		
		log.Debug("Returned: ", intVal)
		
		if err != nil {
			log.Error("Multi command failed: ", err)
		}
		
		return intVal
	}
	return 0
}