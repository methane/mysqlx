// Code generated by protoc-gen-go.
// source: mysqlx_resultset.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ColumnMetaData_FieldType int32

const (
	ColumnMetaData_SINT     ColumnMetaData_FieldType = 1
	ColumnMetaData_UINT     ColumnMetaData_FieldType = 2
	ColumnMetaData_DOUBLE   ColumnMetaData_FieldType = 5
	ColumnMetaData_FLOAT    ColumnMetaData_FieldType = 6
	ColumnMetaData_BYTES    ColumnMetaData_FieldType = 7
	ColumnMetaData_TIME     ColumnMetaData_FieldType = 10
	ColumnMetaData_DATETIME ColumnMetaData_FieldType = 12
	ColumnMetaData_SET      ColumnMetaData_FieldType = 15
	ColumnMetaData_ENUM     ColumnMetaData_FieldType = 16
	ColumnMetaData_BIT      ColumnMetaData_FieldType = 17
	ColumnMetaData_DECIMAL  ColumnMetaData_FieldType = 18
)

var ColumnMetaData_FieldType_name = map[int32]string{
	1:  "SINT",
	2:  "UINT",
	5:  "DOUBLE",
	6:  "FLOAT",
	7:  "BYTES",
	10: "TIME",
	12: "DATETIME",
	15: "SET",
	16: "ENUM",
	17: "BIT",
	18: "DECIMAL",
}
var ColumnMetaData_FieldType_value = map[string]int32{
	"SINT":     1,
	"UINT":     2,
	"DOUBLE":   5,
	"FLOAT":    6,
	"BYTES":    7,
	"TIME":     10,
	"DATETIME": 12,
	"SET":      15,
	"ENUM":     16,
	"BIT":      17,
	"DECIMAL":  18,
}

func (x ColumnMetaData_FieldType) Enum() *ColumnMetaData_FieldType {
	p := new(ColumnMetaData_FieldType)
	*p = x
	return p
}
func (x ColumnMetaData_FieldType) String() string {
	return proto1.EnumName(ColumnMetaData_FieldType_name, int32(x))
}
func (x *ColumnMetaData_FieldType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(ColumnMetaData_FieldType_value, data, "ColumnMetaData_FieldType")
	if err != nil {
		return err
	}
	*x = ColumnMetaData_FieldType(value)
	return nil
}
func (ColumnMetaData_FieldType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{3, 0} }

// resultsets are finished, OUT paramset is next
type FetchDoneMoreOutParams struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *FetchDoneMoreOutParams) Reset()                    { *m = FetchDoneMoreOutParams{} }
func (m *FetchDoneMoreOutParams) String() string            { return proto1.CompactTextString(m) }
func (*FetchDoneMoreOutParams) ProtoMessage()               {}
func (*FetchDoneMoreOutParams) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

// resultset and out-params are finished, but more resultsets available
type FetchDoneMoreResultsets struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *FetchDoneMoreResultsets) Reset()                    { *m = FetchDoneMoreResultsets{} }
func (m *FetchDoneMoreResultsets) String() string            { return proto1.CompactTextString(m) }
func (*FetchDoneMoreResultsets) ProtoMessage()               {}
func (*FetchDoneMoreResultsets) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

// all resultsets are finished
type FetchDone struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *FetchDone) Reset()                    { *m = FetchDone{} }
func (m *FetchDone) String() string            { return proto1.CompactTextString(m) }
func (*FetchDone) ProtoMessage()               {}
func (*FetchDone) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

// meta data of a Column
//
// .. note:: the encoding used for the different ``bytes`` fields in the meta data is externally
//   controlled.
//   .. seealso:: https://dev.mysql.com/doc/refman/5.0/en/charset-connection.html
//
// .. note::
//   The server may not set the ``original_{table|name}`` fields if they are equal to the plain
//   ``{table|name}`` field.
//
//   A client has to reconstruct it like::
//
//     if .original_name is empty and .name is not empty:
//       .original_name = .name
//
//     if .original_table is empty and .table is not empty:
//       .original_table = .table
//
// .. note::
//   ``compact metadata format`` can be requested by the client. In that case only ``.type`` is set and
//   all other fields are empty.
//
//
// :param type:
//   .. table:: Expected Datatype of Mysqlx.Resultset.Row per SQL Type for non NULL values
//
//     ================= ============ ======= ========== ====== ========
//     SQL Type          .type        .length .frac_dig  .flags .charset
//     ================= ============ ======= ========== ====== ========
//     TINY              SINT         x
//     TINY UNSIGNED     UINT         x                  x
//     SHORT             SINT         x
//     SHORT UNSIGNED    UINT         x                  x
//     INT24             SINT         x
//     INT24 UNSIGNED    UINT         x                  x
//     INT               SINT         x
//     INT UNSIGNED      UINT         x                  x
//     LONGLONG          SINT         x
//     LONGLONG UNSIGNED UINT         x                  x
//     DOUBLE            DOUBLE       x       x          x
//     FLOAT             FLOAT        x       x          x
//     DECIMAL           DECIMAL      x       x          x
//     VARCHAR,CHAR,...  BYTES        x                  x      x
//     GEOMETRY          BYTES
//     TIME              TIME         x
//     DATE              DATETIME     x
//     DATETIME          DATETIME     x
//     YEAR              UINT         x                  x
//     TIMESTAMP         DATETIME     x
//     SET               SET                                    x
//     ENUM              ENUM                                   x
//     NULL              BYTES
//     BIT               BIT          x
//     ================= ============ ======= ========== ====== ========
//
//   .. note:: the SQL "NULL" value is sent as an empty field value in :protobuf:msg:`Mysqlx.Resultset::Row`
//   .. seealso:: protobuf encoding of primitive datatypes are decribed in https://developers.google.com/protocol-buffers/docs/encoding
//
//   SINT
//
//     ``.length``
//       maximum number of displayable decimal digits (including minus sign) of the type
//
//       .. note::
//         valid range is 0-255, but usually you'll see 1-20
//
//       =============== ==
//       SQL Type        max digits per type
//       =============== ==
//       TINY SIGNED      4
//       SHORT SIGNED     6
//       INT24 SIGNED     8
//       INT SIGNED      11
//       LONGLONG SIGNED 20
//       =============== ==
//
//       .. seealso:: definition of ``M`` in https://dev.mysql.com/doc/refman/5.5/en/numeric-type-overview.html
//
//     ``value``
//       variable length encoded signed 64 integer
//
//   UINT
//
//     ``.flags & 1`` (zerofill)
//       the client has to left pad with 0's up to .length
//
//     ``.length``
//       maximum number of displayable decimal digits of the type
//
//       .. note::
//         valid range is 0-255, but usually you'll see 1-20
//
//       ================= ==
//       SQL Type          max digits per type
//       ================= ==
//       TINY UNSIGNED      3
//       SHORT UNSIGNED     5
//       INT24 UNSIGNED     8
//       INT UNSIGNED      10
//       LONGLONG UNSIGNED 20
//       ================= ==
//
//       .. seealso:: definition of ``M`` in https://dev.mysql.com/doc/refman/5.5/en/numeric-type-overview.html
//
//     ``value``
//       variable length encoded unsigned 64 integer
//
//   BIT
//
//     ``.length``
//       maximum number of displayable binary digits
//
//       .. note:: valid range for M of the ``BIT`` type is 1 - 64
//       .. seealso:: https://dev.mysql.com/doc/refman/5.5/en/numeric-type-overview.html
//
//     ``value``
//       variable length encoded unsigned 64 integer
//
//   DOUBLE
//
//     ``.length``
//       maximum number of displayable decimal digits (including the decimal point and ``.fractional_digits``)
//
//     ``.fractional_digits``
//       maximum number of displayable decimal digits following the decimal point
//
//     ``value``
//       encoded as Protobuf's 'double'
//
//   FLOAT
//
//     ``.length``
//       maximum number of displayable decimal digits (including the decimal point and ``.fractional_digits``)
//
//     ``.fractional_digits``
//       maximum number of displayable decimal digits following the decimal point
//
//     ``value``
//       encoded as Protobuf's 'float'
//
//   BYTES, ENUM
//     BYTES is used for all opaque byte strings that may have a charset
//
//       * TINYBLOB, BLOB, MEDIUMBLOB, LONGBLOB
//       * TINYTEXT, TEXT, MEDIUMTEXT, LONGTEXT
//       * VARCHAR, VARBINARY
//       * CHAR, BINARY
//       * ENUM
//
//     ``.length``
//       the maximum length of characters of the underlying type
//
//     ``.flags & 1`` (rightpad)
//       if the length of the field is less than ``.length``, the receiver is
//       supposed to add padding characters to the right end of the string.
//       If the ``.charset`` is "binary", the padding character is ``0x00``,
//       otherwise it is a space character as defined by that character set.
//
//       ============= ======= ======== =======
//       SQL Type      .length .charset .flags
//       ============= ======= ======== =======
//       TINYBLOB      256     binary
//       BLOB          65535   binary
//       VARCHAR(32)   32      utf8
//       VARBINARY(32) 32      utf8_bin
//       BINARY(32)    32      binary   rightpad
//       CHAR(32)      32      utf8     rightpad
//       ============= ======= ======== =======
//
//     ``value``
//       sequence of bytes with added one extra '\0' byte at the end. To obtain the
//       original string, the extra '\0' should be removed.
//       .. note:: the length of the string can be acquired with protobuf's field length() method
//         length of sequence-of-bytes = length-of-field - 1
//       .. note:: the extra byte allows to distinguish between a NULL and empty byte sequence
//
//   TIME
//     A time value.
//
//     ``value``
//       the following bytes sequence:
//
//         ``| negate [ | hour | [ | minutes | [ | seconds | [ | useconds | ]]]]``
//
//       * negate - one byte, should be one of: 0x00 for "+", 0x01 for "-"
//       * hour - optional variable length encoded unsigned64 value for the hour
//       * minutes - optional variable length encoded unsigned64 value for the minutes
//       * seconds - optional variable length encoded unsigned64 value for the seconds
//       * useconds - optional variable length encoded unsigned64 value for the microseconds
//
//       .. seealso:: protobuf encoding in https://developers.google.com/protocol-buffers/docs/encoding
//       .. note:: hour, minutes, seconds, useconds are optional if all the values to the right are 0
//
//       Example: 0x00 -> +00:00:00.000000
//
//   DATETIME
//     A date or date and time value.
//
//     ``value``
//       a sequence of variants, arranged as follows:
//
//         ``| year | month | day | [ | hour | [ | minutes | [ | seconds | [ | useconds | ]]]]``
//
//       * year - variable length encoded unsigned64 value for the year
//       * month - variable length encoded unsigned64 value for the month
//       * day - variable length encoded unsigned64 value for the day
//       * hour - optional variable length encoded unsigned64 value for the hour
//       * minutes - optional variable length encoded unsigned64 value for the minutes
//       * seconds - optional variable length encoded unsigned64 value for the seconds
//       * useconds - optional variable length encoded unsigned64 value for the microseconds
//
//       .. note:: hour, minutes, seconds, useconds are optional if all the values to the right are 0
//
//     ``.flags & 1`` (timestamp)
//
//       ============= =======
//       SQL Type      .flags
//       ============= =======
//       DATETIME
//       TIMESTAMP     1
//
//   DECIMAL
//     An arbitrary length number. The number is encoded as a single byte
//     indicating the position of the decimal point followed by the Packed BCD
//     encoded number. Packed BCD is used to simplify conversion to and
//     from strings and other native arbitrary precision math datatypes.
//     .. seealso:: packed BCD in https://en.wikipedia.org/wiki/Binary-coded_decimal
//
//     ``.length``
//       maximum number of displayable decimal digits (*excluding* the decimal point and sign, but including ``.fractional_digits``)
//
//       .. note:: should be in the range of 1 - 65
//
//     ``.fractional_digits``
//       is the decimal digits to display out of length
//
//       .. note:: should be in the range of 0 - 30
//
//     ``value``
//       the following bytes sequence:
//
//         ``| scale | BCD | sign | [0x0] |``
//
//       * scale - 8bit scale value (number of decimal digit after the '.')
//       * BCD - BCD encoded digits (4 bits for each digit)
//       * sign - sign encoded on 4 bits (0xc = "+", 0xd = "-")
//       * 0x0 - last 4bits if length(digits) % 2 == 0
//
//       Example: x04 0x12 0x34 0x01 0xd0 -> -12.3401
//
//   SET
//     A list of strings representing a SET of values.
//
//     ``value``
//       A sequence of 0 or more of protobuf's bytes (length prepended octets) or one of
//       the special sequences with a predefined meaning listed below.
//
//       Example (length of the bytes array shown in brackets):
//         * ``[0]`` - the NULL value
//         * ``[1] 0x00`` - a set containing a blank string ''
//         * ``[1] 0x01`` - this would be an invalid value, but is to be treated as the empty set
//         * ``[2] 0x01 0x00`` - a set with a single item, which is the '\0' character
//         * ``[8] 0x03 F O O 0x03 B A R`` - a set with 2 items: FOO,BAR
//
//
// :param name: name of the column
// :param original_name: name of the column before an alias was applied
// :param table: name of the table the column orginates from
// :param original_table: name of the table the column orginates from before an alias was applied
// :param schema: schema the column originates from
// :param catalog:
//   catalog the schema originates from
//
//   .. note::
//     as there is current no support for catalogs in MySQL, don't expect this field to be set.
//     In the MySQL C/S protocol the field had the value ``def`` all the time.
//
// :param fractional_digits: displayed factional decimal digits for floating point and fixed point numbers
// :param length: maximum count of displayable characters of .type
// :param flags:
//   ``.type`` specific flags
//
//   ======= ====== ===========
//   type    value  description
//   ======= ====== ===========
//   UINT    0x0001 zerofill
//   DOUBLE  0x0001 unsigned
//   FLOAT   0x0001 unsigned
//   DECIMAL 0x0001 unsigned
//   BYTES   0x0001 rightpad
//   ======= ====== ===========
//
//   ====== ================
//   value  description
//   ====== ================
//   0x0010 NOT_NULL
//   0x0020 PRIMARY_KEY
//   0x0040 UNIQUE_KEY
//   0x0080 MULTIPLE_KEY
//   0x0100 AUTO_INCREMENT
//   ====== ================
//
//   default: 0
// :param content_type:
//   a hint about the higher-level encoding of a BYTES field
//
//   ====== ====== ===========
//   type   value  description
//   ====== ====== ===========
//   BYTES  0x0001 GEOMETRY (WKB encoding)
//   BYTES  0x0002 JSON (text encoding)
//   BYTES  0x0003 XML (text encoding)
//   ====== ====== ===========
//
//   .. note::
//     this list isn't comprehensive. As guideline: the field's value is expected
//     to pass a validator check on client and server if this field is set.
//     If the server adds more internal datatypes that rely on BLOB storage
//     like image manipulation, seeking into complex types in BLOBs, ... more
//     types will be added.
//
type ColumnMetaData struct {
	// datatype of the field in a row
	Type             *ColumnMetaData_FieldType `protobuf:"varint,1,req,name=type,enum=Mysqlx.Resultset.ColumnMetaData_FieldType" json:"type,omitempty"`
	Name             []byte                    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	OriginalName     []byte                    `protobuf:"bytes,3,opt,name=original_name" json:"original_name,omitempty"`
	Table            []byte                    `protobuf:"bytes,4,opt,name=table" json:"table,omitempty"`
	OriginalTable    []byte                    `protobuf:"bytes,5,opt,name=original_table" json:"original_table,omitempty"`
	Schema           []byte                    `protobuf:"bytes,6,opt,name=schema" json:"schema,omitempty"`
	Catalog          []byte                    `protobuf:"bytes,7,opt,name=catalog" json:"catalog,omitempty"`
	Collation        *uint64                   `protobuf:"varint,8,opt,name=collation" json:"collation,omitempty"`
	FractionalDigits *uint32                   `protobuf:"varint,9,opt,name=fractional_digits" json:"fractional_digits,omitempty"`
	Length           *uint32                   `protobuf:"varint,10,opt,name=length" json:"length,omitempty"`
	Flags            *uint32                   `protobuf:"varint,11,opt,name=flags" json:"flags,omitempty"`
	ContentType      *uint32                   `protobuf:"varint,12,opt,name=content_type" json:"content_type,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *ColumnMetaData) Reset()                    { *m = ColumnMetaData{} }
func (m *ColumnMetaData) String() string            { return proto1.CompactTextString(m) }
func (*ColumnMetaData) ProtoMessage()               {}
func (*ColumnMetaData) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ColumnMetaData) GetType() ColumnMetaData_FieldType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ColumnMetaData_SINT
}

func (m *ColumnMetaData) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ColumnMetaData) GetOriginalName() []byte {
	if m != nil {
		return m.OriginalName
	}
	return nil
}

func (m *ColumnMetaData) GetTable() []byte {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *ColumnMetaData) GetOriginalTable() []byte {
	if m != nil {
		return m.OriginalTable
	}
	return nil
}

func (m *ColumnMetaData) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *ColumnMetaData) GetCatalog() []byte {
	if m != nil {
		return m.Catalog
	}
	return nil
}

func (m *ColumnMetaData) GetCollation() uint64 {
	if m != nil && m.Collation != nil {
		return *m.Collation
	}
	return 0
}

func (m *ColumnMetaData) GetFractionalDigits() uint32 {
	if m != nil && m.FractionalDigits != nil {
		return *m.FractionalDigits
	}
	return 0
}

func (m *ColumnMetaData) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *ColumnMetaData) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *ColumnMetaData) GetContentType() uint32 {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return 0
}

//   Row in a Resultset
//
//   a row is represented as a list of fields encoded as byte blobs.
//   Blob of size 0 represents the NULL value. Otherwise, if it contains at least
//   one byte, it encodes a non-null value of the field using encoding appropriate for the
//   type of the value given by ``ColumnMetadata``, as specified
//   in the :protobuf:msg:`Mysqlx.Resultset::ColumnMetaData` description.
//
type Row struct {
	Field            [][]byte `protobuf:"bytes,1,rep,name=field" json:"field,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Row) Reset()                    { *m = Row{} }
func (m *Row) String() string            { return proto1.CompactTextString(m) }
func (*Row) ProtoMessage()               {}
func (*Row) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *Row) GetField() [][]byte {
	if m != nil {
		return m.Field
	}
	return nil
}

func init() {
	proto1.RegisterType((*FetchDoneMoreOutParams)(nil), "Mysqlx.Resultset.FetchDoneMoreOutParams")
	proto1.RegisterType((*FetchDoneMoreResultsets)(nil), "Mysqlx.Resultset.FetchDoneMoreResultsets")
	proto1.RegisterType((*FetchDone)(nil), "Mysqlx.Resultset.FetchDone")
	proto1.RegisterType((*ColumnMetaData)(nil), "Mysqlx.Resultset.ColumnMetaData")
	proto1.RegisterType((*Row)(nil), "Mysqlx.Resultset.Row")
	proto1.RegisterEnum("Mysqlx.Resultset.ColumnMetaData_FieldType", ColumnMetaData_FieldType_name, ColumnMetaData_FieldType_value)
}

func init() { proto1.RegisterFile("mysqlx_resultset.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0x41, 0x8f, 0xda, 0x30,
	0x10, 0x85, 0x15, 0x92, 0x10, 0x32, 0x04, 0xd6, 0x58, 0x5b, 0xea, 0x95, 0x7a, 0x88, 0x22, 0x55,
	0x8a, 0x7a, 0xc8, 0xa1, 0xa7, 0x5e, 0x61, 0x13, 0x24, 0x24, 0xb2, 0x54, 0x4b, 0x38, 0xb4, 0x17,
	0xe4, 0xcd, 0x9a, 0x90, 0xca, 0x89, 0x69, 0x6c, 0xd4, 0xe5, 0xda, 0x3f, 0xd2, 0xbf, 0x5a, 0xd9,
	0xa8, 0x54, 0xbb, 0xb7, 0xf7, 0xbe, 0x79, 0xf6, 0xcc, 0x48, 0x03, 0xd3, 0xe6, 0x2c, 0x7f, 0xf2,
	0x97, 0x5d, 0xc7, 0xe4, 0x89, 0x2b, 0xc9, 0x54, 0x72, 0xec, 0x84, 0x12, 0x18, 0xe5, 0x86, 0x27,
	0x8f, 0xff, 0x78, 0x44, 0x60, 0xba, 0x60, 0xaa, 0x3c, 0xa4, 0xa2, 0x65, 0xb9, 0xe8, 0xd8, 0xfa,
	0xa4, 0xbe, 0xd2, 0x8e, 0x36, 0x32, 0xba, 0x83, 0xf7, 0xaf, 0x2a, 0xd7, 0x37, 0x32, 0x1a, 0x82,
	0x7f, 0x2d, 0x45, 0x7f, 0x6c, 0x18, 0xdf, 0x0b, 0x7e, 0x6a, 0xda, 0x9c, 0x29, 0x9a, 0x52, 0x45,
	0xf1, 0x17, 0x70, 0xd4, 0xf9, 0xc8, 0x88, 0x15, 0xf6, 0xe2, 0xf1, 0xe7, 0x4f, 0xc9, 0xdb, 0xae,
	0xc9, 0xeb, 0x7c, 0xb2, 0xa8, 0x19, 0x7f, 0x2e, 0xce, 0x47, 0x86, 0x03, 0x70, 0x5a, 0xda, 0x30,
	0xd2, 0x0b, 0xad, 0x38, 0xc0, 0xef, 0x60, 0x24, 0xba, 0xba, 0xaa, 0x5b, 0xca, 0x77, 0x06, 0xdb,
	0x06, 0x8f, 0xc0, 0x55, 0xf4, 0x89, 0x33, 0xe2, 0x18, 0x3b, 0x85, 0xf1, 0x35, 0x75, 0xe1, 0xae,
	0xe1, 0x63, 0xe8, 0xcb, 0xf2, 0xc0, 0x1a, 0x4a, 0xfa, 0xc6, 0xdf, 0x80, 0x57, 0x52, 0x45, 0xb9,
	0xa8, 0x88, 0x67, 0xc0, 0x04, 0xfc, 0x52, 0x70, 0x4e, 0x55, 0x2d, 0x5a, 0x32, 0x08, 0xad, 0xd8,
	0xc1, 0x77, 0x30, 0xd9, 0x77, 0xb4, 0xd4, 0x84, 0xf2, 0xdd, 0x73, 0x5d, 0xd5, 0x4a, 0x12, 0x3f,
	0xb4, 0xe2, 0x91, 0xfe, 0x8e, 0xb3, 0xb6, 0x52, 0x07, 0x02, 0xc6, 0x8f, 0xc0, 0xdd, 0x73, 0x5a,
	0x49, 0x32, 0x34, 0xf6, 0x16, 0x82, 0x52, 0xb4, 0x8a, 0xb5, 0x6a, 0x67, 0x76, 0x0f, 0x34, 0x8d,
	0x7e, 0x5b, 0xe0, 0xff, 0xdf, 0x6e, 0x00, 0xce, 0x66, 0xf9, 0x50, 0x20, 0x4b, 0xab, 0xad, 0x56,
	0x3d, 0x0c, 0xd0, 0x4f, 0xd7, 0xdb, 0xf9, 0x2a, 0x43, 0x2e, 0xf6, 0xc1, 0x5d, 0xac, 0xd6, 0xb3,
	0x02, 0xf5, 0xb5, 0x9c, 0x7f, 0x2b, 0xb2, 0x0d, 0xf2, 0x74, 0xb6, 0x58, 0xe6, 0x19, 0x02, 0x1c,
	0xc0, 0x20, 0x9d, 0x15, 0x99, 0x71, 0x01, 0xf6, 0xc0, 0xde, 0x64, 0x05, 0xba, 0xd1, 0x81, 0xec,
	0x61, 0x9b, 0x23, 0xa4, 0xd1, 0x7c, 0x59, 0xa0, 0x09, 0x1e, 0x82, 0x97, 0x66, 0xf7, 0xcb, 0x7c,
	0xb6, 0x42, 0x38, 0xba, 0x05, 0xfb, 0x51, 0xfc, 0x32, 0x03, 0xeb, 0x51, 0x88, 0x15, 0xda, 0x71,
	0x30, 0xff, 0x08, 0x1f, 0x4a, 0xd1, 0x24, 0xe6, 0x52, 0x92, 0xf2, 0xc7, 0x45, 0xbc, 0x5c, 0x0e,
	0xe5, 0xe9, 0xb4, 0xff, 0xee, 0x1a, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0x30, 0xd9,
	0x4b, 0x02, 0x00, 0x00,
}
