// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: controller/storage/auth/password/store/v1/argon2.proto

// Package store provides protobufs for storing types in the password package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Argon2Configuration is a configuration for using the argon2id key
// derivation function. It is owned by an AuthMethod.
//
// Iterations, Memory, and Threads are the cost parameters. The cost
// parameters should be increased as memory latency and CPU parallelism
// increases.
//
// For a detailed specification of Argon2 see:
// https://github.com/P-H-C/phc-winner-argon2/blob/master/argon2-specs.pdf
type Argon2Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: `gorm:"primary_key"`
	PrivateId string `protobuf:"bytes,1,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// @inject_tag: `gorm:"not_null"`
	PasswordMethodId string `protobuf:"bytes,3,opt,name=password_method_id,json=passwordMethodId,proto3" json:"password_method_id,omitempty" gorm:"not_null"`
	// Iterations is the time parameter in the Argon2 specification. It
	// specifies the number of passes over the memory. Must be > 0.
	// @inject_tag: `gorm:"default:null"`
	Iterations uint32 `protobuf:"varint,4,opt,name=iterations,proto3" json:"iterations,omitempty" gorm:"default:null"`
	// Memory is the memory parameter in the Argon2 specification. It
	// specifies the size of the memory in KiB. For example Memory=32*1024
	// sets the memory cost to ~32 MB. Must be > 0.
	// @inject_tag: `gorm:"default:null"`
	Memory uint32 `protobuf:"varint,5,opt,name=memory,proto3" json:"memory,omitempty" gorm:"default:null"`
	// Threads is the threads parameter in the Argon2 specification. It can
	// be adjusted to the number of available CPUs. Must be > 0.
	// @inject_tag: `gorm:"default:null"`
	Threads uint32 `protobuf:"varint,6,opt,name=threads,proto3" json:"threads,omitempty" gorm:"default:null"`
	// SaltLength is in bytes. Must be >= 16.
	// @inject_tag: `gorm:"default:null"`
	SaltLength uint32 `protobuf:"varint,7,opt,name=salt_length,json=saltLength,proto3" json:"salt_length,omitempty" gorm:"default:null"`
	// KeyLength is in bytes. Must be >= 16.
	// @inject_tag: `gorm:"default:null"`
	KeyLength uint32 `protobuf:"varint,8,opt,name=key_length,json=keyLength,proto3" json:"key_length,omitempty" gorm:"default:null"`
}

func (x *Argon2Configuration) Reset() {
	*x = Argon2Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Argon2Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Argon2Configuration) ProtoMessage() {}

func (x *Argon2Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Argon2Configuration.ProtoReflect.Descriptor instead.
func (*Argon2Configuration) Descriptor() ([]byte, []int) {
	return file_controller_storage_auth_password_store_v1_argon2_proto_rawDescGZIP(), []int{0}
}

func (x *Argon2Configuration) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *Argon2Configuration) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Argon2Configuration) GetPasswordMethodId() string {
	if x != nil {
		return x.PasswordMethodId
	}
	return ""
}

func (x *Argon2Configuration) GetIterations() uint32 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

func (x *Argon2Configuration) GetMemory() uint32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Argon2Configuration) GetThreads() uint32 {
	if x != nil {
		return x.Threads
	}
	return 0
}

func (x *Argon2Configuration) GetSaltLength() uint32 {
	if x != nil {
		return x.SaltLength
	}
	return 0
}

func (x *Argon2Configuration) GetKeyLength() uint32 {
	if x != nil {
		return x.KeyLength
	}
	return 0
}

type Argon2Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: `gorm:"primary_key"`
	PrivateId string `protobuf:"bytes,1,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// @inject_tag: `gorm:"not_null"`
	PasswordAccountId string `protobuf:"bytes,4,opt,name=password_account_id,json=passwordAccountId,proto3" json:"password_account_id,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"not_null"`
	PasswordConfId string `protobuf:"bytes,5,opt,name=password_conf_id,json=passwordConfId,proto3" json:"password_conf_id,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"not_null"`
	PasswordMethodId string `protobuf:"bytes,6,opt,name=password_method_id,json=passwordMethodId,proto3" json:"password_method_id,omitempty" gorm:"not_null"`
	// ct_salt is the encrypted salt which is stored in the database.
	// @inject_tag: `gorm:"column:salt;not_null" wrapping:"ct,entry_salt"`
	CtSalt []byte `protobuf:"bytes,7,opt,name=ct_salt,json=ctSalt,proto3" json:"ct_salt,omitempty" gorm:"column:salt;not_null" wrapping:"ct,entry_salt"`
	// salt is the unencrypted salt which is not stored in the database.
	// @inject_tag: `gorm:"-" wrapping:"pt,entry_salt"`
	Salt []byte `protobuf:"bytes,8,opt,name=salt,proto3" json:"salt,omitempty" gorm:"-" wrapping:"pt,entry_salt"`
	// derived_key is the derived key produced by the Argon2id key
	// derivation function.
	// @inject_tag: `gorm:"not_null"`
	DerivedKey []byte `protobuf:"bytes,9,opt,name=derived_key,json=derivedKey,proto3" json:"derived_key,omitempty" gorm:"not_null"`
	// key_version_id is the key version ID that was used for the encryption operation. It can be
	// used to identify a specific version of the key needed to decrypt the value,
	// which is useful for caching purposes.
	// @inject_tag: `gorm:"not_null"`
	KeyVersionId string `protobuf:"bytes,10,opt,name=key_version_id,json=keyVersionId,proto3" json:"key_version_id,omitempty" gorm:"not_null"`
}

func (x *Argon2Credential) Reset() {
	*x = Argon2Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Argon2Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Argon2Credential) ProtoMessage() {}

func (x *Argon2Credential) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Argon2Credential.ProtoReflect.Descriptor instead.
func (*Argon2Credential) Descriptor() ([]byte, []int) {
	return file_controller_storage_auth_password_store_v1_argon2_proto_rawDescGZIP(), []int{1}
}

func (x *Argon2Credential) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *Argon2Credential) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Argon2Credential) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Argon2Credential) GetPasswordAccountId() string {
	if x != nil {
		return x.PasswordAccountId
	}
	return ""
}

func (x *Argon2Credential) GetPasswordConfId() string {
	if x != nil {
		return x.PasswordConfId
	}
	return ""
}

func (x *Argon2Credential) GetPasswordMethodId() string {
	if x != nil {
		return x.PasswordMethodId
	}
	return ""
}

func (x *Argon2Credential) GetCtSalt() []byte {
	if x != nil {
		return x.CtSalt
	}
	return nil
}

func (x *Argon2Credential) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *Argon2Credential) GetDerivedKey() []byte {
	if x != nil {
		return x.DerivedKey
	}
	return nil
}

func (x *Argon2Credential) GetKeyVersionId() string {
	if x != nil {
		return x.KeyVersionId
	}
	return ""
}

var File_controller_storage_auth_password_store_v1_argon2_proto protoreflect.FileDescriptor

var file_controller_storage_auth_password_store_v1_argon2_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x67, 0x6f,
	0x6e, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x13, 0x41, 0x72, 0x67, 0x6f, 0x6e, 0x32, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x74,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x61, 0x6c, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x79,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6b,
	0x65, 0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0xc7, 0x03, 0x0a, 0x10, 0x41, 0x72, 0x67,
	0x6f, 0x6e, 0x32, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x74, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x63, 0x74, 0x53, 0x61, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0e,
	0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_auth_password_store_v1_argon2_proto_rawDescOnce sync.Once
	file_controller_storage_auth_password_store_v1_argon2_proto_rawDescData = file_controller_storage_auth_password_store_v1_argon2_proto_rawDesc
)

func file_controller_storage_auth_password_store_v1_argon2_proto_rawDescGZIP() []byte {
	file_controller_storage_auth_password_store_v1_argon2_proto_rawDescOnce.Do(func() {
		file_controller_storage_auth_password_store_v1_argon2_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_auth_password_store_v1_argon2_proto_rawDescData)
	})
	return file_controller_storage_auth_password_store_v1_argon2_proto_rawDescData
}

var file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_auth_password_store_v1_argon2_proto_goTypes = []interface{}{
	(*Argon2Configuration)(nil), // 0: controller.storage.auth.password.store.v1.Argon2Configuration
	(*Argon2Credential)(nil),    // 1: controller.storage.auth.password.store.v1.Argon2Credential
	(*timestamp.Timestamp)(nil), // 2: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_auth_password_store_v1_argon2_proto_depIdxs = []int32{
	2, // 0: controller.storage.auth.password.store.v1.Argon2Configuration.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 1: controller.storage.auth.password.store.v1.Argon2Credential.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 2: controller.storage.auth.password.store.v1.Argon2Credential.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_controller_storage_auth_password_store_v1_argon2_proto_init() }
func file_controller_storage_auth_password_store_v1_argon2_proto_init() {
	if File_controller_storage_auth_password_store_v1_argon2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Argon2Configuration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Argon2Credential); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_storage_auth_password_store_v1_argon2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_auth_password_store_v1_argon2_proto_goTypes,
		DependencyIndexes: file_controller_storage_auth_password_store_v1_argon2_proto_depIdxs,
		MessageInfos:      file_controller_storage_auth_password_store_v1_argon2_proto_msgTypes,
	}.Build()
	File_controller_storage_auth_password_store_v1_argon2_proto = out.File
	file_controller_storage_auth_password_store_v1_argon2_proto_rawDesc = nil
	file_controller_storage_auth_password_store_v1_argon2_proto_goTypes = nil
	file_controller_storage_auth_password_store_v1_argon2_proto_depIdxs = nil
}
