// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: article/article_service.proto

package api

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

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_article_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_article_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_article_article_service_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ArticleShort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title            string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ShortDescription string `protobuf:"bytes,2,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	ImageUrl         string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *ArticleShort) Reset() {
	*x = ArticleShort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_article_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleShort) ProtoMessage() {}

func (x *ArticleShort) ProtoReflect() protoreflect.Message {
	mi := &file_article_article_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleShort.ProtoReflect.Descriptor instead.
func (*ArticleShort) Descriptor() ([]byte, []int) {
	return file_article_article_service_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleShort) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleShort) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *ArticleShort) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*ArticleShort `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_article_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_article_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_article_article_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetListResponse) GetArticles() []*ArticleShort {
	if x != nil {
		return x.Articles
	}
	return nil
}

var File_article_article_service_proto protoreflect.FileDescriptor

var file_article_article_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6d, 0x79, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x42, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50,
	0x61, 0x67, 0x65, 0x22, 0x6e, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x79, 0x5f, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x32, 0x5a, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x6d, 0x79, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x79, 0x5f, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05,
	0x2f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_article_article_service_proto_rawDescOnce sync.Once
	file_article_article_service_proto_rawDescData = file_article_article_service_proto_rawDesc
)

func file_article_article_service_proto_rawDescGZIP() []byte {
	file_article_article_service_proto_rawDescOnce.Do(func() {
		file_article_article_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_article_service_proto_rawDescData)
	})
	return file_article_article_service_proto_rawDescData
}

var file_article_article_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_article_article_service_proto_goTypes = []interface{}{
	(*PaginationRequest)(nil), // 0: my_epoch.proto.PaginationRequest
	(*ArticleShort)(nil),      // 1: my_epoch.proto.ArticleShort
	(*GetListResponse)(nil),   // 2: my_epoch.proto.GetListResponse
}
var file_article_article_service_proto_depIdxs = []int32{
	1, // 0: my_epoch.proto.GetListResponse.articles:type_name -> my_epoch.proto.ArticleShort
	0, // 1: my_epoch.proto.Article.GetList:input_type -> my_epoch.proto.PaginationRequest
	2, // 2: my_epoch.proto.Article.GetList:output_type -> my_epoch.proto.GetListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_article_article_service_proto_init() }
func file_article_article_service_proto_init() {
	if File_article_article_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_article_article_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationRequest); i {
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
		file_article_article_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleShort); i {
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
		file_article_article_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
			RawDescriptor: file_article_article_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_article_service_proto_goTypes,
		DependencyIndexes: file_article_article_service_proto_depIdxs,
		MessageInfos:      file_article_article_service_proto_msgTypes,
	}.Build()
	File_article_article_service_proto = out.File
	file_article_article_service_proto_rawDesc = nil
	file_article_article_service_proto_goTypes = nil
	file_article_article_service_proto_depIdxs = nil
}
