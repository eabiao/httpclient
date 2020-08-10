package main

import (
	"os"
)

var Assets = &fs{}
var assets = map[string][]byte{}

type fs struct{}

// Get Data
func (fs *fs) Get(name string) ([]byte, error) {
	data, ok := assets[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return data, nil
}

func init() {
	assets["page/index.html"] = []byte{
		0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x48, 0x74,
		0x74, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x0a, 0x20, 0x20, 0x20,
		0x20, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x65, 0x71, 0x75, 0x69, 0x76, 0x3d, 0x22, 0x43,
		0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
		0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3b, 0x20, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74,
		0x3d, 0x55, 0x54, 0x46, 0x2d, 0x38, 0x22, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20,
		0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x76, 0x69, 0x65, 0x77, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74,
		0x65, 0x6e, 0x74, 0x3d, 0x22, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x77, 0x69,
		0x64, 0x74, 0x68, 0x2c, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3d, 0x31, 0x2e,
		0x30, 0x2c, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3d, 0x31, 0x2e, 0x30, 0x2c,
		0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x3d, 0x6e, 0x6f, 0x22, 0x2f, 0x3e, 0x0a,
		0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x0a, 0x3c, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x55,
		0x52, 0x4c, 0x3c, 0x62, 0x72, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d,
		0x22, 0x75, 0x72, 0x6c, 0x44, 0x69, 0x76, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73,
		0x65, 0x6c, 0x65, 0x63, 0x74, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x6c, 0x65,
		0x63, 0x74, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6f, 0x70,
		0x74, 0x69, 0x6f, 0x6e, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x47, 0x45, 0x54, 0x22, 0x3e, 0x47, 0x45, 0x54,
		0x3c, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x3c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x50, 0x4f, 0x53,
		0x54, 0x22, 0x3e, 0x50, 0x4f, 0x53, 0x54, 0x3c, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x76, 0x61, 0x6c,
		0x75, 0x65, 0x3d, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x22, 0x3e, 0x50, 0x4f, 0x53, 0x54, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x3c,
		0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x73,
		0x65, 0x6c, 0x65, 0x63, 0x74, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75,
		0x74, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x75, 0x72, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d,
		0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x73, 0x70, 0x65, 0x6c, 0x6c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3d, 0x22, 0x66,
		0x61, 0x6c, 0x73, 0x65, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f,
		0x22, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20,
		0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x3c, 0x62, 0x72, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x65, 0x78,
		0x74, 0x61, 0x72, 0x65, 0x61, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74,
		0x22, 0x20, 0x73, 0x70, 0x65, 0x6c, 0x6c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3d, 0x22, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x22,
		0x3e, 0x3c, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x61, 0x72, 0x65, 0x61, 0x3e, 0x3c, 0x62, 0x72, 0x2f, 0x3e, 0x0a, 0x20, 0x20,
		0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x74,
		0x74, 0x6f, 0x6e, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x22, 0x20, 0x76,
		0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0xe5, 0x8f, 0x91, 0xe9, 0x80, 0x81, 0xe8, 0xaf, 0xb7, 0xe6, 0xb1, 0x82, 0x22, 0x2f,
		0x3e, 0x3c, 0x62, 0x72, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x69, 0x64,
		0x3d, 0x22, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d,
		0x22, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0xe6, 0xb8, 0x85, 0xe7,
		0xa9, 0xba, 0xe7, 0xbb, 0x93, 0xe6, 0x9e, 0x9c, 0x22, 0x2f, 0x3e, 0x3c, 0x62, 0x72, 0x2f, 0x3e, 0x0a, 0x0a, 0x20, 0x20,
		0x20, 0x20, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3c, 0x62, 0x72, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20,
		0x3c, 0x74, 0x65, 0x78, 0x74, 0x61, 0x72, 0x65, 0x61, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x72, 0x65, 0x73, 0x70, 0x54, 0x65,
		0x78, 0x74, 0x22, 0x20, 0x73, 0x70, 0x65, 0x6c, 0x6c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3d, 0x22, 0x66, 0x61, 0x6c, 0x73,
		0x65, 0x22, 0x3e, 0x3c, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x61, 0x72, 0x65, 0x61, 0x3e, 0x3c, 0x62, 0x72, 0x2f, 0x3e, 0x0a,
		0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x0a, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x20, 0x20, 0x20,
		0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
		0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x22, 0x29, 0x2e, 0x61,
		0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x22, 0x63, 0x6c, 0x69,
		0x63, 0x6b, 0x22, 0x2c, 0x20, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
		0x28, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x74, 0x20, 0x75, 0x72, 0x6c,
		0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65,
		0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x75, 0x72, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x22, 0x29, 0x2e, 0x76, 0x61,
		0x6c, 0x75, 0x65, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x69, 0x66, 0x20, 0x28, 0x75, 0x72, 0x6c,
		0x20, 0x3d, 0x3d, 0x3d, 0x20, 0x22, 0x22, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x28, 0x22, 0x75, 0x72, 0x6c, 0xe4, 0xb8, 0x8d, 0xe8, 0x83, 0xbd, 0xe4,
		0xb8, 0xba, 0xe7, 0xa9, 0xba, 0x22, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x74, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x20, 0x3d, 0x20,
		0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42,
		0x79, 0x49, 0x64, 0x28, 0x22, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x22, 0x29, 0x2e,
		0x76, 0x61, 0x6c, 0x75, 0x65, 0x3b, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x74, 0x20,
		0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
		0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x68, 0x65,
		0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x22, 0x29, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3b, 0x0a, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x74, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x72, 0x72, 0x20,
		0x3d, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x22,
		0x5c, 0x6e, 0x22, 0x29, 0x3b, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x74, 0x20, 0x68,
		0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x20, 0x3d, 0x20, 0x7b, 0x7d, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x28, 0x6c, 0x65, 0x74, 0x20, 0x69, 0x20, 0x3d, 0x20, 0x30, 0x3b, 0x20, 0x69,
		0x20, 0x3c, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x72, 0x72, 0x2e, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3b,
		0x20, 0x69, 0x2b, 0x2b, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x6c, 0x65, 0x74, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x73, 0x20, 0x3d, 0x20, 0x68, 0x65,
		0x61, 0x64, 0x65, 0x72, 0x41, 0x72, 0x72, 0x5b, 0x69, 0x5d, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x22, 0x3a, 0x22,
		0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x69, 0x66, 0x20, 0x28, 0x68,
		0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x73, 0x2e, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x3d, 0x3d,
		0x3d, 0x20, 0x32, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x5b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50,
		0x61, 0x72, 0x74, 0x73, 0x5b, 0x30, 0x5d, 0x2e, 0x74, 0x72, 0x69, 0x6d, 0x28, 0x29, 0x5d, 0x20, 0x3d, 0x20, 0x68, 0x65,
		0x61, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x73, 0x5b, 0x31, 0x5d, 0x2e, 0x74, 0x72, 0x69, 0x6d, 0x28, 0x29, 0x0a,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
		0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x72, 0x65,
		0x73, 0x70, 0x54, 0x65, 0x78, 0x74, 0x22, 0x29, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x3d, 0x20, 0x61, 0x77, 0x61,
		0x69, 0x74, 0x20, 0x67, 0x6f, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x6d, 0x65, 0x74,
		0x68, 0x6f, 0x64, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x2c, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x29,
		0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x29, 0x3b, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
		0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x6d,
		0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x22, 0x29, 0x2e, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65,
		0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x22, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x2c,
		0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x28, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x69, 0x66, 0x20, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5b,
		0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x5d, 0x2e,
		0x74, 0x65, 0x78, 0x74, 0x20, 0x3d, 0x3d, 0x3d, 0x20, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x22,
		0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x64, 0x6f, 0x63, 0x75,
		0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28,
		0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x22, 0x29, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20,
		0x3d, 0x20, 0x22, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x20, 0x61, 0x70, 0x70,
		0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3b, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65,
		0x74, 0x3d, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x20, 0x65,
		0x6c, 0x73, 0x65, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x64, 0x6f,
		0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
		0x64, 0x28, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x22, 0x29, 0x2e, 0x76, 0x61, 0x6c, 0x75,
		0x65, 0x20, 0x3d, 0x20, 0x22, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20,
		0x20, 0x7d, 0x29, 0x0a, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x0a, 0x3c, 0x73, 0x74, 0x79, 0x6c,
		0x65, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x23, 0x75, 0x72, 0x6c, 0x44, 0x69, 0x76, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x3b, 0x0a,
		0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x23, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65,
		0x6c, 0x65, 0x63, 0x74, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68,
		0x74, 0x3a, 0x20, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x6d, 0x61, 0x72,
		0x67, 0x69, 0x6e, 0x2d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x32, 0x70, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20,
		0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x23, 0x75, 0x72, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x20, 0x7b, 0x0a, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x3a, 0x20, 0x31, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20,
		0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x23, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x20,
		0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30,
		0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x38,
		0x30, 0x70, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x3a,
		0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x23, 0x73,
		0x65, 0x6e, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x36, 0x30, 0x70, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x32, 0x35, 0x70, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x3a, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3b, 0x0a, 0x20, 0x20, 0x20,
		0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x23, 0x72, 0x65, 0x73, 0x70, 0x54, 0x65, 0x78, 0x74, 0x20, 0x7b, 0x0a,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x25, 0x3b,
		0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x36, 0x30, 0x25,
		0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x6e, 0x6f,
		0x6e, 0x65, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74,
		0x2c, 0x20, 0x74, 0x65, 0x78, 0x74, 0x61, 0x72, 0x65, 0x61, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x35, 0x70, 0x78, 0x3b,
		0x0a, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x3c, 0x2f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e}
}
