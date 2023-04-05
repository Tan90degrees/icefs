//
// @Author: Tan90degrees tangentninetydegrees@gmail.com
// @Date: 2023-03-30 04:19:29
// @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
// @LastEditTime: 2023-04-04 16:13:19
// @FilePath: /icefs/protos/lowlevel/icefsServices.proto
// @Description:
//
// Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: icefsServices.proto

package icefsrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_icefsServices_proto protoreflect.FileDescriptor

var file_icefsServices_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x63, 0x65, 0x66, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x69, 0x63, 0x65, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x86, 0x13, 0x0a, 0x05, 0x49, 0x63, 0x65, 0x66, 0x73, 0x12, 0x2d, 0x0a, 0x0b,
	0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x0d, 0x2e, 0x49, 0x63,
	0x65, 0x66, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x44,
	0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x10, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x6f,
	0x6f, 0x6b, 0x55, 0x70, 0x12, 0x0f, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x6f, 0x6f, 0x6b,
	0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x6f, 0x6f,
	0x6b, 0x55, 0x70, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x44, 0x6f, 0x49, 0x63,
	0x65, 0x66, 0x73, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x49, 0x63, 0x65, 0x66,
	0x73, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0e, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12,
	0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73,
	0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53,
	0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66,
	0x73, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x0f, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x6f, 0x49, 0x63,
	0x65, 0x66, 0x73, 0x4d, 0x6b, 0x6e, 0x6f, 0x64, 0x12, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73,
	0x4d, 0x6b, 0x6e, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73,
	0x4d, 0x6b, 0x6e, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x6f,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x4d, 0x6b, 0x44, 0x69, 0x72, 0x12, 0x0e, 0x2e, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x4d, 0x6b, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x4d, 0x6b, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d,
	0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x0f, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x6d, 0x44, 0x69,
	0x72, 0x12, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x6d, 0x44, 0x69, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x6d, 0x44, 0x69, 0x72, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x79,
	0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x79, 0x6d,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53,
	0x79, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x44,
	0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x49,
	0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x0b, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x0d, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0d,
	0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x0b, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x0d,
	0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2d,
	0x0a, 0x0b, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0d, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x49,
	0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x0c, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x0e, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x0c, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x12,
	0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0e, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x6f, 0x49,
	0x63, 0x65, 0x66, 0x73, 0x46, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66,
	0x73, 0x46, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66,
	0x73, 0x46, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x44,
	0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x69, 0x72, 0x12, 0x10, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x69, 0x72, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x44, 0x69, 0x72, 0x12, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61,
	0x64, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x11, 0x44,
	0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x69, 0x72,
	0x12, 0x13, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x69, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f,
	0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x69, 0x72, 0x12,
	0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x69, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x73, 0x79, 0x6e, 0x63, 0x44,
	0x69, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x44, 0x6f, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x53, 0x74, 0x61, 0x74, 0x46, 0x53, 0x12, 0x0f, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x46, 0x53, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x49, 0x63, 0x65, 0x66,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x46, 0x53, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f,
	0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x65, 0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x12,
	0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x65, 0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x65, 0x74, 0x58, 0x61, 0x74,
	0x74, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x44, 0x6f, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x47, 0x65, 0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x12, 0x11, 0x2e, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x47, 0x65, 0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x47, 0x65, 0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x10, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x12, 0x12, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x58, 0x61, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x12, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x58, 0x61, 0x74, 0x74, 0x72, 0x12, 0x14, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x58, 0x61, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x49,
	0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x58, 0x61, 0x74, 0x74, 0x72, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x44, 0x6f, 0x49,
	0x63, 0x65, 0x66, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x49, 0x63,
	0x65, 0x66, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x0c, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x47, 0x65, 0x74, 0x4c, 0x6b, 0x12, 0x0e,
	0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x47, 0x65, 0x74, 0x4c, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x47, 0x65, 0x74, 0x4c, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x0c, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x65, 0x74, 0x4c, 0x6b,
	0x12, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x65, 0x74, 0x4c, 0x6b, 0x52, 0x65, 0x71,
	0x1a, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x53, 0x65, 0x74, 0x4c, 0x6b, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x42, 0x6d, 0x61,
	0x70, 0x12, 0x0d, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x42, 0x6d, 0x61, 0x70, 0x52, 0x65, 0x71,
	0x1a, 0x0d, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x42, 0x6d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x49, 0x6f, 0x63, 0x74,
	0x6c, 0x12, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x49, 0x6f, 0x63, 0x74, 0x6c, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x49, 0x6f, 0x63, 0x74, 0x6c, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x50, 0x6f,
	0x6c, 0x6c, 0x12, 0x0d, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x1a, 0x0d, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x42, 0x75, 0x66, 0x12, 0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x42, 0x75, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x75, 0x66, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x14, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x12, 0x44, 0x6f, 0x49, 0x63, 0x65,
	0x66, 0x73, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x14, 0x2e,
	0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x6f, 0x72, 0x67, 0x65,
	0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x44,
	0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x2e, 0x49, 0x63,
	0x65, 0x66, 0x73, 0x46, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x49, 0x63,
	0x65, 0x66, 0x73, 0x46, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x10, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x46, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x12, 0x44,
	0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x44, 0x69, 0x72, 0x50, 0x6c, 0x75,
	0x73, 0x12, 0x14, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52, 0x65, 0x61, 0x64, 0x44, 0x69, 0x72,
	0x50, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x44, 0x69, 0x72, 0x50, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x48, 0x0a, 0x14, 0x44, 0x6f, 0x49, 0x63, 0x65, 0x66, 0x73, 0x43, 0x6f, 0x70, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x43,
	0x6f, 0x70, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x49, 0x63, 0x65, 0x66, 0x73, 0x43, 0x6f, 0x70, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x6f, 0x49,
	0x63, 0x65, 0x66, 0x73, 0x4c, 0x73, 0x65, 0x65, 0x6b, 0x12, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66,
	0x73, 0x4c, 0x73, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x49, 0x63, 0x65, 0x66,
	0x73, 0x4c, 0x73, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x3b, 0x69, 0x63, 0x65, 0x66, 0x73, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_icefsServices_proto_goTypes = []interface{}{
	(*IcefsInitReq)(nil),          // 0: IcefsInitReq
	(*IcefsDestroyReq)(nil),       // 1: IcefsDestroyReq
	(*IcefsLookUpReq)(nil),        // 2: IcefsLookUpReq
	(*IcefsForgetReq)(nil),        // 3: IcefsForgetReq
	(*IcefsGetAttrReq)(nil),       // 4: IcefsGetAttrReq
	(*IcefsSetAttrReq)(nil),       // 5: IcefsSetAttrReq
	(*IcefsReadLinkReq)(nil),      // 6: IcefsReadLinkReq
	(*IcefsMknodReq)(nil),         // 7: IcefsMknodReq
	(*IcefsMkDirReq)(nil),         // 8: IcefsMkDirReq
	(*IcefsUnlinkReq)(nil),        // 9: IcefsUnlinkReq
	(*IcefsRmDirReq)(nil),         // 10: IcefsRmDirReq
	(*IcefsSymLinkReq)(nil),       // 11: IcefsSymLinkReq
	(*IcefsRenameReq)(nil),        // 12: IcefsRenameReq
	(*IcefsLinkReq)(nil),          // 13: IcefsLinkReq
	(*IcefsOpenReq)(nil),          // 14: IcefsOpenReq
	(*IcefsReadReq)(nil),          // 15: IcefsReadReq
	(*IcefsWriteReq)(nil),         // 16: IcefsWriteReq
	(*IcefsFlushReq)(nil),         // 17: IcefsFlushReq
	(*IcefsReleaseReq)(nil),       // 18: IcefsReleaseReq
	(*IcefsFsyncReq)(nil),         // 19: IcefsFsyncReq
	(*IcefsOpenDirReq)(nil),       // 20: IcefsOpenDirReq
	(*IcefsReadDirReq)(nil),       // 21: IcefsReadDirReq
	(*IcefsReleaseDirReq)(nil),    // 22: IcefsReleaseDirReq
	(*IcefsFsyncDirReq)(nil),      // 23: IcefsFsyncDirReq
	(*IcefsStatFSReq)(nil),        // 24: IcefsStatFSReq
	(*IcefsSetXattrReq)(nil),      // 25: IcefsSetXattrReq
	(*IcefsGetXattrReq)(nil),      // 26: IcefsGetXattrReq
	(*IcefsListXattrReq)(nil),     // 27: IcefsListXattrReq
	(*IcefsRemoveXattrReq)(nil),   // 28: IcefsRemoveXattrReq
	(*IcefsAccessReq)(nil),        // 29: IcefsAccessReq
	(*IcefsCreateReq)(nil),        // 30: IcefsCreateReq
	(*IcefsGetLkReq)(nil),         // 31: IcefsGetLkReq
	(*IcefsSetLkReq)(nil),         // 32: IcefsSetLkReq
	(*IcefsBmapReq)(nil),          // 33: IcefsBmapReq
	(*IcefsIoctlReq)(nil),         // 34: IcefsIoctlReq
	(*IcefsPollReq)(nil),          // 35: IcefsPollReq
	(*IcefsWriteBufReq)(nil),      // 36: IcefsWriteBufReq
	(*IcefsRetrieveReplyReq)(nil), // 37: IcefsRetrieveReplyReq
	(*IcefsForgetMultiReq)(nil),   // 38: IcefsForgetMultiReq
	(*IcefsFlockReq)(nil),         // 39: IcefsFlockReq
	(*IcefsFallocateReq)(nil),     // 40: IcefsFallocateReq
	(*IcefsReadDirPlusReq)(nil),   // 41: IcefsReadDirPlusReq
	(*IcefsCopyFileRangeReq)(nil), // 42: IcefsCopyFileRangeReq
	(*IcefsLseekReq)(nil),         // 43: IcefsLseekReq
	(*IcefsInitRes)(nil),          // 44: IcefsInitRes
	(*IcefsDestroyRes)(nil),       // 45: IcefsDestroyRes
	(*IcefsLookUpRes)(nil),        // 46: IcefsLookUpRes
	(*IcefsForgetRes)(nil),        // 47: IcefsForgetRes
	(*IcefsGetAttrRes)(nil),       // 48: IcefsGetAttrRes
	(*IcefsSetAttrRes)(nil),       // 49: IcefsSetAttrRes
	(*IcefsReadLinkRes)(nil),      // 50: IcefsReadLinkRes
	(*IcefsMknodRes)(nil),         // 51: IcefsMknodRes
	(*IcefsMkDirRes)(nil),         // 52: IcefsMkDirRes
	(*IcefsUnlinkRes)(nil),        // 53: IcefsUnlinkRes
	(*IcefsRmDirRes)(nil),         // 54: IcefsRmDirRes
	(*IcefsSymLinkRes)(nil),       // 55: IcefsSymLinkRes
	(*IcefsRenameRes)(nil),        // 56: IcefsRenameRes
	(*IcefsLinkRes)(nil),          // 57: IcefsLinkRes
	(*IcefsOpenRes)(nil),          // 58: IcefsOpenRes
	(*IcefsReadRes)(nil),          // 59: IcefsReadRes
	(*IcefsWriteRes)(nil),         // 60: IcefsWriteRes
	(*IcefsFlushRes)(nil),         // 61: IcefsFlushRes
	(*IcefsReleaseRes)(nil),       // 62: IcefsReleaseRes
	(*IcefsFsyncRes)(nil),         // 63: IcefsFsyncRes
	(*IcefsOpenDirRes)(nil),       // 64: IcefsOpenDirRes
	(*IcefsReadDirRes)(nil),       // 65: IcefsReadDirRes
	(*IcefsReleaseDirRes)(nil),    // 66: IcefsReleaseDirRes
	(*IcefsFsyncDirRes)(nil),      // 67: IcefsFsyncDirRes
	(*IcefsStatFSRes)(nil),        // 68: IcefsStatFSRes
	(*IcefsSetXattrRes)(nil),      // 69: IcefsSetXattrRes
	(*IcefsGetXattrRes)(nil),      // 70: IcefsGetXattrRes
	(*IcefsListXattrRes)(nil),     // 71: IcefsListXattrRes
	(*IcefsRemoveXattrRes)(nil),   // 72: IcefsRemoveXattrRes
	(*IcefsAccessRes)(nil),        // 73: IcefsAccessRes
	(*IcefsCreateRes)(nil),        // 74: IcefsCreateRes
	(*IcefsGetLkRes)(nil),         // 75: IcefsGetLkRes
	(*IcefsSetLkRes)(nil),         // 76: IcefsSetLkRes
	(*IcefsBmapRes)(nil),          // 77: IcefsBmapRes
	(*IcefsIoctlRes)(nil),         // 78: IcefsIoctlRes
	(*IcefsPollRes)(nil),          // 79: IcefsPollRes
	(*IcefsWriteBufRes)(nil),      // 80: IcefsWriteBufRes
	(*IcefsRetrieveReplyRes)(nil), // 81: IcefsRetrieveReplyRes
	(*IcefsForgetMultiRes)(nil),   // 82: IcefsForgetMultiRes
	(*IcefsFlockRes)(nil),         // 83: IcefsFlockRes
	(*IcefsFallocateRes)(nil),     // 84: IcefsFallocateRes
	(*IcefsReadDirPlusRes)(nil),   // 85: IcefsReadDirPlusRes
	(*IcefsCopyFileRangeRes)(nil), // 86: IcefsCopyFileRangeRes
	(*IcefsLseekRes)(nil),         // 87: IcefsLseekRes
}
var file_icefsServices_proto_depIdxs = []int32{
	0,  // 0: Icefs.DoIcefsInit:input_type -> IcefsInitReq
	1,  // 1: Icefs.DoIcefsDestroy:input_type -> IcefsDestroyReq
	2,  // 2: Icefs.DoIcefsLookUp:input_type -> IcefsLookUpReq
	3,  // 3: Icefs.DoIcefsForget:input_type -> IcefsForgetReq
	4,  // 4: Icefs.DoIcefsGetAttr:input_type -> IcefsGetAttrReq
	5,  // 5: Icefs.DoIcefsSetAttr:input_type -> IcefsSetAttrReq
	6,  // 6: Icefs.DoIcefsReadLink:input_type -> IcefsReadLinkReq
	7,  // 7: Icefs.DoIcefsMknod:input_type -> IcefsMknodReq
	8,  // 8: Icefs.DoIcefsMkDir:input_type -> IcefsMkDirReq
	9,  // 9: Icefs.DoIcefsUnlink:input_type -> IcefsUnlinkReq
	10, // 10: Icefs.DoIcefsRmDir:input_type -> IcefsRmDirReq
	11, // 11: Icefs.DoIcefsSymLink:input_type -> IcefsSymLinkReq
	12, // 12: Icefs.DoIcefsRename:input_type -> IcefsRenameReq
	13, // 13: Icefs.DoIcefsLink:input_type -> IcefsLinkReq
	14, // 14: Icefs.DoIcefsOpen:input_type -> IcefsOpenReq
	15, // 15: Icefs.DoIcefsRead:input_type -> IcefsReadReq
	16, // 16: Icefs.DoIcefsWrite:input_type -> IcefsWriteReq
	17, // 17: Icefs.DoIcefsFlush:input_type -> IcefsFlushReq
	18, // 18: Icefs.DoIcefsRelease:input_type -> IcefsReleaseReq
	19, // 19: Icefs.DoIcefsFsync:input_type -> IcefsFsyncReq
	20, // 20: Icefs.DoIcefsOpenDir:input_type -> IcefsOpenDirReq
	21, // 21: Icefs.DoIcefsReadDir:input_type -> IcefsReadDirReq
	22, // 22: Icefs.DoIcefsReleaseDir:input_type -> IcefsReleaseDirReq
	23, // 23: Icefs.DoIcefsFsyncDir:input_type -> IcefsFsyncDirReq
	24, // 24: Icefs.DoIcefsStatFS:input_type -> IcefsStatFSReq
	25, // 25: Icefs.DoIcefsSetXattr:input_type -> IcefsSetXattrReq
	26, // 26: Icefs.DoIcefsGetXattr:input_type -> IcefsGetXattrReq
	27, // 27: Icefs.DoIcefsListXattr:input_type -> IcefsListXattrReq
	28, // 28: Icefs.DoIcefsRemoveXattr:input_type -> IcefsRemoveXattrReq
	29, // 29: Icefs.DoIcefsAccess:input_type -> IcefsAccessReq
	30, // 30: Icefs.DoIcefsCreate:input_type -> IcefsCreateReq
	31, // 31: Icefs.DoIcefsGetLk:input_type -> IcefsGetLkReq
	32, // 32: Icefs.DoIcefsSetLk:input_type -> IcefsSetLkReq
	33, // 33: Icefs.DoIcefsBmap:input_type -> IcefsBmapReq
	34, // 34: Icefs.DoIcefsIoctl:input_type -> IcefsIoctlReq
	35, // 35: Icefs.DoIcefsPoll:input_type -> IcefsPollReq
	36, // 36: Icefs.DoIcefsWriteBuf:input_type -> IcefsWriteBufReq
	37, // 37: Icefs.DoIcefsRetrieveReply:input_type -> IcefsRetrieveReplyReq
	38, // 38: Icefs.DoIcefsForgetMulti:input_type -> IcefsForgetMultiReq
	39, // 39: Icefs.DoIcefsFlock:input_type -> IcefsFlockReq
	40, // 40: Icefs.DoIcefsFallocate:input_type -> IcefsFallocateReq
	41, // 41: Icefs.DoIcefsReadDirPlus:input_type -> IcefsReadDirPlusReq
	42, // 42: Icefs.DoIcefsCopyFileRange:input_type -> IcefsCopyFileRangeReq
	43, // 43: Icefs.DoIcefsLseek:input_type -> IcefsLseekReq
	44, // 44: Icefs.DoIcefsInit:output_type -> IcefsInitRes
	45, // 45: Icefs.DoIcefsDestroy:output_type -> IcefsDestroyRes
	46, // 46: Icefs.DoIcefsLookUp:output_type -> IcefsLookUpRes
	47, // 47: Icefs.DoIcefsForget:output_type -> IcefsForgetRes
	48, // 48: Icefs.DoIcefsGetAttr:output_type -> IcefsGetAttrRes
	49, // 49: Icefs.DoIcefsSetAttr:output_type -> IcefsSetAttrRes
	50, // 50: Icefs.DoIcefsReadLink:output_type -> IcefsReadLinkRes
	51, // 51: Icefs.DoIcefsMknod:output_type -> IcefsMknodRes
	52, // 52: Icefs.DoIcefsMkDir:output_type -> IcefsMkDirRes
	53, // 53: Icefs.DoIcefsUnlink:output_type -> IcefsUnlinkRes
	54, // 54: Icefs.DoIcefsRmDir:output_type -> IcefsRmDirRes
	55, // 55: Icefs.DoIcefsSymLink:output_type -> IcefsSymLinkRes
	56, // 56: Icefs.DoIcefsRename:output_type -> IcefsRenameRes
	57, // 57: Icefs.DoIcefsLink:output_type -> IcefsLinkRes
	58, // 58: Icefs.DoIcefsOpen:output_type -> IcefsOpenRes
	59, // 59: Icefs.DoIcefsRead:output_type -> IcefsReadRes
	60, // 60: Icefs.DoIcefsWrite:output_type -> IcefsWriteRes
	61, // 61: Icefs.DoIcefsFlush:output_type -> IcefsFlushRes
	62, // 62: Icefs.DoIcefsRelease:output_type -> IcefsReleaseRes
	63, // 63: Icefs.DoIcefsFsync:output_type -> IcefsFsyncRes
	64, // 64: Icefs.DoIcefsOpenDir:output_type -> IcefsOpenDirRes
	65, // 65: Icefs.DoIcefsReadDir:output_type -> IcefsReadDirRes
	66, // 66: Icefs.DoIcefsReleaseDir:output_type -> IcefsReleaseDirRes
	67, // 67: Icefs.DoIcefsFsyncDir:output_type -> IcefsFsyncDirRes
	68, // 68: Icefs.DoIcefsStatFS:output_type -> IcefsStatFSRes
	69, // 69: Icefs.DoIcefsSetXattr:output_type -> IcefsSetXattrRes
	70, // 70: Icefs.DoIcefsGetXattr:output_type -> IcefsGetXattrRes
	71, // 71: Icefs.DoIcefsListXattr:output_type -> IcefsListXattrRes
	72, // 72: Icefs.DoIcefsRemoveXattr:output_type -> IcefsRemoveXattrRes
	73, // 73: Icefs.DoIcefsAccess:output_type -> IcefsAccessRes
	74, // 74: Icefs.DoIcefsCreate:output_type -> IcefsCreateRes
	75, // 75: Icefs.DoIcefsGetLk:output_type -> IcefsGetLkRes
	76, // 76: Icefs.DoIcefsSetLk:output_type -> IcefsSetLkRes
	77, // 77: Icefs.DoIcefsBmap:output_type -> IcefsBmapRes
	78, // 78: Icefs.DoIcefsIoctl:output_type -> IcefsIoctlRes
	79, // 79: Icefs.DoIcefsPoll:output_type -> IcefsPollRes
	80, // 80: Icefs.DoIcefsWriteBuf:output_type -> IcefsWriteBufRes
	81, // 81: Icefs.DoIcefsRetrieveReply:output_type -> IcefsRetrieveReplyRes
	82, // 82: Icefs.DoIcefsForgetMulti:output_type -> IcefsForgetMultiRes
	83, // 83: Icefs.DoIcefsFlock:output_type -> IcefsFlockRes
	84, // 84: Icefs.DoIcefsFallocate:output_type -> IcefsFallocateRes
	85, // 85: Icefs.DoIcefsReadDirPlus:output_type -> IcefsReadDirPlusRes
	86, // 86: Icefs.DoIcefsCopyFileRange:output_type -> IcefsCopyFileRangeRes
	87, // 87: Icefs.DoIcefsLseek:output_type -> IcefsLseekRes
	44, // [44:88] is the sub-list for method output_type
	0,  // [0:44] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_icefsServices_proto_init() }
func file_icefsServices_proto_init() {
	if File_icefsServices_proto != nil {
		return
	}
	file_icefs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_icefsServices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_icefsServices_proto_goTypes,
		DependencyIndexes: file_icefsServices_proto_depIdxs,
	}.Build()
	File_icefsServices_proto = out.File
	file_icefsServices_proto_rawDesc = nil
	file_icefsServices_proto_goTypes = nil
	file_icefsServices_proto_depIdxs = nil
}
