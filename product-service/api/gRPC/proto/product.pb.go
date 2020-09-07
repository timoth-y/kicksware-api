// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: product.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SneakerProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId       string               `protobuf:"bytes,1,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	BrandName      string               `protobuf:"bytes,2,opt,name=brandName,proto3" json:"brandName,omitempty"`
	ModelName      string               `protobuf:"bytes,3,opt,name=modelName,proto3" json:"modelName,omitempty"`
	ModelSKU       string               `protobuf:"bytes,4,opt,name=modelSKU,proto3" json:"modelSKU,omitempty"`
	ReferenceId    string               `protobuf:"bytes,5,opt,name=referenceId,proto3" json:"referenceId,omitempty"`
	Price          float64              `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Type           string               `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Size           *SneakerSize         `protobuf:"bytes,8,opt,name=size,proto3" json:"size,omitempty"`
	Color          string               `protobuf:"bytes,9,opt,name=color,proto3" json:"color,omitempty"`
	Condition      string               `protobuf:"bytes,10,opt,name=condition,proto3" json:"condition,omitempty"`
	Description    string               `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Owner          string               `protobuf:"bytes,12,opt,name=owner,proto3" json:"owner,omitempty"`
	ConditionIndex float64              `protobuf:"fixed64,14,opt,name=conditionIndex,proto3" json:"conditionIndex,omitempty"`
	AddedAt        *timestamp.Timestamp `protobuf:"bytes,15,opt,name=addedAt,proto3" json:"addedAt,omitempty"`
}

func (x *SneakerProduct) Reset() {
	*x = SneakerProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SneakerProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SneakerProduct) ProtoMessage() {}

func (x *SneakerProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SneakerProduct.ProtoReflect.Descriptor instead.
func (*SneakerProduct) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *SneakerProduct) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *SneakerProduct) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *SneakerProduct) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *SneakerProduct) GetModelSKU() string {
	if x != nil {
		return x.ModelSKU
	}
	return ""
}

func (x *SneakerProduct) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *SneakerProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SneakerProduct) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SneakerProduct) GetSize() *SneakerSize {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *SneakerProduct) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *SneakerProduct) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *SneakerProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SneakerProduct) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SneakerProduct) GetConditionIndex() float64 {
	if x != nil {
		return x.ConditionIndex
	}
	return 0
}

func (x *SneakerProduct) GetAddedAt() *timestamp.Timestamp {
	if x != nil {
		return x.AddedAt
	}
	return nil
}

type SneakerSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Europe        float64 `protobuf:"fixed64,1,opt,name=europe,proto3" json:"europe,omitempty"`
	UnitedStates  float64 `protobuf:"fixed64,2,opt,name=unitedStates,proto3" json:"unitedStates,omitempty"`
	UnitedKingdom float64 `protobuf:"fixed64,3,opt,name=unitedKingdom,proto3" json:"unitedKingdom,omitempty"`
	Centimeters   float64 `protobuf:"fixed64,4,opt,name=centimeters,proto3" json:"centimeters,omitempty"`
}

func (x *SneakerSize) Reset() {
	*x = SneakerSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SneakerSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SneakerSize) ProtoMessage() {}

func (x *SneakerSize) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SneakerSize.ProtoReflect.Descriptor instead.
func (*SneakerSize) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *SneakerSize) GetEurope() float64 {
	if x != nil {
		return x.Europe
	}
	return 0
}

func (x *SneakerSize) GetUnitedStates() float64 {
	if x != nil {
		return x.UnitedStates
	}
	return 0
}

func (x *SneakerSize) GetUnitedKingdom() float64 {
	if x != nil {
		return x.UnitedKingdom
	}
	return 0
}

func (x *SneakerSize) GetCentimeters() float64 {
	if x != nil {
		return x.Centimeters
	}
	return 0
}

type ProductFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID     []string        `protobuf:"bytes,1,rep,name=productID,proto3" json:"productID,omitempty"`
	RequestQuery  *_struct.Struct `protobuf:"bytes,2,opt,name=requestQuery,proto3" json:"requestQuery,omitempty"`
	RequestParams *RequestParams  `protobuf:"bytes,3,opt,name=requestParams,proto3" json:"requestParams,omitempty"`
}

func (x *ProductFilter) Reset() {
	*x = ProductFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductFilter) ProtoMessage() {}

func (x *ProductFilter) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductFilter.ProtoReflect.Descriptor instead.
func (*ProductFilter) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductFilter) GetProductID() []string {
	if x != nil {
		return x.ProductID
	}
	return nil
}

func (x *ProductFilter) GetRequestQuery() *_struct.Struct {
	if x != nil {
		return x.RequestQuery
	}
	return nil
}

func (x *ProductFilter) GetRequestParams() *RequestParams {
	if x != nil {
		return x.RequestParams
	}
	return nil
}

type ProductInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products      []*SneakerProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	RequestParams *RequestParams    `protobuf:"bytes,2,opt,name=requestParams,proto3" json:"requestParams,omitempty"`
}

func (x *ProductInput) Reset() {
	*x = ProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInput) ProtoMessage() {}

func (x *ProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInput.ProtoReflect.Descriptor instead.
func (*ProductInput) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductInput) GetProducts() []*SneakerProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductInput) GetRequestParams() *RequestParams {
	if x != nil {
		return x.RequestParams
	}
	return nil
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*SneakerProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Count    int64             `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductResponse) GetProducts() []*SneakerProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x03, 0x0a, 0x0e, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x4b, 0x55, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x4b, 0x55, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6e, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x53, 0x6e, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x75, 0x72, 0x6f,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x4b, 0x69,
	0x6e, 0x67, 0x64, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x69,
	0x74, 0x65, 0x64, 0x4b, 0x69, 0x6e, 0x67, 0x64, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xa6, 0x01, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x3b, 0x0a, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x7d, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x5a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0xcf, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x11, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x05,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_proto_goTypes = []interface{}{
	(*SneakerProduct)(nil),      // 0: proto.SneakerProduct
	(*SneakerSize)(nil),         // 1: proto.SneakerSize
	(*ProductFilter)(nil),       // 2: proto.ProductFilter
	(*ProductInput)(nil),        // 3: proto.ProductInput
	(*ProductResponse)(nil),     // 4: proto.ProductResponse
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*_struct.Struct)(nil),      // 6: google.protobuf.Struct
	(*RequestParams)(nil),       // 7: proto.RequestParams
}
var file_product_proto_depIdxs = []int32{
	1,  // 0: proto.SneakerProduct.size:type_name -> proto.SneakerSize
	5,  // 1: proto.SneakerProduct.addedAt:type_name -> google.protobuf.Timestamp
	6,  // 2: proto.ProductFilter.requestQuery:type_name -> google.protobuf.Struct
	7,  // 3: proto.ProductFilter.requestParams:type_name -> proto.RequestParams
	0,  // 4: proto.ProductInput.products:type_name -> proto.SneakerProduct
	7,  // 5: proto.ProductInput.requestParams:type_name -> proto.RequestParams
	0,  // 6: proto.ProductResponse.products:type_name -> proto.SneakerProduct
	2,  // 7: proto.ProductService.GetProducts:input_type -> proto.ProductFilter
	2,  // 8: proto.ProductService.CountProducts:input_type -> proto.ProductFilter
	3,  // 9: proto.ProductService.AddProducts:input_type -> proto.ProductInput
	3,  // 10: proto.ProductService.EditProducts:input_type -> proto.ProductInput
	2,  // 11: proto.ProductService.DeleteProducts:input_type -> proto.ProductFilter
	4,  // 12: proto.ProductService.GetProducts:output_type -> proto.ProductResponse
	4,  // 13: proto.ProductService.CountProducts:output_type -> proto.ProductResponse
	4,  // 14: proto.ProductService.AddProducts:output_type -> proto.ProductResponse
	4,  // 15: proto.ProductService.EditProducts:output_type -> proto.ProductResponse
	4,  // 16: proto.ProductService.DeleteProducts:output_type -> proto.ProductResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SneakerProduct); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SneakerSize); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductFilter); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInput); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error)
	CountProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error)
	AddProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (*ProductResponse, error)
	EditProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (*ProductResponse, error)
	DeleteProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CountProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductService/CountProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductService/AddProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) EditProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductService/EditProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductService/DeleteProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProducts(context.Context, *ProductFilter) (*ProductResponse, error)
	CountProducts(context.Context, *ProductFilter) (*ProductResponse, error)
	AddProducts(context.Context, *ProductInput) (*ProductResponse, error)
	EditProducts(context.Context, *ProductInput) (*ProductResponse, error)
	DeleteProducts(context.Context, *ProductFilter) (*ProductResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProducts(context.Context, *ProductFilter) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (*UnimplementedProductServiceServer) CountProducts(context.Context, *ProductFilter) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountProducts not implemented")
}
func (*UnimplementedProductServiceServer) AddProducts(context.Context, *ProductInput) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProducts not implemented")
}
func (*UnimplementedProductServiceServer) EditProducts(context.Context, *ProductInput) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProducts not implemented")
}
func (*UnimplementedProductServiceServer) DeleteProducts(context.Context, *ProductFilter) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProducts not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*ProductFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CountProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CountProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/CountProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CountProducts(ctx, req.(*ProductFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/AddProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProducts(ctx, req.(*ProductInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_EditProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).EditProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/EditProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).EditProducts(ctx, req.(*ProductInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/DeleteProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProducts(ctx, req.(*ProductFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
		{
			MethodName: "CountProducts",
			Handler:    _ProductService_CountProducts_Handler,
		},
		{
			MethodName: "AddProducts",
			Handler:    _ProductService_AddProducts_Handler,
		},
		{
			MethodName: "EditProducts",
			Handler:    _ProductService_EditProducts_Handler,
		},
		{
			MethodName: "DeleteProducts",
			Handler:    _ProductService_DeleteProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}