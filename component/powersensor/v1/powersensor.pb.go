// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: component/powersensor/v1/powersensor.proto

package v1

import (
	v1 "go.viam.com/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetVoltageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a power sensor
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Additional arguments to the method
	Extra *structpb.Struct `protobuf:"bytes,99,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *GetVoltageRequest) Reset() {
	*x = GetVoltageRequest{}
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoltageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoltageRequest) ProtoMessage() {}

func (x *GetVoltageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoltageRequest.ProtoReflect.Descriptor instead.
func (*GetVoltageRequest) Descriptor() ([]byte, []int) {
	return file_component_powersensor_v1_powersensor_proto_rawDescGZIP(), []int{0}
}

func (x *GetVoltageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetVoltageRequest) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

type GetVoltageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Voltage in volts
	Volts float64 `protobuf:"fixed64,1,opt,name=volts,proto3" json:"volts,omitempty"`
	// Bool describing whether the voltage is DC or AC
	IsAc bool `protobuf:"varint,2,opt,name=is_ac,json=isAc,proto3" json:"is_ac,omitempty"`
}

func (x *GetVoltageResponse) Reset() {
	*x = GetVoltageResponse{}
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoltageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoltageResponse) ProtoMessage() {}

func (x *GetVoltageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoltageResponse.ProtoReflect.Descriptor instead.
func (*GetVoltageResponse) Descriptor() ([]byte, []int) {
	return file_component_powersensor_v1_powersensor_proto_rawDescGZIP(), []int{1}
}

func (x *GetVoltageResponse) GetVolts() float64 {
	if x != nil {
		return x.Volts
	}
	return 0
}

func (x *GetVoltageResponse) GetIsAc() bool {
	if x != nil {
		return x.IsAc
	}
	return false
}

type GetCurrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a power sensor
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Additional arguments to the method
	Extra *structpb.Struct `protobuf:"bytes,99,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *GetCurrentRequest) Reset() {
	*x = GetCurrentRequest{}
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentRequest) ProtoMessage() {}

func (x *GetCurrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentRequest) Descriptor() ([]byte, []int) {
	return file_component_powersensor_v1_powersensor_proto_rawDescGZIP(), []int{2}
}

func (x *GetCurrentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCurrentRequest) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

type GetCurrentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current in amperes
	Amperes float64 `protobuf:"fixed64,1,opt,name=amperes,proto3" json:"amperes,omitempty"`
	// Bool descibing whether the current is DC or AC
	IsAc bool `protobuf:"varint,2,opt,name=is_ac,json=isAc,proto3" json:"is_ac,omitempty"`
}

func (x *GetCurrentResponse) Reset() {
	*x = GetCurrentResponse{}
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentResponse) ProtoMessage() {}

func (x *GetCurrentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentResponse) Descriptor() ([]byte, []int) {
	return file_component_powersensor_v1_powersensor_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrentResponse) GetAmperes() float64 {
	if x != nil {
		return x.Amperes
	}
	return 0
}

func (x *GetCurrentResponse) GetIsAc() bool {
	if x != nil {
		return x.IsAc
	}
	return false
}

type GetPowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a power sensor
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Additional arguments to the method
	Extra *structpb.Struct `protobuf:"bytes,99,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *GetPowerRequest) Reset() {
	*x = GetPowerRequest{}
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerRequest) ProtoMessage() {}

func (x *GetPowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerRequest.ProtoReflect.Descriptor instead.
func (*GetPowerRequest) Descriptor() ([]byte, []int) {
	return file_component_powersensor_v1_powersensor_proto_rawDescGZIP(), []int{4}
}

func (x *GetPowerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPowerRequest) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

type GetPowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Power in watts
	Watts float64 `protobuf:"fixed64,1,opt,name=watts,proto3" json:"watts,omitempty"`
}

func (x *GetPowerResponse) Reset() {
	*x = GetPowerResponse{}
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerResponse) ProtoMessage() {}

func (x *GetPowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_component_powersensor_v1_powersensor_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerResponse.ProtoReflect.Descriptor instead.
func (*GetPowerResponse) Descriptor() ([]byte, []int) {
	return file_component_powersensor_v1_powersensor_proto_rawDescGZIP(), []int{5}
}

func (x *GetPowerResponse) GetWatts() float64 {
	if x != nil {
		return x.Watts
	}
	return 0
}

var File_component_powersensor_v1_powersensor_proto protoreflect.FileDescriptor

var file_component_powersensor_v1_powersensor_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x76, 0x69,
	0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x6f,
	0x6c, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x6f, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x6f,
	0x6c, 0x74, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x41, 0x63, 0x22, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6d, 0x70, 0x65, 0x72, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x6d, 0x70, 0x65, 0x72, 0x65, 0x73,
	0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x69, 0x73, 0x41, 0x63, 0x22, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x28, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x77, 0x61, 0x74, 0x74, 0x73, 0x32, 0xc4, 0x06, 0x0a, 0x12, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xad, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x30, 0x2e, 0x76, 0x69,
	0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0xad, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x76, 0x69,
	0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0xa5, 0x01, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x76, 0x69, 0x61, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x76, 0x69, 0x61, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x93, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x22, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x09, 0x44,
	0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x37, 0x22, 0x35, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x64, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x4d, 0x0a, 0x21,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_component_powersensor_v1_powersensor_proto_rawDescOnce sync.Once
	file_component_powersensor_v1_powersensor_proto_rawDescData = file_component_powersensor_v1_powersensor_proto_rawDesc
)

func file_component_powersensor_v1_powersensor_proto_rawDescGZIP() []byte {
	file_component_powersensor_v1_powersensor_proto_rawDescOnce.Do(func() {
		file_component_powersensor_v1_powersensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_component_powersensor_v1_powersensor_proto_rawDescData)
	})
	return file_component_powersensor_v1_powersensor_proto_rawDescData
}

var file_component_powersensor_v1_powersensor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_component_powersensor_v1_powersensor_proto_goTypes = []any{
	(*GetVoltageRequest)(nil),      // 0: viam.component.powersensor.v1.GetVoltageRequest
	(*GetVoltageResponse)(nil),     // 1: viam.component.powersensor.v1.GetVoltageResponse
	(*GetCurrentRequest)(nil),      // 2: viam.component.powersensor.v1.GetCurrentRequest
	(*GetCurrentResponse)(nil),     // 3: viam.component.powersensor.v1.GetCurrentResponse
	(*GetPowerRequest)(nil),        // 4: viam.component.powersensor.v1.GetPowerRequest
	(*GetPowerResponse)(nil),       // 5: viam.component.powersensor.v1.GetPowerResponse
	(*structpb.Struct)(nil),        // 6: google.protobuf.Struct
	(*v1.GetReadingsRequest)(nil),  // 7: viam.common.v1.GetReadingsRequest
	(*v1.DoCommandRequest)(nil),    // 8: viam.common.v1.DoCommandRequest
	(*v1.GetReadingsResponse)(nil), // 9: viam.common.v1.GetReadingsResponse
	(*v1.DoCommandResponse)(nil),   // 10: viam.common.v1.DoCommandResponse
}
var file_component_powersensor_v1_powersensor_proto_depIdxs = []int32{
	6,  // 0: viam.component.powersensor.v1.GetVoltageRequest.extra:type_name -> google.protobuf.Struct
	6,  // 1: viam.component.powersensor.v1.GetCurrentRequest.extra:type_name -> google.protobuf.Struct
	6,  // 2: viam.component.powersensor.v1.GetPowerRequest.extra:type_name -> google.protobuf.Struct
	0,  // 3: viam.component.powersensor.v1.PowerSensorService.GetVoltage:input_type -> viam.component.powersensor.v1.GetVoltageRequest
	2,  // 4: viam.component.powersensor.v1.PowerSensorService.GetCurrent:input_type -> viam.component.powersensor.v1.GetCurrentRequest
	4,  // 5: viam.component.powersensor.v1.PowerSensorService.GetPower:input_type -> viam.component.powersensor.v1.GetPowerRequest
	7,  // 6: viam.component.powersensor.v1.PowerSensorService.GetReadings:input_type -> viam.common.v1.GetReadingsRequest
	8,  // 7: viam.component.powersensor.v1.PowerSensorService.DoCommand:input_type -> viam.common.v1.DoCommandRequest
	1,  // 8: viam.component.powersensor.v1.PowerSensorService.GetVoltage:output_type -> viam.component.powersensor.v1.GetVoltageResponse
	3,  // 9: viam.component.powersensor.v1.PowerSensorService.GetCurrent:output_type -> viam.component.powersensor.v1.GetCurrentResponse
	5,  // 10: viam.component.powersensor.v1.PowerSensorService.GetPower:output_type -> viam.component.powersensor.v1.GetPowerResponse
	9,  // 11: viam.component.powersensor.v1.PowerSensorService.GetReadings:output_type -> viam.common.v1.GetReadingsResponse
	10, // 12: viam.component.powersensor.v1.PowerSensorService.DoCommand:output_type -> viam.common.v1.DoCommandResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_component_powersensor_v1_powersensor_proto_init() }
func file_component_powersensor_v1_powersensor_proto_init() {
	if File_component_powersensor_v1_powersensor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_component_powersensor_v1_powersensor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_component_powersensor_v1_powersensor_proto_goTypes,
		DependencyIndexes: file_component_powersensor_v1_powersensor_proto_depIdxs,
		MessageInfos:      file_component_powersensor_v1_powersensor_proto_msgTypes,
	}.Build()
	File_component_powersensor_v1_powersensor_proto = out.File
	file_component_powersensor_v1_powersensor_proto_rawDesc = nil
	file_component_powersensor_v1_powersensor_proto_goTypes = nil
	file_component_powersensor_v1_powersensor_proto_depIdxs = nil
}
