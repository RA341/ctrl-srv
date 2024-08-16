// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: protos/filesystem.proto

package fs

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

type InputFolders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcPath  *string `protobuf:"bytes,1,req,name=srcPath" json:"srcPath,omitempty"`
	DestPath *string `protobuf:"bytes,2,req,name=destPath" json:"destPath,omitempty"`
}

func (x *InputFolders) Reset() {
	*x = InputFolders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputFolders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputFolders) ProtoMessage() {}

func (x *InputFolders) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputFolders.ProtoReflect.Descriptor instead.
func (*InputFolders) Descriptor() ([]byte, []int) {
	return file_protos_filesystem_proto_rawDescGZIP(), []int{0}
}

func (x *InputFolders) GetSrcPath() string {
	if x != nil && x.SrcPath != nil {
		return *x.SrcPath
	}
	return ""
}

func (x *InputFolders) GetDestPath() string {
	if x != nil && x.DestPath != nil {
		return *x.DestPath
	}
	return ""
}

type LinkResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (x *LinkResult) Reset() {
	*x = LinkResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filesystem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkResult) ProtoMessage() {}

func (x *LinkResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filesystem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkResult.ProtoReflect.Descriptor instead.
func (*LinkResult) Descriptor() ([]byte, []int) {
	return file_protos_filesystem_proto_rawDescGZIP(), []int{1}
}

func (x *LinkResult) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filesystem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filesystem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_protos_filesystem_proto_rawDescGZIP(), []int{2}
}

func (x *Path) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullPath *string   `protobuf:"bytes,1,req,name=fullPath" json:"fullPath,omitempty"`
	Files    []*File   `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
	Folders  []*Folder `protobuf:"bytes,3,rep,name=folders" json:"folders,omitempty"`
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filesystem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filesystem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_protos_filesystem_proto_rawDescGZIP(), []int{3}
}

func (x *Folder) GetFullPath() string {
	if x != nil && x.FullPath != nil {
		return *x.FullPath
	}
	return ""
}

func (x *Folder) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Folder) GetFolders() []*Folder {
	if x != nil {
		return x.Folders
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filesystem_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filesystem_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_protos_filesystem_proto_rawDescGZIP(), []int{4}
}

func (x *File) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_protos_filesystem_proto protoreflect.FileDescriptor

var file_protos_filesystem_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x66, 0x73, 0x22, 0x44, 0x0a,
	0x0c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x72, 0x63, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x72, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x22, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x6a, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x66, 0x73, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x73, 0x2e,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x1a, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x8f, 0x01, 0x0a, 0x0a,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x08, 0x2e, 0x66, 0x73, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x1a, 0x0a, 0x2e, 0x66, 0x73, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x10, 0x2e,
	0x66, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x1a,
	0x0e, 0x2e, 0x66, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x2a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x08, 0x2e, 0x66, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x0e, 0x2e, 0x66, 0x73,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x73,
}

var (
	file_protos_filesystem_proto_rawDescOnce sync.Once
	file_protos_filesystem_proto_rawDescData = file_protos_filesystem_proto_rawDesc
)

func file_protos_filesystem_proto_rawDescGZIP() []byte {
	file_protos_filesystem_proto_rawDescOnce.Do(func() {
		file_protos_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_filesystem_proto_rawDescData)
	})
	return file_protos_filesystem_proto_rawDescData
}

var file_protos_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_filesystem_proto_goTypes = []any{
	(*InputFolders)(nil), // 0: fs.InputFolders
	(*LinkResult)(nil),   // 1: fs.LinkResult
	(*Path)(nil),         // 2: fs.Path
	(*Folder)(nil),       // 3: fs.Folder
	(*File)(nil),         // 4: fs.File
}
var file_protos_filesystem_proto_depIdxs = []int32{
	4, // 0: fs.Folder.files:type_name -> fs.File
	3, // 1: fs.Folder.folders:type_name -> fs.Folder
	2, // 2: fs.Filesystem.ListFiles:input_type -> fs.Path
	0, // 3: fs.Filesystem.LinkFolder:input_type -> fs.InputFolders
	2, // 4: fs.Filesystem.CreateFolder:input_type -> fs.Path
	3, // 5: fs.Filesystem.ListFiles:output_type -> fs.Folder
	1, // 6: fs.Filesystem.LinkFolder:output_type -> fs.LinkResult
	1, // 7: fs.Filesystem.CreateFolder:output_type -> fs.LinkResult
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_filesystem_proto_init() }
func file_protos_filesystem_proto_init() {
	if File_protos_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_filesystem_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InputFolders); i {
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
		file_protos_filesystem_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LinkResult); i {
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
		file_protos_filesystem_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Path); i {
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
		file_protos_filesystem_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Folder); i {
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
		file_protos_filesystem_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*File); i {
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
			RawDescriptor: file_protos_filesystem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_filesystem_proto_goTypes,
		DependencyIndexes: file_protos_filesystem_proto_depIdxs,
		MessageInfos:      file_protos_filesystem_proto_msgTypes,
	}.Build()
	File_protos_filesystem_proto = out.File
	file_protos_filesystem_proto_rawDesc = nil
	file_protos_filesystem_proto_goTypes = nil
	file_protos_filesystem_proto_depIdxs = nil
}
