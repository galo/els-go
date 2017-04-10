// Code generated by protoc-gen-go.
// source: google/api/log.proto
// DO NOT EDIT!

package serviceconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api2 "google.golang.org/genproto/googleapis/api/label"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A description of a log type. Example in YAML format:
//
//     - name: library.googleapis.com/activity_history
//       description: The history of borrowing and returning library items.
//       display_name: Activity
//       labels:
//       - key: /customer_id
//         description: Identifier of a library customer
type LogDescriptor struct {
	// The name of the log. It must be less than 512 characters long and can
	// include the following characters: upper- and lower-case alphanumeric
	// characters [A-Za-z0-9], and punctuation characters including
	// slash, underscore, hyphen, period [/_-.].
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The set of labels that are available to describe a specific log entry.
	// Runtime requests that contain labels not specified here are
	// considered invalid.
	Labels []*google_api2.LabelDescriptor `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	// A human-readable description of this log. This information appears in
	// the documentation and can contain details.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// The human-readable name for this log. This information appears on
	// the user interface and should be concise.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
}

func (m *LogDescriptor) Reset()                    { *m = LogDescriptor{} }
func (m *LogDescriptor) String() string            { return proto.CompactTextString(m) }
func (*LogDescriptor) ProtoMessage()               {}
func (*LogDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *LogDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogDescriptor) GetLabels() []*google_api2.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LogDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterType((*LogDescriptor)(nil), "google.api.LogDescriptor")
}

func init() { proto.RegisterFile("google/api/log.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x1b, 0x8a, 0x6e, 0xd5, 0xc3, 0x22, 0x12, 0xf4, 0x12, 0x3d, 0xf5, 0xb4, 0x01,
	0x7b, 0xf4, 0x64, 0x51, 0x44, 0x08, 0x12, 0x7a, 0xf4, 0x22, 0xd3, 0x74, 0x1c, 0x46, 0x36, 0x3b,
	0xcb, 0x6e, 0x11, 0x7c, 0x18, 0x2f, 0x3e, 0xa9, 0x74, 0x13, 0x68, 0x7a, 0xdb, 0xfd, 0xe6, 0x9b,
	0x7f, 0x66, 0xd4, 0x25, 0x89, 0x90, 0xc5, 0x0a, 0x3c, 0x57, 0x56, 0xc8, 0xf8, 0x20, 0x3b, 0xd1,
	0xaa, 0xa7, 0x06, 0x3c, 0x5f, 0x5f, 0x8d, 0x0d, 0xd8, 0xa0, 0xed, 0x9d, 0xbb, 0xdf, 0x4c, 0x9d,
	0xd7, 0x42, 0x4f, 0x18, 0xdb, 0xc0, 0x7e, 0x27, 0x41, 0x6b, 0x95, 0x3b, 0xe8, 0xb0, 0xc8, 0xca,
	0x6c, 0x71, 0xba, 0x4e, 0x6f, 0xbd, 0x54, 0xb3, 0xd4, 0x14, 0x8b, 0x49, 0x39, 0x5d, 0xcc, 0xef,
	0x6f, 0xcc, 0x21, 0xda, 0xd4, 0xfb, 0xca, 0x21, 0x60, 0x3d, 0xa8, 0xba, 0x54, 0xf3, 0xed, 0x40,
	0x59, 0x5c, 0x31, 0x4d, 0x79, 0x63, 0xa4, 0x6f, 0xd5, 0xd9, 0x96, 0xa3, 0xb7, 0xf0, 0xf3, 0x91,
	0x46, 0xe6, 0x83, 0xd2, 0xb3, 0x37, 0xe8, 0x70, 0xf5, 0xa5, 0x2e, 0x5a, 0xe9, 0x46, 0xe3, 0x56,
	0x27, 0xb5, 0x50, 0xb3, 0xdf, 0xbd, 0xc9, 0xde, 0x9f, 0x07, 0x4e, 0x62, 0xc1, 0x91, 0x91, 0x40,
	0x15, 0xa1, 0x4b, 0x97, 0x55, 0x7d, 0x09, 0x3c, 0xc7, 0x74, 0x74, 0xc4, 0xf0, 0xcd, 0x2d, 0xb6,
	0xe2, 0x3e, 0x99, 0x1e, 0x8e, 0x7e, 0x7f, 0x93, 0xfc, 0xe5, 0xb1, 0x79, 0xdd, 0xcc, 0x52, 0xe3,
	0xf2, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x25, 0x6c, 0x32, 0xff, 0x4e, 0x01, 0x00, 0x00,
}
