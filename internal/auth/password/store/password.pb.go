// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: controller/storage/auth/password/store/v1/password.proto

// Package store provides protobufs for storing types in the password package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	_ "github.com/hashicorp/boundary/internal/gen/controller/protooptions"
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

type AuthMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// name is optional. If set, it must be unique within scope_id.
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// description is optional.
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// The scope_id of the owning scope. Must be set.
	// @inject_tag: `gorm:"not_null"`
	ScopeId string `protobuf:"bytes,6,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// @inject_tag: `gorm:"not_null"`
	PasswordConfId string `protobuf:"bytes,8,opt,name=password_conf_id,json=passwordConfId,proto3" json:"password_conf_id,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"default:null"`
	MinLoginNameLength uint32 `protobuf:"varint,9,opt,name=min_login_name_length,json=minLoginNameLength,proto3" json:"min_login_name_length,omitempty" gorm:"default:null"`
	// @inject_tag: `gorm:"default:null"`
	MinPasswordLength uint32 `protobuf:"varint,10,opt,name=min_password_length,json=minPasswordLength,proto3" json:"min_password_length,omitempty" gorm:"default:null"`
	// is_primary_auth_method is a read-only output field which indicates if the
	// auth method is set as the scope's primary auth method.
	// @inject_tag: `gorm:"-"`
	IsPrimaryAuthMethod bool `protobuf:"varint,20,opt,name=is_primary_auth_method,json=isPrimaryAuthMethod,proto3" json:"is_primary_auth_method,omitempty" gorm:"-"`
}

func (x *AuthMethod) Reset() {
	*x = AuthMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_auth_password_store_v1_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMethod) ProtoMessage() {}

func (x *AuthMethod) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_auth_password_store_v1_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMethod.ProtoReflect.Descriptor instead.
func (*AuthMethod) Descriptor() ([]byte, []int) {
	return file_controller_storage_auth_password_store_v1_password_proto_rawDescGZIP(), []int{0}
}

func (x *AuthMethod) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *AuthMethod) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AuthMethod) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *AuthMethod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthMethod) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AuthMethod) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *AuthMethod) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AuthMethod) GetPasswordConfId() string {
	if x != nil {
		return x.PasswordConfId
	}
	return ""
}

func (x *AuthMethod) GetMinLoginNameLength() uint32 {
	if x != nil {
		return x.MinLoginNameLength
	}
	return 0
}

func (x *AuthMethod) GetMinPasswordLength() uint32 {
	if x != nil {
		return x.MinPasswordLength
	}
	return 0
}

func (x *AuthMethod) GetIsPrimaryAuthMethod() bool {
	if x != nil {
		return x.IsPrimaryAuthMethod
	}
	return false
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// name is optional. If set, it must be unique within scope_id.
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// description is optional.
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// @inject_tag: `gorm:"not_null"`
	AuthMethodId string `protobuf:"bytes,7,opt,name=auth_method_id,json=authMethodId,proto3" json:"auth_method_id,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"not_null"`
	LoginName string `protobuf:"bytes,8,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty" gorm:"not_null"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_auth_password_store_v1_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_auth_password_store_v1_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_controller_storage_auth_password_store_v1_password_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Account) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Account) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Account) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Account) GetAuthMethodId() string {
	if x != nil {
		return x.AuthMethodId
	}
	return ""
}

func (x *Account) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: `gorm:"primary_key"`
	PrivateId string `protobuf:"bytes,1,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// @inject_tag: `gorm:"not_null"`
	PasswordAccountId string `protobuf:"bytes,2,opt,name=password_account_id,json=passwordAccountId,proto3" json:"password_account_id,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"not_null"`
	PasswordConfId string `protobuf:"bytes,3,opt,name=password_conf_id,json=passwordConfId,proto3" json:"password_conf_id,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"not_null"`
	PasswordMethodId string `protobuf:"bytes,4,opt,name=password_method_id,json=passwordMethodId,proto3" json:"password_method_id,omitempty" gorm:"not_null"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_auth_password_store_v1_password_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_auth_password_store_v1_password_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_controller_storage_auth_password_store_v1_password_proto_rawDescGZIP(), []int{2}
}

func (x *Credential) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *Credential) GetPasswordAccountId() string {
	if x != nil {
		return x.PasswordAccountId
	}
	return ""
}

func (x *Credential) GetPasswordConfId() string {
	if x != nil {
		return x.PasswordConfId
	}
	return ""
}

func (x *Credential) GetPasswordMethodId() string {
	if x != nil {
		return x.PasswordMethodId
	}
	return ""
}

var File_controller_storage_auth_password_store_v1_password_proto protoreflect.FileDescriptor

var file_controller_storage_auth_password_store_v1_password_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x97, 0x05, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x4b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xc2, 0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x49, 0x64, 0x12,
	0x6d, 0x0a, 0x15, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x3a,
	0xc2, 0xdd, 0x29, 0x36, 0x0a, 0x12, 0x4d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x20, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x67,
	0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x37, 0xc2, 0xdd, 0x29,
	0x33, 0x0a, 0x11, 0x4d, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x1e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x33, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xaf, 0x03, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xc2,
	0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xdd, 0x29, 0x1a,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xc2, 0xdd, 0x29,
	0x22, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb3,
	0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x49, 0x64, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_auth_password_store_v1_password_proto_rawDescOnce sync.Once
	file_controller_storage_auth_password_store_v1_password_proto_rawDescData = file_controller_storage_auth_password_store_v1_password_proto_rawDesc
)

func file_controller_storage_auth_password_store_v1_password_proto_rawDescGZIP() []byte {
	file_controller_storage_auth_password_store_v1_password_proto_rawDescOnce.Do(func() {
		file_controller_storage_auth_password_store_v1_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_auth_password_store_v1_password_proto_rawDescData)
	})
	return file_controller_storage_auth_password_store_v1_password_proto_rawDescData
}

var file_controller_storage_auth_password_store_v1_password_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controller_storage_auth_password_store_v1_password_proto_goTypes = []interface{}{
	(*AuthMethod)(nil),          // 0: controller.storage.auth.password.store.v1.AuthMethod
	(*Account)(nil),             // 1: controller.storage.auth.password.store.v1.Account
	(*Credential)(nil),          // 2: controller.storage.auth.password.store.v1.Credential
	(*timestamp.Timestamp)(nil), // 3: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_auth_password_store_v1_password_proto_depIdxs = []int32{
	3, // 0: controller.storage.auth.password.store.v1.AuthMethod.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	3, // 1: controller.storage.auth.password.store.v1.AuthMethod.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	3, // 2: controller.storage.auth.password.store.v1.Account.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	3, // 3: controller.storage.auth.password.store.v1.Account.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_controller_storage_auth_password_store_v1_password_proto_init() }
func file_controller_storage_auth_password_store_v1_password_proto_init() {
	if File_controller_storage_auth_password_store_v1_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_auth_password_store_v1_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMethod); i {
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
		file_controller_storage_auth_password_store_v1_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_controller_storage_auth_password_store_v1_password_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
			RawDescriptor: file_controller_storage_auth_password_store_v1_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_auth_password_store_v1_password_proto_goTypes,
		DependencyIndexes: file_controller_storage_auth_password_store_v1_password_proto_depIdxs,
		MessageInfos:      file_controller_storage_auth_password_store_v1_password_proto_msgTypes,
	}.Build()
	File_controller_storage_auth_password_store_v1_password_proto = out.File
	file_controller_storage_auth_password_store_v1_password_proto_rawDesc = nil
	file_controller_storage_auth_password_store_v1_password_proto_goTypes = nil
	file_controller_storage_auth_password_store_v1_password_proto_depIdxs = nil
}
