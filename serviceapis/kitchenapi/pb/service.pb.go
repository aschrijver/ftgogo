// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: kitchenapi/pb/service.proto

package kitchenpb

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

type GetRestaurantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantID string `protobuf:"bytes,1,opt,name=RestaurantID,proto3" json:"RestaurantID,omitempty"`
}

func (x *GetRestaurantRequest) Reset() {
	*x = GetRestaurantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchenapi_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRestaurantRequest) ProtoMessage() {}

func (x *GetRestaurantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitchenapi_pb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRestaurantRequest.ProtoReflect.Descriptor instead.
func (*GetRestaurantRequest) Descriptor() ([]byte, []int) {
	return file_kitchenapi_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRestaurantRequest) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

type GetRestaurantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantID string `protobuf:"bytes,1,opt,name=RestaurantID,proto3" json:"RestaurantID,omitempty"`
}

func (x *GetRestaurantResponse) Reset() {
	*x = GetRestaurantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchenapi_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRestaurantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRestaurantResponse) ProtoMessage() {}

func (x *GetRestaurantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kitchenapi_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRestaurantResponse.ProtoReflect.Descriptor instead.
func (*GetRestaurantResponse) Descriptor() ([]byte, []int) {
	return file_kitchenapi_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRestaurantResponse) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

type AcceptTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketID string                 `protobuf:"bytes,1,opt,name=TicketID,proto3" json:"TicketID,omitempty"`
	ReadyBy  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=ReadyBy,proto3" json:"ReadyBy,omitempty"`
}

func (x *AcceptTicketRequest) Reset() {
	*x = AcceptTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchenapi_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptTicketRequest) ProtoMessage() {}

func (x *AcceptTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitchenapi_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptTicketRequest.ProtoReflect.Descriptor instead.
func (*AcceptTicketRequest) Descriptor() ([]byte, []int) {
	return file_kitchenapi_pb_service_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptTicketRequest) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

func (x *AcceptTicketRequest) GetReadyBy() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadyBy
	}
	return nil
}

type AcceptTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketID string `protobuf:"bytes,1,opt,name=TicketID,proto3" json:"TicketID,omitempty"`
}

func (x *AcceptTicketResponse) Reset() {
	*x = AcceptTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchenapi_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptTicketResponse) ProtoMessage() {}

func (x *AcceptTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kitchenapi_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptTicketResponse.ProtoReflect.Descriptor instead.
func (*AcceptTicketResponse) Descriptor() ([]byte, []int) {
	return file_kitchenapi_pb_service_proto_rawDescGZIP(), []int{3}
}

func (x *AcceptTicketResponse) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

var File_kitchenapi_pb_service_proto protoreflect.FileDescriptor

var file_kitchenapi_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6b,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x22, 0x67, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x79, 0x42, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x52, 0x65, 0x61, 0x64, 0x79, 0x42, 0x79, 0x22, 0x32, 0x0a, 0x14, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x32,
	0xb5, 0x01, 0x0a, 0x0e, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x75, 0x73, 0x2f, 0x66, 0x74,
	0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x3b, 0x6b,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kitchenapi_pb_service_proto_rawDescOnce sync.Once
	file_kitchenapi_pb_service_proto_rawDescData = file_kitchenapi_pb_service_proto_rawDesc
)

func file_kitchenapi_pb_service_proto_rawDescGZIP() []byte {
	file_kitchenapi_pb_service_proto_rawDescOnce.Do(func() {
		file_kitchenapi_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kitchenapi_pb_service_proto_rawDescData)
	})
	return file_kitchenapi_pb_service_proto_rawDescData
}

var file_kitchenapi_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kitchenapi_pb_service_proto_goTypes = []interface{}{
	(*GetRestaurantRequest)(nil),  // 0: kitchenpb.GetRestaurantRequest
	(*GetRestaurantResponse)(nil), // 1: kitchenpb.GetRestaurantResponse
	(*AcceptTicketRequest)(nil),   // 2: kitchenpb.AcceptTicketRequest
	(*AcceptTicketResponse)(nil),  // 3: kitchenpb.AcceptTicketResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_kitchenapi_pb_service_proto_depIdxs = []int32{
	4, // 0: kitchenpb.AcceptTicketRequest.ReadyBy:type_name -> google.protobuf.Timestamp
	0, // 1: kitchenpb.KitchenService.GetRestaurant:input_type -> kitchenpb.GetRestaurantRequest
	2, // 2: kitchenpb.KitchenService.AcceptTicket:input_type -> kitchenpb.AcceptTicketRequest
	1, // 3: kitchenpb.KitchenService.GetRestaurant:output_type -> kitchenpb.GetRestaurantResponse
	3, // 4: kitchenpb.KitchenService.AcceptTicket:output_type -> kitchenpb.AcceptTicketResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kitchenapi_pb_service_proto_init() }
func file_kitchenapi_pb_service_proto_init() {
	if File_kitchenapi_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kitchenapi_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRestaurantRequest); i {
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
		file_kitchenapi_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRestaurantResponse); i {
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
		file_kitchenapi_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptTicketRequest); i {
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
		file_kitchenapi_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptTicketResponse); i {
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
			RawDescriptor: file_kitchenapi_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kitchenapi_pb_service_proto_goTypes,
		DependencyIndexes: file_kitchenapi_pb_service_proto_depIdxs,
		MessageInfos:      file_kitchenapi_pb_service_proto_msgTypes,
	}.Build()
	File_kitchenapi_pb_service_proto = out.File
	file_kitchenapi_pb_service_proto_rawDesc = nil
	file_kitchenapi_pb_service_proto_goTypes = nil
	file_kitchenapi_pb_service_proto_depIdxs = nil
}