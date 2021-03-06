// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_query.proto

package mgmt

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type BioHealthReq struct {
	DevUuid              string   `protobuf:"bytes,1,opt,name=dev_uuid,json=devUuid,proto3" json:"dev_uuid,omitempty"`
	TgtId                string   `protobuf:"bytes,2,opt,name=tgt_id,json=tgtId,proto3" json:"tgt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BioHealthReq) Reset()         { *m = BioHealthReq{} }
func (m *BioHealthReq) String() string { return proto.CompactTextString(m) }
func (*BioHealthReq) ProtoMessage()    {}
func (*BioHealthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{0}
}

func (m *BioHealthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BioHealthReq.Unmarshal(m, b)
}
func (m *BioHealthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BioHealthReq.Marshal(b, m, deterministic)
}
func (m *BioHealthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BioHealthReq.Merge(m, src)
}
func (m *BioHealthReq) XXX_Size() int {
	return xxx_messageInfo_BioHealthReq.Size(m)
}
func (m *BioHealthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BioHealthReq.DiscardUnknown(m)
}

var xxx_messageInfo_BioHealthReq proto.InternalMessageInfo

func (m *BioHealthReq) GetDevUuid() string {
	if m != nil {
		return m.DevUuid
	}
	return ""
}

func (m *BioHealthReq) GetTgtId() string {
	if m != nil {
		return m.TgtId
	}
	return ""
}

type BioHealthResp struct {
	Status                int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	DevUuid               string   `protobuf:"bytes,2,opt,name=dev_uuid,json=devUuid,proto3" json:"dev_uuid,omitempty"`
	ErrorCount            uint64   `protobuf:"varint,3,opt,name=error_count,json=errorCount,proto3" json:"error_count,omitempty"`
	Temperature           uint32   `protobuf:"varint,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
	MediaErrors           uint64   `protobuf:"varint,5,opt,name=media_errors,json=mediaErrors,proto3" json:"media_errors,omitempty"`
	ReadErrors            uint32   `protobuf:"varint,6,opt,name=read_errors,json=readErrors,proto3" json:"read_errors,omitempty"`
	WriteErrors           uint32   `protobuf:"varint,7,opt,name=write_errors,json=writeErrors,proto3" json:"write_errors,omitempty"`
	UnmapErrors           uint32   `protobuf:"varint,8,opt,name=unmap_errors,json=unmapErrors,proto3" json:"unmap_errors,omitempty"`
	ChecksumErrors        uint32   `protobuf:"varint,9,opt,name=checksum_errors,json=checksumErrors,proto3" json:"checksum_errors,omitempty"`
	TempWarn              bool     `protobuf:"varint,10,opt,name=temp_warn,json=tempWarn,proto3" json:"temp_warn,omitempty"`
	SpareWarn             bool     `protobuf:"varint,11,opt,name=spare_warn,json=spareWarn,proto3" json:"spare_warn,omitempty"`
	ReadonlyWarn          bool     `protobuf:"varint,12,opt,name=readonly_warn,json=readonlyWarn,proto3" json:"readonly_warn,omitempty"`
	DeviceReliabilityWarn bool     `protobuf:"varint,13,opt,name=device_reliability_warn,json=deviceReliabilityWarn,proto3" json:"device_reliability_warn,omitempty"`
	VolatileMemoryWarn    bool     `protobuf:"varint,14,opt,name=volatile_memory_warn,json=volatileMemoryWarn,proto3" json:"volatile_memory_warn,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *BioHealthResp) Reset()         { *m = BioHealthResp{} }
func (m *BioHealthResp) String() string { return proto.CompactTextString(m) }
func (*BioHealthResp) ProtoMessage()    {}
func (*BioHealthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{1}
}

func (m *BioHealthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BioHealthResp.Unmarshal(m, b)
}
func (m *BioHealthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BioHealthResp.Marshal(b, m, deterministic)
}
func (m *BioHealthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BioHealthResp.Merge(m, src)
}
func (m *BioHealthResp) XXX_Size() int {
	return xxx_messageInfo_BioHealthResp.Size(m)
}
func (m *BioHealthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BioHealthResp.DiscardUnknown(m)
}

var xxx_messageInfo_BioHealthResp proto.InternalMessageInfo

func (m *BioHealthResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BioHealthResp) GetDevUuid() string {
	if m != nil {
		return m.DevUuid
	}
	return ""
}

func (m *BioHealthResp) GetErrorCount() uint64 {
	if m != nil {
		return m.ErrorCount
	}
	return 0
}

func (m *BioHealthResp) GetTemperature() uint32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *BioHealthResp) GetMediaErrors() uint64 {
	if m != nil {
		return m.MediaErrors
	}
	return 0
}

func (m *BioHealthResp) GetReadErrors() uint32 {
	if m != nil {
		return m.ReadErrors
	}
	return 0
}

func (m *BioHealthResp) GetWriteErrors() uint32 {
	if m != nil {
		return m.WriteErrors
	}
	return 0
}

func (m *BioHealthResp) GetUnmapErrors() uint32 {
	if m != nil {
		return m.UnmapErrors
	}
	return 0
}

func (m *BioHealthResp) GetChecksumErrors() uint32 {
	if m != nil {
		return m.ChecksumErrors
	}
	return 0
}

func (m *BioHealthResp) GetTempWarn() bool {
	if m != nil {
		return m.TempWarn
	}
	return false
}

func (m *BioHealthResp) GetSpareWarn() bool {
	if m != nil {
		return m.SpareWarn
	}
	return false
}

func (m *BioHealthResp) GetReadonlyWarn() bool {
	if m != nil {
		return m.ReadonlyWarn
	}
	return false
}

func (m *BioHealthResp) GetDeviceReliabilityWarn() bool {
	if m != nil {
		return m.DeviceReliabilityWarn
	}
	return false
}

func (m *BioHealthResp) GetVolatileMemoryWarn() bool {
	if m != nil {
		return m.VolatileMemoryWarn
	}
	return false
}

type SmdDevReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmdDevReq) Reset()         { *m = SmdDevReq{} }
func (m *SmdDevReq) String() string { return proto.CompactTextString(m) }
func (*SmdDevReq) ProtoMessage()    {}
func (*SmdDevReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{2}
}

func (m *SmdDevReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdDevReq.Unmarshal(m, b)
}
func (m *SmdDevReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdDevReq.Marshal(b, m, deterministic)
}
func (m *SmdDevReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdDevReq.Merge(m, src)
}
func (m *SmdDevReq) XXX_Size() int {
	return xxx_messageInfo_SmdDevReq.Size(m)
}
func (m *SmdDevReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdDevReq.DiscardUnknown(m)
}

var xxx_messageInfo_SmdDevReq proto.InternalMessageInfo

type SmdDevResp struct {
	Status               int32                `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Devices              []*SmdDevResp_Device `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SmdDevResp) Reset()         { *m = SmdDevResp{} }
func (m *SmdDevResp) String() string { return proto.CompactTextString(m) }
func (*SmdDevResp) ProtoMessage()    {}
func (*SmdDevResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{3}
}

func (m *SmdDevResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdDevResp.Unmarshal(m, b)
}
func (m *SmdDevResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdDevResp.Marshal(b, m, deterministic)
}
func (m *SmdDevResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdDevResp.Merge(m, src)
}
func (m *SmdDevResp) XXX_Size() int {
	return xxx_messageInfo_SmdDevResp.Size(m)
}
func (m *SmdDevResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdDevResp.DiscardUnknown(m)
}

var xxx_messageInfo_SmdDevResp proto.InternalMessageInfo

func (m *SmdDevResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SmdDevResp) GetDevices() []*SmdDevResp_Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type SmdDevResp_Device struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TgtIds               []int32  `protobuf:"varint,2,rep,packed,name=tgt_ids,json=tgtIds,proto3" json:"tgt_ids,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmdDevResp_Device) Reset()         { *m = SmdDevResp_Device{} }
func (m *SmdDevResp_Device) String() string { return proto.CompactTextString(m) }
func (*SmdDevResp_Device) ProtoMessage()    {}
func (*SmdDevResp_Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{3, 0}
}

func (m *SmdDevResp_Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdDevResp_Device.Unmarshal(m, b)
}
func (m *SmdDevResp_Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdDevResp_Device.Marshal(b, m, deterministic)
}
func (m *SmdDevResp_Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdDevResp_Device.Merge(m, src)
}
func (m *SmdDevResp_Device) XXX_Size() int {
	return xxx_messageInfo_SmdDevResp_Device.Size(m)
}
func (m *SmdDevResp_Device) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdDevResp_Device.DiscardUnknown(m)
}

var xxx_messageInfo_SmdDevResp_Device proto.InternalMessageInfo

func (m *SmdDevResp_Device) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SmdDevResp_Device) GetTgtIds() []int32 {
	if m != nil {
		return m.TgtIds
	}
	return nil
}

func (m *SmdDevResp_Device) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type SmdPoolReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmdPoolReq) Reset()         { *m = SmdPoolReq{} }
func (m *SmdPoolReq) String() string { return proto.CompactTextString(m) }
func (*SmdPoolReq) ProtoMessage()    {}
func (*SmdPoolReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{4}
}

func (m *SmdPoolReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdPoolReq.Unmarshal(m, b)
}
func (m *SmdPoolReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdPoolReq.Marshal(b, m, deterministic)
}
func (m *SmdPoolReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdPoolReq.Merge(m, src)
}
func (m *SmdPoolReq) XXX_Size() int {
	return xxx_messageInfo_SmdPoolReq.Size(m)
}
func (m *SmdPoolReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdPoolReq.DiscardUnknown(m)
}

var xxx_messageInfo_SmdPoolReq proto.InternalMessageInfo

type SmdPoolResp struct {
	Status               int32               `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Pools                []*SmdPoolResp_Pool `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SmdPoolResp) Reset()         { *m = SmdPoolResp{} }
func (m *SmdPoolResp) String() string { return proto.CompactTextString(m) }
func (*SmdPoolResp) ProtoMessage()    {}
func (*SmdPoolResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{5}
}

func (m *SmdPoolResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdPoolResp.Unmarshal(m, b)
}
func (m *SmdPoolResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdPoolResp.Marshal(b, m, deterministic)
}
func (m *SmdPoolResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdPoolResp.Merge(m, src)
}
func (m *SmdPoolResp) XXX_Size() int {
	return xxx_messageInfo_SmdPoolResp.Size(m)
}
func (m *SmdPoolResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdPoolResp.DiscardUnknown(m)
}

var xxx_messageInfo_SmdPoolResp proto.InternalMessageInfo

func (m *SmdPoolResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SmdPoolResp) GetPools() []*SmdPoolResp_Pool {
	if m != nil {
		return m.Pools
	}
	return nil
}

type SmdPoolResp_Pool struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TgtIds               []int32  `protobuf:"varint,2,rep,packed,name=tgt_ids,json=tgtIds,proto3" json:"tgt_ids,omitempty"`
	Blobs                []uint64 `protobuf:"varint,3,rep,packed,name=blobs,proto3" json:"blobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmdPoolResp_Pool) Reset()         { *m = SmdPoolResp_Pool{} }
func (m *SmdPoolResp_Pool) String() string { return proto.CompactTextString(m) }
func (*SmdPoolResp_Pool) ProtoMessage()    {}
func (*SmdPoolResp_Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{5, 0}
}

func (m *SmdPoolResp_Pool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdPoolResp_Pool.Unmarshal(m, b)
}
func (m *SmdPoolResp_Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdPoolResp_Pool.Marshal(b, m, deterministic)
}
func (m *SmdPoolResp_Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdPoolResp_Pool.Merge(m, src)
}
func (m *SmdPoolResp_Pool) XXX_Size() int {
	return xxx_messageInfo_SmdPoolResp_Pool.Size(m)
}
func (m *SmdPoolResp_Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdPoolResp_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_SmdPoolResp_Pool proto.InternalMessageInfo

func (m *SmdPoolResp_Pool) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SmdPoolResp_Pool) GetTgtIds() []int32 {
	if m != nil {
		return m.TgtIds
	}
	return nil
}

func (m *SmdPoolResp_Pool) GetBlobs() []uint64 {
	if m != nil {
		return m.Blobs
	}
	return nil
}

type DevStateReq struct {
	DevUuid              string   `protobuf:"bytes,1,opt,name=dev_uuid,json=devUuid,proto3" json:"dev_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DevStateReq) Reset()         { *m = DevStateReq{} }
func (m *DevStateReq) String() string { return proto.CompactTextString(m) }
func (*DevStateReq) ProtoMessage()    {}
func (*DevStateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{6}
}

func (m *DevStateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevStateReq.Unmarshal(m, b)
}
func (m *DevStateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevStateReq.Marshal(b, m, deterministic)
}
func (m *DevStateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevStateReq.Merge(m, src)
}
func (m *DevStateReq) XXX_Size() int {
	return xxx_messageInfo_DevStateReq.Size(m)
}
func (m *DevStateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DevStateReq.DiscardUnknown(m)
}

var xxx_messageInfo_DevStateReq proto.InternalMessageInfo

func (m *DevStateReq) GetDevUuid() string {
	if m != nil {
		return m.DevUuid
	}
	return ""
}

type DevStateResp struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	DevUuid              string   `protobuf:"bytes,2,opt,name=dev_uuid,json=devUuid,proto3" json:"dev_uuid,omitempty"`
	DevState             string   `protobuf:"bytes,3,opt,name=dev_state,json=devState,proto3" json:"dev_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DevStateResp) Reset()         { *m = DevStateResp{} }
func (m *DevStateResp) String() string { return proto.CompactTextString(m) }
func (*DevStateResp) ProtoMessage()    {}
func (*DevStateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{7}
}

func (m *DevStateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevStateResp.Unmarshal(m, b)
}
func (m *DevStateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevStateResp.Marshal(b, m, deterministic)
}
func (m *DevStateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevStateResp.Merge(m, src)
}
func (m *DevStateResp) XXX_Size() int {
	return xxx_messageInfo_DevStateResp.Size(m)
}
func (m *DevStateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DevStateResp.DiscardUnknown(m)
}

var xxx_messageInfo_DevStateResp proto.InternalMessageInfo

func (m *DevStateResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DevStateResp) GetDevUuid() string {
	if m != nil {
		return m.DevUuid
	}
	return ""
}

func (m *DevStateResp) GetDevState() string {
	if m != nil {
		return m.DevState
	}
	return ""
}

type SmdQueryReq struct {
	OmitDevices          bool     `protobuf:"varint,1,opt,name=omitDevices,proto3" json:"omitDevices,omitempty"`
	OmitPools            bool     `protobuf:"varint,2,opt,name=omitPools,proto3" json:"omitPools,omitempty"`
	IncludeBioHealth     bool     `protobuf:"varint,3,opt,name=includeBioHealth,proto3" json:"includeBioHealth,omitempty"`
	SetFaulty            bool     `protobuf:"varint,4,opt,name=setFaulty,proto3" json:"setFaulty,omitempty"`
	Uuid                 string   `protobuf:"bytes,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Rank                 uint32   `protobuf:"varint,6,opt,name=rank,proto3" json:"rank,omitempty"`
	Target               string   `protobuf:"bytes,7,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmdQueryReq) Reset()         { *m = SmdQueryReq{} }
func (m *SmdQueryReq) String() string { return proto.CompactTextString(m) }
func (*SmdQueryReq) ProtoMessage()    {}
func (*SmdQueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{8}
}

func (m *SmdQueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdQueryReq.Unmarshal(m, b)
}
func (m *SmdQueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdQueryReq.Marshal(b, m, deterministic)
}
func (m *SmdQueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdQueryReq.Merge(m, src)
}
func (m *SmdQueryReq) XXX_Size() int {
	return xxx_messageInfo_SmdQueryReq.Size(m)
}
func (m *SmdQueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdQueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_SmdQueryReq proto.InternalMessageInfo

func (m *SmdQueryReq) GetOmitDevices() bool {
	if m != nil {
		return m.OmitDevices
	}
	return false
}

func (m *SmdQueryReq) GetOmitPools() bool {
	if m != nil {
		return m.OmitPools
	}
	return false
}

func (m *SmdQueryReq) GetIncludeBioHealth() bool {
	if m != nil {
		return m.IncludeBioHealth
	}
	return false
}

func (m *SmdQueryReq) GetSetFaulty() bool {
	if m != nil {
		return m.SetFaulty
	}
	return false
}

func (m *SmdQueryReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SmdQueryReq) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *SmdQueryReq) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type SmdQueryResp struct {
	Status               int32                    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Ranks                []*SmdQueryResp_RankResp `protobuf:"bytes,2,rep,name=ranks,proto3" json:"ranks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SmdQueryResp) Reset()         { *m = SmdQueryResp{} }
func (m *SmdQueryResp) String() string { return proto.CompactTextString(m) }
func (*SmdQueryResp) ProtoMessage()    {}
func (*SmdQueryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{9}
}

func (m *SmdQueryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdQueryResp.Unmarshal(m, b)
}
func (m *SmdQueryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdQueryResp.Marshal(b, m, deterministic)
}
func (m *SmdQueryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdQueryResp.Merge(m, src)
}
func (m *SmdQueryResp) XXX_Size() int {
	return xxx_messageInfo_SmdQueryResp.Size(m)
}
func (m *SmdQueryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdQueryResp.DiscardUnknown(m)
}

var xxx_messageInfo_SmdQueryResp proto.InternalMessageInfo

func (m *SmdQueryResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SmdQueryResp) GetRanks() []*SmdQueryResp_RankResp {
	if m != nil {
		return m.Ranks
	}
	return nil
}

type SmdQueryResp_Device struct {
	Uuid                 string         `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TgtIds               []int32        `protobuf:"varint,2,rep,packed,name=tgt_ids,json=tgtIds,proto3" json:"tgt_ids,omitempty"`
	State                string         `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Health               *BioHealthResp `protobuf:"bytes,4,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SmdQueryResp_Device) Reset()         { *m = SmdQueryResp_Device{} }
func (m *SmdQueryResp_Device) String() string { return proto.CompactTextString(m) }
func (*SmdQueryResp_Device) ProtoMessage()    {}
func (*SmdQueryResp_Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{9, 0}
}

func (m *SmdQueryResp_Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdQueryResp_Device.Unmarshal(m, b)
}
func (m *SmdQueryResp_Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdQueryResp_Device.Marshal(b, m, deterministic)
}
func (m *SmdQueryResp_Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdQueryResp_Device.Merge(m, src)
}
func (m *SmdQueryResp_Device) XXX_Size() int {
	return xxx_messageInfo_SmdQueryResp_Device.Size(m)
}
func (m *SmdQueryResp_Device) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdQueryResp_Device.DiscardUnknown(m)
}

var xxx_messageInfo_SmdQueryResp_Device proto.InternalMessageInfo

func (m *SmdQueryResp_Device) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SmdQueryResp_Device) GetTgtIds() []int32 {
	if m != nil {
		return m.TgtIds
	}
	return nil
}

func (m *SmdQueryResp_Device) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *SmdQueryResp_Device) GetHealth() *BioHealthResp {
	if m != nil {
		return m.Health
	}
	return nil
}

type SmdQueryResp_Pool struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TgtIds               []int32  `protobuf:"varint,2,rep,packed,name=tgt_ids,json=tgtIds,proto3" json:"tgt_ids,omitempty"`
	Blobs                []uint64 `protobuf:"varint,3,rep,packed,name=blobs,proto3" json:"blobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmdQueryResp_Pool) Reset()         { *m = SmdQueryResp_Pool{} }
func (m *SmdQueryResp_Pool) String() string { return proto.CompactTextString(m) }
func (*SmdQueryResp_Pool) ProtoMessage()    {}
func (*SmdQueryResp_Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{9, 1}
}

func (m *SmdQueryResp_Pool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdQueryResp_Pool.Unmarshal(m, b)
}
func (m *SmdQueryResp_Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdQueryResp_Pool.Marshal(b, m, deterministic)
}
func (m *SmdQueryResp_Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdQueryResp_Pool.Merge(m, src)
}
func (m *SmdQueryResp_Pool) XXX_Size() int {
	return xxx_messageInfo_SmdQueryResp_Pool.Size(m)
}
func (m *SmdQueryResp_Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdQueryResp_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_SmdQueryResp_Pool proto.InternalMessageInfo

func (m *SmdQueryResp_Pool) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SmdQueryResp_Pool) GetTgtIds() []int32 {
	if m != nil {
		return m.TgtIds
	}
	return nil
}

func (m *SmdQueryResp_Pool) GetBlobs() []uint64 {
	if m != nil {
		return m.Blobs
	}
	return nil
}

type SmdQueryResp_RankResp struct {
	Rank                 uint32                 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Devices              []*SmdQueryResp_Device `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
	Pools                []*SmdQueryResp_Pool   `protobuf:"bytes,3,rep,name=pools,proto3" json:"pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SmdQueryResp_RankResp) Reset()         { *m = SmdQueryResp_RankResp{} }
func (m *SmdQueryResp_RankResp) String() string { return proto.CompactTextString(m) }
func (*SmdQueryResp_RankResp) ProtoMessage()    {}
func (*SmdQueryResp_RankResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87a8d20722a9416, []int{9, 2}
}

func (m *SmdQueryResp_RankResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmdQueryResp_RankResp.Unmarshal(m, b)
}
func (m *SmdQueryResp_RankResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmdQueryResp_RankResp.Marshal(b, m, deterministic)
}
func (m *SmdQueryResp_RankResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmdQueryResp_RankResp.Merge(m, src)
}
func (m *SmdQueryResp_RankResp) XXX_Size() int {
	return xxx_messageInfo_SmdQueryResp_RankResp.Size(m)
}
func (m *SmdQueryResp_RankResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SmdQueryResp_RankResp.DiscardUnknown(m)
}

var xxx_messageInfo_SmdQueryResp_RankResp proto.InternalMessageInfo

func (m *SmdQueryResp_RankResp) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *SmdQueryResp_RankResp) GetDevices() []*SmdQueryResp_Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *SmdQueryResp_RankResp) GetPools() []*SmdQueryResp_Pool {
	if m != nil {
		return m.Pools
	}
	return nil
}

func init() {
	proto.RegisterType((*BioHealthReq)(nil), "mgmt.BioHealthReq")
	proto.RegisterType((*BioHealthResp)(nil), "mgmt.BioHealthResp")
	proto.RegisterType((*SmdDevReq)(nil), "mgmt.SmdDevReq")
	proto.RegisterType((*SmdDevResp)(nil), "mgmt.SmdDevResp")
	proto.RegisterType((*SmdDevResp_Device)(nil), "mgmt.SmdDevResp.Device")
	proto.RegisterType((*SmdPoolReq)(nil), "mgmt.SmdPoolReq")
	proto.RegisterType((*SmdPoolResp)(nil), "mgmt.SmdPoolResp")
	proto.RegisterType((*SmdPoolResp_Pool)(nil), "mgmt.SmdPoolResp.Pool")
	proto.RegisterType((*DevStateReq)(nil), "mgmt.DevStateReq")
	proto.RegisterType((*DevStateResp)(nil), "mgmt.DevStateResp")
	proto.RegisterType((*SmdQueryReq)(nil), "mgmt.SmdQueryReq")
	proto.RegisterType((*SmdQueryResp)(nil), "mgmt.SmdQueryResp")
	proto.RegisterType((*SmdQueryResp_Device)(nil), "mgmt.SmdQueryResp.Device")
	proto.RegisterType((*SmdQueryResp_Pool)(nil), "mgmt.SmdQueryResp.Pool")
	proto.RegisterType((*SmdQueryResp_RankResp)(nil), "mgmt.SmdQueryResp.RankResp")
}

func init() {
	proto.RegisterFile("storage_query.proto", fileDescriptor_d87a8d20722a9416)
}

var fileDescriptor_d87a8d20722a9416 = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6f, 0xd3, 0x48,
	0x14, 0x95, 0x1b, 0x3b, 0x89, 0xaf, 0x9d, 0xee, 0x6a, 0xfa, 0xe5, 0xa6, 0xbb, 0x5a, 0xd7, 0xfb,
	0xb0, 0xd1, 0x02, 0x11, 0xb4, 0x12, 0xcf, 0x08, 0x02, 0xa2, 0x42, 0x48, 0x65, 0x2a, 0xc4, 0x1b,
	0xd6, 0x24, 0x1e, 0xa5, 0x56, 0xfd, 0x91, 0x8e, 0xc7, 0x29, 0x7d, 0x85, 0xff, 0xc1, 0x0b, 0xfc,
	0x27, 0xfe, 0x0e, 0x9a, 0x3b, 0xb6, 0xe3, 0x16, 0x14, 0x09, 0xd4, 0xb7, 0xb9, 0xe7, 0x9e, 0x7b,
	0xe7, 0xfa, 0xe4, 0xdc, 0x09, 0x6c, 0x15, 0x32, 0x17, 0x6c, 0xce, 0xc3, 0xcb, 0x92, 0x8b, 0xeb,
	0xf1, 0x42, 0xe4, 0x32, 0x27, 0x66, 0x3a, 0x4f, 0x65, 0xf0, 0x04, 0xdc, 0xa7, 0x71, 0xfe, 0x92,
	0xb3, 0x44, 0x9e, 0x53, 0x7e, 0x49, 0xf6, 0xa1, 0x1f, 0xf1, 0x65, 0x58, 0x96, 0x71, 0xe4, 0x19,
	0xbe, 0x31, 0xb2, 0x69, 0x2f, 0xe2, 0xcb, 0xb7, 0x65, 0x1c, 0x91, 0x1d, 0xe8, 0xca, 0xb9, 0x0c,
	0xe3, 0xc8, 0xdb, 0xc0, 0x84, 0x25, 0xe7, 0xf2, 0x24, 0x0a, 0x3e, 0x99, 0x30, 0x68, 0xb5, 0x28,
	0x16, 0x64, 0x17, 0xba, 0x85, 0x64, 0xb2, 0x2c, 0xb0, 0x83, 0x45, 0xab, 0xe8, 0x46, 0xef, 0x8d,
	0x9b, 0xbd, 0xff, 0x01, 0x87, 0x0b, 0x91, 0x8b, 0x70, 0x96, 0x97, 0x99, 0xf4, 0x3a, 0xbe, 0x31,
	0x32, 0x29, 0x20, 0xf4, 0x4c, 0x21, 0xc4, 0x07, 0x47, 0xf2, 0x74, 0xc1, 0x05, 0x93, 0xa5, 0xe0,
	0x9e, 0xe9, 0x1b, 0xa3, 0x01, 0x6d, 0x43, 0xe4, 0x10, 0xdc, 0x94, 0x47, 0x31, 0x0b, 0xb1, 0xaa,
	0xf0, 0x2c, 0xec, 0xe1, 0x20, 0xf6, 0x1c, 0x21, 0x75, 0x8b, 0xe0, 0x2c, 0xaa, 0x19, 0x5d, 0x6c,
	0x02, 0x0a, 0xaa, 0x08, 0x87, 0xe0, 0x5e, 0x89, 0x58, 0xf2, 0x9a, 0xd1, 0xd3, 0xd7, 0x20, 0xb6,
	0xa2, 0x94, 0x59, 0xca, 0x16, 0x35, 0xa5, 0xaf, 0x29, 0x88, 0x55, 0x94, 0xff, 0xe0, 0x8f, 0xd9,
	0x39, 0x9f, 0x5d, 0x14, 0x65, 0x5a, 0xb3, 0x6c, 0x64, 0x6d, 0xd6, 0x70, 0x45, 0x3c, 0x00, 0x5b,
	0x7d, 0x41, 0x78, 0xc5, 0x44, 0xe6, 0x81, 0x6f, 0x8c, 0xfa, 0xb4, 0xaf, 0x80, 0x77, 0x4c, 0x64,
	0xe4, 0x6f, 0x80, 0x62, 0xc1, 0x04, 0xd7, 0x59, 0x07, 0xb3, 0x36, 0x22, 0x98, 0xfe, 0x17, 0x06,
	0x6a, 0xf0, 0x3c, 0x4b, 0xae, 0x35, 0xc3, 0x45, 0x86, 0x5b, 0x83, 0x48, 0x7a, 0x0c, 0x7b, 0x11,
	0x5f, 0xc6, 0x33, 0x1e, 0x0a, 0x9e, 0xc4, 0x6c, 0x1a, 0x27, 0xb1, 0xac, 0xe8, 0x03, 0xa4, 0xef,
	0xe8, 0x34, 0x5d, 0x65, 0xb1, 0xee, 0x21, 0x6c, 0x2f, 0xf3, 0x84, 0xc9, 0x38, 0xe1, 0x61, 0xca,
	0xd3, 0x5c, 0x54, 0x45, 0x9b, 0x58, 0x44, 0xea, 0xdc, 0x6b, 0x4c, 0xa9, 0x8a, 0xc0, 0x01, 0xfb,
	0x2c, 0x8d, 0x26, 0x7c, 0x49, 0xf9, 0x65, 0xf0, 0xd5, 0x00, 0xa8, 0xa3, 0x35, 0x7e, 0x78, 0x04,
	0x3d, 0x7d, 0x7d, 0xe1, 0x6d, 0xf8, 0x9d, 0x91, 0x73, 0xb4, 0x37, 0x56, 0x9e, 0x1c, 0xaf, 0x4a,
	0xc7, 0x13, 0x3d, 0x5e, 0xcd, 0x1b, 0xbe, 0x82, 0xae, 0x86, 0x08, 0x01, 0xb3, 0x65, 0x52, 0x3c,
	0x93, 0x3d, 0xe8, 0x69, 0x87, 0xea, 0x86, 0x16, 0xed, 0xa2, 0x45, 0x0b, 0xb2, 0x0d, 0x96, 0xba,
	0x93, 0xa3, 0xb1, 0x6c, 0xaa, 0x83, 0xc0, 0xc5, 0x29, 0x4f, 0xf3, 0x3c, 0x51, 0x43, 0x7f, 0x36,
	0xc0, 0x69, 0xc2, 0x35, 0x53, 0xdf, 0x07, 0x6b, 0x91, 0xe7, 0x49, 0x3d, 0xf3, 0x6e, 0x33, 0x73,
	0x5d, 0x39, 0xc6, 0x83, 0x26, 0x0d, 0x4f, 0xc0, 0x54, 0xe1, 0x2f, 0x8f, 0x3b, 0x4d, 0xf2, 0x69,
	0xe1, 0x75, 0xfc, 0xce, 0xc8, 0xa4, 0x3a, 0x08, 0x46, 0xe0, 0x4c, 0xf8, 0xf2, 0x4c, 0x8d, 0xbe,
	0x7e, 0x53, 0x83, 0xf7, 0xe0, 0xae, 0x98, 0xbf, 0xb7, 0x90, 0x07, 0x60, 0xab, 0x54, 0x5b, 0x35,
	0xc5, 0xc5, 0x9e, 0xc1, 0x37, 0x2d, 0xd5, 0x1b, 0xf5, 0x9a, 0xa8, 0x51, 0x7c, 0x70, 0xf2, 0x34,
	0x96, 0x93, 0xea, 0xc7, 0x34, 0xd0, 0x25, 0x6d, 0x88, 0xfc, 0x05, 0xb6, 0x0a, 0x4f, 0x2b, 0xe1,
	0xd0, 0xcb, 0x0d, 0x40, 0xfe, 0x87, 0x3f, 0xe3, 0x6c, 0x96, 0x94, 0x11, 0x6f, 0x1e, 0x12, 0xbc,
	0xb3, 0x4f, 0x7f, 0xc0, 0x55, 0xa7, 0x82, 0xcb, 0x17, 0xac, 0x4c, 0xe4, 0x35, 0x3e, 0x03, 0x6a,
	0x2b, 0x6a, 0xa0, 0x91, 0xd9, 0x6a, 0xc9, 0x4c, 0xc0, 0x14, 0x2c, 0xbb, 0xa8, 0xd6, 0x1d, 0xcf,
	0x4a, 0x11, 0xc9, 0xc4, 0x9c, 0x4b, 0x5c, 0x71, 0x9b, 0x56, 0x51, 0xf0, 0xa5, 0x03, 0xee, 0xea,
	0xcb, 0xd6, 0x7a, 0xd7, 0x52, 0x8d, 0x6a, 0x17, 0x1c, 0x34, 0x2e, 0x68, 0x4a, 0xc7, 0x94, 0x65,
	0x17, 0xea, 0x40, 0x35, 0x73, 0xf8, 0xe1, 0x0e, 0xbd, 0x4b, 0xee, 0x41, 0xf7, 0x5c, 0x0b, 0xa5,
	0x34, 0x70, 0x8e, 0xb6, 0xf4, 0x00, 0x37, 0x1e, 0x62, 0x5a, 0x51, 0xee, 0xd0, 0x84, 0xc3, 0x8f,
	0x06, 0xf4, 0xeb, 0x0f, 0x6b, 0x94, 0x35, 0x5a, 0xca, 0x1e, 0xdf, 0x5e, 0xea, 0xfd, 0x9f, 0x48,
	0x73, 0x6b, 0xad, 0xc9, 0x83, 0x7a, 0xa7, 0x3a, 0xb7, 0xde, 0x81, 0x55, 0x49, 0x6b, 0xa9, 0xa6,
	0x5d, 0xfc, 0x07, 0x3b, 0xfe, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x45, 0x67, 0xcd, 0xd8, 0x06,
	0x00, 0x00,
}
