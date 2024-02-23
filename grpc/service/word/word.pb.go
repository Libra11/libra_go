// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: word.proto

package word

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

type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Word       string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	Definition string `protobuf:"bytes,3,opt,name=definition,proto3" json:"definition,omitempty"`
	Example    string `protobuf:"bytes,4,opt,name=example,proto3" json:"example,omitempty"`
	Phrase     string `protobuf:"bytes,5,opt,name=phrase,proto3" json:"phrase,omitempty"`
	CreateAt   int64  `protobuf:"varint,6,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt   int64  `protobuf:"varint,7,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	Phonetic   string `protobuf:"bytes,8,opt,name=phonetic,proto3" json:"phonetic,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{0}
}

func (x *Word) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Word) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *Word) GetDefinition() string {
	if x != nil {
		return x.Definition
	}
	return ""
}

func (x *Word) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *Word) GetPhrase() string {
	if x != nil {
		return x.Phrase
	}
	return ""
}

func (x *Word) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Word) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Word) GetPhonetic() string {
	if x != nil {
		return x.Phonetic
	}
	return ""
}

type AddWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word *Word `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *AddWordRequest) Reset() {
	*x = AddWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWordRequest) ProtoMessage() {}

func (x *AddWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWordRequest.ProtoReflect.Descriptor instead.
func (*AddWordRequest) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{1}
}

func (x *AddWordRequest) GetWord() *Word {
	if x != nil {
		return x.Word
	}
	return nil
}

type AddWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddWordResponse) Reset() {
	*x = AddWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWordResponse) ProtoMessage() {}

func (x *AddWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWordResponse.ProtoReflect.Descriptor instead.
func (*AddWordResponse) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{2}
}

func (x *AddWordResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Word     string `protobuf:"bytes,3,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *GetWordsRequest) Reset() {
	*x = GetWordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordsRequest) ProtoMessage() {}

func (x *GetWordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordsRequest.ProtoReflect.Descriptor instead.
func (*GetWordsRequest) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{3}
}

func (x *GetWordsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetWordsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetWordsRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type GetWordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words []*Word `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetWordsResponse) Reset() {
	*x = GetWordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordsResponse) ProtoMessage() {}

func (x *GetWordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordsResponse.ProtoReflect.Descriptor instead.
func (*GetWordsResponse) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{4}
}

func (x *GetWordsResponse) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

func (x *GetWordsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWordRequest) Reset() {
	*x = DeleteWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWordRequest) ProtoMessage() {}

func (x *DeleteWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWordRequest.ProtoReflect.Descriptor instead.
func (*DeleteWordRequest) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteWordRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWordResponse) Reset() {
	*x = DeleteWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWordResponse) ProtoMessage() {}

func (x *DeleteWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWordResponse.ProtoReflect.Descriptor instead.
func (*DeleteWordResponse) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{6}
}

type GetWordByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWordByIdRequest) Reset() {
	*x = GetWordByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordByIdRequest) ProtoMessage() {}

func (x *GetWordByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordByIdRequest.ProtoReflect.Descriptor instead.
func (*GetWordByIdRequest) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{7}
}

func (x *GetWordByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWordByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word *Word `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *GetWordByIdResponse) Reset() {
	*x = GetWordByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_word_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordByIdResponse) ProtoMessage() {}

func (x *GetWordByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_word_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordByIdResponse.ProtoReflect.Descriptor instead.
func (*GetWordByIdResponse) Descriptor() ([]byte, []int) {
	return file_word_proto_rawDescGZIP(), []int{8}
}

func (x *GetWordByIdResponse) GetWord() *Word {
	if x != nil {
		return x.Word
	}
	return nil
}

var File_word_proto protoreflect.FileDescriptor

var file_word_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0xd0, 0x01, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x72,
	0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x72, 0x61, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x74, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x74, 0x69, 0x63, 0x22, 0x30, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x57, 0x6f, 0x72,
	0x64, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x57, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x4a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x57, 0x6f, 0x72, 0x64,
	0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52,
	0x04, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x85, 0x02, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64,
	0x12, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x64,
	0x64, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x15, 0x2e, 0x77, 0x6f, 0x72, 0x64,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x3b, 0x77, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_word_proto_rawDescOnce sync.Once
	file_word_proto_rawDescData = file_word_proto_rawDesc
)

func file_word_proto_rawDescGZIP() []byte {
	file_word_proto_rawDescOnce.Do(func() {
		file_word_proto_rawDescData = protoimpl.X.CompressGZIP(file_word_proto_rawDescData)
	})
	return file_word_proto_rawDescData
}

var file_word_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_word_proto_goTypes = []interface{}{
	(*Word)(nil),                // 0: word.Word
	(*AddWordRequest)(nil),      // 1: word.AddWordRequest
	(*AddWordResponse)(nil),     // 2: word.AddWordResponse
	(*GetWordsRequest)(nil),     // 3: word.GetWordsRequest
	(*GetWordsResponse)(nil),    // 4: word.GetWordsResponse
	(*DeleteWordRequest)(nil),   // 5: word.DeleteWordRequest
	(*DeleteWordResponse)(nil),  // 6: word.DeleteWordResponse
	(*GetWordByIdRequest)(nil),  // 7: word.GetWordByIdRequest
	(*GetWordByIdResponse)(nil), // 8: word.GetWordByIdResponse
}
var file_word_proto_depIdxs = []int32{
	0, // 0: word.AddWordRequest.word:type_name -> word.Word
	0, // 1: word.GetWordsResponse.words:type_name -> word.Word
	0, // 2: word.GetWordByIdResponse.word:type_name -> word.Word
	1, // 3: word.WordService.AddWord:input_type -> word.AddWordRequest
	3, // 4: word.WordService.GetWords:input_type -> word.GetWordsRequest
	5, // 5: word.WordService.DeleteWord:input_type -> word.DeleteWordRequest
	7, // 6: word.WordService.GetWordById:input_type -> word.GetWordByIdRequest
	2, // 7: word.WordService.AddWord:output_type -> word.AddWordResponse
	4, // 8: word.WordService.GetWords:output_type -> word.GetWordsResponse
	6, // 9: word.WordService.DeleteWord:output_type -> word.DeleteWordResponse
	8, // 10: word.WordService.GetWordById:output_type -> word.GetWordByIdResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_word_proto_init() }
func file_word_proto_init() {
	if File_word_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_word_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
		file_word_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWordRequest); i {
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
		file_word_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWordResponse); i {
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
		file_word_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordsRequest); i {
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
		file_word_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordsResponse); i {
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
		file_word_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWordRequest); i {
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
		file_word_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWordResponse); i {
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
		file_word_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordByIdRequest); i {
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
		file_word_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordByIdResponse); i {
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
			RawDescriptor: file_word_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_word_proto_goTypes,
		DependencyIndexes: file_word_proto_depIdxs,
		MessageInfos:      file_word_proto_msgTypes,
	}.Build()
	File_word_proto = out.File
	file_word_proto_rawDesc = nil
	file_word_proto_goTypes = nil
	file_word_proto_depIdxs = nil
}
