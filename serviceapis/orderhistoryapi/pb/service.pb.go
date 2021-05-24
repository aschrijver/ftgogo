// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: orderhistoryapi/pb/service.proto

package orderhistorypb

import (
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

type GetConsumerOrderHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID string                                 `protobuf:"bytes,1,opt,name=ConsumerID,proto3" json:"ConsumerID,omitempty"`
	Filter     *GetConsumerOrderHistoryRequestFilters `protobuf:"bytes,2,opt,name=Filter,proto3" json:"Filter,omitempty"`
	Next       string                                 `protobuf:"bytes,3,opt,name=Next,proto3" json:"Next,omitempty"`
	Limit      int64                                  `protobuf:"varint,4,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetConsumerOrderHistoryRequest) Reset() {
	*x = GetConsumerOrderHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderhistoryapi_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerOrderHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerOrderHistoryRequest) ProtoMessage() {}

func (x *GetConsumerOrderHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderhistoryapi_pb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerOrderHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetConsumerOrderHistoryRequest) Descriptor() ([]byte, []int) {
	return file_orderhistoryapi_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetConsumerOrderHistoryRequest) GetConsumerID() string {
	if x != nil {
		return x.ConsumerID
	}
	return ""
}

func (x *GetConsumerOrderHistoryRequest) GetFilter() *GetConsumerOrderHistoryRequestFilters {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetConsumerOrderHistoryRequest) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

func (x *GetConsumerOrderHistoryRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetConsumerOrderHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*GetConsumerOrderHistoryResponseOrderHistory `protobuf:"bytes,1,rep,name=Orders,proto3" json:"Orders,omitempty"`
	Next   string                                         `protobuf:"bytes,2,opt,name=Next,proto3" json:"Next,omitempty"`
}

func (x *GetConsumerOrderHistoryResponse) Reset() {
	*x = GetConsumerOrderHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderhistoryapi_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerOrderHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerOrderHistoryResponse) ProtoMessage() {}

func (x *GetConsumerOrderHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderhistoryapi_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerOrderHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetConsumerOrderHistoryResponse) Descriptor() ([]byte, []int) {
	return file_orderhistoryapi_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetConsumerOrderHistoryResponse) GetOrders() []*GetConsumerOrderHistoryResponseOrderHistory {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *GetConsumerOrderHistoryResponse) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

type GetOrderHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *GetOrderHistoryRequest) Reset() {
	*x = GetOrderHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderhistoryapi_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryRequest) ProtoMessage() {}

func (x *GetOrderHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderhistoryapi_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryRequest) Descriptor() ([]byte, []int) {
	return file_orderhistoryapi_pb_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderHistoryRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type GetOrderHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID        string                 `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	Status         string                 `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	RestaurantID   string                 `protobuf:"bytes,3,opt,name=RestaurantID,proto3" json:"RestaurantID,omitempty"`
	RestaurantName string                 `protobuf:"bytes,4,opt,name=RestaurantName,proto3" json:"RestaurantName,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *GetOrderHistoryResponse) Reset() {
	*x = GetOrderHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderhistoryapi_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryResponse) ProtoMessage() {}

func (x *GetOrderHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderhistoryapi_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryResponse) Descriptor() ([]byte, []int) {
	return file_orderhistoryapi_pb_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderHistoryResponse) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *GetOrderHistoryResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetOrderHistoryResponse) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

func (x *GetOrderHistoryResponse) GetRestaurantName() string {
	if x != nil {
		return x.RestaurantName
	}
	return ""
}

func (x *GetOrderHistoryResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetConsumerOrderHistoryRequestFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Since    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=Since,proto3" json:"Since,omitempty"`
	Keywords []string               `protobuf:"bytes,2,rep,name=Keywords,proto3" json:"Keywords,omitempty"`
	Status   int64                  `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GetConsumerOrderHistoryRequestFilters) Reset() {
	*x = GetConsumerOrderHistoryRequestFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderhistoryapi_pb_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerOrderHistoryRequestFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerOrderHistoryRequestFilters) ProtoMessage() {}

func (x *GetConsumerOrderHistoryRequestFilters) ProtoReflect() protoreflect.Message {
	mi := &file_orderhistoryapi_pb_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerOrderHistoryRequestFilters.ProtoReflect.Descriptor instead.
func (*GetConsumerOrderHistoryRequestFilters) Descriptor() ([]byte, []int) {
	return file_orderhistoryapi_pb_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetConsumerOrderHistoryRequestFilters) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *GetConsumerOrderHistoryRequestFilters) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *GetConsumerOrderHistoryRequestFilters) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetConsumerOrderHistoryResponseOrderHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID        string                 `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	Status         string                 `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	RestaurantID   string                 `protobuf:"bytes,3,opt,name=RestaurantID,proto3" json:"RestaurantID,omitempty"`
	RestaurantName string                 `protobuf:"bytes,4,opt,name=RestaurantName,proto3" json:"RestaurantName,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *GetConsumerOrderHistoryResponseOrderHistory) Reset() {
	*x = GetConsumerOrderHistoryResponseOrderHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderhistoryapi_pb_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerOrderHistoryResponseOrderHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerOrderHistoryResponseOrderHistory) ProtoMessage() {}

func (x *GetConsumerOrderHistoryResponseOrderHistory) ProtoReflect() protoreflect.Message {
	mi := &file_orderhistoryapi_pb_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerOrderHistoryResponseOrderHistory.ProtoReflect.Descriptor instead.
func (*GetConsumerOrderHistoryResponseOrderHistory) Descriptor() ([]byte, []int) {
	return file_orderhistoryapi_pb_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetConsumerOrderHistoryResponseOrderHistory) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *GetConsumerOrderHistoryResponseOrderHistory) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetConsumerOrderHistoryResponseOrderHistory) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

func (x *GetConsumerOrderHistoryResponseOrderHistory) GetRestaurantName() string {
	if x != nil {
		return x.RestaurantName
	}
	return ""
}

func (x *GetConsumerOrderHistoryResponseOrderHistory) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_orderhistoryapi_pb_service_proto protoreflect.FileDescriptor

var file_orderhistoryapi_pb_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x4e, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x1a, 0x6f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x53,
	0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xd4, 0x02, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x1a,
	0xc6, 0x01, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0xd1, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x26,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x32, 0xf5, 0x01, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x2e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x75, 0x73, 0x2f, 0x66,
	0x74, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x62, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderhistoryapi_pb_service_proto_rawDescOnce sync.Once
	file_orderhistoryapi_pb_service_proto_rawDescData = file_orderhistoryapi_pb_service_proto_rawDesc
)

func file_orderhistoryapi_pb_service_proto_rawDescGZIP() []byte {
	file_orderhistoryapi_pb_service_proto_rawDescOnce.Do(func() {
		file_orderhistoryapi_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderhistoryapi_pb_service_proto_rawDescData)
	})
	return file_orderhistoryapi_pb_service_proto_rawDescData
}

var file_orderhistoryapi_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orderhistoryapi_pb_service_proto_goTypes = []interface{}{
	(*GetConsumerOrderHistoryRequest)(nil),              // 0: orderhistorypb.GetConsumerOrderHistoryRequest
	(*GetConsumerOrderHistoryResponse)(nil),             // 1: orderhistorypb.GetConsumerOrderHistoryResponse
	(*GetOrderHistoryRequest)(nil),                      // 2: orderhistorypb.GetOrderHistoryRequest
	(*GetOrderHistoryResponse)(nil),                     // 3: orderhistorypb.GetOrderHistoryResponse
	(*GetConsumerOrderHistoryRequestFilters)(nil),       // 4: orderhistorypb.GetConsumerOrderHistoryRequest.filters
	(*GetConsumerOrderHistoryResponseOrderHistory)(nil), // 5: orderhistorypb.GetConsumerOrderHistoryResponse.orderHistory
	(*timestamppb.Timestamp)(nil),                       // 6: google.protobuf.Timestamp
}
var file_orderhistoryapi_pb_service_proto_depIdxs = []int32{
	4, // 0: orderhistorypb.GetConsumerOrderHistoryRequest.Filter:type_name -> orderhistorypb.GetConsumerOrderHistoryRequest.filters
	5, // 1: orderhistorypb.GetConsumerOrderHistoryResponse.Orders:type_name -> orderhistorypb.GetConsumerOrderHistoryResponse.orderHistory
	6, // 2: orderhistorypb.GetOrderHistoryResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	6, // 3: orderhistorypb.GetConsumerOrderHistoryRequest.filters.Since:type_name -> google.protobuf.Timestamp
	6, // 4: orderhistorypb.GetConsumerOrderHistoryResponse.orderHistory.CreatedAt:type_name -> google.protobuf.Timestamp
	0, // 5: orderhistorypb.OrderHistoryService.GetConsumerOrderHistory:input_type -> orderhistorypb.GetConsumerOrderHistoryRequest
	2, // 6: orderhistorypb.OrderHistoryService.GetOrderHistory:input_type -> orderhistorypb.GetOrderHistoryRequest
	1, // 7: orderhistorypb.OrderHistoryService.GetConsumerOrderHistory:output_type -> orderhistorypb.GetConsumerOrderHistoryResponse
	3, // 8: orderhistorypb.OrderHistoryService.GetOrderHistory:output_type -> orderhistorypb.GetOrderHistoryResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orderhistoryapi_pb_service_proto_init() }
func file_orderhistoryapi_pb_service_proto_init() {
	if File_orderhistoryapi_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderhistoryapi_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerOrderHistoryRequest); i {
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
		file_orderhistoryapi_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerOrderHistoryResponse); i {
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
		file_orderhistoryapi_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryRequest); i {
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
		file_orderhistoryapi_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryResponse); i {
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
		file_orderhistoryapi_pb_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerOrderHistoryRequestFilters); i {
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
		file_orderhistoryapi_pb_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerOrderHistoryResponseOrderHistory); i {
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
			RawDescriptor: file_orderhistoryapi_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderhistoryapi_pb_service_proto_goTypes,
		DependencyIndexes: file_orderhistoryapi_pb_service_proto_depIdxs,
		MessageInfos:      file_orderhistoryapi_pb_service_proto_msgTypes,
	}.Build()
	File_orderhistoryapi_pb_service_proto = out.File
	file_orderhistoryapi_pb_service_proto_rawDesc = nil
	file_orderhistoryapi_pb_service_proto_goTypes = nil
	file_orderhistoryapi_pb_service_proto_depIdxs = nil
}