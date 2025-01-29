// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: vendor.proto

package vendor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	mi := &file_vendor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_vendor_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendMessageReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         string                 `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	mi := &file_vendor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_vendor_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_vendor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_vendor_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type InventoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShoppingCart  []*Product             `protobuf:"bytes,1,rep,name=shoppingCart,proto3" json:"shoppingCart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryRequest) Reset() {
	*x = InventoryRequest{}
	mi := &file_vendor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryRequest) ProtoMessage() {}

func (x *InventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryRequest.ProtoReflect.Descriptor instead.
func (*InventoryRequest) Descriptor() ([]byte, []int) {
	return file_vendor_proto_rawDescGZIP(), []int{3}
}

func (x *InventoryRequest) GetShoppingCart() []*Product {
	if x != nil {
		return x.ShoppingCart
	}
	return nil
}

type StockItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price         int32                  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StockItem) Reset() {
	*x = StockItem{}
	mi := &file_vendor_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockItem) ProtoMessage() {}

func (x *StockItem) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockItem.ProtoReflect.Descriptor instead.
func (*StockItem) Descriptor() ([]byte, []int) {
	return file_vendor_proto_rawDescGZIP(), []int{4}
}

func (x *StockItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockItem) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type InventoryReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StockItems    []*StockItem           `protobuf:"bytes,1,rep,name=stockItems,proto3" json:"stockItems,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryReply) Reset() {
	*x = InventoryReply{}
	mi := &file_vendor_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryReply) ProtoMessage() {}

func (x *InventoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryReply.ProtoReflect.Descriptor instead.
func (*InventoryReply) Descriptor() ([]byte, []int) {
	return file_vendor_proto_rawDescGZIP(), []int{5}
}

func (x *InventoryReply) GetStockItems() []*StockItem {
	if x != nil {
		return x.StockItems
	}
	return nil
}

var File_vendor_proto protoreflect.FileDescriptor

var file_vendor_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x47, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43,
	0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x22, 0x35, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x43, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x32, 0x92, 0x01, 0x0a, 0x0d, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3e, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x4f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0xaa, 0x02, 0x0b, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_vendor_proto_rawDescOnce sync.Once
	file_vendor_proto_rawDescData []byte
)

func file_vendor_proto_rawDescGZIP() []byte {
	file_vendor_proto_rawDescOnce.Do(func() {
		file_vendor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_vendor_proto_rawDesc), len(file_vendor_proto_rawDesc)))
	})
	return file_vendor_proto_rawDescData
}

var file_vendor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vendor_proto_goTypes = []any{
	(*SendMessageRequest)(nil), // 0: vendor.SendMessageRequest
	(*SendMessageReply)(nil),   // 1: vendor.SendMessageReply
	(*Product)(nil),            // 2: vendor.Product
	(*InventoryRequest)(nil),   // 3: vendor.InventoryRequest
	(*StockItem)(nil),          // 4: vendor.StockItem
	(*InventoryReply)(nil),     // 5: vendor.InventoryReply
}
var file_vendor_proto_depIdxs = []int32{
	2, // 0: vendor.InventoryRequest.shoppingCart:type_name -> vendor.Product
	4, // 1: vendor.InventoryReply.stockItems:type_name -> vendor.StockItem
	0, // 2: vendor.VendorService.SendMessage:input_type -> vendor.SendMessageRequest
	3, // 3: vendor.VendorService.Products:input_type -> vendor.InventoryRequest
	1, // 4: vendor.VendorService.SendMessage:output_type -> vendor.SendMessageReply
	5, // 5: vendor.VendorService.Products:output_type -> vendor.InventoryReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vendor_proto_init() }
func file_vendor_proto_init() {
	if File_vendor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vendor_proto_rawDesc), len(file_vendor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vendor_proto_goTypes,
		DependencyIndexes: file_vendor_proto_depIdxs,
		MessageInfos:      file_vendor_proto_msgTypes,
	}.Build()
	File_vendor_proto = out.File
	file_vendor_proto_goTypes = nil
	file_vendor_proto_depIdxs = nil
}
