// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: accountingapi/service.proto

package accountingapi

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

type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountingapi_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accountingapi_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_accountingapi_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type GetAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *GetAccountResponse) Reset() {
	*x = GetAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountingapi_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResponse) ProtoMessage() {}

func (x *GetAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accountingapi_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return file_accountingapi_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountResponse) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type DisableAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *DisableAccountRequest) Reset() {
	*x = DisableAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountingapi_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableAccountRequest) ProtoMessage() {}

func (x *DisableAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accountingapi_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableAccountRequest.ProtoReflect.Descriptor instead.
func (*DisableAccountRequest) Descriptor() ([]byte, []int) {
	return file_accountingapi_service_proto_rawDescGZIP(), []int{2}
}

func (x *DisableAccountRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type DisableAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *DisableAccountResponse) Reset() {
	*x = DisableAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountingapi_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableAccountResponse) ProtoMessage() {}

func (x *DisableAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accountingapi_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableAccountResponse.ProtoReflect.Descriptor instead.
func (*DisableAccountResponse) Descriptor() ([]byte, []int) {
	return file_accountingapi_service_proto_rawDescGZIP(), []int{3}
}

func (x *DisableAccountResponse) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type EnableAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *EnableAccountRequest) Reset() {
	*x = EnableAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountingapi_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableAccountRequest) ProtoMessage() {}

func (x *EnableAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accountingapi_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableAccountRequest.ProtoReflect.Descriptor instead.
func (*EnableAccountRequest) Descriptor() ([]byte, []int) {
	return file_accountingapi_service_proto_rawDescGZIP(), []int{4}
}

func (x *EnableAccountRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type EnableAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *EnableAccountResponse) Reset() {
	*x = EnableAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountingapi_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableAccountResponse) ProtoMessage() {}

func (x *EnableAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accountingapi_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableAccountResponse.ProtoReflect.Descriptor instead.
func (*EnableAccountResponse) Descriptor() ([]byte, []int) {
	return file_accountingapi_service_proto_rawDescGZIP(), []int{5}
}

func (x *EnableAccountResponse) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

var File_accountingapi_service_proto protoreflect.FileDescriptor

var file_accountingapi_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x22, 0x31, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22,
	0x32, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x15, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x16, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x22, 0x34, 0x0a, 0x14, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x15, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x32,
	0xa1, 0x02, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x75, 0x73, 0x2f, 0x66, 0x74, 0x67, 0x6f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_accountingapi_service_proto_rawDescOnce sync.Once
	file_accountingapi_service_proto_rawDescData = file_accountingapi_service_proto_rawDesc
)

func file_accountingapi_service_proto_rawDescGZIP() []byte {
	file_accountingapi_service_proto_rawDescOnce.Do(func() {
		file_accountingapi_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_accountingapi_service_proto_rawDescData)
	})
	return file_accountingapi_service_proto_rawDescData
}

var file_accountingapi_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accountingapi_service_proto_goTypes = []interface{}{
	(*GetAccountRequest)(nil),      // 0: accountingapi.GetAccountRequest
	(*GetAccountResponse)(nil),     // 1: accountingapi.GetAccountResponse
	(*DisableAccountRequest)(nil),  // 2: accountingapi.DisableAccountRequest
	(*DisableAccountResponse)(nil), // 3: accountingapi.DisableAccountResponse
	(*EnableAccountRequest)(nil),   // 4: accountingapi.EnableAccountRequest
	(*EnableAccountResponse)(nil),  // 5: accountingapi.EnableAccountResponse
}
var file_accountingapi_service_proto_depIdxs = []int32{
	0, // 0: accountingapi.AccountingService.GetAccount:input_type -> accountingapi.GetAccountRequest
	2, // 1: accountingapi.AccountingService.DisableAccount:input_type -> accountingapi.DisableAccountRequest
	4, // 2: accountingapi.AccountingService.EnableAccount:input_type -> accountingapi.EnableAccountRequest
	1, // 3: accountingapi.AccountingService.GetAccount:output_type -> accountingapi.GetAccountResponse
	3, // 4: accountingapi.AccountingService.DisableAccount:output_type -> accountingapi.DisableAccountResponse
	5, // 5: accountingapi.AccountingService.EnableAccount:output_type -> accountingapi.EnableAccountResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accountingapi_service_proto_init() }
func file_accountingapi_service_proto_init() {
	if File_accountingapi_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accountingapi_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountRequest); i {
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
		file_accountingapi_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountResponse); i {
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
		file_accountingapi_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableAccountRequest); i {
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
		file_accountingapi_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableAccountResponse); i {
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
		file_accountingapi_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableAccountRequest); i {
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
		file_accountingapi_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableAccountResponse); i {
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
			RawDescriptor: file_accountingapi_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accountingapi_service_proto_goTypes,
		DependencyIndexes: file_accountingapi_service_proto_depIdxs,
		MessageInfos:      file_accountingapi_service_proto_msgTypes,
	}.Build()
	File_accountingapi_service_proto = out.File
	file_accountingapi_service_proto_rawDesc = nil
	file_accountingapi_service_proto_goTypes = nil
	file_accountingapi_service_proto_depIdxs = nil
}
