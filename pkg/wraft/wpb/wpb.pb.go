// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: pkg/wraft/wpb/wpb.proto

package wpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_Error      Status = 0
	Status_WillJoin   Status = 1 // 将要加入集群
	Status_Joining    Status = 3 // 加入中
	Status_Joined     Status = 5 // 已加入集群
	Status_WillRemove Status = 6 // 将要从集群中移除
	Status_Removed    Status = 7 // 已被移除集群
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Error",
		1: "WillJoin",
		3: "Joining",
		5: "Joined",
		6: "WillRemove",
		7: "Removed",
	}
	Status_value = map[string]int32{
		"Error":      0,
		"WillJoin":   1,
		"Joining":    3,
		"Joined":     5,
		"WillRemove": 6,
		"Removed":    7,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_wraft_wpb_wpb_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_wraft_wpb_wpb_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_wraft_wpb_wpb_proto_rawDescGZIP(), []int{0}
}

type Role int32

const (
	Role_Unknown      Role = 0
	Role_Follower     Role = 1
	Role_Candidate    Role = 2
	Role_PreCandidate Role = 3
	Role_Leader       Role = 4
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "Unknown",
		1: "Follower",
		2: "Candidate",
		3: "PreCandidate",
		4: "Leader",
	}
	Role_value = map[string]int32{
		"Unknown":      0,
		"Follower":     1,
		"Candidate":    2,
		"PreCandidate": 3,
		"Leader":       4,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_wraft_wpb_wpb_proto_enumTypes[1].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_pkg_wraft_wpb_wpb_proto_enumTypes[1]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_pkg_wraft_wpb_wpb_proto_rawDescGZIP(), []int{1}
}

type CMDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cmd  uint32 `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CMDReq) Reset() {
	*x = CMDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMDReq) ProtoMessage() {}

func (x *CMDReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMDReq.ProtoReflect.Descriptor instead.
func (*CMDReq) Descriptor() ([]byte, []int) {
	return file_pkg_wraft_wpb_wpb_proto_rawDescGZIP(), []int{0}
}

func (x *CMDReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CMDReq) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *CMDReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CMDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=wpb.Status" json:"status,omitempty"` // 请求状态
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                      // 返回数据
}

func (x *CMDResp) Reset() {
	*x = CMDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMDResp) ProtoMessage() {}

func (x *CMDResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMDResp.ProtoReflect.Descriptor instead.
func (*CMDResp) Descriptor() ([]byte, []int) {
	return file_pkg_wraft_wpb_wpb_proto_rawDescGZIP(), []int{1}
}

func (x *CMDResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CMDResp) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Error
}

func (x *CMDResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers   []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	Version uint64  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_pkg_wraft_wpb_wpb_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterConfig) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *ClusterConfig) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr      string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Applicant bool   `protobuf:"varint,3,opt,name=applicant,proto3" json:"applicant,omitempty"` // 是否是加入集群申请人
	Status    Status `protobuf:"varint,4,opt,name=status,proto3,enum=wpb.Status" json:"status,omitempty"`
	Role      Role   `protobuf:"varint,5,opt,name=role,proto3,enum=wpb.Role" json:"role,omitempty"`
	Term      uint64 `protobuf:"varint,6,opt,name=term,proto3" json:"term,omitempty"`
	Vote      uint64 `protobuf:"varint,7,opt,name=vote,proto3" json:"vote,omitempty"`
	Lead      uint64 `protobuf:"varint,8,opt,name=lead,proto3" json:"lead,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wraft_wpb_wpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_pkg_wraft_wpb_wpb_proto_rawDescGZIP(), []int{3}
}

func (x *Peer) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Peer) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Peer) GetApplicant() bool {
	if x != nil {
		return x.Applicant
	}
	return false
}

func (x *Peer) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Error
}

func (x *Peer) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_Unknown
}

func (x *Peer) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Peer) GetVote() uint64 {
	if x != nil {
		return x.Vote
	}
	return 0
}

func (x *Peer) GetLead() uint64 {
	if x != nil {
		return x.Lead
	}
	return 0
}

var File_pkg_wraft_wpb_wpb_proto protoreflect.FileDescriptor

var file_pkg_wraft_wpb_wpb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x77, 0x70, 0x62, 0x2f,
	0x77, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x77, 0x70, 0x62, 0x22, 0x3e,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x52,
	0x0a, 0x07, 0x43, 0x4d, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x77, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x4a, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x77, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc8,
	0x01, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x77, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x77,
	0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x64, 0x2a, 0x57, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x57, 0x69, 0x6c, 0x6c, 0x4a, 0x6f, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x4a, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x6f, 0x69,
	0x6e, 0x65, 0x64, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x10, 0x07, 0x2a, 0x4e, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x10, 0x04, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x77, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_wraft_wpb_wpb_proto_rawDescOnce sync.Once
	file_pkg_wraft_wpb_wpb_proto_rawDescData = file_pkg_wraft_wpb_wpb_proto_rawDesc
)

func file_pkg_wraft_wpb_wpb_proto_rawDescGZIP() []byte {
	file_pkg_wraft_wpb_wpb_proto_rawDescOnce.Do(func() {
		file_pkg_wraft_wpb_wpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_wraft_wpb_wpb_proto_rawDescData)
	})
	return file_pkg_wraft_wpb_wpb_proto_rawDescData
}

var file_pkg_wraft_wpb_wpb_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_wraft_wpb_wpb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_wraft_wpb_wpb_proto_goTypes = []interface{}{
	(Status)(0),           // 0: wpb.Status
	(Role)(0),             // 1: wpb.Role
	(*CMDReq)(nil),        // 2: wpb.CMDReq
	(*CMDResp)(nil),       // 3: wpb.CMDResp
	(*ClusterConfig)(nil), // 4: wpb.ClusterConfig
	(*Peer)(nil),          // 5: wpb.Peer
}
var file_pkg_wraft_wpb_wpb_proto_depIdxs = []int32{
	0, // 0: wpb.CMDResp.status:type_name -> wpb.Status
	5, // 1: wpb.ClusterConfig.peers:type_name -> wpb.Peer
	0, // 2: wpb.Peer.status:type_name -> wpb.Status
	1, // 3: wpb.Peer.role:type_name -> wpb.Role
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_wraft_wpb_wpb_proto_init() }
func file_pkg_wraft_wpb_wpb_proto_init() {
	if File_pkg_wraft_wpb_wpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_wraft_wpb_wpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMDReq); i {
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
		file_pkg_wraft_wpb_wpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMDResp); i {
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
		file_pkg_wraft_wpb_wpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterConfig); i {
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
		file_pkg_wraft_wpb_wpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
			RawDescriptor: file_pkg_wraft_wpb_wpb_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_wraft_wpb_wpb_proto_goTypes,
		DependencyIndexes: file_pkg_wraft_wpb_wpb_proto_depIdxs,
		EnumInfos:         file_pkg_wraft_wpb_wpb_proto_enumTypes,
		MessageInfos:      file_pkg_wraft_wpb_wpb_proto_msgTypes,
	}.Build()
	File_pkg_wraft_wpb_wpb_proto = out.File
	file_pkg_wraft_wpb_wpb_proto_rawDesc = nil
	file_pkg_wraft_wpb_wpb_proto_goTypes = nil
	file_pkg_wraft_wpb_wpb_proto_depIdxs = nil
}