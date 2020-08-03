// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: proto/product.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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
		mi := &file_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SneakerProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SneakerProduct) ProtoMessage() {}

func (x *SneakerProduct) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[0]
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
	return file_proto_product_proto_rawDescGZIP(), []int{0}
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
		mi := &file_proto_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SneakerSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SneakerSize) ProtoMessage() {}

func (x *SneakerSize) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[1]
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
	return file_proto_product_proto_rawDescGZIP(), []int{1}
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

	ProductID     []string          `protobuf:"bytes,1,rep,name=ProductID,proto3" json:"ProductID,omitempty"`
	RequestQuery  map[string]string `protobuf:"bytes,2,rep,name=RequestQuery,proto3" json:"RequestQuery,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequestParams *RequestParams    `protobuf:"bytes,3,opt,name=RequestParams,proto3" json:"RequestParams,omitempty"`
}

func (x *ProductFilter) Reset() {
	*x = ProductFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductFilter) ProtoMessage() {}

func (x *ProductFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[2]
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
	return file_proto_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductFilter) GetProductID() []string {
	if x != nil {
		return x.ProductID
	}
	return nil
}

func (x *ProductFilter) GetRequestQuery() map[string]string {
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

	Products      []*SneakerProduct `protobuf:"bytes,1,rep,name=Products,proto3" json:"Products,omitempty"`
	RequestParams *RequestParams    `protobuf:"bytes,2,opt,name=RequestParams,proto3" json:"RequestParams,omitempty"`
}

func (x *ProductInput) Reset() {
	*x = ProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInput) ProtoMessage() {}

func (x *ProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[3]
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
	return file_proto_product_proto_rawDescGZIP(), []int{3}
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

	Products []*SneakerProduct `protobuf:"bytes,1,rep,name=Products,proto3" json:"Products,omitempty"`
	Count    int64             `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[4]
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
	return file_proto_product_proto_rawDescGZIP(), []int{4}
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

type RequestParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit         int32  `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset        int32  `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	SortBy        string `protobuf:"bytes,3,opt,name=SortBy,proto3" json:"SortBy,omitempty"`
	SortDirection string `protobuf:"bytes,4,opt,name=SortDirection,proto3" json:"SortDirection,omitempty"`
}

func (x *RequestParams) Reset() {
	*x = RequestParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestParams) ProtoMessage() {}

func (x *RequestParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestParams.ProtoReflect.Descriptor instead.
func (*RequestParams) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{5}
}

func (x *RequestParams) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RequestParams) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *RequestParams) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *RequestParams) GetSortDirection() string {
	if x != nil {
		return x.SortDirection
	}
	return ""
}

var File_proto_product_proto protoreflect.FileDescriptor

var file_proto_product_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x03,
	0x0a, 0x0e, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x53, 0x4b, 0x55, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x53, 0x4b, 0x55, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x34, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x6e,
	0x69, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x4b, 0x69, 0x6e, 0x67, 0x64, 0x6f, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x64, 0x4b, 0x69, 0x6e,
	0x67, 0x64, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x65, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x4a, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x3a, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x3f,
	0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x7d, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x31, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6e, 0x65, 0x61, 0x6b, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x5a,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6e, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7b, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x95, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3f,
	0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_proto_rawDescOnce sync.Once
	file_proto_product_proto_rawDescData = file_proto_product_proto_rawDesc
)

func file_proto_product_proto_rawDescGZIP() []byte {
	file_proto_product_proto_rawDescOnce.Do(func() {
		file_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_proto_rawDescData)
	})
	return file_proto_product_proto_rawDescData
}

var file_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_product_proto_goTypes = []interface{}{
	(*SneakerProduct)(nil),      // 0: proto.SneakerProduct
	(*SneakerSize)(nil),         // 1: proto.SneakerSize
	(*ProductFilter)(nil),       // 2: proto.ProductFilter
	(*ProductInput)(nil),        // 3: proto.ProductInput
	(*ProductResponse)(nil),     // 4: proto.ProductResponse
	(*RequestParams)(nil),       // 5: proto.RequestParams
	nil,                         // 6: proto.ProductFilter.RequestQueryEntry
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_proto_product_proto_depIdxs = []int32{
	1,  // 0: proto.SneakerProduct.size:type_name -> proto.SneakerSize
	7,  // 1: proto.SneakerProduct.addedAt:type_name -> google.protobuf.Timestamp
	6,  // 2: proto.ProductFilter.RequestQuery:type_name -> proto.ProductFilter.RequestQueryEntry
	5,  // 3: proto.ProductFilter.RequestParams:type_name -> proto.RequestParams
	0,  // 4: proto.ProductInput.Products:type_name -> proto.SneakerProduct
	5,  // 5: proto.ProductInput.RequestParams:type_name -> proto.RequestParams
	0,  // 6: proto.ProductResponse.Products:type_name -> proto.SneakerProduct
	2,  // 7: proto.ProductService.GetProducts:input_type -> proto.ProductFilter
	2,  // 8: proto.ProductService.CountProducts:input_type -> proto.ProductFilter
	3,  // 9: proto.ProductService.AddProducts:input_type -> proto.ProductInput
	3,  // 10: proto.ProductService.EditProducts:input_type -> proto.ProductInput
	4,  // 11: proto.ProductService.GetProducts:output_type -> proto.ProductResponse
	4,  // 12: proto.ProductService.CountProducts:output_type -> proto.ProductResponse
	4,  // 13: proto.ProductService.AddProducts:output_type -> proto.ProductResponse
	4,  // 14: proto.ProductService.EditProducts:output_type -> proto.ProductResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_product_proto_init() }
func file_proto_product_proto_init() {
	if File_proto_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestParams); i {
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
			RawDescriptor: file_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_proto_goTypes,
		DependencyIndexes: file_proto_product_proto_depIdxs,
		MessageInfos:      file_proto_product_proto_msgTypes,
	}.Build()
	File_proto_product_proto = out.File
	file_proto_product_proto_rawDesc = nil
	file_proto_product_proto_goTypes = nil
	file_proto_product_proto_depIdxs = nil
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
	GetProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (ProductService_GetProductsClient, error)
	CountProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (ProductService_CountProductsClient, error)
	AddProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (ProductService_AddProductsClient, error)
	EditProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (ProductService_EditProductsClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (ProductService_GetProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[0], "/proto.ProductService/GetProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceGetProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_GetProductsClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type productServiceGetProductsClient struct {
	grpc.ClientStream
}

func (x *productServiceGetProductsClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) CountProducts(ctx context.Context, in *ProductFilter, opts ...grpc.CallOption) (ProductService_CountProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[1], "/proto.ProductService/CountProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceCountProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_CountProductsClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type productServiceCountProductsClient struct {
	grpc.ClientStream
}

func (x *productServiceCountProductsClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) AddProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (ProductService_AddProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[2], "/proto.ProductService/AddProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceAddProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_AddProductsClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type productServiceAddProductsClient struct {
	grpc.ClientStream
}

func (x *productServiceAddProductsClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) EditProducts(ctx context.Context, in *ProductInput, opts ...grpc.CallOption) (ProductService_EditProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[3], "/proto.ProductService/EditProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceEditProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_EditProductsClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type productServiceEditProductsClient struct {
	grpc.ClientStream
}

func (x *productServiceEditProductsClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProducts(*ProductFilter, ProductService_GetProductsServer) error
	CountProducts(*ProductFilter, ProductService_CountProductsServer) error
	AddProducts(*ProductInput, ProductService_AddProductsServer) error
	EditProducts(*ProductInput, ProductService_EditProductsServer) error
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProducts(*ProductFilter, ProductService_GetProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (*UnimplementedProductServiceServer) CountProducts(*ProductFilter, ProductService_CountProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method CountProducts not implemented")
}
func (*UnimplementedProductServiceServer) AddProducts(*ProductInput, ProductService_AddProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method AddProducts not implemented")
}
func (*UnimplementedProductServiceServer) EditProducts(*ProductInput, ProductService_EditProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method EditProducts not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).GetProducts(m, &productServiceGetProductsServer{stream})
}

type ProductService_GetProductsServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type productServiceGetProductsServer struct {
	grpc.ServerStream
}

func (x *productServiceGetProductsServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_CountProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).CountProducts(m, &productServiceCountProductsServer{stream})
}

type ProductService_CountProductsServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type productServiceCountProductsServer struct {
	grpc.ServerStream
}

func (x *productServiceCountProductsServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_AddProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).AddProducts(m, &productServiceAddProductsServer{stream})
}

type ProductService_AddProductsServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type productServiceAddProductsServer struct {
	grpc.ServerStream
}

func (x *productServiceAddProductsServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_EditProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).EditProducts(m, &productServiceEditProductsServer{stream})
}

type ProductService_EditProductsServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type productServiceEditProductsServer struct {
	grpc.ServerStream
}

func (x *productServiceEditProductsServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProducts",
			Handler:       _ProductService_GetProducts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CountProducts",
			Handler:       _ProductService_CountProducts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddProducts",
			Handler:       _ProductService_AddProducts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EditProducts",
			Handler:       _ProductService_EditProducts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}
