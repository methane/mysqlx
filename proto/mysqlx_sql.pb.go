// Code generated by protoc-gen-go.
// source: mysqlx_sql.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// execute a statement in the given namespace
//
// .. uml::
//
//   client -> server: StmtExecute
//   ... zero or more Resultsets ...
//   server --> client: StmtExecuteOk
//
// Notices:
//   This message may generate a notice containing WARNINGs generated by its execution.
//   This message may generate a notice containing INFO messages generated by its execution.
//
// :param namespace: namespace of the statement to be executed
// :param stmt: statement that shall be executed.
// :param args: values for wildcard replacements
// :param compact_metadata: send only type information for :protobuf:msg:`Mysqlx.Resultset::ColumnMetadata`, skipping names and others
// :returns:
//    * zero or one :protobuf:msg:`Mysqlx.Resultset::` followed by :protobuf:msg:`Mysqlx.Sql::StmtExecuteOk`
type StmtExecute struct {
	Namespace        *string `protobuf:"bytes,3,opt,name=namespace,def=sql" json:"namespace,omitempty"`
	Stmt             []byte  `protobuf:"bytes,1,req,name=stmt" json:"stmt,omitempty"`
	Args             []*Any  `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	CompactMetadata  *bool   `protobuf:"varint,4,opt,name=compact_metadata,def=0" json:"compact_metadata,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StmtExecute) Reset()                    { *m = StmtExecute{} }
func (m *StmtExecute) String() string            { return proto1.CompactTextString(m) }
func (*StmtExecute) ProtoMessage()               {}
func (*StmtExecute) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

const Default_StmtExecute_Namespace string = "sql"
const Default_StmtExecute_CompactMetadata bool = false

func (m *StmtExecute) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return Default_StmtExecute_Namespace
}

func (m *StmtExecute) GetStmt() []byte {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *StmtExecute) GetArgs() []*Any {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *StmtExecute) GetCompactMetadata() bool {
	if m != nil && m.CompactMetadata != nil {
		return *m.CompactMetadata
	}
	return Default_StmtExecute_CompactMetadata
}

// statement executed successful
type StmtExecuteOk struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StmtExecuteOk) Reset()                    { *m = StmtExecuteOk{} }
func (m *StmtExecuteOk) String() string            { return proto1.CompactTextString(m) }
func (*StmtExecuteOk) ProtoMessage()               {}
func (*StmtExecuteOk) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func init() {
	proto1.RegisterType((*StmtExecute)(nil), "Mysqlx.Sql.StmtExecute")
	proto1.RegisterType((*StmtExecuteOk)(nil), "Mysqlx.Sql.StmtExecuteOk")
}

func init() { proto1.RegisterFile("mysqlx_sql.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x40, 0xd9, 0xee, 0x16, 0xec, 0xb4, 0x62, 0x09, 0x58, 0x82, 0x08, 0x86, 0x8a, 0x90, 0x53,
	0x0e, 0x1e, 0x7b, 0x53, 0xf4, 0x28, 0x1e, 0x7a, 0xf3, 0x52, 0xc6, 0x98, 0x0a, 0x9a, 0x34, 0xc9,
	0xce, 0x2c, 0x6c, 0xfe, 0x5e, 0xdc, 0x5d, 0xc4, 0xdb, 0x63, 0x66, 0x98, 0xf7, 0x60, 0x1d, 0x0a,
	0x65, 0xdf, 0x1f, 0x28, 0x7b, 0x93, 0xda, 0xc8, 0x51, 0xc0, 0xcb, 0x30, 0x31, 0xfb, 0xec, 0xaf,
	0x36, 0xd3, 0xf6, 0x03, 0x19, 0xb9, 0x24, 0x47, 0xe3, 0xcd, 0xb6, 0xc0, 0x72, 0xcf, 0x81, 0x9f,
	0x7b, 0x67, 0x3b, 0x76, 0x62, 0x03, 0x8b, 0x13, 0x06, 0x47, 0x09, 0xad, 0x93, 0xb5, 0xaa, 0xf4,
	0x62, 0x57, 0x53, 0xf6, 0x62, 0x05, 0x0d, 0x71, 0x60, 0x59, 0xa9, 0x99, 0x5e, 0x89, 0x5b, 0x68,
	0xb0, 0xfd, 0x24, 0x39, 0x53, 0xb5, 0x5e, 0xde, 0x5f, 0x9a, 0xc9, 0xf3, 0xf4, 0xf7, 0xfb, 0xe1,
	0x54, 0xc4, 0x0d, 0xac, 0x6d, 0x0c, 0x09, 0x2d, 0x1f, 0x82, 0x63, 0xfc, 0x15, 0xcb, 0x46, 0x55,
	0xfa, 0x6c, 0x37, 0x3f, 0xa2, 0x27, 0xb7, 0xbd, 0x80, 0xf3, 0x7f, 0xea, 0xd7, 0xef, 0xc7, 0x3b,
	0xb8, 0xb6, 0x31, 0x98, 0xa1, 0xd4, 0xd8, 0xaf, 0x11, 0xfa, 0x31, 0xf4, 0xbd, 0x3b, 0xbe, 0xcd,
	0x07, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x63, 0x8b, 0x1b, 0xe9, 0x00, 0x00, 0x00,
}