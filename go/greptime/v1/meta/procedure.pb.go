// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/meta/procedure.proto

package meta

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

type ProcedureStatus int32

const (
	ProcedureStatus_Executing ProcedureStatus = 0
	ProcedureStatus_Suspended ProcedureStatus = 1
	ProcedureStatus_Done      ProcedureStatus = 2
)

// Enum value maps for ProcedureStatus.
var (
	ProcedureStatus_name = map[int32]string{
		0: "Executing",
		1: "Suspended",
		2: "Done",
	}
	ProcedureStatus_value = map[string]int32{
		"Executing": 0,
		"Suspended": 1,
		"Done":      2,
	}
)

func (x ProcedureStatus) Enum() *ProcedureStatus {
	p := new(ProcedureStatus)
	*p = x
	return p
}

func (x ProcedureStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcedureStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_meta_procedure_proto_enumTypes[0].Descriptor()
}

func (ProcedureStatus) Type() protoreflect.EnumType {
	return &file_greptime_v1_meta_procedure_proto_enumTypes[0]
}

func (x ProcedureStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcedureStatus.Descriptor instead.
func (ProcedureStatus) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{0}
}

type QueryProcedureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid *ProcedureId `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *QueryProcedureRequest) Reset() {
	*x = QueryProcedureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_procedure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProcedureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProcedureRequest) ProtoMessage() {}

func (x *QueryProcedureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_procedure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProcedureRequest.ProtoReflect.Descriptor instead.
func (*QueryProcedureRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{0}
}

func (x *QueryProcedureRequest) GetPid() *ProcedureId {
	if x != nil {
		return x.Pid
	}
	return nil
}

type ProcedureStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ProcedureStatus `protobuf:"varint,1,opt,name=status,proto3,enum=greptime.v1.meta.ProcedureStatus" json:"status,omitempty"`
}

func (x *ProcedureStateResponse) Reset() {
	*x = ProcedureStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_procedure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureStateResponse) ProtoMessage() {}

func (x *ProcedureStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_procedure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureStateResponse.ProtoReflect.Descriptor instead.
func (*ProcedureStateResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{1}
}

func (x *ProcedureStateResponse) GetStatus() ProcedureStatus {
	if x != nil {
		return x.Status
	}
	return ProcedureStatus_Executing
}

var File_greptime_v1_meta_procedure_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_procedure_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1d, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48,
	0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x49, 0x64, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x39, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0d, 0x0a, 0x09, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x32, 0x8f, 0x02, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x12, 0x5a, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x27, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x03, 0x64, 0x64, 0x6c, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x64, 0x6c,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x44,
	0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a,
	0x0a, 0x07, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x65, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_procedure_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_procedure_proto_rawDescData = file_greptime_v1_meta_procedure_proto_rawDesc
)

func file_greptime_v1_meta_procedure_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_procedure_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_procedure_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_procedure_proto_rawDescData)
	})
	return file_greptime_v1_meta_procedure_proto_rawDescData
}

var file_greptime_v1_meta_procedure_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_greptime_v1_meta_procedure_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greptime_v1_meta_procedure_proto_goTypes = []interface{}{
	(ProcedureStatus)(0),           // 0: greptime.v1.meta.ProcedureStatus
	(*QueryProcedureRequest)(nil),  // 1: greptime.v1.meta.QueryProcedureRequest
	(*ProcedureStateResponse)(nil), // 2: greptime.v1.meta.ProcedureStateResponse
	(*ProcedureId)(nil),            // 3: greptime.v1.meta.ProcedureId
	(*DdlTaskRequest)(nil),         // 4: greptime.v1.meta.DdlTaskRequest
	(*MigrateRegionRequest)(nil),   // 5: greptime.v1.meta.MigrateRegionRequest
	(*DdlTaskResponse)(nil),        // 6: greptime.v1.meta.DdlTaskResponse
	(*MigrateRegionResponse)(nil),  // 7: greptime.v1.meta.MigrateRegionResponse
}
var file_greptime_v1_meta_procedure_proto_depIdxs = []int32{
	3, // 0: greptime.v1.meta.QueryProcedureRequest.pid:type_name -> greptime.v1.meta.ProcedureId
	0, // 1: greptime.v1.meta.ProcedureStateResponse.status:type_name -> greptime.v1.meta.ProcedureStatus
	1, // 2: greptime.v1.meta.Procedure.query:input_type -> greptime.v1.meta.QueryProcedureRequest
	4, // 3: greptime.v1.meta.Procedure.ddl:input_type -> greptime.v1.meta.DdlTaskRequest
	5, // 4: greptime.v1.meta.Procedure.migrate:input_type -> greptime.v1.meta.MigrateRegionRequest
	2, // 5: greptime.v1.meta.Procedure.query:output_type -> greptime.v1.meta.ProcedureStateResponse
	6, // 6: greptime.v1.meta.Procedure.ddl:output_type -> greptime.v1.meta.DdlTaskResponse
	7, // 7: greptime.v1.meta.Procedure.migrate:output_type -> greptime.v1.meta.MigrateRegionResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_procedure_proto_init() }
func file_greptime_v1_meta_procedure_proto_init() {
	if File_greptime_v1_meta_procedure_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	file_greptime_v1_meta_ddl_proto_init()
	file_greptime_v1_meta_region_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_procedure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProcedureRequest); i {
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
		file_greptime_v1_meta_procedure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureStateResponse); i {
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
			RawDescriptor: file_greptime_v1_meta_procedure_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_procedure_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_procedure_proto_depIdxs,
		EnumInfos:         file_greptime_v1_meta_procedure_proto_enumTypes,
		MessageInfos:      file_greptime_v1_meta_procedure_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_procedure_proto = out.File
	file_greptime_v1_meta_procedure_proto_rawDesc = nil
	file_greptime_v1_meta_procedure_proto_goTypes = nil
	file_greptime_v1_meta_procedure_proto_depIdxs = nil
}
