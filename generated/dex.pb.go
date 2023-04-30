// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: dex.proto

package generated

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

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol  string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimal int64  `protobuf:"varint,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_dex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_dex_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Token) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Token) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Token) GetDecimal() int64 {
	if x != nil {
		return x.Decimal
	}
	return 0
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token0      *Token `protobuf:"bytes,1,opt,name=token0,proto3" json:"token0,omitempty"`
	Token1      *Token `protobuf:"bytes,2,opt,name=token1,proto3" json:"token1,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Reserve0    []byte `protobuf:"bytes,4,opt,name=reserve0,proto3" json:"reserve0,omitempty"`
	Reserve1    []byte `protobuf:"bytes,5,opt,name=reserve1,proto3" json:"reserve1,omitempty"`
	LastUpdated int64  `protobuf:"varint,6,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_dex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_dex_proto_rawDescGZIP(), []int{1}
}

func (x *Pair) GetToken0() *Token {
	if x != nil {
		return x.Token0
	}
	return nil
}

func (x *Pair) GetToken1() *Token {
	if x != nil {
		return x.Token1
	}
	return nil
}

func (x *Pair) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Pair) GetReserve0() []byte {
	if x != nil {
		return x.Reserve0
	}
	return nil
}

func (x *Pair) GetReserve1() []byte {
	if x != nil {
		return x.Reserve1
	}
	return nil
}

func (x *Pair) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

type Dex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FactoryAddress string `protobuf:"bytes,1,opt,name=factoryAddress,proto3" json:"factoryAddress,omitempty"`
	RouterAddress  string `protobuf:"bytes,2,opt,name=routerAddress,proto3" json:"routerAddress,omitempty"`
	NumPairs       int64  `protobuf:"varint,3,opt,name=numPairs,proto3" json:"numPairs,omitempty"`
}

func (x *Dex) Reset() {
	*x = Dex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dex) ProtoMessage() {}

func (x *Dex) ProtoReflect() protoreflect.Message {
	mi := &file_dex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dex.ProtoReflect.Descriptor instead.
func (*Dex) Descriptor() ([]byte, []int) {
	return file_dex_proto_rawDescGZIP(), []int{2}
}

func (x *Dex) GetFactoryAddress() string {
	if x != nil {
		return x.FactoryAddress
	}
	return ""
}

func (x *Dex) GetRouterAddress() string {
	if x != nil {
		return x.RouterAddress
	}
	return ""
}

func (x *Dex) GetNumPairs() int64 {
	if x != nil {
		return x.NumPairs
	}
	return 0
}

type GetAllDexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllDexRequest) Reset() {
	*x = GetAllDexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDexRequest) ProtoMessage() {}

func (x *GetAllDexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDexRequest.ProtoReflect.Descriptor instead.
func (*GetAllDexRequest) Descriptor() ([]byte, []int) {
	return file_dex_proto_rawDescGZIP(), []int{3}
}

type GetAllDexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dexs []*Dex `protobuf:"bytes,1,rep,name=dexs,proto3" json:"dexs,omitempty"`
}

func (x *GetAllDexResponse) Reset() {
	*x = GetAllDexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dex_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDexResponse) ProtoMessage() {}

func (x *GetAllDexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dex_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDexResponse.ProtoReflect.Descriptor instead.
func (*GetAllDexResponse) Descriptor() ([]byte, []int) {
	return file_dex_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllDexResponse) GetDexs() []*Dex {
	if x != nil {
		return x.Dexs
	}
	return nil
}

type InsertDexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InsertDexResponse) Reset() {
	*x = InsertDexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dex_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertDexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertDexResponse) ProtoMessage() {}

func (x *InsertDexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dex_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertDexResponse.ProtoReflect.Descriptor instead.
func (*InsertDexResponse) Descriptor() ([]byte, []int) {
	return file_dex_proto_rawDescGZIP(), []int{5}
}

var File_dex_proto protoreflect.FileDescriptor

var file_dex_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x22, 0xba, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1e, 0x0a,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x12, 0x1e, 0x0a,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x31, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x30, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x31, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x31, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x6f, 0x0a, 0x03, 0x44, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x69,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x69,
	0x72, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x64,
	0x65, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x44, 0x65, 0x78, 0x52,
	0x04, 0x64, 0x65, 0x78, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x44,
	0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x69, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x44, 0x65, 0x78, 0x12, 0x04, 0x2e, 0x44, 0x65, 0x78, 0x1a, 0x12, 0x2e, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x44, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x78, 0x12, 0x11, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dex_proto_rawDescOnce sync.Once
	file_dex_proto_rawDescData = file_dex_proto_rawDesc
)

func file_dex_proto_rawDescGZIP() []byte {
	file_dex_proto_rawDescOnce.Do(func() {
		file_dex_proto_rawDescData = protoimpl.X.CompressGZIP(file_dex_proto_rawDescData)
	})
	return file_dex_proto_rawDescData
}

var file_dex_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dex_proto_goTypes = []interface{}{
	(*Token)(nil),             // 0: Token
	(*Pair)(nil),              // 1: Pair
	(*Dex)(nil),               // 2: Dex
	(*GetAllDexRequest)(nil),  // 3: GetAllDexRequest
	(*GetAllDexResponse)(nil), // 4: GetAllDexResponse
	(*InsertDexResponse)(nil), // 5: InsertDexResponse
}
var file_dex_proto_depIdxs = []int32{
	0, // 0: Pair.token0:type_name -> Token
	0, // 1: Pair.token1:type_name -> Token
	2, // 2: GetAllDexResponse.dexs:type_name -> Dex
	2, // 3: Database.InsertDex:input_type -> Dex
	3, // 4: Database.GetAllDex:input_type -> GetAllDexRequest
	5, // 5: Database.InsertDex:output_type -> InsertDexResponse
	4, // 6: Database.GetAllDex:output_type -> GetAllDexResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dex_proto_init() }
func file_dex_proto_init() {
	if File_dex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_dex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_dex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dex); i {
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
		file_dex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDexRequest); i {
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
		file_dex_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDexResponse); i {
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
		file_dex_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertDexResponse); i {
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
			RawDescriptor: file_dex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dex_proto_goTypes,
		DependencyIndexes: file_dex_proto_depIdxs,
		MessageInfos:      file_dex_proto_msgTypes,
	}.Build()
	File_dex_proto = out.File
	file_dex_proto_rawDesc = nil
	file_dex_proto_goTypes = nil
	file_dex_proto_depIdxs = nil
}
