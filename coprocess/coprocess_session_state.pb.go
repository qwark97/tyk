// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: coprocess_session_state.proto

package coprocess

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccessSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *AccessSpec) Reset() {
	*x = AccessSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_session_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessSpec) ProtoMessage() {}

func (x *AccessSpec) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_session_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessSpec.ProtoReflect.Descriptor instead.
func (*AccessSpec) Descriptor() ([]byte, []int) {
	return file_coprocess_session_state_proto_rawDescGZIP(), []int{0}
}

func (x *AccessSpec) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AccessSpec) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type AccessDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiName     string        `protobuf:"bytes,1,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	ApiId       string        `protobuf:"bytes,2,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	Versions    []string      `protobuf:"bytes,3,rep,name=versions,proto3" json:"versions,omitempty"`
	AllowedUrls []*AccessSpec `protobuf:"bytes,4,rep,name=allowed_urls,json=allowedUrls,proto3" json:"allowed_urls,omitempty"`
}

func (x *AccessDefinition) Reset() {
	*x = AccessDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_session_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessDefinition) ProtoMessage() {}

func (x *AccessDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_session_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessDefinition.ProtoReflect.Descriptor instead.
func (*AccessDefinition) Descriptor() ([]byte, []int) {
	return file_coprocess_session_state_proto_rawDescGZIP(), []int{1}
}

func (x *AccessDefinition) GetApiName() string {
	if x != nil {
		return x.ApiName
	}
	return ""
}

func (x *AccessDefinition) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

func (x *AccessDefinition) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *AccessDefinition) GetAllowedUrls() []*AccessSpec {
	if x != nil {
		return x.AllowedUrls
	}
	return nil
}

type BasicAuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Hash     string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BasicAuthData) Reset() {
	*x = BasicAuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_session_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAuthData) ProtoMessage() {}

func (x *BasicAuthData) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_session_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAuthData.ProtoReflect.Descriptor instead.
func (*BasicAuthData) Descriptor() ([]byte, []int) {
	return file_coprocess_session_state_proto_rawDescGZIP(), []int{2}
}

func (x *BasicAuthData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *BasicAuthData) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type JWTData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *JWTData) Reset() {
	*x = JWTData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_session_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTData) ProtoMessage() {}

func (x *JWTData) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_session_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTData.ProtoReflect.Descriptor instead.
func (*JWTData) Descriptor() ([]byte, []int) {
	return file_coprocess_session_state_proto_rawDescGZIP(), []int{3}
}

func (x *JWTData) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type Monitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TriggerLimits []float64 `protobuf:"fixed64,1,rep,packed,name=trigger_limits,json=triggerLimits,proto3" json:"trigger_limits,omitempty"`
}

func (x *Monitor) Reset() {
	*x = Monitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_session_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitor) ProtoMessage() {}

func (x *Monitor) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_session_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitor.ProtoReflect.Descriptor instead.
func (*Monitor) Descriptor() ([]byte, []int) {
	return file_coprocess_session_state_proto_rawDescGZIP(), []int{4}
}

func (x *Monitor) GetTriggerLimits() []float64 {
	if x != nil {
		return x.TriggerLimits
	}
	return nil
}

type SessionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastCheck               int64                        `protobuf:"varint,1,opt,name=last_check,json=lastCheck,proto3" json:"last_check,omitempty"`
	Allowance               float64                      `protobuf:"fixed64,2,opt,name=allowance,proto3" json:"allowance,omitempty"`
	Rate                    float64                      `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
	Per                     float64                      `protobuf:"fixed64,4,opt,name=per,proto3" json:"per,omitempty"`
	Expires                 int64                        `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"`
	QuotaMax                int64                        `protobuf:"varint,6,opt,name=quota_max,json=quotaMax,proto3" json:"quota_max,omitempty"`
	QuotaRenews             int64                        `protobuf:"varint,7,opt,name=quota_renews,json=quotaRenews,proto3" json:"quota_renews,omitempty"`
	QuotaRemaining          int64                        `protobuf:"varint,8,opt,name=quota_remaining,json=quotaRemaining,proto3" json:"quota_remaining,omitempty"`
	QuotaRenewalRate        int64                        `protobuf:"varint,9,opt,name=quota_renewal_rate,json=quotaRenewalRate,proto3" json:"quota_renewal_rate,omitempty"`
	AccessRights            map[string]*AccessDefinition `protobuf:"bytes,10,rep,name=access_rights,json=accessRights,proto3" json:"access_rights,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrgId                   string                       `protobuf:"bytes,11,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	OauthClientId           string                       `protobuf:"bytes,12,opt,name=oauth_client_id,json=oauthClientId,proto3" json:"oauth_client_id,omitempty"`
	OauthKeys               map[string]string            `protobuf:"bytes,13,rep,name=oauth_keys,json=oauthKeys,proto3" json:"oauth_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BasicAuthData           *BasicAuthData               `protobuf:"bytes,14,opt,name=basic_auth_data,json=basicAuthData,proto3" json:"basic_auth_data,omitempty"`
	JwtData                 *JWTData                     `protobuf:"bytes,15,opt,name=jwt_data,json=jwtData,proto3" json:"jwt_data,omitempty"`
	HmacEnabled             bool                         `protobuf:"varint,16,opt,name=hmac_enabled,json=hmacEnabled,proto3" json:"hmac_enabled,omitempty"`
	HmacSecret              string                       `protobuf:"bytes,17,opt,name=hmac_secret,json=hmacSecret,proto3" json:"hmac_secret,omitempty"`
	IsInactive              bool                         `protobuf:"varint,18,opt,name=is_inactive,json=isInactive,proto3" json:"is_inactive,omitempty"`
	ApplyPolicyId           string                       `protobuf:"bytes,19,opt,name=apply_policy_id,json=applyPolicyId,proto3" json:"apply_policy_id,omitempty"`
	DataExpires             int64                        `protobuf:"varint,20,opt,name=data_expires,json=dataExpires,proto3" json:"data_expires,omitempty"`
	Monitor                 *Monitor                     `protobuf:"bytes,21,opt,name=monitor,proto3" json:"monitor,omitempty"`
	EnableDetailedRecording bool                         `protobuf:"varint,22,opt,name=enable_detailed_recording,json=enableDetailedRecording,proto3" json:"enable_detailed_recording,omitempty"`
	Metadata                map[string]string            `protobuf:"bytes,23,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags                    []string                     `protobuf:"bytes,24,rep,name=tags,proto3" json:"tags,omitempty"`
	Alias                   string                       `protobuf:"bytes,25,opt,name=alias,proto3" json:"alias,omitempty"`
	LastUpdated             string                       `protobuf:"bytes,26,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	IdExtractorDeadline     int64                        `protobuf:"varint,27,opt,name=id_extractor_deadline,json=idExtractorDeadline,proto3" json:"id_extractor_deadline,omitempty"`
	SessionLifetime         int64                        `protobuf:"varint,28,opt,name=session_lifetime,json=sessionLifetime,proto3" json:"session_lifetime,omitempty"`
	ApplyPolicies           []string                     `protobuf:"bytes,29,rep,name=apply_policies,json=applyPolicies,proto3" json:"apply_policies,omitempty"`
	Certificate             string                       `protobuf:"bytes,30,opt,name=certificate,proto3" json:"certificate,omitempty"`
	MaxQueryDepth           int64                        `protobuf:"varint,31,opt,name=max_query_depth,json=maxQueryDepth,proto3" json:"max_query_depth,omitempty"`
}

func (x *SessionState) Reset() {
	*x = SessionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_session_state_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionState) ProtoMessage() {}

func (x *SessionState) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_session_state_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionState.ProtoReflect.Descriptor instead.
func (*SessionState) Descriptor() ([]byte, []int) {
	return file_coprocess_session_state_proto_rawDescGZIP(), []int{5}
}

func (x *SessionState) GetLastCheck() int64 {
	if x != nil {
		return x.LastCheck
	}
	return 0
}

func (x *SessionState) GetAllowance() float64 {
	if x != nil {
		return x.Allowance
	}
	return 0
}

func (x *SessionState) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *SessionState) GetPer() float64 {
	if x != nil {
		return x.Per
	}
	return 0
}

func (x *SessionState) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *SessionState) GetQuotaMax() int64 {
	if x != nil {
		return x.QuotaMax
	}
	return 0
}

func (x *SessionState) GetQuotaRenews() int64 {
	if x != nil {
		return x.QuotaRenews
	}
	return 0
}

func (x *SessionState) GetQuotaRemaining() int64 {
	if x != nil {
		return x.QuotaRemaining
	}
	return 0
}

func (x *SessionState) GetQuotaRenewalRate() int64 {
	if x != nil {
		return x.QuotaRenewalRate
	}
	return 0
}

func (x *SessionState) GetAccessRights() map[string]*AccessDefinition {
	if x != nil {
		return x.AccessRights
	}
	return nil
}

func (x *SessionState) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *SessionState) GetOauthClientId() string {
	if x != nil {
		return x.OauthClientId
	}
	return ""
}

func (x *SessionState) GetOauthKeys() map[string]string {
	if x != nil {
		return x.OauthKeys
	}
	return nil
}

func (x *SessionState) GetBasicAuthData() *BasicAuthData {
	if x != nil {
		return x.BasicAuthData
	}
	return nil
}

func (x *SessionState) GetJwtData() *JWTData {
	if x != nil {
		return x.JwtData
	}
	return nil
}

func (x *SessionState) GetHmacEnabled() bool {
	if x != nil {
		return x.HmacEnabled
	}
	return false
}

func (x *SessionState) GetHmacSecret() string {
	if x != nil {
		return x.HmacSecret
	}
	return ""
}

func (x *SessionState) GetIsInactive() bool {
	if x != nil {
		return x.IsInactive
	}
	return false
}

func (x *SessionState) GetApplyPolicyId() string {
	if x != nil {
		return x.ApplyPolicyId
	}
	return ""
}

func (x *SessionState) GetDataExpires() int64 {
	if x != nil {
		return x.DataExpires
	}
	return 0
}

func (x *SessionState) GetMonitor() *Monitor {
	if x != nil {
		return x.Monitor
	}
	return nil
}

func (x *SessionState) GetEnableDetailedRecording() bool {
	if x != nil {
		return x.EnableDetailedRecording
	}
	return false
}

func (x *SessionState) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SessionState) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *SessionState) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *SessionState) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *SessionState) GetIdExtractorDeadline() int64 {
	if x != nil {
		return x.IdExtractorDeadline
	}
	return 0
}

func (x *SessionState) GetSessionLifetime() int64 {
	if x != nil {
		return x.SessionLifetime
	}
	return 0
}

func (x *SessionState) GetApplyPolicies() []string {
	if x != nil {
		return x.ApplyPolicies
	}
	return nil
}

func (x *SessionState) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *SessionState) GetMaxQueryDepth() int64 {
	if x != nil {
		return x.MaxQueryDepth
	}
	return 0
}

var File_coprocess_session_state_proto protoreflect.FileDescriptor

var file_coprocess_session_state_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x69,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x73, 0x22, 0x3f, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x22, 0x21, 0x0a, 0x07, 0x4a, 0x57, 0x54, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x30, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0xbc, 0x0b, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x6d,
	0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x4d,
	0x61, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x6e, 0x65,
	0x77, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x52,
	0x65, 0x6e, 0x65, 0x77, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2c,
	0x0a, 0x12, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x0d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x08, 0x6a, 0x77, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6a, 0x77, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x6d, 0x61, 0x63, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x6d, 0x61,
	0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x69, 0x6e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x12, 0x3a, 0x0a, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x41, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a,
	0x15, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x69, 0x64,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x1d,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x6d, 0x61, 0x78, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x70, 0x74, 0x68, 0x1a, 0x5c, 0x0a,
	0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x4f,
	0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x79, 0x6b, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f,
	0x67, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x6b, 0x2f, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coprocess_session_state_proto_rawDescOnce sync.Once
	file_coprocess_session_state_proto_rawDescData = file_coprocess_session_state_proto_rawDesc
)

func file_coprocess_session_state_proto_rawDescGZIP() []byte {
	file_coprocess_session_state_proto_rawDescOnce.Do(func() {
		file_coprocess_session_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_coprocess_session_state_proto_rawDescData)
	})
	return file_coprocess_session_state_proto_rawDescData
}

var file_coprocess_session_state_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_coprocess_session_state_proto_goTypes = []interface{}{
	(*AccessSpec)(nil),       // 0: coprocess.AccessSpec
	(*AccessDefinition)(nil), // 1: coprocess.AccessDefinition
	(*BasicAuthData)(nil),    // 2: coprocess.BasicAuthData
	(*JWTData)(nil),          // 3: coprocess.JWTData
	(*Monitor)(nil),          // 4: coprocess.Monitor
	(*SessionState)(nil),     // 5: coprocess.SessionState
	nil,                      // 6: coprocess.SessionState.AccessRightsEntry
	nil,                      // 7: coprocess.SessionState.OauthKeysEntry
	nil,                      // 8: coprocess.SessionState.MetadataEntry
}
var file_coprocess_session_state_proto_depIdxs = []int32{
	0, // 0: coprocess.AccessDefinition.allowed_urls:type_name -> coprocess.AccessSpec
	6, // 1: coprocess.SessionState.access_rights:type_name -> coprocess.SessionState.AccessRightsEntry
	7, // 2: coprocess.SessionState.oauth_keys:type_name -> coprocess.SessionState.OauthKeysEntry
	2, // 3: coprocess.SessionState.basic_auth_data:type_name -> coprocess.BasicAuthData
	3, // 4: coprocess.SessionState.jwt_data:type_name -> coprocess.JWTData
	4, // 5: coprocess.SessionState.monitor:type_name -> coprocess.Monitor
	8, // 6: coprocess.SessionState.metadata:type_name -> coprocess.SessionState.MetadataEntry
	1, // 7: coprocess.SessionState.AccessRightsEntry.value:type_name -> coprocess.AccessDefinition
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_coprocess_session_state_proto_init() }
func file_coprocess_session_state_proto_init() {
	if File_coprocess_session_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coprocess_session_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coprocess_session_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessDefinition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coprocess_session_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAuthData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coprocess_session_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coprocess_session_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coprocess_session_state_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coprocess_session_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coprocess_session_state_proto_goTypes,
		DependencyIndexes: file_coprocess_session_state_proto_depIdxs,
		MessageInfos:      file_coprocess_session_state_proto_msgTypes,
	}.Build()
	File_coprocess_session_state_proto = out.File
	file_coprocess_session_state_proto_rawDesc = nil
	file_coprocess_session_state_proto_goTypes = nil
	file_coprocess_session_state_proto_depIdxs = nil
}
