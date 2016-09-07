// Code generated by protoc-gen-go.
// source: mysqlx_session.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// the initial message send from the client to the server to start the
// authentication proccess
//
// :param mech_name: authentication mechanism name
// :param auth_data: authentication data
// :param initial_response: initial response
// :Returns: :protobuf:msg:`Mysqlx.Session::AuthenticateContinue`
type AuthenticateStart struct {
	MechName         *string `protobuf:"bytes,1,req,name=mech_name" json:"mech_name,omitempty"`
	AuthData         []byte  `protobuf:"bytes,2,opt,name=auth_data" json:"auth_data,omitempty"`
	InitialResponse  []byte  `protobuf:"bytes,3,opt,name=initial_response" json:"initial_response,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthenticateStart) Reset()                    { *m = AuthenticateStart{} }
func (m *AuthenticateStart) String() string            { return proto1.CompactTextString(m) }
func (*AuthenticateStart) ProtoMessage()               {}
func (*AuthenticateStart) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *AuthenticateStart) GetMechName() string {
	if m != nil && m.MechName != nil {
		return *m.MechName
	}
	return ""
}

func (m *AuthenticateStart) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

func (m *AuthenticateStart) GetInitialResponse() []byte {
	if m != nil {
		return m.InitialResponse
	}
	return nil
}

// send by client or server after a :protobuf:msg:`Mysqlx.Session::AuthenticateStart` to
// exchange more auth data
//
// :param auth_data: authentication data
// :Returns: :protobuf:msg:`Mysqlx.Session::AuthenticateContinue`
type AuthenticateContinue struct {
	AuthData         []byte `protobuf:"bytes,1,req,name=auth_data" json:"auth_data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AuthenticateContinue) Reset()                    { *m = AuthenticateContinue{} }
func (m *AuthenticateContinue) String() string            { return proto1.CompactTextString(m) }
func (*AuthenticateContinue) ProtoMessage()               {}
func (*AuthenticateContinue) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *AuthenticateContinue) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

// sent by the server after successful authentication
//
// :param auth_data: authentication data
type AuthenticateOk struct {
	AuthData         []byte `protobuf:"bytes,1,opt,name=auth_data" json:"auth_data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AuthenticateOk) Reset()                    { *m = AuthenticateOk{} }
func (m *AuthenticateOk) String() string            { return proto1.CompactTextString(m) }
func (*AuthenticateOk) ProtoMessage()               {}
func (*AuthenticateOk) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *AuthenticateOk) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

// reset the current session
//
// :Returns: :protobuf:msg:`Mysqlx::Ok`
type Reset struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Reset) Reset()                    { *m = Reset{} }
func (m *Reset) String() string            { return proto1.CompactTextString(m) }
func (*Reset) ProtoMessage()               {}
func (*Reset) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

// close the current session
//
// :Returns: :protobuf:msg:`Mysqlx::Ok`
type Close struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Close) Reset()                    { *m = Close{} }
func (m *Close) String() string            { return proto1.CompactTextString(m) }
func (*Close) ProtoMessage()               {}
func (*Close) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func init() {
	proto1.RegisterType((*AuthenticateStart)(nil), "Mysqlx.Session.AuthenticateStart")
	proto1.RegisterType((*AuthenticateContinue)(nil), "Mysqlx.Session.AuthenticateContinue")
	proto1.RegisterType((*AuthenticateOk)(nil), "Mysqlx.Session.AuthenticateOk")
	proto1.RegisterType((*Reset)(nil), "Mysqlx.Session.Reset")
	proto1.RegisterType((*Close)(nil), "Mysqlx.Session.Close")
}

func init() { proto1.RegisterFile("mysqlx_session.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x46, 0xd9, 0x95, 0x45, 0x76, 0x28, 0xc5, 0x96, 0x1e, 0x72, 0xf0, 0x50, 0x2a, 0x42, 0xbd,
	0xe4, 0x3f, 0x68, 0xcf, 0x22, 0xd8, 0x9b, 0x97, 0x10, 0xe3, 0x48, 0xa3, 0x4d, 0x52, 0x3b, 0x13,
	0xa8, 0xff, 0x5e, 0xd2, 0x5e, 0xea, 0xde, 0x1e, 0x6f, 0x86, 0x07, 0x1f, 0x54, 0xee, 0x97, 0x7e,
	0xc6, 0x45, 0x11, 0x12, 0xd9, 0xe0, 0xe5, 0x34, 0x07, 0x0e, 0x65, 0xfe, 0xbc, 0x5a, 0xd9, 0x6f,
	0xb6, 0xe9, 0xa1, 0x78, 0x8c, 0x3c, 0xa0, 0x67, 0x6b, 0x34, 0x63, 0xcf, 0x7a, 0xe6, 0xb2, 0x80,
	0xb3, 0x43, 0x33, 0x28, 0xaf, 0x1d, 0x8a, 0x43, 0x7d, 0x6c, 0xcf, 0x49, 0xe9, 0xc8, 0x83, 0xfa,
	0xd0, 0xac, 0xc5, 0xb1, 0x3e, 0xb4, 0x59, 0x29, 0xe0, 0xc6, 0x7a, 0xcb, 0x56, 0x8f, 0x6a, 0x46,
	0x9a, 0x82, 0x27, 0x14, 0x57, 0xe9, 0xd2, 0x3c, 0x40, 0xb5, 0x8f, 0x76, 0xc1, 0xb3, 0xf5, 0x11,
	0xff, 0x47, 0x52, 0x37, 0x6b, 0xee, 0x20, 0xdf, 0xbf, 0xbe, 0x7c, 0x5f, 0x3e, 0xa5, 0xde, 0x35,
	0x9c, 0x5e, 0x91, 0x90, 0x13, 0x74, 0x63, 0x20, 0x7c, 0xba, 0x87, 0x5b, 0x13, 0x9c, 0x5c, 0x27,
	0x4a, 0xf3, 0xb5, 0xc1, 0xb2, 0x6d, 0x7c, 0x8f, 0x9f, 0x6f, 0xa7, 0x95, 0xfe, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0xf2, 0xad, 0x6e, 0x04, 0x01, 0x00, 0x00,
}
