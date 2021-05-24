// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: orderapi/pb/service.proto

package orderpb

import (
	pb "github.com/stackus/ftgogo/serviceapis/commonapi/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderState int32

const (
	OrderState_ApprovalPending OrderState = 0
	OrderState_Approved        OrderState = 1
	OrderState_Rejected        OrderState = 2
	OrderState_CancelPending   OrderState = 3
	OrderState_Cancelled       OrderState = 4
	OrderState_RevisionPending OrderState = 5
	OrderState_Unknown         OrderState = 6
)

// Enum value maps for OrderState.
var (
	OrderState_name = map[int32]string{
		0: "ApprovalPending",
		1: "Approved",
		2: "Rejected",
		3: "CancelPending",
		4: "Cancelled",
		5: "RevisionPending",
		6: "Unknown",
	}
	OrderState_value = map[string]int32{
		"ApprovalPending": 0,
		"Approved":        1,
		"Rejected":        2,
		"CancelPending":   3,
		"Cancelled":       4,
		"RevisionPending": 5,
		"Unknown":         6,
	}
)

func (x OrderState) Enum() *OrderState {
	p := new(OrderState)
	*p = x
	return p
}

func (x OrderState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderState) Descriptor() protoreflect.EnumDescriptor {
	return file_orderapi_pb_service_proto_enumTypes[0].Descriptor()
}

func (OrderState) Type() protoreflect.EnumType {
	return &file_orderapi_pb_service_proto_enumTypes[0]
}

func (x OrderState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderState.Descriptor instead.
func (OrderState) EnumDescriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID      string     `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	ConsumerID   string     `protobuf:"bytes,2,opt,name=ConsumerID,proto3" json:"ConsumerID,omitempty"`
	RestaurantID string     `protobuf:"bytes,3,opt,name=RestaurantID,proto3" json:"RestaurantID,omitempty"`
	OrderTotal   int64      `protobuf:"varint,4,opt,name=OrderTotal,proto3" json:"OrderTotal,omitempty"`
	Status       OrderState `protobuf:"varint,5,opt,name=Status,proto3,enum=orderpb.OrderState" json:"Status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[0]
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
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *Order) GetConsumerID() string {
	if x != nil {
		return x.ConsumerID
	}
	return ""
}

func (x *Order) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

func (x *Order) GetOrderTotal() int64 {
	if x != nil {
		return x.OrderTotal
	}
	return 0
}

func (x *Order) GetStatus() OrderState {
	if x != nil {
		return x.Status
	}
	return OrderState_ApprovalPending
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID   string                 `protobuf:"bytes,1,opt,name=ConsumerID,proto3" json:"ConsumerID,omitempty"`
	RestaurantID string                 `protobuf:"bytes,2,opt,name=RestaurantID,proto3" json:"RestaurantID,omitempty"`
	DeliverAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=DeliverAt,proto3" json:"DeliverAt,omitempty"`
	DeliverTo    *pb.Address            `protobuf:"bytes,4,opt,name=DeliverTo,proto3" json:"DeliverTo,omitempty"`
	LineItems    *pb.MenuItemQuantities `protobuf:"bytes,5,opt,name=LineItems,proto3" json:"LineItems,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetConsumerID() string {
	if x != nil {
		return x.ConsumerID
	}
	return ""
}

func (x *CreateOrderRequest) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

func (x *CreateOrderRequest) GetDeliverAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliverAt
	}
	return nil
}

func (x *CreateOrderRequest) GetDeliverTo() *pb.Address {
	if x != nil {
		return x.DeliverTo
	}
	return nil
}

func (x *CreateOrderRequest) GetLineItems() *pb.MenuItemQuantities {
	if x != nil {
		return x.LineItems
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type GetOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=Order,proto3" json:"Order,omitempty"`
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{5}
}

func (x *CancelOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type CancelOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OrderState `protobuf:"varint,1,opt,name=Status,proto3,enum=orderpb.OrderState" json:"Status,omitempty"`
}

func (x *CancelOrderResponse) Reset() {
	*x = CancelOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderResponse) ProtoMessage() {}

func (x *CancelOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderResponse.ProtoReflect.Descriptor instead.
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{6}
}

func (x *CancelOrderResponse) GetStatus() OrderState {
	if x != nil {
		return x.Status
	}
	return OrderState_ApprovalPending
}

type ReviseOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID           string                 `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	RevisedQuantities *pb.MenuItemQuantities `protobuf:"bytes,2,opt,name=RevisedQuantities,proto3" json:"RevisedQuantities,omitempty"`
}

func (x *ReviseOrderRequest) Reset() {
	*x = ReviseOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviseOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviseOrderRequest) ProtoMessage() {}

func (x *ReviseOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviseOrderRequest.ProtoReflect.Descriptor instead.
func (*ReviseOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{7}
}

func (x *ReviseOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *ReviseOrderRequest) GetRevisedQuantities() *pb.MenuItemQuantities {
	if x != nil {
		return x.RevisedQuantities
	}
	return nil
}

type ReviseOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OrderState `protobuf:"varint,1,opt,name=Status,proto3,enum=orderpb.OrderState" json:"Status,omitempty"`
}

func (x *ReviseOrderResponse) Reset() {
	*x = ReviseOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderapi_pb_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviseOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviseOrderResponse) ProtoMessage() {}

func (x *ReviseOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderapi_pb_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviseOrderResponse.ProtoReflect.Descriptor instead.
func (*ReviseOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderapi_pb_service_proto_rawDescGZIP(), []int{8}
}

func (x *ReviseOrderResponse) GetStatus() OrderState {
	if x != nil {
		return x.Status
	}
	return OrderState_ApprovalPending
}

var File_orderapi_pb_service_proto protoreflect.FileDescriptor

var file_orderapi_pb_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xff, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x38, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x12, 0x3a, 0x0a, 0x09,
	0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49,
	0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x09, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x2e, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x42, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x7a, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x11, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x64, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49,
	0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x11, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x42, 0x0a, 0x13, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0x81, 0x01, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x06, 0x32, 0xad, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x75, 0x73, 0x2f, 0x66,
	0x74, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x3b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderapi_pb_service_proto_rawDescOnce sync.Once
	file_orderapi_pb_service_proto_rawDescData = file_orderapi_pb_service_proto_rawDesc
)

func file_orderapi_pb_service_proto_rawDescGZIP() []byte {
	file_orderapi_pb_service_proto_rawDescOnce.Do(func() {
		file_orderapi_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderapi_pb_service_proto_rawDescData)
	})
	return file_orderapi_pb_service_proto_rawDescData
}

var file_orderapi_pb_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orderapi_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_orderapi_pb_service_proto_goTypes = []interface{}{
	(OrderState)(0),               // 0: orderpb.OrderState
	(*Order)(nil),                 // 1: orderpb.Order
	(*CreateOrderRequest)(nil),    // 2: orderpb.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 3: orderpb.CreateOrderResponse
	(*GetOrderRequest)(nil),       // 4: orderpb.GetOrderRequest
	(*GetOrderResponse)(nil),      // 5: orderpb.GetOrderResponse
	(*CancelOrderRequest)(nil),    // 6: orderpb.CancelOrderRequest
	(*CancelOrderResponse)(nil),   // 7: orderpb.CancelOrderResponse
	(*ReviseOrderRequest)(nil),    // 8: orderpb.ReviseOrderRequest
	(*ReviseOrderResponse)(nil),   // 9: orderpb.ReviseOrderResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*pb.Address)(nil),            // 11: commonpb.Address
	(*pb.MenuItemQuantities)(nil), // 12: commonpb.MenuItemQuantities
}
var file_orderapi_pb_service_proto_depIdxs = []int32{
	0,  // 0: orderpb.Order.Status:type_name -> orderpb.OrderState
	10, // 1: orderpb.CreateOrderRequest.DeliverAt:type_name -> google.protobuf.Timestamp
	11, // 2: orderpb.CreateOrderRequest.DeliverTo:type_name -> commonpb.Address
	12, // 3: orderpb.CreateOrderRequest.LineItems:type_name -> commonpb.MenuItemQuantities
	1,  // 4: orderpb.GetOrderResponse.Order:type_name -> orderpb.Order
	0,  // 5: orderpb.CancelOrderResponse.Status:type_name -> orderpb.OrderState
	12, // 6: orderpb.ReviseOrderRequest.RevisedQuantities:type_name -> commonpb.MenuItemQuantities
	0,  // 7: orderpb.ReviseOrderResponse.Status:type_name -> orderpb.OrderState
	2,  // 8: orderpb.OrderService.CreateOrder:input_type -> orderpb.CreateOrderRequest
	4,  // 9: orderpb.OrderService.GetOrder:input_type -> orderpb.GetOrderRequest
	6,  // 10: orderpb.OrderService.CancelOrder:input_type -> orderpb.CancelOrderRequest
	8,  // 11: orderpb.OrderService.ReviseOrder:input_type -> orderpb.ReviseOrderRequest
	3,  // 12: orderpb.OrderService.CreateOrder:output_type -> orderpb.CreateOrderResponse
	5,  // 13: orderpb.OrderService.GetOrder:output_type -> orderpb.GetOrderResponse
	7,  // 14: orderpb.OrderService.CancelOrder:output_type -> orderpb.CancelOrderResponse
	9,  // 15: orderpb.OrderService.ReviseOrder:output_type -> orderpb.ReviseOrderResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_orderapi_pb_service_proto_init() }
func file_orderapi_pb_service_proto_init() {
	if File_orderapi_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderapi_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_orderapi_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_orderapi_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_orderapi_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderRequest); i {
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
		file_orderapi_pb_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderResponse); i {
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
		file_orderapi_pb_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRequest); i {
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
		file_orderapi_pb_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderResponse); i {
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
		file_orderapi_pb_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviseOrderRequest); i {
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
		file_orderapi_pb_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviseOrderResponse); i {
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
			RawDescriptor: file_orderapi_pb_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderapi_pb_service_proto_goTypes,
		DependencyIndexes: file_orderapi_pb_service_proto_depIdxs,
		EnumInfos:         file_orderapi_pb_service_proto_enumTypes,
		MessageInfos:      file_orderapi_pb_service_proto_msgTypes,
	}.Build()
	File_orderapi_pb_service_proto = out.File
	file_orderapi_pb_service_proto_rawDesc = nil
	file_orderapi_pb_service_proto_goTypes = nil
	file_orderapi_pb_service_proto_depIdxs = nil
}