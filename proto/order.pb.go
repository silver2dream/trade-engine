// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: order.proto

package proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     uint32 `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TradeId  uint32 `protobuf:"varint,2,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	StockId  uint64 `protobuf:"varint,3,opt,name=stockId,proto3" json:"stockId,omitempty"`
	Kind     int32  `protobuf:"varint,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Quantity uint64 `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price    uint64 `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetUuid() uint32 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *Order) GetTradeId() uint32 {
	if x != nil {
		return x.TradeId
	}
	return 0
}

func (x *Order) GetStockId() uint64 {
	if x != nil {
		return x.StockId
	}
	return 0
}

func (x *Order) GetKind() int32 {
	if x != nil {
		return x.Kind
	}
	return 0
}

func (x *Order) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type TradeSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraderId uint32 `protobuf:"varint,1,opt,name=trader_id,json=traderId,proto3" json:"trader_id,omitempty"`
}

func (x *TradeSession) Reset() {
	*x = TradeSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeSession) ProtoMessage() {}

func (x *TradeSession) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeSession.ProtoReflect.Descriptor instead.
func (*TradeSession) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *TradeSession) GetTraderId() uint32 {
	if x != nil {
		return x.TraderId
	}
	return 0
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2b, 0x0a,
	0x0c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_order_proto_goTypes = []interface{}{
	(*Order)(nil),        // 0: proto.Order
	(*TradeSession)(nil), // 1: proto.TradeSession
}
var file_order_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeSession); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}