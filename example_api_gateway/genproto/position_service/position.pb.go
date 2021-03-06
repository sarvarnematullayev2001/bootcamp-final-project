// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: position.proto

package position_service

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

type PositionAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeId string `protobuf:"bytes,1,opt,name=attribute_id,json=attributeId,proto3" json:"attribute_id,omitempty"`
	Value       string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PositionAttribute) Reset() {
	*x = PositionAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionAttribute) ProtoMessage() {}

func (x *PositionAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionAttribute.ProtoReflect.Descriptor instead.
func (*PositionAttribute) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{0}
}

func (x *PositionAttribute) GetAttributeId() string {
	if x != nil {
		return x.AttributeId
	}
	return ""
}

func (x *PositionAttribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PositionAttribute2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PositionAttribute2) Reset() {
	*x = PositionAttribute2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionAttribute2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionAttribute2) ProtoMessage() {}

func (x *PositionAttribute2) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionAttribute2.ProtoReflect.Descriptor instead.
func (*PositionAttribute2) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{1}
}

func (x *PositionAttribute2) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PositionAttribute2) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetPositionAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AttributeId string     `protobuf:"bytes,2,opt,name=attribute_id,json=attributeId,proto3" json:"attribute_id,omitempty"`
	PositionId  string     `protobuf:"bytes,3,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Value       string     `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Attribute   *Attribute `protobuf:"bytes,5,opt,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *GetPositionAttribute) Reset() {
	*x = GetPositionAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPositionAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionAttribute) ProtoMessage() {}

func (x *GetPositionAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionAttribute.ProtoReflect.Descriptor instead.
func (*GetPositionAttribute) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{2}
}

func (x *GetPositionAttribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPositionAttribute) GetAttributeId() string {
	if x != nil {
		return x.AttributeId
	}
	return ""
}

func (x *GetPositionAttribute) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *GetPositionAttribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetPositionAttribute) GetAttribute() *Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProfessionId       string                  `protobuf:"bytes,3,opt,name=profession_id,json=professionId,proto3" json:"profession_id,omitempty"`
	CompanyId          string                  `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	PositionAttributes []*GetPositionAttribute `protobuf:"bytes,5,rep,name=position_attributes,json=positionAttributes,proto3" json:"position_attributes,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{3}
}

func (x *Position) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Position) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Position) GetProfessionId() string {
	if x != nil {
		return x.ProfessionId
	}
	return ""
}

func (x *Position) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *Position) GetPositionAttributes() []*GetPositionAttribute {
	if x != nil {
		return x.PositionAttributes
	}
	return nil
}

type CreatePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProfessionId       string               `protobuf:"bytes,2,opt,name=profession_id,json=professionId,proto3" json:"profession_id,omitempty"`
	CompanyId          string               `protobuf:"bytes,3,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	PositionAttributes []*PositionAttribute `protobuf:"bytes,4,rep,name=position_attributes,json=positionAttributes,proto3" json:"position_attributes,omitempty"`
}

func (x *CreatePositionRequest) Reset() {
	*x = CreatePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePositionRequest) ProtoMessage() {}

func (x *CreatePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePositionRequest.ProtoReflect.Descriptor instead.
func (*CreatePositionRequest) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePositionRequest) GetProfessionId() string {
	if x != nil {
		return x.ProfessionId
	}
	return ""
}

func (x *CreatePositionRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CreatePositionRequest) GetPositionAttributes() []*PositionAttribute {
	if x != nil {
		return x.PositionAttributes
	}
	return nil
}

type PositionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PositionId) Reset() {
	*x = PositionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionId) ProtoMessage() {}

func (x *PositionId) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionId.ProtoReflect.Descriptor instead.
func (*PositionId) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{5}
}

func (x *PositionId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit        uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset       uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ProfessionId string `protobuf:"bytes,4,opt,name=profession_id,json=professionId,proto3" json:"profession_id,omitempty"`
	CompanyId    string `protobuf:"bytes,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *GetAllPositionRequest) Reset() {
	*x = GetAllPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPositionRequest) ProtoMessage() {}

func (x *GetAllPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPositionRequest.ProtoReflect.Descriptor instead.
func (*GetAllPositionRequest) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllPositionRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllPositionRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllPositionRequest) GetProfessionId() string {
	if x != nil {
		return x.ProfessionId
	}
	return ""
}

func (x *GetAllPositionRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

type GetAllPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*Position `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
	Count     uint32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllPositionResponse) Reset() {
	*x = GetAllPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPositionResponse) ProtoMessage() {}

func (x *GetAllPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPositionResponse.ProtoReflect.Descriptor instead.
func (*GetAllPositionResponse) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllPositionResponse) GetPositions() []*Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *GetAllPositionResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdatePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProfessionId      string                `protobuf:"bytes,3,opt,name=profession_id,json=professionId,proto3" json:"profession_id,omitempty"`
	CompanyId         string                `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	PositionAttribute []*PositionAttribute2 `protobuf:"bytes,5,rep,name=position_attribute,json=positionAttribute,proto3" json:"position_attribute,omitempty"`
	Attribute         []*Attribute          `protobuf:"bytes,6,rep,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *UpdatePositionRequest) Reset() {
	*x = UpdatePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePositionRequest) ProtoMessage() {}

func (x *UpdatePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePositionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePositionRequest) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePositionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePositionRequest) GetProfessionId() string {
	if x != nil {
		return x.ProfessionId
	}
	return ""
}

func (x *UpdatePositionRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *UpdatePositionRequest) GetPositionAttribute() []*PositionAttribute2 {
	if x != nil {
		return x.PositionAttribute
	}
	return nil
}

func (x *UpdatePositionRequest) GetAttribute() []*Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type UpdatePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp string `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *UpdatePosition) Reset() {
	*x = UpdatePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosition) ProtoMessage() {}

func (x *UpdatePosition) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosition.ProtoReflect.Descriptor instead.
func (*UpdatePosition) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePosition) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

var File_position_proto protoreflect.FileDescriptor

var file_position_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x11, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x32, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x08,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x12, 0x4f, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x12, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x22, 0xbd, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x12, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x22, 0x1c, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x9d, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22,
	0x60, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xff, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x12, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x32, 0x52, 0x11, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x31, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_position_proto_rawDescOnce sync.Once
	file_position_proto_rawDescData = file_position_proto_rawDesc
)

func file_position_proto_rawDescGZIP() []byte {
	file_position_proto_rawDescOnce.Do(func() {
		file_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_position_proto_rawDescData)
	})
	return file_position_proto_rawDescData
}

var file_position_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_position_proto_goTypes = []interface{}{
	(*PositionAttribute)(nil),      // 0: genproto.PositionAttribute
	(*PositionAttribute2)(nil),     // 1: genproto.PositionAttribute2
	(*GetPositionAttribute)(nil),   // 2: genproto.GetPositionAttribute
	(*Position)(nil),               // 3: genproto.Position
	(*CreatePositionRequest)(nil),  // 4: genproto.CreatePositionRequest
	(*PositionId)(nil),             // 5: genproto.PositionId
	(*GetAllPositionRequest)(nil),  // 6: genproto.GetAllPositionRequest
	(*GetAllPositionResponse)(nil), // 7: genproto.GetAllPositionResponse
	(*UpdatePositionRequest)(nil),  // 8: genproto.UpdatePositionRequest
	(*UpdatePosition)(nil),         // 9: genproto.UpdatePosition
	(*Attribute)(nil),              // 10: genproto.Attribute
}
var file_position_proto_depIdxs = []int32{
	10, // 0: genproto.GetPositionAttribute.attribute:type_name -> genproto.Attribute
	2,  // 1: genproto.Position.position_attributes:type_name -> genproto.GetPositionAttribute
	0,  // 2: genproto.CreatePositionRequest.position_attributes:type_name -> genproto.PositionAttribute
	3,  // 3: genproto.GetAllPositionResponse.positions:type_name -> genproto.Position
	1,  // 4: genproto.UpdatePositionRequest.position_attribute:type_name -> genproto.PositionAttribute2
	10, // 5: genproto.UpdatePositionRequest.attribute:type_name -> genproto.Attribute
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_position_proto_init() }
func file_position_proto_init() {
	if File_position_proto != nil {
		return
	}
	file_attribute_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionAttribute); i {
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
		file_position_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionAttribute2); i {
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
		file_position_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPositionAttribute); i {
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
		file_position_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_position_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePositionRequest); i {
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
		file_position_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionId); i {
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
		file_position_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPositionRequest); i {
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
		file_position_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPositionResponse); i {
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
		file_position_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePositionRequest); i {
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
		file_position_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosition); i {
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
			RawDescriptor: file_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_position_proto_goTypes,
		DependencyIndexes: file_position_proto_depIdxs,
		MessageInfos:      file_position_proto_msgTypes,
	}.Build()
	File_position_proto = out.File
	file_position_proto_rawDesc = nil
	file_position_proto_goTypes = nil
	file_position_proto_depIdxs = nil
}
