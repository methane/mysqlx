// Code generated by protoc-gen-go.
// source: mysqlx_notice.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Frame_Scope int32

const (
	Frame_GLOBAL Frame_Scope = 1
	Frame_LOCAL  Frame_Scope = 2
)

var Frame_Scope_name = map[int32]string{
	1: "GLOBAL",
	2: "LOCAL",
}
var Frame_Scope_value = map[string]int32{
	"GLOBAL": 1,
	"LOCAL":  2,
}

func (x Frame_Scope) Enum() *Frame_Scope {
	p := new(Frame_Scope)
	*p = x
	return p
}
func (x Frame_Scope) String() string {
	return proto1.EnumName(Frame_Scope_name, int32(x))
}
func (x *Frame_Scope) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Frame_Scope_value, data, "Frame_Scope")
	if err != nil {
		return err
	}
	*x = Frame_Scope(value)
	return nil
}
func (Frame_Scope) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

type Warning_Level int32

const (
	Warning_NOTE    Warning_Level = 1
	Warning_WARNING Warning_Level = 2
	Warning_ERROR   Warning_Level = 3
)

var Warning_Level_name = map[int32]string{
	1: "NOTE",
	2: "WARNING",
	3: "ERROR",
}
var Warning_Level_value = map[string]int32{
	"NOTE":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x Warning_Level) Enum() *Warning_Level {
	p := new(Warning_Level)
	*p = x
	return p
}
func (x Warning_Level) String() string {
	return proto1.EnumName(Warning_Level_name, int32(x))
}
func (x *Warning_Level) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Warning_Level_value, data, "Warning_Level")
	if err != nil {
		return err
	}
	*x = Warning_Level(value)
	return nil
}
func (Warning_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{1, 0} }

type SessionStateChanged_Parameter int32

const (
	SessionStateChanged_CURRENT_SCHEMA      SessionStateChanged_Parameter = 1
	SessionStateChanged_ACCOUNT_EXPIRED     SessionStateChanged_Parameter = 2
	SessionStateChanged_GENERATED_INSERT_ID SessionStateChanged_Parameter = 3
	SessionStateChanged_ROWS_AFFECTED       SessionStateChanged_Parameter = 4
	SessionStateChanged_ROWS_FOUND          SessionStateChanged_Parameter = 5
	SessionStateChanged_ROWS_MATCHED        SessionStateChanged_Parameter = 6
	SessionStateChanged_TRX_COMMITTED       SessionStateChanged_Parameter = 7
	SessionStateChanged_TRX_ROLLEDBACK      SessionStateChanged_Parameter = 9
	SessionStateChanged_PRODUCED_MESSAGE    SessionStateChanged_Parameter = 10
	SessionStateChanged_CLIENT_ID_ASSIGNED  SessionStateChanged_Parameter = 11
)

var SessionStateChanged_Parameter_name = map[int32]string{
	1:  "CURRENT_SCHEMA",
	2:  "ACCOUNT_EXPIRED",
	3:  "GENERATED_INSERT_ID",
	4:  "ROWS_AFFECTED",
	5:  "ROWS_FOUND",
	6:  "ROWS_MATCHED",
	7:  "TRX_COMMITTED",
	9:  "TRX_ROLLEDBACK",
	10: "PRODUCED_MESSAGE",
	11: "CLIENT_ID_ASSIGNED",
}
var SessionStateChanged_Parameter_value = map[string]int32{
	"CURRENT_SCHEMA":      1,
	"ACCOUNT_EXPIRED":     2,
	"GENERATED_INSERT_ID": 3,
	"ROWS_AFFECTED":       4,
	"ROWS_FOUND":          5,
	"ROWS_MATCHED":        6,
	"TRX_COMMITTED":       7,
	"TRX_ROLLEDBACK":      9,
	"PRODUCED_MESSAGE":    10,
	"CLIENT_ID_ASSIGNED":  11,
}

func (x SessionStateChanged_Parameter) Enum() *SessionStateChanged_Parameter {
	p := new(SessionStateChanged_Parameter)
	*p = x
	return p
}
func (x SessionStateChanged_Parameter) String() string {
	return proto1.EnumName(SessionStateChanged_Parameter_name, int32(x))
}
func (x *SessionStateChanged_Parameter) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(SessionStateChanged_Parameter_value, data, "SessionStateChanged_Parameter")
	if err != nil {
		return err
	}
	*x = SessionStateChanged_Parameter(value)
	return nil
}
func (SessionStateChanged_Parameter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{3, 0}
}

// Common Frame for all Notices
//
// ===================================================== =====
// .type                                                 value
// ===================================================== =====
// :protobuf:msg:`Mysqlx.Notice::Warning`                1
// :protobuf:msg:`Mysqlx.Notice::SessionVariableChanged` 2
// :protobuf:msg:`Mysqlx.Notice::SessionStateChanged`    3
// ===================================================== =====
//
// :param type: the type of the payload
// :param payload: the payload of the notification
// :param scope: global or local notification
//
type Frame struct {
	Type             *uint32      `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Scope            *Frame_Scope `protobuf:"varint,2,opt,name=scope,enum=Mysqlx.Notice.Frame_Scope,def=1" json:"scope,omitempty"`
	Payload          []byte       `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto1.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

const Default_Frame_Scope Frame_Scope = Frame_GLOBAL

func (m *Frame) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Frame) GetScope() Frame_Scope {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return Default_Frame_Scope
}

func (m *Frame) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Server-side warnings and notes
//
// ``.scope`` == ``local``
//   ``.level``, ``.code`` and ``.msg`` map the content of
//
//   .. code-block:: sql
//
//     SHOW WARNINGS
//
// ``.scope`` == ``global``
//   (undefined) will be used for global, unstructured messages like:
//
//   * server is shutting down
//   * a node disconnected from group
//   * schema or table dropped
//
// ========================================== =======================
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =======================
// ``.type``                                  1
// ``.scope``                                 ``local`` or ``global``
// ========================================== =======================
//
// :param level: warning level: Note or Warning
// :param code: warning code
// :param msg: warning message
type Warning struct {
	Level            *Warning_Level `protobuf:"varint,1,opt,name=level,enum=Mysqlx.Notice.Warning_Level,def=2" json:"level,omitempty"`
	Code             *uint32        `protobuf:"varint,2,req,name=code" json:"code,omitempty"`
	Msg              *string        `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Warning) Reset()                    { *m = Warning{} }
func (m *Warning) String() string            { return proto1.CompactTextString(m) }
func (*Warning) ProtoMessage()               {}
func (*Warning) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

const Default_Warning_Level Warning_Level = Warning_WARNING

func (m *Warning) GetLevel() Warning_Level {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Default_Warning_Level
}

func (m *Warning) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Warning) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// Notify clients about changes to the current session variables
//
// Every change to a variable that is accessable through:
//
// .. code-block:: sql
//
//   SHOW SESSION VARIABLES
//
// ========================================== =========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =========
// ``.type``                                  2
// ``.scope``                                 ``local``
// ========================================== =========
//
// :param namespace: namespace that param belongs to
// :param param: name of the variable
// :param value: the changed value of param
type SessionVariableChanged struct {
	Param            *string `protobuf:"bytes,1,req,name=param" json:"param,omitempty"`
	Value            *Scalar `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SessionVariableChanged) Reset()                    { *m = SessionVariableChanged{} }
func (m *SessionVariableChanged) String() string            { return proto1.CompactTextString(m) }
func (*SessionVariableChanged) ProtoMessage()               {}
func (*SessionVariableChanged) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *SessionVariableChanged) GetParam() string {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return ""
}

func (m *SessionVariableChanged) GetValue() *Scalar {
	if m != nil {
		return m.Value
	}
	return nil
}

// Notify clients about changes to the internal session state
//
// ========================================== =========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =========
// ``.type``                                  3
// ``.scope``                                 ``local``
// ========================================== =========
//
// :param param: parameter key
// :param value: updated value
type SessionStateChanged struct {
	Param            *SessionStateChanged_Parameter `protobuf:"varint,1,req,name=param,enum=Mysqlx.Notice.SessionStateChanged_Parameter" json:"param,omitempty"`
	Value            *Scalar                        `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *SessionStateChanged) Reset()                    { *m = SessionStateChanged{} }
func (m *SessionStateChanged) String() string            { return proto1.CompactTextString(m) }
func (*SessionStateChanged) ProtoMessage()               {}
func (*SessionStateChanged) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *SessionStateChanged) GetParam() SessionStateChanged_Parameter {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return SessionStateChanged_CURRENT_SCHEMA
}

func (m *SessionStateChanged) GetValue() *Scalar {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto1.RegisterType((*Frame)(nil), "Mysqlx.Notice.Frame")
	proto1.RegisterType((*Warning)(nil), "Mysqlx.Notice.Warning")
	proto1.RegisterType((*SessionVariableChanged)(nil), "Mysqlx.Notice.SessionVariableChanged")
	proto1.RegisterType((*SessionStateChanged)(nil), "Mysqlx.Notice.SessionStateChanged")
	proto1.RegisterEnum("Mysqlx.Notice.Frame_Scope", Frame_Scope_name, Frame_Scope_value)
	proto1.RegisterEnum("Mysqlx.Notice.Warning_Level", Warning_Level_name, Warning_Level_value)
	proto1.RegisterEnum("Mysqlx.Notice.SessionStateChanged_Parameter", SessionStateChanged_Parameter_name, SessionStateChanged_Parameter_value)
}

func init() { proto1.RegisterFile("mysqlx_notice.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x49, 0xda, 0x6c, 0xed, 0xed, 0x9f, 0x1d, 0xa7, 0x52, 0xc3, 0xb2, 0x48, 0x09, 0x88,
	0x15, 0x24, 0x60, 0xf1, 0x69, 0x7d, 0x4a, 0x67, 0xa6, 0xdd, 0x60, 0xfe, 0x94, 0x49, 0x6a, 0x17,
	0x5f, 0xc2, 0x6c, 0x1a, 0x6b, 0x25, 0x6d, 0x6a, 0x92, 0x5d, 0xb6, 0x1f, 0xc0, 0x8f, 0x28, 0x7e,
	0x1d, 0x99, 0xb4, 0xbb, 0xea, 0xe2, 0x83, 0x6f, 0x93, 0xc3, 0xef, 0x9c, 0x7b, 0x72, 0xb9, 0xd0,
	0xdb, 0xec, 0x8b, 0x6f, 0xe9, 0x5d, 0xb4, 0xcd, 0xca, 0x75, 0x9c, 0x98, 0xbb, 0x3c, 0x2b, 0x33,
	0xdc, 0x71, 0x2b, 0xd1, 0xf4, 0x2a, 0xf1, 0xac, 0x7f, 0x64, 0x96, 0xa2, 0x14, 0xe5, 0x7e, 0x97,
	0x14, 0x07, 0xcc, 0xd8, 0x83, 0x36, 0xc9, 0xc5, 0x26, 0xc1, 0x6d, 0xa8, 0x4b, 0x5d, 0x57, 0x06,
	0xea, 0xb0, 0x83, 0xdf, 0x82, 0x56, 0xc4, 0xd9, 0x2e, 0xd1, 0xd5, 0x81, 0x32, 0xec, 0x8e, 0xce,
	0xcc, 0xbf, 0xd2, 0xcc, 0xca, 0x62, 0x06, 0x92, 0xb8, 0x38, 0x99, 0x3a, 0xfe, 0xd8, 0x72, 0xf0,
	0x29, 0x34, 0x76, 0x62, 0x9f, 0x66, 0x62, 0xa9, 0xd7, 0x06, 0xca, 0xb0, 0x6d, 0xbc, 0x00, 0xad,
	0x22, 0x30, 0xc0, 0x91, 0x41, 0x0a, 0x6e, 0x82, 0xe6, 0xf8, 0xc4, 0x72, 0x90, 0x6a, 0x7c, 0x57,
	0xa0, 0xb1, 0x10, 0xf9, 0x76, 0xbd, 0x5d, 0xe1, 0x77, 0xa0, 0xa5, 0xc9, 0x6d, 0x92, 0xea, 0x4a,
	0x35, 0xef, 0xfc, 0xd1, 0xbc, 0x23, 0x66, 0x3a, 0x92, 0xb9, 0x68, 0x2c, 0x2c, 0xee, 0xd9, 0xde,
	0x54, 0x76, 0x8e, 0xb3, 0xa5, 0x2c, 0x29, 0x3b, 0xb7, 0xa0, 0xb6, 0x29, 0x56, 0x7a, 0x6d, 0xa0,
	0x0e, 0x9b, 0xc6, 0x6b, 0xd0, 0x2a, 0x18, 0x3f, 0x81, 0xba, 0xe7, 0x87, 0x0c, 0x29, 0xb8, 0x05,
	0xf7, 0x46, 0xa4, 0xca, 0x1e, 0x8c, 0x73, 0x9f, 0xa3, 0x9a, 0x31, 0x83, 0x7e, 0x90, 0x14, 0xc5,
	0x3a, 0xdb, 0x7e, 0x14, 0xf9, 0x5a, 0x5c, 0xa7, 0x09, 0xf9, 0x22, 0xb6, 0xab, 0x64, 0x89, 0x3b,
	0xa0, 0xed, 0x44, 0x2e, 0x36, 0xd5, 0x52, 0x9a, 0xf8, 0x15, 0x68, 0xb7, 0x22, 0xbd, 0x39, 0x2c,
	0xa5, 0x35, 0xd2, 0xef, 0x4b, 0xd2, 0x87, 0x9d, 0x06, 0xb1, 0x48, 0x45, 0x6e, 0xfc, 0x50, 0xa1,
	0x77, 0x8c, 0x0c, 0x4a, 0x51, 0x3e, 0xe4, 0xbd, 0xff, 0x33, 0xaf, 0x3b, 0x7a, 0xf3, 0xe8, 0x2f,
	0xff, 0x61, 0x31, 0x67, 0x92, 0x4f, 0xca, 0x24, 0xff, 0xff, 0xe9, 0x3f, 0x15, 0x68, 0xfe, 0xb6,
	0x61, 0xe8, 0x92, 0x39, 0xe7, 0xcc, 0x0b, 0xa3, 0x80, 0x5c, 0x32, 0xd7, 0x42, 0x0a, 0xee, 0xc1,
	0xa9, 0x45, 0x88, 0x3f, 0xf7, 0xc2, 0x88, 0x5d, 0xcd, 0x6c, 0xce, 0x28, 0x52, 0xf1, 0x73, 0xe8,
	0x4d, 0x99, 0xc7, 0xb8, 0x15, 0x32, 0x1a, 0xd9, 0x5e, 0xc0, 0x78, 0x18, 0xd9, 0x14, 0xd5, 0xf0,
	0x53, 0xe8, 0x70, 0x7f, 0x11, 0x44, 0xd6, 0x64, 0xc2, 0x48, 0xc8, 0x28, 0xaa, 0xe3, 0x2e, 0x40,
	0x25, 0x4d, 0xfc, 0xb9, 0x47, 0x91, 0x86, 0x11, 0xb4, 0xab, 0x6f, 0xd7, 0x0a, 0xc9, 0x25, 0xa3,
	0xe8, 0x44, 0x9a, 0x42, 0x7e, 0x15, 0x11, 0xdf, 0x75, 0xed, 0x50, 0x9a, 0x1a, 0xb2, 0x89, 0x94,
	0xb8, 0xef, 0x38, 0x8c, 0x8e, 0x2d, 0xf2, 0x01, 0x35, 0xf1, 0x33, 0x40, 0x33, 0xee, 0xd3, 0x39,
	0x61, 0x34, 0x72, 0x59, 0x10, 0x58, 0x53, 0x86, 0x00, 0xf7, 0x01, 0x13, 0xc7, 0x96, 0x95, 0x6d,
	0x1a, 0x59, 0x41, 0x60, 0x4f, 0x3d, 0x46, 0x51, 0x6b, 0xfc, 0x12, 0xce, 0xe3, 0x6c, 0x63, 0x56,
	0xa7, 0x6c, 0xc6, 0x5f, 0x0f, 0x8f, 0xbb, 0xc3, 0x25, 0x5f, 0xdf, 0x7c, 0xfe, 0xa4, 0x55, 0xaf,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0xfd, 0x5d, 0xf9, 0x10, 0x03, 0x00, 0x00,
}
