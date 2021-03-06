// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronization/scan_mode.proto

package synchronization

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ScanMode specifies the mode for synchronization root scanning.
type ScanMode int32

const (
	// ScanMode_ScanModeDefault represents an unspecified scan mode. It should
	// be converted to one of the following values based on the desired default
	// behavior.
	ScanMode_ScanModeDefault ScanMode = 0
	// ScanMode_ScanModeFull specifies that full scans should be performed on
	// each synchronization cycle.
	ScanMode_ScanModeFull ScanMode = 1
	// ScanMode_ScanModeAccelerated specifies that scans should attempt to use
	// watch-based acceleration.
	ScanMode_ScanModeAccelerated ScanMode = 2
)

var ScanMode_name = map[int32]string{
	0: "ScanModeDefault",
	1: "ScanModeFull",
	2: "ScanModeAccelerated",
}

var ScanMode_value = map[string]int32{
	"ScanModeDefault":     0,
	"ScanModeFull":        1,
	"ScanModeAccelerated": 2,
}

func (x ScanMode) String() string {
	return proto.EnumName(ScanMode_name, int32(x))
}

func (ScanMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f8aeecf10d3570d9, []int{0}
}

func init() {
	proto.RegisterEnum("synchronization.ScanMode", ScanMode_name, ScanMode_value)
}

func init() { proto.RegisterFile("synchronization/scan_mode.proto", fileDescriptor_f8aeecf10d3570d9) }

var fileDescriptor_f8aeecf10d3570d9 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xae, 0xcc, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x4e, 0x4e, 0xcc,
	0x8b, 0xcf, 0xcd, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x47, 0x53, 0xa0,
	0xe5, 0xc5, 0xc5, 0x11, 0x9c, 0x9c, 0x98, 0xe7, 0x9b, 0x9f, 0x92, 0x2a, 0x24, 0xcc, 0xc5, 0x0f,
	0x63, 0xbb, 0xa4, 0xa6, 0x25, 0x96, 0xe6, 0x94, 0x08, 0x30, 0x08, 0x09, 0x70, 0xf1, 0xc0, 0x04,
	0xdd, 0x4a, 0x73, 0x72, 0x04, 0x18, 0x85, 0xc4, 0xb9, 0x84, 0x61, 0x22, 0x8e, 0xc9, 0xc9, 0xa9,
	0x39, 0xa9, 0x45, 0x89, 0x25, 0xa9, 0x29, 0x02, 0x4c, 0x4e, 0xc6, 0x51, 0x86, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0xa5, 0x25, 0x89, 0xe9, 0xa9, 0x79, 0xba,
	0x99, 0xf9, 0x30, 0xa6, 0x7e, 0x41, 0x76, 0xba, 0x3e, 0x9a, 0x03, 0x92, 0xd8, 0xc0, 0x0e, 0x33,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xca, 0x33, 0x52, 0x41, 0xbb, 0x00, 0x00, 0x00,
}
