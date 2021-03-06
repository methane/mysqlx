// Code generated by protoc-gen-go.
// source: mysqlx_expr.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Expr_Type int32

const (
	Expr_IDENT       Expr_Type = 1
	Expr_LITERAL     Expr_Type = 2
	Expr_VARIABLE    Expr_Type = 3
	Expr_FUNC_CALL   Expr_Type = 4
	Expr_OPERATOR    Expr_Type = 5
	Expr_PLACEHOLDER Expr_Type = 6
	Expr_OBJECT      Expr_Type = 7
	Expr_ARRAY       Expr_Type = 8
)

var Expr_Type_name = map[int32]string{
	1: "IDENT",
	2: "LITERAL",
	3: "VARIABLE",
	4: "FUNC_CALL",
	5: "OPERATOR",
	6: "PLACEHOLDER",
	7: "OBJECT",
	8: "ARRAY",
}
var Expr_Type_value = map[string]int32{
	"IDENT":       1,
	"LITERAL":     2,
	"VARIABLE":    3,
	"FUNC_CALL":   4,
	"OPERATOR":    5,
	"PLACEHOLDER": 6,
	"OBJECT":      7,
	"ARRAY":       8,
}

func (x Expr_Type) Enum() *Expr_Type {
	p := new(Expr_Type)
	*p = x
	return p
}
func (x Expr_Type) String() string {
	return proto1.EnumName(Expr_Type_name, int32(x))
}
func (x *Expr_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Expr_Type_value, data, "Expr_Type")
	if err != nil {
		return err
	}
	*x = Expr_Type(value)
	return nil
}
func (Expr_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0, 0} }

type DocumentPathItem_Type int32

const (
	DocumentPathItem_MEMBER               DocumentPathItem_Type = 1
	DocumentPathItem_MEMBER_ASTERISK      DocumentPathItem_Type = 2
	DocumentPathItem_ARRAY_INDEX          DocumentPathItem_Type = 3
	DocumentPathItem_ARRAY_INDEX_ASTERISK DocumentPathItem_Type = 4
	DocumentPathItem_DOUBLE_ASTERISK      DocumentPathItem_Type = 5
)

var DocumentPathItem_Type_name = map[int32]string{
	1: "MEMBER",
	2: "MEMBER_ASTERISK",
	3: "ARRAY_INDEX",
	4: "ARRAY_INDEX_ASTERISK",
	5: "DOUBLE_ASTERISK",
}
var DocumentPathItem_Type_value = map[string]int32{
	"MEMBER":               1,
	"MEMBER_ASTERISK":      2,
	"ARRAY_INDEX":          3,
	"ARRAY_INDEX_ASTERISK": 4,
	"DOUBLE_ASTERISK":      5,
}

func (x DocumentPathItem_Type) Enum() *DocumentPathItem_Type {
	p := new(DocumentPathItem_Type)
	*p = x
	return p
}
func (x DocumentPathItem_Type) String() string {
	return proto1.EnumName(DocumentPathItem_Type_name, int32(x))
}
func (x *DocumentPathItem_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(DocumentPathItem_Type_value, data, "DocumentPathItem_Type")
	if err != nil {
		return err
	}
	*x = DocumentPathItem_Type(value)
	return nil
}
func (DocumentPathItem_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2, 0} }

// Expressions
//
// the "root" of the expression tree
//
// .. productionlist::
//   expr: `operator` |
//       : `identifier` |
//       : `function_call` |
//       : variable |
//       : `literal` |
//       : placeholder
//
// If expression type is PLACEHOLDER then it refers to the value of a parameter
// specified when executing a statement (see `args` field of `StmtExecute` command).
// Field `position` (which must be present for such an expression) gives 0-based
// position of the parameter in the parameter list.
//
type Expr struct {
	Type             *Expr_Type        `protobuf:"varint,1,req,name=type,enum=Mysqlx.Expr.Expr_Type" json:"type,omitempty"`
	Identifier       *ColumnIdentifier `protobuf:"bytes,2,opt,name=identifier" json:"identifier,omitempty"`
	Variable         *string           `protobuf:"bytes,3,opt,name=variable" json:"variable,omitempty"`
	Literal          *Scalar           `protobuf:"bytes,4,opt,name=literal" json:"literal,omitempty"`
	FunctionCall     *FunctionCall     `protobuf:"bytes,5,opt,name=function_call" json:"function_call,omitempty"`
	Operator         *Operator         `protobuf:"bytes,6,opt,name=operator" json:"operator,omitempty"`
	Position         *uint32           `protobuf:"varint,7,opt,name=position" json:"position,omitempty"`
	Object           *Object           `protobuf:"bytes,8,opt,name=object" json:"object,omitempty"`
	Array            *Array            `protobuf:"bytes,9,opt,name=array" json:"array,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Expr) Reset()                    { *m = Expr{} }
func (m *Expr) String() string            { return proto1.CompactTextString(m) }
func (*Expr) ProtoMessage()               {}
func (*Expr) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Expr) GetType() Expr_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Expr_IDENT
}

func (m *Expr) GetIdentifier() *ColumnIdentifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Expr) GetVariable() string {
	if m != nil && m.Variable != nil {
		return *m.Variable
	}
	return ""
}

func (m *Expr) GetLiteral() *Scalar {
	if m != nil {
		return m.Literal
	}
	return nil
}

func (m *Expr) GetFunctionCall() *FunctionCall {
	if m != nil {
		return m.FunctionCall
	}
	return nil
}

func (m *Expr) GetOperator() *Operator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *Expr) GetPosition() uint32 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

func (m *Expr) GetObject() *Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *Expr) GetArray() *Array {
	if m != nil {
		return m.Array
	}
	return nil
}

// identifier: name, schame.name
//
// .. productionlist::
//   identifier: string "." string |
//             : string
type Identifier struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	SchemaName       *string `protobuf:"bytes,2,opt,name=schema_name" json:"schema_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Identifier) Reset()                    { *m = Identifier{} }
func (m *Identifier) String() string            { return proto1.CompactTextString(m) }
func (*Identifier) ProtoMessage()               {}
func (*Identifier) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Identifier) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Identifier) GetSchemaName() string {
	if m != nil && m.SchemaName != nil {
		return *m.SchemaName
	}
	return ""
}

// DocumentPathItem
//
// .. productionlist::
//    document_path: path_item | path_item document_path
//    path_item    : member | array_index | "**"
//    member       : "." string | "." "*"
//    array_index  : "[" number "]" | "[" "*" "]"
//
type DocumentPathItem struct {
	Type             *DocumentPathItem_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Expr.DocumentPathItem_Type" json:"type,omitempty"`
	Value            *string                `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Index            *uint32                `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *DocumentPathItem) Reset()                    { *m = DocumentPathItem{} }
func (m *DocumentPathItem) String() string            { return proto1.CompactTextString(m) }
func (*DocumentPathItem) ProtoMessage()               {}
func (*DocumentPathItem) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *DocumentPathItem) GetType() DocumentPathItem_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return DocumentPathItem_MEMBER
}

func (m *DocumentPathItem) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *DocumentPathItem) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

// col_identifier (table): col@doc_path, tbl.col@doc_path col, tbl.col, schema.tbl.col
// col_identifier (document): doc_path
//
// .. productionlist::
//   col_identifier: string "." string "." string |
//             : string "." string |
//             : string |
//             : string "." string "." string "@" document_path |
//             : string "." string "@" document_path |
//             : string "@" document_path |
//             : document_path
//    document_path: member | arrayLocation | doubleAsterisk
//    member = "." string | "." "*"
//    arrayLocation = "[" index "]" | "[" "*" "]"
//    doubleAsterisk = "**"
//
type ColumnIdentifier struct {
	DocumentPath     []*DocumentPathItem `protobuf:"bytes,1,rep,name=document_path" json:"document_path,omitempty"`
	Name             *string             `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	TableName        *string             `protobuf:"bytes,3,opt,name=table_name" json:"table_name,omitempty"`
	SchemaName       *string             `protobuf:"bytes,4,opt,name=schema_name" json:"schema_name,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ColumnIdentifier) Reset()                    { *m = ColumnIdentifier{} }
func (m *ColumnIdentifier) String() string            { return proto1.CompactTextString(m) }
func (*ColumnIdentifier) ProtoMessage()               {}
func (*ColumnIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *ColumnIdentifier) GetDocumentPath() []*DocumentPathItem {
	if m != nil {
		return m.DocumentPath
	}
	return nil
}

func (m *ColumnIdentifier) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ColumnIdentifier) GetTableName() string {
	if m != nil && m.TableName != nil {
		return *m.TableName
	}
	return ""
}

func (m *ColumnIdentifier) GetSchemaName() string {
	if m != nil && m.SchemaName != nil {
		return *m.SchemaName
	}
	return ""
}

// function call: ``func(a, b, "1", 3)``
//
// .. productionlist::
//   function_call: `identifier` "(" [ `expr` ["," `expr` ]* ] ")"
type FunctionCall struct {
	Name             *Identifier `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Param            []*Expr     `protobuf:"bytes,2,rep,name=param" json:"param,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *FunctionCall) Reset()                    { *m = FunctionCall{} }
func (m *FunctionCall) String() string            { return proto1.CompactTextString(m) }
func (*FunctionCall) ProtoMessage()               {}
func (*FunctionCall) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *FunctionCall) GetName() *Identifier {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *FunctionCall) GetParam() []*Expr {
	if m != nil {
		return m.Param
	}
	return nil
}

// operator: ``<<(a, b)``
//
// .. note::
//
//   Non-authoritative list of operators implemented (case sensitive):
//
//   Nullary
//     * ``*``
//     * ``default``
//
//   Unary
//     * ``!``
//     * ``sign_plus``
//     * ``sign_minus``
//     * ``~``
//
//   Binary
//     * ``&&``
//     * ``||``
//     * ``xor``
//     * ``==``
//     * ``!=``
//     * ``>``
//     * ``>=``
//     * ``<``
//     * ``<=``
//     * ``&``
//     * ``|``
//     * ``^``
//     * ``<<``
//     * ``>>``
//     * ``+``
//     * ``-``
//     * ``*``
//     * ``/``
//     * ``div``
//     * ``%``
//     * ``is``
//     * ``is_not``
//     * ``regexp``
//     * ``not_regexp``
//     * ``like``
//     * ``not_like``
//     * ``cast``
//
//   Using special representation, with more than 2 params
//     * ``in`` (param[0] IN (param[1], param[2], ...))
//     * ``not_in`` (param[0] NOT IN (param[1], param[2], ...))
//
//   Ternary
//     * ``between``
//     * ``between_not``
//     * ``date_add``
//     * ``date_sub``
//
//   Units for date_add/date_sub
//     * ``MICROSECOND``
//     * ``SECOND``
//     * ``MINUTE``
//     * ``HOUR``
//     * ``DAY``
//     * ``WEEK``
//     * ``MONTH``
//     * ``QUARTER``
//     * ``YEAR``
//     * ``SECOND_MICROSECOND``
//     * ``MINUTE_MICROSECOND``
//     * ``MINUTE_SECOND``
//     * ``HOUR_MICROSECOND``
//     * ``HOUR_SECOND``
//     * ``HOUR_MINUTE``
//     * ``DAY_MICROSECOND``
//     * ``DAY_SECOND``
//     * ``DAY_MINUTE``
//     * ``DAY_HOUR``
//
//   Types for cast
//     * ``BINARY[(N)]``
//     * ``CHAR[(N)]``
//     * ``DATE``
//     * ``DATETIME``
//     * ``DECIMAL[(M[,D])]``
//     * ``JSON``
//     * ``SIGNED [INTEGER]``
//     * ``TIME``
//     * ``UNSIGNED [INTEGER]``
//
// .. productionlist::
//   operator: `name` "(" [ `expr` ["," `expr` ]* ] ")"
type Operator struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Param            []*Expr `protobuf:"bytes,2,rep,name=param" json:"param,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Operator) Reset()                    { *m = Operator{} }
func (m *Operator) String() string            { return proto1.CompactTextString(m) }
func (*Operator) ProtoMessage()               {}
func (*Operator) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *Operator) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Operator) GetParam() []*Expr {
	if m != nil {
		return m.Param
	}
	return nil
}

// an object (with expression values)
type Object struct {
	Fld              []*Object_ObjectField `protobuf:"bytes,1,rep,name=fld" json:"fld,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto1.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *Object) GetFld() []*Object_ObjectField {
	if m != nil {
		return m.Fld
	}
	return nil
}

type Object_ObjectField struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *Expr   `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Object_ObjectField) Reset()                    { *m = Object_ObjectField{} }
func (m *Object_ObjectField) String() string            { return proto1.CompactTextString(m) }
func (*Object_ObjectField) ProtoMessage()               {}
func (*Object_ObjectField) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6, 0} }

func (m *Object_ObjectField) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Object_ObjectField) GetValue() *Expr {
	if m != nil {
		return m.Value
	}
	return nil
}

// a Array of expressions
type Array struct {
	Value            []*Expr `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Array) Reset()                    { *m = Array{} }
func (m *Array) String() string            { return proto1.CompactTextString(m) }
func (*Array) ProtoMessage()               {}
func (*Array) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *Array) GetValue() []*Expr {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto1.RegisterType((*Expr)(nil), "Mysqlx.Expr.Expr")
	proto1.RegisterType((*Identifier)(nil), "Mysqlx.Expr.Identifier")
	proto1.RegisterType((*DocumentPathItem)(nil), "Mysqlx.Expr.DocumentPathItem")
	proto1.RegisterType((*ColumnIdentifier)(nil), "Mysqlx.Expr.ColumnIdentifier")
	proto1.RegisterType((*FunctionCall)(nil), "Mysqlx.Expr.FunctionCall")
	proto1.RegisterType((*Operator)(nil), "Mysqlx.Expr.Operator")
	proto1.RegisterType((*Object)(nil), "Mysqlx.Expr.Object")
	proto1.RegisterType((*Object_ObjectField)(nil), "Mysqlx.Expr.Object.ObjectField")
	proto1.RegisterType((*Array)(nil), "Mysqlx.Expr.Array")
	proto1.RegisterEnum("Mysqlx.Expr.Expr_Type", Expr_Type_name, Expr_Type_value)
	proto1.RegisterEnum("Mysqlx.Expr.DocumentPathItem_Type", DocumentPathItem_Type_name, DocumentPathItem_Type_value)
}

func init() { proto1.RegisterFile("mysqlx_expr.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x96, 0x63, 0x3b, 0x8f, 0x71, 0x43, 0xb7, 0x5b, 0x28, 0xa6, 0x02, 0x11, 0x0c, 0x15, 0xa9,
	0x84, 0x42, 0xa9, 0x38, 0x21, 0x2e, 0x4e, 0xe2, 0x0a, 0x83, 0xdb, 0x54, 0x6e, 0xca, 0xeb, 0x62,
	0x6d, 0x9d, 0x0d, 0x75, 0x59, 0x3f, 0xea, 0x6c, 0xaa, 0xe4, 0xc4, 0x9f, 0xe3, 0xc6, 0x9f, 0x42,
	0x5e, 0xbb, 0x4d, 0xdc, 0x14, 0x71, 0x49, 0xd6, 0x33, 0xdf, 0xbc, 0xbe, 0xf9, 0x76, 0x61, 0x23,
	0x9c, 0x4f, 0x2e, 0xd9, 0xcc, 0xa3, 0xb3, 0x24, 0xed, 0x24, 0x69, 0xcc, 0x63, 0xac, 0x1d, 0x0a,
	0x53, 0xc7, 0x9a, 0x25, 0xe9, 0xf6, 0x56, 0xe1, 0x1f, 0x11, 0x4e, 0xf8, 0x3c, 0xa1, 0x93, 0x1c,
	0x64, 0xfc, 0x96, 0x41, 0xc9, 0x00, 0xf8, 0x05, 0x28, 0x99, 0x5d, 0x97, 0x5a, 0x95, 0xf6, 0xbd,
	0xfd, 0xad, 0xce, 0x52, 0x70, 0xfe, 0x33, 0x9c, 0x27, 0x14, 0xbf, 0x01, 0x08, 0x46, 0x34, 0xe2,
	0xc1, 0x38, 0xa0, 0xa9, 0x5e, 0x69, 0x49, 0x6d, 0x6d, 0xff, 0x49, 0x09, 0xdb, 0x8b, 0xd9, 0x34,
	0x8c, 0xec, 0x1b, 0x10, 0x46, 0x50, 0xbf, 0x22, 0x69, 0x40, 0xce, 0x18, 0xd5, 0xe5, 0x96, 0xd4,
	0x6e, 0xe0, 0x5d, 0xa8, 0xb1, 0x80, 0xd3, 0x94, 0x30, 0x5d, 0x11, 0x19, 0xf4, 0xeb, 0x0c, 0xfd,
	0x9b, 0xee, 0x4e, 0x7c, 0xc2, 0x48, 0x8a, 0xf7, 0xa0, 0x39, 0x9e, 0x46, 0x3e, 0x0f, 0xe2, 0xc8,
	0xf3, 0x09, 0x63, 0xba, 0x2a, 0x02, 0x1e, 0x95, 0x4a, 0x1e, 0x14, 0x88, 0x1e, 0x61, 0x0c, 0xbf,
	0x84, 0x7a, 0x9c, 0xd0, 0x94, 0xf0, 0x38, 0xd5, 0xab, 0x02, 0xfc, 0xa0, 0x04, 0x1e, 0x14, 0xce,
	0xac, 0xaf, 0x24, 0x9e, 0x04, 0x59, 0xa0, 0x5e, 0x6b, 0x49, 0xed, 0x26, 0x7e, 0x0e, 0xd5, 0xf8,
	0xec, 0x82, 0xfa, 0x5c, 0xaf, 0x8b, 0xc0, 0xcd, 0x72, 0xa0, 0x70, 0xe1, 0x67, 0xa0, 0x92, 0x34,
	0x25, 0x73, 0xbd, 0x21, 0x30, 0xb8, 0x84, 0x31, 0x33, 0x8f, 0x71, 0x09, 0x8a, 0x20, 0xab, 0x01,
	0xaa, 0xdd, 0xb7, 0x8e, 0x86, 0x48, 0xc2, 0x1a, 0xd4, 0x1c, 0x7b, 0x68, 0xb9, 0xa6, 0x83, 0x2a,
	0x78, 0x0d, 0xea, 0x9f, 0x4d, 0xd7, 0x36, 0xbb, 0x8e, 0x85, 0x64, 0xdc, 0x84, 0xc6, 0xc1, 0xe9,
	0x51, 0xcf, 0xeb, 0x99, 0x8e, 0x83, 0x94, 0xcc, 0x39, 0x38, 0xb6, 0x5c, 0x73, 0x38, 0x70, 0x91,
	0x8a, 0xd7, 0x41, 0x3b, 0x76, 0xcc, 0x9e, 0xf5, 0x61, 0xe0, 0xf4, 0x2d, 0x17, 0x55, 0x31, 0x40,
	0x75, 0xd0, 0xfd, 0x68, 0xf5, 0x86, 0xa8, 0x96, 0xe5, 0x37, 0x5d, 0xd7, 0xfc, 0x86, 0xea, 0xc6,
	0x6b, 0x80, 0x25, 0xca, 0xd7, 0x40, 0x89, 0x48, 0x98, 0xef, 0xb2, 0x81, 0x37, 0x41, 0x9b, 0xf8,
	0xe7, 0x34, 0x24, 0x9e, 0x30, 0x66, 0x4b, 0x6b, 0x18, 0x7f, 0x24, 0x40, 0xfd, 0xd8, 0x9f, 0x86,
	0x34, 0xe2, 0xc7, 0x84, 0x9f, 0xdb, 0x9c, 0x86, 0x78, 0xaf, 0xa4, 0x01, 0xa3, 0x34, 0xda, 0x6d,
	0x70, 0xae, 0x87, 0x26, 0xa8, 0x57, 0x84, 0x4d, 0x8b, 0xac, 0xd9, 0x67, 0x10, 0x8d, 0xe8, 0x4c,
	0x2c, 0xba, 0x69, 0xfc, 0x28, 0x88, 0x00, 0xa8, 0x1e, 0x5a, 0x87, 0x5d, 0xcb, 0x45, 0x12, 0xde,
	0x84, 0xf5, 0xfc, 0xec, 0x99, 0x27, 0x43, 0xcb, 0xb5, 0x4f, 0x3e, 0xa1, 0x4a, 0x36, 0xa6, 0x98,
	0xc4, 0xb3, 0x8f, 0xfa, 0xd6, 0x57, 0x24, 0x63, 0x1d, 0xee, 0x2f, 0x19, 0x16, 0x50, 0x25, 0x8b,
	0xef, 0x0f, 0x4e, 0xbb, 0x8e, 0xb5, 0x30, 0xaa, 0xc6, 0x2f, 0x40, 0x2b, 0xba, 0x7b, 0x0b, 0xcd,
	0x51, 0xd1, 0xb3, 0x97, 0x10, 0x7e, 0xae, 0x4b, 0x2d, 0x79, 0x45, 0xad, 0x2b, 0x14, 0x5c, 0x53,
	0x97, 0xcf, 0x83, 0x01, 0x78, 0x26, 0xdc, 0x9c, 0xb9, 0x5c, 0xbd, 0xb7, 0xe8, 0x54, 0x04, 0x9d,
	0x5f, 0x60, 0xad, 0xa4, 0xc2, 0x9d, 0xa5, 0x0d, 0x68, 0xfb, 0x0f, 0x4b, 0x35, 0x97, 0x7a, 0x6c,
	0x81, 0x9a, 0x90, 0x94, 0x84, 0x7a, 0x45, 0xf4, 0xb6, 0xb1, 0x72, 0xeb, 0x8c, 0x77, 0x50, 0xbf,
	0x51, 0x6c, 0x79, 0xad, 0xff, 0x8f, 0xe5, 0x50, 0x2d, 0x44, 0xfb, 0x0a, 0xe4, 0x31, 0x1b, 0x15,
	0x0c, 0x3c, 0xbd, 0x43, 0xd6, 0xc5, 0xdf, 0x41, 0x40, 0xd9, 0x68, 0xfb, 0x3d, 0x68, 0x4b, 0x9f,
	0x58, 0x03, 0xf9, 0x27, 0x9d, 0x2f, 0xaa, 0x5e, 0x2f, 0xbc, 0x72, 0x77, 0xd5, 0x5d, 0x50, 0xc5,
	0x35, 0x58, 0x40, 0xa5, 0x7f, 0x34, 0xd8, 0xdd, 0x81, 0xc7, 0x7e, 0x1c, 0x76, 0xc4, 0xd3, 0xd4,
	0xf1, 0x2f, 0xf2, 0xc3, 0x2c, 0x7f, 0x99, 0xce, 0xa6, 0xe3, 0xef, 0xaa, 0x38, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x5d, 0xa7, 0x02, 0xdc, 0x04, 0x00, 0x00,
}
