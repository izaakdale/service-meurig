// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: api/commerce/v1/commerce.proto

package v1

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

type ShoppingCart_ShoppingCartStatus int32

const (
	ShoppingCart_ACTIVE    ShoppingCart_ShoppingCartStatus = 0
	ShoppingCart_COMPLETED ShoppingCart_ShoppingCartStatus = 1
	ShoppingCart_CANCELLED ShoppingCart_ShoppingCartStatus = 2
)

// Enum value maps for ShoppingCart_ShoppingCartStatus.
var (
	ShoppingCart_ShoppingCartStatus_name = map[int32]string{
		0: "ACTIVE",
		1: "COMPLETED",
		2: "CANCELLED",
	}
	ShoppingCart_ShoppingCartStatus_value = map[string]int32{
		"ACTIVE":    0,
		"COMPLETED": 1,
		"CANCELLED": 2,
	}
)

func (x ShoppingCart_ShoppingCartStatus) Enum() *ShoppingCart_ShoppingCartStatus {
	p := new(ShoppingCart_ShoppingCartStatus)
	*p = x
	return p
}

func (x ShoppingCart_ShoppingCartStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShoppingCart_ShoppingCartStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commerce_v1_commerce_proto_enumTypes[0].Descriptor()
}

func (ShoppingCart_ShoppingCartStatus) Type() protoreflect.EnumType {
	return &file_api_commerce_v1_commerce_proto_enumTypes[0]
}

func (x ShoppingCart_ShoppingCartStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShoppingCart_ShoppingCartStatus.Descriptor instead.
func (ShoppingCart_ShoppingCartStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{2, 0}
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl    string  `protobuf:"bytes,5,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Aisle       string  `protobuf:"bytes,6,opt,name=aisle,proto3" json:"aisle,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commerce_v1_commerce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_api_commerce_v1_commerce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Product) GetAisle() string {
	if x != nil {
		return x.Aisle
	}
	return ""
}

type ProductWithQuantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product  *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Quantity int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ProductWithQuantity) Reset() {
	*x = ProductWithQuantity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commerce_v1_commerce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductWithQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductWithQuantity) ProtoMessage() {}

func (x *ProductWithQuantity) ProtoReflect() protoreflect.Message {
	mi := &file_api_commerce_v1_commerce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductWithQuantity.ProtoReflect.Descriptor instead.
func (*ProductWithQuantity) Descriptor() ([]byte, []int) {
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{1}
}

func (x *ProductWithQuantity) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ProductWithQuantity) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ShoppingCart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Products   []*ProductWithQuantity          `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Status     ShoppingCart_ShoppingCartStatus `protobuf:"varint,3,opt,name=status,proto3,enum=ShoppingCart_ShoppingCartStatus" json:"status,omitempty"`
	TotalPrice float64                         `protobuf:"fixed64,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *ShoppingCart) Reset() {
	*x = ShoppingCart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commerce_v1_commerce_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShoppingCart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingCart) ProtoMessage() {}

func (x *ShoppingCart) ProtoReflect() protoreflect.Message {
	mi := &file_api_commerce_v1_commerce_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingCart.ProtoReflect.Descriptor instead.
func (*ShoppingCart) Descriptor() ([]byte, []int) {
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{2}
}

func (x *ShoppingCart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShoppingCart) GetProducts() []*ProductWithQuantity {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ShoppingCart) GetStatus() ShoppingCart_ShoppingCartStatus {
	if x != nil {
		return x.Status
	}
	return ShoppingCart_ACTIVE
}

func (x *ShoppingCart) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShoppingCart          *ShoppingCart `protobuf:"bytes,2,opt,name=shopping_cart,json=shoppingCart,proto3" json:"shopping_cart,omitempty"`
	CustomerName          string        `protobuf:"bytes,3,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	DeliveryAddress       string        `protobuf:"bytes,4,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	DeliveryInstructions  string        `protobuf:"bytes,5,opt,name=delivery_instructions,json=deliveryInstructions,proto3" json:"delivery_instructions,omitempty"`
	OrderPlacedDate       int64         `protobuf:"varint,6,opt,name=order_placed_date,json=orderPlacedDate,proto3" json:"order_placed_date,omitempty"`
	EstimatedDeliveryDate int64         `protobuf:"varint,7,opt,name=estimated_delivery_date,json=estimatedDeliveryDate,proto3" json:"estimated_delivery_date,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commerce_v1_commerce_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_commerce_v1_commerce_proto_msgTypes[3]
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
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetShoppingCart() *ShoppingCart {
	if x != nil {
		return x.ShoppingCart
	}
	return nil
}

func (x *Order) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *Order) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

func (x *Order) GetDeliveryInstructions() string {
	if x != nil {
		return x.DeliveryInstructions
	}
	return ""
}

func (x *Order) GetOrderPlacedDate() int64 {
	if x != nil {
		return x.OrderPlacedDate
	}
	return 0
}

func (x *Order) GetEstimatedDeliveryDate() int64 {
	if x != nil {
		return x.EstimatedDeliveryDate
	}
	return 0
}

type StoreSection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AisleNumber int32  `protobuf:"varint,3,opt,name=aisle_number,json=aisleNumber,proto3" json:"aisle_number,omitempty"`
}

func (x *StoreSection) Reset() {
	*x = StoreSection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commerce_v1_commerce_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreSection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreSection) ProtoMessage() {}

func (x *StoreSection) ProtoReflect() protoreflect.Message {
	mi := &file_api_commerce_v1_commerce_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreSection.ProtoReflect.Descriptor instead.
func (*StoreSection) Descriptor() ([]byte, []int) {
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{4}
}

func (x *StoreSection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoreSection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreSection) GetAisleNumber() int32 {
	if x != nil {
		return x.AisleNumber
	}
	return 0
}

type ProductListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProductListRequest) Reset() {
	*x = ProductListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commerce_v1_commerce_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListRequest) ProtoMessage() {}

func (x *ProductListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_commerce_v1_commerce_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListRequest.ProtoReflect.Descriptor instead.
func (*ProductListRequest) Descriptor() ([]byte, []int) {
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{5}
}

type ProductListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ProductListResponse) Reset() {
	*x = ProductListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commerce_v1_commerce_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListResponse) ProtoMessage() {}

func (x *ProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_commerce_v1_commerce_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListResponse.ProtoReflect.Descriptor instead.
func (*ProductListResponse) Descriptor() ([]byte, []int) {
	return file_api_commerce_v1_commerce_proto_rawDescGZIP(), []int{6}
}

func (x *ProductListResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_api_commerce_v1_commerce_proto protoreflect.FileDescriptor

var file_api_commerce_v1_commerce_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x69, 0x73, 0x6c, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x69, 0x73, 0x6c, 0x65, 0x22, 0x55, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0xeb, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43,
	0x61, 0x72, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x3e, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22,
	0xb4, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x0c, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a,
	0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x36,
	0x0a, 0x17, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x15, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x69,
	0x73, 0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x61, 0x69, 0x73, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x14, 0x0a,
	0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x32, 0x43, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commerce_v1_commerce_proto_rawDescOnce sync.Once
	file_api_commerce_v1_commerce_proto_rawDescData = file_api_commerce_v1_commerce_proto_rawDesc
)

func file_api_commerce_v1_commerce_proto_rawDescGZIP() []byte {
	file_api_commerce_v1_commerce_proto_rawDescOnce.Do(func() {
		file_api_commerce_v1_commerce_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commerce_v1_commerce_proto_rawDescData)
	})
	return file_api_commerce_v1_commerce_proto_rawDescData
}

var file_api_commerce_v1_commerce_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commerce_v1_commerce_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_commerce_v1_commerce_proto_goTypes = []interface{}{
	(ShoppingCart_ShoppingCartStatus)(0), // 0: ShoppingCart.ShoppingCartStatus
	(*Product)(nil),                      // 1: Product
	(*ProductWithQuantity)(nil),          // 2: ProductWithQuantity
	(*ShoppingCart)(nil),                 // 3: ShoppingCart
	(*Order)(nil),                        // 4: Order
	(*StoreSection)(nil),                 // 5: StoreSection
	(*ProductListRequest)(nil),           // 6: ProductListRequest
	(*ProductListResponse)(nil),          // 7: ProductListResponse
}
var file_api_commerce_v1_commerce_proto_depIdxs = []int32{
	1, // 0: ProductWithQuantity.product:type_name -> Product
	2, // 1: ShoppingCart.products:type_name -> ProductWithQuantity
	0, // 2: ShoppingCart.status:type_name -> ShoppingCart.ShoppingCartStatus
	3, // 3: Order.shopping_cart:type_name -> ShoppingCart
	1, // 4: ProductListResponse.products:type_name -> Product
	6, // 5: Store.FetchProducts:input_type -> ProductListRequest
	7, // 6: Store.FetchProducts:output_type -> ProductListResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_commerce_v1_commerce_proto_init() }
func file_api_commerce_v1_commerce_proto_init() {
	if File_api_commerce_v1_commerce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commerce_v1_commerce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_api_commerce_v1_commerce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductWithQuantity); i {
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
		file_api_commerce_v1_commerce_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShoppingCart); i {
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
		file_api_commerce_v1_commerce_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_commerce_v1_commerce_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreSection); i {
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
		file_api_commerce_v1_commerce_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListRequest); i {
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
		file_api_commerce_v1_commerce_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListResponse); i {
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
			RawDescriptor: file_api_commerce_v1_commerce_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_commerce_v1_commerce_proto_goTypes,
		DependencyIndexes: file_api_commerce_v1_commerce_proto_depIdxs,
		EnumInfos:         file_api_commerce_v1_commerce_proto_enumTypes,
		MessageInfos:      file_api_commerce_v1_commerce_proto_msgTypes,
	}.Build()
	File_api_commerce_v1_commerce_proto = out.File
	file_api_commerce_v1_commerce_proto_rawDesc = nil
	file_api_commerce_v1_commerce_proto_goTypes = nil
	file_api_commerce_v1_commerce_proto_depIdxs = nil
}
