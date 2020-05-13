// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hcs.proto

package hcs

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HcsMessageRegular_Class int32

const (
	HcsMessageRegular_NORMAL HcsMessageRegular_Class = 0
	HcsMessageRegular_CONFIG HcsMessageRegular_Class = 1
)

var HcsMessageRegular_Class_name = map[int32]string{
	0: "NORMAL",
	1: "CONFIG",
}

var HcsMessageRegular_Class_value = map[string]int32{
	"NORMAL": 0,
	"CONFIG": 1,
}

func (x HcsMessageRegular_Class) String() string {
	return proto.EnumName(HcsMessageRegular_Class_name, int32(x))
}

func (HcsMessageRegular_Class) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{1, 0}
}

// HcsMessage is a wrapper type for the messages that the HCS-based
// orderer deals with.
type HcsMessage struct {
	// Types that are valid to be assigned to Type:
	//	*HcsMessage_Regular
	//	*HcsMessage_TimeToCut
	//	*HcsMessage_OrdererStarted
	Type                 isHcsMessage_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HcsMessage) Reset()         { *m = HcsMessage{} }
func (m *HcsMessage) String() string { return proto.CompactTextString(m) }
func (*HcsMessage) ProtoMessage()    {}
func (*HcsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{0}
}

func (m *HcsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessage.Unmarshal(m, b)
}
func (m *HcsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessage.Marshal(b, m, deterministic)
}
func (m *HcsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessage.Merge(m, src)
}
func (m *HcsMessage) XXX_Size() int {
	return xxx_messageInfo_HcsMessage.Size(m)
}
func (m *HcsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessage proto.InternalMessageInfo

type isHcsMessage_Type interface {
	isHcsMessage_Type()
}

type HcsMessage_Regular struct {
	Regular *HcsMessageRegular `protobuf:"bytes,1,opt,name=regular,proto3,oneof"`
}

type HcsMessage_TimeToCut struct {
	TimeToCut *HcsMessageTimeToCut `protobuf:"bytes,2,opt,name=time_to_cut,json=timeToCut,proto3,oneof"`
}

type HcsMessage_OrdererStarted struct {
	OrdererStarted *HcsMessageOrdererStarted `protobuf:"bytes,3,opt,name=orderer_started,json=ordererStarted,proto3,oneof"`
}

func (*HcsMessage_Regular) isHcsMessage_Type() {}

func (*HcsMessage_TimeToCut) isHcsMessage_Type() {}

func (*HcsMessage_OrdererStarted) isHcsMessage_Type() {}

func (m *HcsMessage) GetType() isHcsMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *HcsMessage) GetRegular() *HcsMessageRegular {
	if x, ok := m.GetType().(*HcsMessage_Regular); ok {
		return x.Regular
	}
	return nil
}

func (m *HcsMessage) GetTimeToCut() *HcsMessageTimeToCut {
	if x, ok := m.GetType().(*HcsMessage_TimeToCut); ok {
		return x.TimeToCut
	}
	return nil
}

func (m *HcsMessage) GetOrdererStarted() *HcsMessageOrdererStarted {
	if x, ok := m.GetType().(*HcsMessage_OrdererStarted); ok {
		return x.OrdererStarted
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HcsMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HcsMessage_Regular)(nil),
		(*HcsMessage_TimeToCut)(nil),
		(*HcsMessage_OrdererStarted)(nil),
	}
}

// HcsMessageRegular wraps a marshalled envelope.
type HcsMessageRegular struct {
	Payload              []byte                  `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	ConfigSeq            uint64                  `protobuf:"varint,2,opt,name=config_seq,json=configSeq,proto3" json:"config_seq,omitempty"`
	Class                HcsMessageRegular_Class `protobuf:"varint,3,opt,name=class,proto3,enum=hcs.HcsMessageRegular_Class" json:"class,omitempty"`
	OriginalSeq          uint64                  `protobuf:"varint,4,opt,name=original_seq,json=originalSeq,proto3" json:"original_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HcsMessageRegular) Reset()         { *m = HcsMessageRegular{} }
func (m *HcsMessageRegular) String() string { return proto.CompactTextString(m) }
func (*HcsMessageRegular) ProtoMessage()    {}
func (*HcsMessageRegular) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{1}
}

func (m *HcsMessageRegular) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessageRegular.Unmarshal(m, b)
}
func (m *HcsMessageRegular) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessageRegular.Marshal(b, m, deterministic)
}
func (m *HcsMessageRegular) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessageRegular.Merge(m, src)
}
func (m *HcsMessageRegular) XXX_Size() int {
	return xxx_messageInfo_HcsMessageRegular.Size(m)
}
func (m *HcsMessageRegular) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessageRegular.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessageRegular proto.InternalMessageInfo

func (m *HcsMessageRegular) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HcsMessageRegular) GetConfigSeq() uint64 {
	if m != nil {
		return m.ConfigSeq
	}
	return 0
}

func (m *HcsMessageRegular) GetClass() HcsMessageRegular_Class {
	if m != nil {
		return m.Class
	}
	return HcsMessageRegular_NORMAL
}

func (m *HcsMessageRegular) GetOriginalSeq() uint64 {
	if m != nil {
		return m.OriginalSeq
	}
	return 0
}

// HcsMessageTimeToCut is used to signal to the orderers that it is time to
// cut block <block_number>.
type HcsMessageTimeToCut struct {
	BlockNumber          uint64   `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HcsMessageTimeToCut) Reset()         { *m = HcsMessageTimeToCut{} }
func (m *HcsMessageTimeToCut) String() string { return proto.CompactTextString(m) }
func (*HcsMessageTimeToCut) ProtoMessage()    {}
func (*HcsMessageTimeToCut) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{2}
}

func (m *HcsMessageTimeToCut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessageTimeToCut.Unmarshal(m, b)
}
func (m *HcsMessageTimeToCut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessageTimeToCut.Marshal(b, m, deterministic)
}
func (m *HcsMessageTimeToCut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessageTimeToCut.Merge(m, src)
}
func (m *HcsMessageTimeToCut) XXX_Size() int {
	return xxx_messageInfo_HcsMessageTimeToCut.Size(m)
}
func (m *HcsMessageTimeToCut) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessageTimeToCut.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessageTimeToCut proto.InternalMessageInfo

func (m *HcsMessageTimeToCut) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

type HcsMessageOrdererStarted struct {
	OrdererIdentity      []byte   `protobuf:"bytes,1,opt,name=orderer_identity,json=ordererIdentity,proto3" json:"orderer_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HcsMessageOrdererStarted) Reset()         { *m = HcsMessageOrdererStarted{} }
func (m *HcsMessageOrdererStarted) String() string { return proto.CompactTextString(m) }
func (*HcsMessageOrdererStarted) ProtoMessage()    {}
func (*HcsMessageOrdererStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{3}
}

func (m *HcsMessageOrdererStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessageOrdererStarted.Unmarshal(m, b)
}
func (m *HcsMessageOrdererStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessageOrdererStarted.Marshal(b, m, deterministic)
}
func (m *HcsMessageOrdererStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessageOrdererStarted.Merge(m, src)
}
func (m *HcsMessageOrdererStarted) XXX_Size() int {
	return xxx_messageInfo_HcsMessageOrdererStarted.Size(m)
}
func (m *HcsMessageOrdererStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessageOrdererStarted.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessageOrdererStarted proto.InternalMessageInfo

func (m *HcsMessageOrdererStarted) GetOrdererIdentity() []byte {
	if m != nil {
		return m.OrdererIdentity
	}
	return nil
}

// HcsMetadata is encoded into the ORDERER block to keep track of HCS-related
// metadata associated with this block.
type HcsMetadata struct {
	// lastTimestampPersisted is the used to keep track of the timestamp of
	// the last ordered message in the last block so when an HCS-based orderer
	// boots up, it knows from when to retrieve ordererd messages
	LastConsensusTimestampPersisted *timestamp.Timestamp `protobuf:"bytes,1,opt,name=last_consensus_timestamp_persisted,json=lastConsensusTimestampPersisted,proto3" json:"last_consensus_timestamp_persisted,omitempty"`
	LastOriginalSequenceProcessed   uint64               `protobuf:"varint,2,opt,name=last_original_sequence_processed,json=lastOriginalSequenceProcessed,proto3" json:"last_original_sequence_processed,omitempty"`
	LastResubmittedConfigSequence   uint64               `protobuf:"varint,3,opt,name=last_resubmitted_config_sequence,json=lastResubmittedConfigSequence,proto3" json:"last_resubmitted_config_sequence,omitempty"`
	// last chunk free consensus timestamp is the consensus timestamp of the last
	// chunk free block
	LastChunkFreeConsensusTimestampPersisted *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_chunk_free_consensus_timestamp_persisted,json=lastChunkFreeConsensusTimestampPersisted,proto3" json:"last_chunk_free_consensus_timestamp_persisted,omitempty"`
	XXX_NoUnkeyedLiteral                     struct{}             `json:"-"`
	XXX_unrecognized                         []byte               `json:"-"`
	XXX_sizecache                            int32                `json:"-"`
}

func (m *HcsMetadata) Reset()         { *m = HcsMetadata{} }
func (m *HcsMetadata) String() string { return proto.CompactTextString(m) }
func (*HcsMetadata) ProtoMessage()    {}
func (*HcsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{4}
}

func (m *HcsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMetadata.Unmarshal(m, b)
}
func (m *HcsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMetadata.Marshal(b, m, deterministic)
}
func (m *HcsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMetadata.Merge(m, src)
}
func (m *HcsMetadata) XXX_Size() int {
	return xxx_messageInfo_HcsMetadata.Size(m)
}
func (m *HcsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMetadata proto.InternalMessageInfo

func (m *HcsMetadata) GetLastConsensusTimestampPersisted() *timestamp.Timestamp {
	if m != nil {
		return m.LastConsensusTimestampPersisted
	}
	return nil
}

func (m *HcsMetadata) GetLastOriginalSequenceProcessed() uint64 {
	if m != nil {
		return m.LastOriginalSequenceProcessed
	}
	return 0
}

func (m *HcsMetadata) GetLastResubmittedConfigSequence() uint64 {
	if m != nil {
		return m.LastResubmittedConfigSequence
	}
	return 0
}

func (m *HcsMetadata) GetLastChunkFreeConsensusTimestampPersisted() *timestamp.Timestamp {
	if m != nil {
		return m.LastChunkFreeConsensusTimestampPersisted
	}
	return nil
}

type HcsConfigPublicKey struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HcsConfigPublicKey) Reset()         { *m = HcsConfigPublicKey{} }
func (m *HcsConfigPublicKey) String() string { return proto.CompactTextString(m) }
func (*HcsConfigPublicKey) ProtoMessage()    {}
func (*HcsConfigPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{5}
}

func (m *HcsConfigPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsConfigPublicKey.Unmarshal(m, b)
}
func (m *HcsConfigPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsConfigPublicKey.Marshal(b, m, deterministic)
}
func (m *HcsConfigPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsConfigPublicKey.Merge(m, src)
}
func (m *HcsConfigPublicKey) XXX_Size() int {
	return xxx_messageInfo_HcsConfigPublicKey.Size(m)
}
func (m *HcsConfigPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsConfigPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_HcsConfigPublicKey proto.InternalMessageInfo

func (m *HcsConfigPublicKey) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *HcsConfigPublicKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type HcsConfigMetadata struct {
	// HCS topic ID in the format of shard.realm.num
	TopicId string `protobuf:"bytes,1,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	// public keys of orderer nodes for signature verification
	PublicKeys           []*HcsConfigPublicKey `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HcsConfigMetadata) Reset()         { *m = HcsConfigMetadata{} }
func (m *HcsConfigMetadata) String() string { return proto.CompactTextString(m) }
func (*HcsConfigMetadata) ProtoMessage()    {}
func (*HcsConfigMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{6}
}

func (m *HcsConfigMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsConfigMetadata.Unmarshal(m, b)
}
func (m *HcsConfigMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsConfigMetadata.Marshal(b, m, deterministic)
}
func (m *HcsConfigMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsConfigMetadata.Merge(m, src)
}
func (m *HcsConfigMetadata) XXX_Size() int {
	return xxx_messageInfo_HcsConfigMetadata.Size(m)
}
func (m *HcsConfigMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsConfigMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HcsConfigMetadata proto.InternalMessageInfo

func (m *HcsConfigMetadata) GetTopicId() string {
	if m != nil {
		return m.TopicId
	}
	return ""
}

func (m *HcsConfigMetadata) GetPublicKeys() []*HcsConfigPublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type AccountID struct {
	ShardNum             int64    `protobuf:"varint,1,opt,name=shardNum,proto3" json:"shardNum,omitempty"`
	RealmNum             int64    `protobuf:"varint,2,opt,name=realmNum,proto3" json:"realmNum,omitempty"`
	AccountNum           int64    `protobuf:"varint,3,opt,name=accountNum,proto3" json:"accountNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountID) Reset()         { *m = AccountID{} }
func (m *AccountID) String() string { return proto.CompactTextString(m) }
func (*AccountID) ProtoMessage()    {}
func (*AccountID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{7}
}

func (m *AccountID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountID.Unmarshal(m, b)
}
func (m *AccountID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountID.Marshal(b, m, deterministic)
}
func (m *AccountID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountID.Merge(m, src)
}
func (m *AccountID) XXX_Size() int {
	return xxx_messageInfo_AccountID.Size(m)
}
func (m *AccountID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountID.DiscardUnknown(m)
}

var xxx_messageInfo_AccountID proto.InternalMessageInfo

func (m *AccountID) GetShardNum() int64 {
	if m != nil {
		return m.ShardNum
	}
	return 0
}

func (m *AccountID) GetRealmNum() int64 {
	if m != nil {
		return m.RealmNum
	}
	return 0
}

func (m *AccountID) GetAccountNum() int64 {
	if m != nil {
		return m.AccountNum
	}
	return 0
}

type ApplicationMessageID struct {
	ValidStart           *timestamp.Timestamp `protobuf:"bytes,1,opt,name=validStart,proto3" json:"validStart,omitempty"`
	AccountID            *AccountID           `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Metadata             *any.Any             `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApplicationMessageID) Reset()         { *m = ApplicationMessageID{} }
func (m *ApplicationMessageID) String() string { return proto.CompactTextString(m) }
func (*ApplicationMessageID) ProtoMessage()    {}
func (*ApplicationMessageID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{8}
}

func (m *ApplicationMessageID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMessageID.Unmarshal(m, b)
}
func (m *ApplicationMessageID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMessageID.Marshal(b, m, deterministic)
}
func (m *ApplicationMessageID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMessageID.Merge(m, src)
}
func (m *ApplicationMessageID) XXX_Size() int {
	return xxx_messageInfo_ApplicationMessageID.Size(m)
}
func (m *ApplicationMessageID) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMessageID.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMessageID proto.InternalMessageInfo

func (m *ApplicationMessageID) GetValidStart() *timestamp.Timestamp {
	if m != nil {
		return m.ValidStart
	}
	return nil
}

func (m *ApplicationMessageID) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *ApplicationMessageID) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// a complete message containing data provided by an app
type ApplicationMessage struct {
	ApplicationMessageId                  *ApplicationMessageID `protobuf:"bytes,1,opt,name=applicationMessageId,proto3" json:"applicationMessageId,omitempty"`
	BusinessProcessMessage                []byte                `protobuf:"bytes,2,opt,name=businessProcessMessage,proto3" json:"businessProcessMessage,omitempty"`
	UnencryptedBusinessProcessMessageHash []byte                `protobuf:"bytes,3,opt,name=unencryptedBusinessProcessMessageHash,proto3" json:"unencryptedBusinessProcessMessageHash,omitempty"`
	BusinessProcessSignatureOnHash        []byte                `protobuf:"bytes,4,opt,name=businessProcessSignatureOnHash,proto3" json:"businessProcessSignatureOnHash,omitempty"`
	EncryptionRandom                      []byte                `protobuf:"bytes,5,opt,name=encryptionRandom,proto3" json:"encryptionRandom,omitempty"`
	Metadata                              *any.Any              `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	MetadataHash                          []byte                `protobuf:"bytes,7,opt,name=metadataHash,proto3" json:"metadataHash,omitempty"`
	MetadataSignatureOnHash               []byte                `protobuf:"bytes,8,opt,name=metadataSignatureOnHash,proto3" json:"metadataSignatureOnHash,omitempty"`
	XXX_NoUnkeyedLiteral                  struct{}              `json:"-"`
	XXX_unrecognized                      []byte                `json:"-"`
	XXX_sizecache                         int32                 `json:"-"`
}

func (m *ApplicationMessage) Reset()         { *m = ApplicationMessage{} }
func (m *ApplicationMessage) String() string { return proto.CompactTextString(m) }
func (*ApplicationMessage) ProtoMessage()    {}
func (*ApplicationMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{9}
}

func (m *ApplicationMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMessage.Unmarshal(m, b)
}
func (m *ApplicationMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMessage.Marshal(b, m, deterministic)
}
func (m *ApplicationMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMessage.Merge(m, src)
}
func (m *ApplicationMessage) XXX_Size() int {
	return xxx_messageInfo_ApplicationMessage.Size(m)
}
func (m *ApplicationMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMessage proto.InternalMessageInfo

func (m *ApplicationMessage) GetApplicationMessageId() *ApplicationMessageID {
	if m != nil {
		return m.ApplicationMessageId
	}
	return nil
}

func (m *ApplicationMessage) GetBusinessProcessMessage() []byte {
	if m != nil {
		return m.BusinessProcessMessage
	}
	return nil
}

func (m *ApplicationMessage) GetUnencryptedBusinessProcessMessageHash() []byte {
	if m != nil {
		return m.UnencryptedBusinessProcessMessageHash
	}
	return nil
}

func (m *ApplicationMessage) GetBusinessProcessSignatureOnHash() []byte {
	if m != nil {
		return m.BusinessProcessSignatureOnHash
	}
	return nil
}

func (m *ApplicationMessage) GetEncryptionRandom() []byte {
	if m != nil {
		return m.EncryptionRandom
	}
	return nil
}

func (m *ApplicationMessage) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ApplicationMessage) GetMetadataHash() []byte {
	if m != nil {
		return m.MetadataHash
	}
	return nil
}

func (m *ApplicationMessage) GetMetadataSignatureOnHash() []byte {
	if m != nil {
		return m.MetadataSignatureOnHash
	}
	return nil
}

// parts of a complete message
type ApplicationMessageChunk struct {
	ApplicationMessageId *ApplicationMessageID `protobuf:"bytes,1,opt,name=applicationMessageId,proto3" json:"applicationMessageId,omitempty"`
	ChunksCount          int32                 `protobuf:"varint,2,opt,name=chunksCount,proto3" json:"chunksCount,omitempty"`
	ChunkIndex           int32                 `protobuf:"varint,3,opt,name=chunkIndex,proto3" json:"chunkIndex,omitempty"`
	MessageChunk         []byte                `protobuf:"bytes,4,opt,name=messageChunk,proto3" json:"messageChunk,omitempty"`
	Metadata             *any.Any              `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ApplicationMessageChunk) Reset()         { *m = ApplicationMessageChunk{} }
func (m *ApplicationMessageChunk) String() string { return proto.CompactTextString(m) }
func (*ApplicationMessageChunk) ProtoMessage()    {}
func (*ApplicationMessageChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{10}
}

func (m *ApplicationMessageChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMessageChunk.Unmarshal(m, b)
}
func (m *ApplicationMessageChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMessageChunk.Marshal(b, m, deterministic)
}
func (m *ApplicationMessageChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMessageChunk.Merge(m, src)
}
func (m *ApplicationMessageChunk) XXX_Size() int {
	return xxx_messageInfo_ApplicationMessageChunk.Size(m)
}
func (m *ApplicationMessageChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMessageChunk.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMessageChunk proto.InternalMessageInfo

func (m *ApplicationMessageChunk) GetApplicationMessageId() *ApplicationMessageID {
	if m != nil {
		return m.ApplicationMessageId
	}
	return nil
}

func (m *ApplicationMessageChunk) GetChunksCount() int32 {
	if m != nil {
		return m.ChunksCount
	}
	return 0
}

func (m *ApplicationMessageChunk) GetChunkIndex() int32 {
	if m != nil {
		return m.ChunkIndex
	}
	return 0
}

func (m *ApplicationMessageChunk) GetMessageChunk() []byte {
	if m != nil {
		return m.MessageChunk
	}
	return nil
}

func (m *ApplicationMessageChunk) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// signature with public key
type ApplicationSignature struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationSignature) Reset()         { *m = ApplicationSignature{} }
func (m *ApplicationSignature) String() string { return proto.CompactTextString(m) }
func (*ApplicationSignature) ProtoMessage()    {}
func (*ApplicationSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a1dde7c60075de7, []int{11}
}

func (m *ApplicationSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationSignature.Unmarshal(m, b)
}
func (m *ApplicationSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationSignature.Marshal(b, m, deterministic)
}
func (m *ApplicationSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationSignature.Merge(m, src)
}
func (m *ApplicationSignature) XXX_Size() int {
	return xxx_messageInfo_ApplicationSignature.Size(m)
}
func (m *ApplicationSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationSignature.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationSignature proto.InternalMessageInfo

func (m *ApplicationSignature) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *ApplicationSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("hcs.HcsMessageRegular_Class", HcsMessageRegular_Class_name, HcsMessageRegular_Class_value)
	proto.RegisterType((*HcsMessage)(nil), "hcs.HcsMessage")
	proto.RegisterType((*HcsMessageRegular)(nil), "hcs.HcsMessageRegular")
	proto.RegisterType((*HcsMessageTimeToCut)(nil), "hcs.HcsMessageTimeToCut")
	proto.RegisterType((*HcsMessageOrdererStarted)(nil), "hcs.HcsMessageOrdererStarted")
	proto.RegisterType((*HcsMetadata)(nil), "hcs.HcsMetadata")
	proto.RegisterType((*HcsConfigPublicKey)(nil), "hcs.HcsConfigPublicKey")
	proto.RegisterType((*HcsConfigMetadata)(nil), "hcs.HcsConfigMetadata")
	proto.RegisterType((*AccountID)(nil), "hcs.AccountID")
	proto.RegisterType((*ApplicationMessageID)(nil), "hcs.ApplicationMessageID")
	proto.RegisterType((*ApplicationMessage)(nil), "hcs.ApplicationMessage")
	proto.RegisterType((*ApplicationMessageChunk)(nil), "hcs.ApplicationMessageChunk")
	proto.RegisterType((*ApplicationSignature)(nil), "hcs.ApplicationSignature")
}

func init() { proto.RegisterFile("hcs.proto", fileDescriptor_4a1dde7c60075de7) }

var fileDescriptor_4a1dde7c60075de7 = []byte{
	// 951 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x4e, 0xe4, 0x46,
	0x10, 0x66, 0x98, 0x19, 0x60, 0x6a, 0x10, 0x21, 0x1d, 0xb4, 0x0c, 0x68, 0x59, 0x58, 0x47, 0x91,
	0x48, 0x94, 0x0c, 0x11, 0x91, 0x22, 0xc4, 0x8d, 0x9d, 0x0d, 0x3b, 0xa3, 0x84, 0x1f, 0x35, 0x73,
	0xca, 0xc5, 0xea, 0x69, 0x37, 0xb6, 0x85, 0xdd, 0x6d, 0xdc, 0xed, 0x68, 0xad, 0x5c, 0x72, 0xc9,
	0x0b, 0xe5, 0x19, 0x72, 0xce, 0xdb, 0xe4, 0x1e, 0xb9, 0xec, 0x9e, 0x7f, 0x02, 0x87, 0xbd, 0xb9,
	0xbf, 0xfa, 0xea, 0xef, 0xab, 0xea, 0x36, 0xb4, 0x02, 0xae, 0xbb, 0x49, 0xaa, 0x8c, 0x22, 0xf5,
	0x80, 0xeb, 0xfd, 0x43, 0x5f, 0x29, 0x3f, 0x12, 0x27, 0x08, 0x8d, 0xb2, 0xfb, 0x13, 0x13, 0xc6,
	0x42, 0x1b, 0x16, 0x27, 0x25, 0x6b, 0x7f, 0x6f, 0x9e, 0xc0, 0x64, 0x5e, 0x9a, 0x9c, 0x7f, 0x6a,
	0x00, 0x7d, 0xae, 0xaf, 0x84, 0xd6, 0xcc, 0x17, 0xe4, 0x14, 0xd6, 0x53, 0xe1, 0x67, 0x11, 0x4b,
	0x3b, 0xb5, 0xa3, 0xda, 0x71, 0xfb, 0xf4, 0x55, 0xb7, 0x48, 0x36, 0x61, 0xd0, 0xd2, 0xda, 0x5f,
	0xa1, 0x96, 0x48, 0xce, 0xa1, 0x5d, 0x24, 0x74, 0x8d, 0x72, 0x79, 0x66, 0x3a, 0xab, 0xe8, 0xd7,
	0x99, 0xf3, 0x1b, 0x86, 0xb1, 0x18, 0xaa, 0x5e, 0x66, 0xfa, 0x2b, 0xb4, 0x65, 0xec, 0x81, 0xf4,
	0xe1, 0x33, 0x95, 0x7a, 0x22, 0x15, 0xa9, 0xab, 0x0d, 0x4b, 0x8d, 0xf0, 0x3a, 0x75, 0xf4, 0x3f,
	0x98, 0xf3, 0xbf, 0x29, 0x59, 0x77, 0x25, 0xa9, 0xbf, 0x42, 0xb7, 0xd4, 0x0c, 0xf2, 0x6e, 0x0d,
	0x1a, 0xc3, 0x3c, 0x11, 0xce, 0xdf, 0x35, 0xf8, 0x7c, 0xa1, 0x5c, 0xd2, 0x81, 0xf5, 0x84, 0xe5,
	0x91, 0x62, 0x1e, 0xf6, 0xb5, 0x49, 0xed, 0x91, 0x1c, 0x00, 0x70, 0x25, 0xef, 0x43, 0xdf, 0xd5,
	0xe2, 0x11, 0x8b, 0x6f, 0xd0, 0x56, 0x89, 0xdc, 0x89, 0x47, 0x72, 0x0a, 0x4d, 0x1e, 0x31, 0xad,
	0xb1, 0xac, 0xad, 0xd3, 0xd7, 0xcb, 0xe5, 0xe8, 0xf6, 0x0a, 0x0e, 0x2d, 0xa9, 0xe4, 0x2d, 0x6c,
	0xaa, 0x34, 0xf4, 0x43, 0xc9, 0x22, 0x0c, 0xda, 0xc0, 0xa0, 0x6d, 0x8b, 0xdd, 0x89, 0x47, 0xe7,
	0x10, 0x9a, 0xe8, 0x42, 0x00, 0xd6, 0xae, 0x6f, 0xe8, 0xd5, 0xc5, 0x2f, 0xdb, 0x2b, 0xc5, 0x77,
	0xef, 0xe6, 0xfa, 0x72, 0xf0, 0x61, 0xbb, 0xe6, 0x9c, 0xc1, 0x17, 0x4b, 0xc4, 0x2b, 0x42, 0x8f,
	0x22, 0xc5, 0x1f, 0x5c, 0x99, 0xc5, 0x23, 0x51, 0x0e, 0xa9, 0x41, 0xdb, 0x88, 0x5d, 0x23, 0xe4,
	0xfc, 0x04, 0x9d, 0xa7, 0x64, 0x23, 0x5f, 0xc3, 0xb6, 0x95, 0x3b, 0xf4, 0x84, 0x34, 0xa1, 0xc9,
	0x2b, 0x3d, 0xec, 0x18, 0x06, 0x15, 0xec, 0xfc, 0x59, 0x87, 0x36, 0xc6, 0x31, 0xcc, 0x63, 0x86,
	0x11, 0x1f, 0x9c, 0x88, 0x69, 0xe3, 0x72, 0x25, 0xb5, 0x90, 0x3a, 0xd3, 0xee, 0x78, 0xcb, 0xdc,
	0x44, 0xa4, 0x3a, 0xd4, 0xc5, 0xf0, 0xca, 0xa5, 0xd9, 0xef, 0x96, 0x0b, 0xd7, 0xb5, 0x0b, 0xd7,
	0x1d, 0x5a, 0x2e, 0x3d, 0x2c, 0xa2, 0xf4, 0x6c, 0x90, 0x31, 0x7e, 0x6b, 0x43, 0x90, 0x0f, 0x70,
	0x84, 0x89, 0xa6, 0x25, 0xcc, 0x84, 0xe4, 0xc2, 0x4d, 0x52, 0xc5, 0x85, 0xd6, 0xc2, 0xab, 0xc6,
	0x74, 0x50, 0xf0, 0x6e, 0x26, 0xaa, 0x22, 0xeb, 0xd6, 0x92, 0xc6, 0x81, 0x52, 0xa1, 0xb3, 0x51,
	0x1c, 0x1a, 0x23, 0x3c, 0x77, 0x32, 0x6a, 0x24, 0xe3, 0x54, 0xab, 0x40, 0x74, 0x42, 0xeb, 0xd9,
	0xf1, 0x23, 0x89, 0xfc, 0x0e, 0xdf, 0x95, 0xad, 0x07, 0x99, 0x7c, 0x70, 0xef, 0x53, 0x21, 0x9e,
	0x51, 0xa1, 0xf1, 0xac, 0x0a, 0xc7, 0xa8, 0x42, 0x11, 0xef, 0x32, 0x15, 0xe2, 0x7f, 0xe4, 0x70,
	0xce, 0x81, 0xf4, 0xb9, 0x2e, 0x2b, 0xba, 0xcd, 0x46, 0x51, 0xc8, 0x7f, 0x16, 0x39, 0x21, 0xd0,
	0x30, 0x79, 0x22, 0x50, 0xef, 0x16, 0xc5, 0x6f, 0xb2, 0x0d, 0xf5, 0x07, 0x91, 0xa3, 0x36, 0x2d,
	0x5a, 0x7c, 0x3a, 0x01, 0x5e, 0x85, 0xd2, 0x77, 0x3c, 0xc8, 0x3d, 0xd8, 0x30, 0x2a, 0x09, 0xb9,
	0x1b, 0x7a, 0x95, 0xfb, 0x3a, 0x9e, 0x07, 0x1e, 0x39, 0x83, 0x76, 0x82, 0x29, 0xdc, 0x07, 0x91,
	0xeb, 0xce, 0xea, 0x51, 0xfd, 0xb8, 0x7d, 0xba, 0x6b, 0x57, 0x7e, 0xae, 0x06, 0x0a, 0x89, 0xfd,
	0xd4, 0x0e, 0x87, 0xd6, 0x05, 0xe7, 0x2a, 0x93, 0x66, 0xf0, 0x9e, 0xec, 0xc3, 0x86, 0x0e, 0x58,
	0xea, 0x5d, 0x67, 0x31, 0x66, 0xa8, 0xd3, 0xf1, 0xb9, 0xb0, 0xa5, 0x82, 0x45, 0x71, 0x61, 0x5b,
	0x2d, 0x6d, 0xf6, 0x4c, 0xde, 0x00, 0xb0, 0x32, 0x48, 0x61, 0xad, 0xa3, 0x75, 0x0a, 0x71, 0xfe,
	0xaa, 0xc1, 0xce, 0x45, 0x92, 0x44, 0x21, 0x67, 0x26, 0x54, 0xb2, 0x5a, 0xf1, 0xc1, 0x7b, 0x72,
	0x0e, 0xf0, 0x1b, 0x8b, 0x42, 0x0f, 0xd7, 0xfc, 0x05, 0x3b, 0x38, 0xc5, 0x26, 0xdf, 0x42, 0x8b,
	0xd9, 0xca, 0xab, 0xb7, 0x6b, 0x0b, 0x3b, 0x1e, 0xf7, 0x43, 0x27, 0x04, 0xf2, 0x3d, 0x6c, 0xc4,
	0x95, 0x90, 0xd5, 0x43, 0xb5, 0xb3, 0x90, 0xe7, 0x42, 0xe6, 0x74, 0xcc, 0x72, 0xfe, 0xad, 0x03,
	0x59, 0x2c, 0x9a, 0x5c, 0xc1, 0x0e, 0x5b, 0x6c, 0xc5, 0x5e, 0xa0, 0xbd, 0xb2, 0x82, 0x25, 0xbd,
	0xd2, 0xa5, 0x6e, 0xe4, 0x47, 0x78, 0x35, 0xca, 0x74, 0x28, 0x85, 0xd6, 0xd5, 0x05, 0xa8, 0x6c,
	0xd8, 0xd2, 0x26, 0x7d, 0xc2, 0x4a, 0x86, 0xf0, 0x55, 0x26, 0x85, 0xe4, 0x69, 0x9e, 0x14, 0x8f,
	0xe8, 0x52, 0x52, 0x9f, 0xe9, 0x00, 0x9b, 0xdd, 0xa4, 0x2f, 0x23, 0x93, 0x4b, 0x78, 0x33, 0x97,
	0xef, 0x2e, 0xf4, 0x25, 0x33, 0x59, 0x2a, 0x6e, 0x24, 0x86, 0x6b, 0x60, 0xb8, 0x67, 0x58, 0xe4,
	0x1b, 0xd8, 0xae, 0xd2, 0x85, 0x4a, 0x52, 0x26, 0x3d, 0x15, 0x77, 0x9a, 0xe8, 0xb9, 0x80, 0xcf,
	0x4c, 0x66, 0xed, 0x25, 0x93, 0x21, 0x0e, 0x6c, 0xda, 0x6f, 0xac, 0x69, 0x1d, 0x23, 0xcf, 0x60,
	0xe4, 0x0c, 0x76, 0xed, 0x79, 0xbe, 0x85, 0x0d, 0xa4, 0x3f, 0x65, 0x76, 0xfe, 0x58, 0x85, 0xdd,
	0xc5, 0x01, 0xe2, 0x95, 0xff, 0xd4, 0xc3, 0x3f, 0x82, 0x36, 0x3e, 0x4d, 0xba, 0x57, 0x6c, 0x29,
	0x4e, 0xbc, 0x49, 0xa7, 0xa1, 0xe2, 0x66, 0xe1, 0x71, 0x20, 0x3d, 0xf1, 0x11, 0x67, 0xd9, 0xa4,
	0x53, 0x48, 0x29, 0xc5, 0xa4, 0xc0, 0x6a, 0x3c, 0x33, 0xd8, 0x8c, 0xc0, 0xcd, 0x17, 0xad, 0x3e,
	0x9d, 0xb9, 0xae, 0x63, 0x81, 0xc8, 0x6b, 0x68, 0x8d, 0x9f, 0x8e, 0xea, 0xf7, 0x33, 0x01, 0x0a,
	0xab, 0xb6, 0xd4, 0x6a, 0x7b, 0x27, 0xc0, 0xbb, 0x2f, 0x7f, 0x7d, 0xeb, 0x87, 0x26, 0xc8, 0x46,
	0x5d, 0xae, 0xe2, 0x93, 0x80, 0xe9, 0xc0, 0x4f, 0x59, 0x12, 0x9c, 0x98, 0x8f, 0x31, 0x4b, 0x12,
	0x91, 0x9e, 0x04, 0x5c, 0x8f, 0xd6, 0xb0, 0xa0, 0x1f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0xc2, 0x33, 0x2c, 0x29, 0x09, 0x00, 0x00,
}
