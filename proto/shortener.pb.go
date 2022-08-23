// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/shortener.proto

package proto

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

type GetOrigURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetOrigURLRequest) Reset() {
	*x = GetOrigURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrigURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrigURLRequest) ProtoMessage() {}

func (x *GetOrigURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrigURLRequest.ProtoReflect.Descriptor instead.
func (*GetOrigURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrigURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetOrigURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrigUrl string `protobuf:"bytes,1,opt,name=orig_url,json=origUrl,proto3" json:"orig_url,omitempty"`
}

func (x *GetOrigURLResponse) Reset() {
	*x = GetOrigURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrigURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrigURLResponse) ProtoMessage() {}

func (x *GetOrigURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrigURLResponse.ProtoReflect.Descriptor instead.
func (*GetOrigURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrigURLResponse) GetOrigUrl() string {
	if x != nil {
		return x.OrigUrl
	}
	return ""
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OrigUrl  string `protobuf:"bytes,2,opt,name=orig_url,json=origUrl,proto3" json:"orig_url,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *Pair) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *Pair) GetOrigUrl() string {
	if x != nil {
		return x.OrigUrl
	}
	return ""
}

type GetGetAllPairsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGetAllPairsRequest) Reset() {
	*x = GetGetAllPairsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGetAllPairsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGetAllPairsRequest) ProtoMessage() {}

func (x *GetGetAllPairsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGetAllPairsRequest.ProtoReflect.Descriptor instead.
func (*GetGetAllPairsRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{3}
}

type GetGetAllPairsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*Pair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *GetGetAllPairsResponse) Reset() {
	*x = GetGetAllPairsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGetAllPairsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGetAllPairsResponse) ProtoMessage() {}

func (x *GetGetAllPairsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGetAllPairsResponse.ProtoReflect.Descriptor instead.
func (*GetGetAllPairsResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *GetGetAllPairsResponse) GetPairs() []*Pair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type PostProvideShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrigUrl string `protobuf:"bytes,1,opt,name=orig_url,json=origUrl,proto3" json:"orig_url,omitempty"`
}

func (x *PostProvideShortURLRequest) Reset() {
	*x = PostProvideShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostProvideShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostProvideShortURLRequest) ProtoMessage() {}

func (x *PostProvideShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostProvideShortURLRequest.ProtoReflect.Descriptor instead.
func (*PostProvideShortURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *PostProvideShortURLRequest) GetOrigUrl() string {
	if x != nil {
		return x.OrigUrl
	}
	return ""
}

type PostProvideShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *PostProvideShortURLResponse) Reset() {
	*x = PostProvideShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostProvideShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostProvideShortURLResponse) ProtoMessage() {}

func (x *PostProvideShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostProvideShortURLResponse.ProtoReflect.Descriptor instead.
func (*PostProvideShortURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *PostProvideShortURLResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type PostBatchRequestPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OrigUrl       string `protobuf:"bytes,2,opt,name=orig_url,json=origUrl,proto3" json:"orig_url,omitempty"`
}

func (x *PostBatchRequestPair) Reset() {
	*x = PostBatchRequestPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBatchRequestPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBatchRequestPair) ProtoMessage() {}

func (x *PostBatchRequestPair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBatchRequestPair.ProtoReflect.Descriptor instead.
func (*PostBatchRequestPair) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *PostBatchRequestPair) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *PostBatchRequestPair) GetOrigUrl() string {
	if x != nil {
		return x.OrigUrl
	}
	return ""
}

type PostBatchResponsePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *PostBatchResponsePair) Reset() {
	*x = PostBatchResponsePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBatchResponsePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBatchResponsePair) ProtoMessage() {}

func (x *PostBatchResponsePair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBatchResponsePair.ProtoReflect.Descriptor instead.
func (*PostBatchResponsePair) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *PostBatchResponsePair) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *PostBatchResponsePair) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type PostBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqPairs []*PostBatchRequestPair `protobuf:"bytes,1,rep,name=req_pairs,json=reqPairs,proto3" json:"req_pairs,omitempty"`
}

func (x *PostBatchRequest) Reset() {
	*x = PostBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBatchRequest) ProtoMessage() {}

func (x *PostBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBatchRequest.ProtoReflect.Descriptor instead.
func (*PostBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *PostBatchRequest) GetReqPairs() []*PostBatchRequestPair {
	if x != nil {
		return x.ReqPairs
	}
	return nil
}

type PostBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespPairs []*PostBatchResponsePair `protobuf:"bytes,1,rep,name=resp_pairs,json=respPairs,proto3" json:"resp_pairs,omitempty"`
}

func (x *PostBatchResponse) Reset() {
	*x = PostBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBatchResponse) ProtoMessage() {}

func (x *PostBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBatchResponse.ProtoReflect.Descriptor instead.
func (*PostBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *PostBatchResponse) GetRespPairs() []*PostBatchResponsePair {
	if x != nil {
		return x.RespPairs
	}
	return nil
}

type DelBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrls []string `protobuf:"bytes,1,rep,name=short_urls,json=shortUrls,proto3" json:"short_urls,omitempty"`
}

func (x *DelBatchRequest) Reset() {
	*x = DelBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelBatchRequest) ProtoMessage() {}

func (x *DelBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelBatchRequest.ProtoReflect.Descriptor instead.
func (*DelBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *DelBatchRequest) GetShortUrls() []string {
	if x != nil {
		return x.ShortUrls
	}
	return nil
}

type DelBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DelBatchResponse) Reset() {
	*x = DelBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelBatchResponse) ProtoMessage() {}

func (x *DelBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelBatchResponse.ProtoReflect.Descriptor instead.
func (*DelBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{12}
}

func (x *DelBatchResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatRequest) Reset() {
	*x = GetStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatRequest) ProtoMessage() {}

func (x *GetStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatRequest.ProtoReflect.Descriptor instead.
func (*GetStatRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{13}
}

type GetStatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls  int32 `protobuf:"varint,1,opt,name=urls,proto3" json:"urls,omitempty"`
	Users int32 `protobuf:"varint,2,opt,name=users,proto3" json:"users,omitempty"`
}

func (x *GetStatResponse) Reset() {
	*x = GetStatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatResponse) ProtoMessage() {}

func (x *GetStatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatResponse.ProtoReflect.Descriptor instead.
func (*GetStatResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{14}
}

func (x *GetStatResponse) GetUrls() int32 {
	if x != nil {
		return x.Urls
	}
	return 0
}

func (x *GetStatResponse) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

var File_proto_shortener_proto protoreflect.FileDescriptor

var file_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x69, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x69, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x3e, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x69, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x69, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4f,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x22,
	0x37, 0x0a, 0x1a, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x69, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x3a, 0x0a, 0x1b, 0x50, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x58, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x69, 0x72, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x5b,
	0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x60, 0x0a, 0x10, 0x50,
	0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x08, 0x72, 0x65, 0x71, 0x50, 0x61, 0x69, 0x72, 0x73, 0x22, 0x64, 0x0a,
	0x11, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xa3,
	0x05, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x55, 0x52, 0x4c, 0x12, 0x2c, 0x2e, 0x55, 0x52, 0x4c,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x30, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x69, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x13,
	0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x12, 0x35, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x55, 0x52, 0x4c,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x66, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x2b, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x55,
	0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2a, 0x2e, 0x55, 0x52, 0x4c, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x12, 0x29, 0x2e,
	0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x45, 0x65, 0x73, 0x74, 0x69, 0x43, 0x68, 0x61, 0x6d, 0x65, 0x6c, 0x65, 0x6f,
	0x6e, 0x2f, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_shortener_proto_rawDescOnce sync.Once
	file_proto_shortener_proto_rawDescData = file_proto_shortener_proto_rawDesc
)

func file_proto_shortener_proto_rawDescGZIP() []byte {
	file_proto_shortener_proto_rawDescOnce.Do(func() {
		file_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shortener_proto_rawDescData)
	})
	return file_proto_shortener_proto_rawDescData
}

var file_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_shortener_proto_goTypes = []interface{}{
	(*GetOrigURLRequest)(nil),           // 0: URLShortenerService.proto.GetOrigURLRequest
	(*GetOrigURLResponse)(nil),          // 1: URLShortenerService.proto.GetOrigURLResponse
	(*Pair)(nil),                        // 2: URLShortenerService.proto.Pair
	(*GetGetAllPairsRequest)(nil),       // 3: URLShortenerService.proto.GetGetAllPairsRequest
	(*GetGetAllPairsResponse)(nil),      // 4: URLShortenerService.proto.GetGetAllPairsResponse
	(*PostProvideShortURLRequest)(nil),  // 5: URLShortenerService.proto.PostProvideShortURLRequest
	(*PostProvideShortURLResponse)(nil), // 6: URLShortenerService.proto.PostProvideShortURLResponse
	(*PostBatchRequestPair)(nil),        // 7: URLShortenerService.proto.PostBatchRequestPair
	(*PostBatchResponsePair)(nil),       // 8: URLShortenerService.proto.PostBatchResponsePair
	(*PostBatchRequest)(nil),            // 9: URLShortenerService.proto.PostBatchRequest
	(*PostBatchResponse)(nil),           // 10: URLShortenerService.proto.PostBatchResponse
	(*DelBatchRequest)(nil),             // 11: URLShortenerService.proto.DelBatchRequest
	(*DelBatchResponse)(nil),            // 12: URLShortenerService.proto.DelBatchResponse
	(*GetStatRequest)(nil),              // 13: URLShortenerService.proto.GetStatRequest
	(*GetStatResponse)(nil),             // 14: URLShortenerService.proto.GetStatResponse
}
var file_proto_shortener_proto_depIdxs = []int32{
	2,  // 0: URLShortenerService.proto.GetGetAllPairsResponse.pairs:type_name -> URLShortenerService.proto.Pair
	7,  // 1: URLShortenerService.proto.PostBatchRequest.req_pairs:type_name -> URLShortenerService.proto.PostBatchRequestPair
	8,  // 2: URLShortenerService.proto.PostBatchResponse.resp_pairs:type_name -> URLShortenerService.proto.PostBatchResponsePair
	0,  // 3: URLShortenerService.proto.Shortener.GetOrigURL:input_type -> URLShortenerService.proto.GetOrigURLRequest
	3,  // 4: URLShortenerService.proto.Shortener.GetAllPairs:input_type -> URLShortenerService.proto.GetGetAllPairsRequest
	5,  // 5: URLShortenerService.proto.Shortener.PostProvideShortURL:input_type -> URLShortenerService.proto.PostProvideShortURLRequest
	9,  // 6: URLShortenerService.proto.Shortener.PostBatch:input_type -> URLShortenerService.proto.PostBatchRequest
	11, // 7: URLShortenerService.proto.Shortener.DeleteBatch:input_type -> URLShortenerService.proto.DelBatchRequest
	13, // 8: URLShortenerService.proto.Shortener.GetStat:input_type -> URLShortenerService.proto.GetStatRequest
	1,  // 9: URLShortenerService.proto.Shortener.GetOrigURL:output_type -> URLShortenerService.proto.GetOrigURLResponse
	4,  // 10: URLShortenerService.proto.Shortener.GetAllPairs:output_type -> URLShortenerService.proto.GetGetAllPairsResponse
	6,  // 11: URLShortenerService.proto.Shortener.PostProvideShortURL:output_type -> URLShortenerService.proto.PostProvideShortURLResponse
	10, // 12: URLShortenerService.proto.Shortener.PostBatch:output_type -> URLShortenerService.proto.PostBatchResponse
	12, // 13: URLShortenerService.proto.Shortener.DeleteBatch:output_type -> URLShortenerService.proto.DelBatchResponse
	14, // 14: URLShortenerService.proto.Shortener.GetStat:output_type -> URLShortenerService.proto.GetStatResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_shortener_proto_init() }
func file_proto_shortener_proto_init() {
	if File_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrigURLRequest); i {
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
		file_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrigURLResponse); i {
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
		file_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGetAllPairsRequest); i {
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
		file_proto_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGetAllPairsResponse); i {
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
		file_proto_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostProvideShortURLRequest); i {
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
		file_proto_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostProvideShortURLResponse); i {
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
		file_proto_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBatchRequestPair); i {
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
		file_proto_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBatchResponsePair); i {
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
		file_proto_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBatchRequest); i {
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
		file_proto_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBatchResponse); i {
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
		file_proto_shortener_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelBatchRequest); i {
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
		file_proto_shortener_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelBatchResponse); i {
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
		file_proto_shortener_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatRequest); i {
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
		file_proto_shortener_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatResponse); i {
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
			RawDescriptor: file_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shortener_proto_goTypes,
		DependencyIndexes: file_proto_shortener_proto_depIdxs,
		MessageInfos:      file_proto_shortener_proto_msgTypes,
	}.Build()
	File_proto_shortener_proto = out.File
	file_proto_shortener_proto_rawDesc = nil
	file_proto_shortener_proto_goTypes = nil
	file_proto_shortener_proto_depIdxs = nil
}
