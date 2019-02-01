// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/sitelink_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for Sitelink placeholder fields.
type SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField int32

const (
	// Not specified.
	SitelinkPlaceholderFieldEnum_UNSPECIFIED SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	SitelinkPlaceholderFieldEnum_UNKNOWN SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 1
	// Data Type: STRING. The link text for your sitelink.
	SitelinkPlaceholderFieldEnum_TEXT SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 2
	// Data Type: STRING. First line of the sitelink description.
	SitelinkPlaceholderFieldEnum_LINE_1 SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 3
	// Data Type: STRING. Second line of the sitelink description.
	SitelinkPlaceholderFieldEnum_LINE_2 SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 4
	// Data Type: URL_LIST. Final URLs for the sitelink when using Upgraded
	// URLs.
	SitelinkPlaceholderFieldEnum_FINAL_URLS SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 5
	// Data Type: URL_LIST. Final Mobile URLs for the sitelink when using
	// Upgraded URLs.
	SitelinkPlaceholderFieldEnum_FINAL_MOBILE_URLS SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 6
	// Data Type: URL. Tracking template for the sitelink when using Upgraded
	// URLs.
	SitelinkPlaceholderFieldEnum_TRACKING_URL SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 7
	// Data Type: STRING. Final URL suffix for sitelink when using parallel
	// tracking.
	SitelinkPlaceholderFieldEnum_FINAL_URL_SUFFIX SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 8
)

var SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "TEXT",
	3: "LINE_1",
	4: "LINE_2",
	5: "FINAL_URLS",
	6: "FINAL_MOBILE_URLS",
	7: "TRACKING_URL",
	8: "FINAL_URL_SUFFIX",
}
var SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"TEXT":              2,
	"LINE_1":            3,
	"LINE_2":            4,
	"FINAL_URLS":        5,
	"FINAL_MOBILE_URLS": 6,
	"TRACKING_URL":      7,
	"FINAL_URL_SUFFIX":  8,
}

func (x SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) String() string {
	return proto.EnumName(SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField_name, int32(x))
}
func (SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sitelink_placeholder_field_1dbac5ed08bf4cb2, []int{0, 0}
}

// Values for Sitelink placeholder fields.
type SitelinkPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SitelinkPlaceholderFieldEnum) Reset()         { *m = SitelinkPlaceholderFieldEnum{} }
func (m *SitelinkPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*SitelinkPlaceholderFieldEnum) ProtoMessage()    {}
func (*SitelinkPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_sitelink_placeholder_field_1dbac5ed08bf4cb2, []int{0}
}
func (m *SitelinkPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SitelinkPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *SitelinkPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SitelinkPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *SitelinkPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SitelinkPlaceholderFieldEnum.Merge(dst, src)
}
func (m *SitelinkPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_SitelinkPlaceholderFieldEnum.Size(m)
}
func (m *SitelinkPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SitelinkPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SitelinkPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SitelinkPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.SitelinkPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField", SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField_name, SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/sitelink_placeholder_field.proto", fileDescriptor_sitelink_placeholder_field_1dbac5ed08bf4cb2)
}

var fileDescriptor_sitelink_placeholder_field_1dbac5ed08bf4cb2 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0x81, 0x0b, 0xe4, 0x70, 0x73, 0xef, 0xdc, 0x89, 0x26, 0x2e, 0x64, 0x21, 0x0f,
	0x30, 0xad, 0xba, 0x34, 0x31, 0x69, 0xb1, 0x25, 0x13, 0xea, 0xd0, 0x50, 0x8a, 0xc4, 0x34, 0x69,
	0x2a, 0x1d, 0x6b, 0x63, 0xe9, 0x10, 0x06, 0x78, 0x20, 0x77, 0xfa, 0x28, 0x2c, 0x7c, 0x26, 0xd3,
	0x16, 0xea, 0x0a, 0x37, 0x93, 0x7f, 0xce, 0x99, 0xef, 0x9f, 0x9c, 0xff, 0xc0, 0x6d, 0x2c, 0x44,
	0x9c, 0x72, 0x35, 0x8c, 0xa4, 0x5a, 0xca, 0x5c, 0x6d, 0x35, 0x95, 0x67, 0x9b, 0x85, 0x54, 0x65,
	0xb2, 0xe6, 0x69, 0x92, 0xbd, 0x06, 0xcb, 0x34, 0x9c, 0xf3, 0x17, 0x91, 0x46, 0x7c, 0x15, 0x3c,
	0x27, 0x3c, 0x8d, 0xc8, 0x72, 0x25, 0xd6, 0x02, 0x77, 0x4b, 0x88, 0x84, 0x91, 0x24, 0x15, 0x4f,
	0xb6, 0x1a, 0x29, 0xf8, 0xde, 0x4e, 0x81, 0x73, 0x77, 0xef, 0xe1, 0x7c, 0x5b, 0x58, 0xb9, 0x83,
	0x99, 0x6d, 0x16, 0xbd, 0x77, 0x05, 0xce, 0x8e, 0x3d, 0xc0, 0xff, 0xa0, 0xe3, 0x31, 0xd7, 0x31,
	0xfb, 0xd4, 0xa2, 0xe6, 0x1d, 0xfa, 0x85, 0x3b, 0xd0, 0xf2, 0xd8, 0x90, 0x8d, 0x1e, 0x18, 0x52,
	0x70, 0x1b, 0x1a, 0x13, 0x73, 0x36, 0x41, 0x35, 0x0c, 0xd0, 0xb4, 0x29, 0x33, 0x83, 0x4b, 0x54,
	0xaf, 0xf4, 0x15, 0x6a, 0xe0, 0xbf, 0x00, 0x16, 0x65, 0xba, 0x1d, 0x78, 0x63, 0xdb, 0x45, 0xbf,
	0xf1, 0x29, 0xfc, 0x2f, 0xef, 0xf7, 0x23, 0x83, 0xda, 0x66, 0x59, 0x6e, 0x62, 0x04, 0x7f, 0x26,
	0x63, 0xbd, 0x3f, 0xa4, 0x6c, 0x90, 0x97, 0x50, 0x0b, 0x9f, 0x00, 0xaa, 0xc0, 0xc0, 0xf5, 0x2c,
	0x8b, 0xce, 0x50, 0xdb, 0xf8, 0x54, 0xe0, 0x62, 0x2e, 0x16, 0xe4, 0xc7, 0x91, 0x8d, 0xee, 0xb1,
	0x71, 0x9c, 0x3c, 0x30, 0x47, 0x79, 0x34, 0xf6, 0x7c, 0x2c, 0xd2, 0x30, 0x8b, 0x89, 0x58, 0xc5,
	0x6a, 0xcc, 0xb3, 0x22, 0xce, 0xc3, 0x0a, 0x96, 0x89, 0x3c, 0xb2, 0x91, 0x9b, 0xe2, 0x7c, 0xab,
	0xd5, 0x07, 0xba, 0xfe, 0x51, 0xeb, 0x0e, 0x4a, 0x2b, 0x3d, 0x92, 0xa4, 0x94, 0xb9, 0x9a, 0x6a,
	0x24, 0xcf, 0x56, 0xee, 0x0e, 0x7d, 0x5f, 0x8f, 0xa4, 0x5f, 0xf5, 0xfd, 0xa9, 0xe6, 0x17, 0xfd,
	0xa7, 0x66, 0xf1, 0xe9, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xa3, 0x37, 0xa7, 0x05,
	0x02, 0x00, 0x00,
}