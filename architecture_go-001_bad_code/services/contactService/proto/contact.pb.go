// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: contact.proto

package proto

import (
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

type Contact_Gender int32

const (
	Contact_NONE   Contact_Gender = 0
	Contact_MALE   Contact_Gender = 1
	Contact_FEMALE Contact_Gender = 2
)

// Enum value maps for Contact_Gender.
var (
	Contact_Gender_name = map[int32]string{
		0: "NONE",
		1: "MALE",
		2: "FEMALE",
	}
	Contact_Gender_value = map[string]int32{
		"NONE":   0,
		"MALE":   1,
		"FEMALE": 2,
	}
)

func (x Contact_Gender) Enum() *Contact_Gender {
	p := new(Contact_Gender)
	*p = x
	return p
}

func (x Contact_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Contact_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_contact_proto_enumTypes[0].Descriptor()
}

func (Contact_Gender) Type() protoreflect.EnumType {
	return &file_contact_proto_enumTypes[0]
}

func (x Contact_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Contact_Gender.Descriptor instead.
func (Contact_Gender) EnumDescriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0, 0}
}

// CONTACT ----------------------------------
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName     string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email        string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber  string                 `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Age          int32                  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Gender       int32                  `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	City         string                 `protobuf:"bytes,7,opt,name=city,proto3" json:"city,omitempty"`
	FileName     string                 `protobuf:"bytes,8,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	CreatedAt    int64                  `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt   int64                  `protobuf:"varint,10,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	CreatedBy    string                 `protobuf:"bytes,11,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Attempts     int32                  `protobuf:"varint,12,opt,name=attempts,proto3" json:"attempts,omitempty"`
	InGroups     string                 `protobuf:"bytes,13,opt,name=in_groups,json=inGroups,proto3" json:"in_groups,omitempty"`
	Attributes   *Contact_AttributeList `protobuf:"bytes,14,opt,name=attributes,proto3" json:"attributes,omitempty"`
	IsArchived   bool                   `protobuf:"varint,15,opt,name=is_archived,json=isArchived,proto3" json:"is_archived,omitempty"`
	Gtm          string                 `protobuf:"bytes,16,opt,name=gtm,proto3" json:"gtm,omitempty"`
	IsBlacklist  bool                   `protobuf:"varint,17,opt,name=is_blacklist,json=isBlacklist,proto3" json:"is_blacklist,omitempty"`
	Name         string                 `protobuf:"bytes,18,opt,name=name,proto3" json:"name,omitempty"`
	Surname      string                 `protobuf:"bytes,19,opt,name=surname,proto3" json:"surname,omitempty"`
	Patronymic   string                 `protobuf:"bytes,20,opt,name=patronymic,proto3" json:"patronymic,omitempty"`
	Parameter_1  string                 `protobuf:"bytes,21,opt,name=parameter_1,json=parameter1,proto3" json:"parameter_1,omitempty"`
	Parameter_2  string                 `protobuf:"bytes,22,opt,name=parameter_2,json=parameter2,proto3" json:"parameter_2,omitempty"`
	Timezone     string                 `protobuf:"bytes,23,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Error        string                 `protobuf:"bytes,24,opt,name=error,proto3" json:"error,omitempty"`
	LastCallTime int64                  `protobuf:"varint,25,opt,name=last_call_time,json=lastCallTime,proto3" json:"last_call_time,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Contact) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Contact) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Contact) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Contact) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Contact) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Contact) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Contact) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Contact) GetModifiedAt() int64 {
	if x != nil {
		return x.ModifiedAt
	}
	return 0
}

func (x *Contact) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Contact) GetAttempts() int32 {
	if x != nil {
		return x.Attempts
	}
	return 0
}

func (x *Contact) GetInGroups() string {
	if x != nil {
		return x.InGroups
	}
	return ""
}

func (x *Contact) GetAttributes() *Contact_AttributeList {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Contact) GetIsArchived() bool {
	if x != nil {
		return x.IsArchived
	}
	return false
}

func (x *Contact) GetGtm() string {
	if x != nil {
		return x.Gtm
	}
	return ""
}

func (x *Contact) GetIsBlacklist() bool {
	if x != nil {
		return x.IsBlacklist
	}
	return false
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Contact) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

func (x *Contact) GetParameter_1() string {
	if x != nil {
		return x.Parameter_1
	}
	return ""
}

func (x *Contact) GetParameter_2() string {
	if x != nil {
		return x.Parameter_2
	}
	return ""
}

func (x *Contact) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Contact) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Contact) GetLastCallTime() int64 {
	if x != nil {
		return x.LastCallTime
	}
	return 0
}

type ContactIdAndCreatedBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserUuid    string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	IsBlacklist bool   `protobuf:"varint,3,opt,name=is_blacklist,json=isBlacklist,proto3" json:"is_blacklist,omitempty"`
}

func (x *ContactIdAndCreatedBy) Reset() {
	*x = ContactIdAndCreatedBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactIdAndCreatedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactIdAndCreatedBy) ProtoMessage() {}

func (x *ContactIdAndCreatedBy) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactIdAndCreatedBy.ProtoReflect.Descriptor instead.
func (*ContactIdAndCreatedBy) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{1}
}

func (x *ContactIdAndCreatedBy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContactIdAndCreatedBy) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *ContactIdAndCreatedBy) GetIsBlacklist() bool {
	if x != nil {
		return x.IsBlacklist
	}
	return false
}

type AttributeOwnerId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AttributeOwnerId) Reset() {
	*x = AttributeOwnerId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeOwnerId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeOwnerId) ProtoMessage() {}

func (x *AttributeOwnerId) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeOwnerId.ProtoReflect.Descriptor instead.
func (*AttributeOwnerId) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{2}
}

func (x *AttributeOwnerId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ContactAttributeKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	KeyTitle   string `protobuf:"bytes,2,opt,name=key_title,json=keyTitle,proto3" json:"key_title,omitempty"`
	CreatedBy  string `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt  int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt int64  `protobuf:"varint,5,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
}

func (x *ContactAttributeKey) Reset() {
	*x = ContactAttributeKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactAttributeKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactAttributeKey) ProtoMessage() {}

func (x *ContactAttributeKey) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactAttributeKey.ProtoReflect.Descriptor instead.
func (*ContactAttributeKey) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{3}
}

func (x *ContactAttributeKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContactAttributeKey) GetKeyTitle() string {
	if x != nil {
		return x.KeyTitle
	}
	return ""
}

func (x *ContactAttributeKey) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *ContactAttributeKey) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ContactAttributeKey) GetModifiedAt() int64 {
	if x != nil {
		return x.ModifiedAt
	}
	return 0
}

type ContactAttributeKeyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*ContactAttributeKey `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ContactAttributeKeyList) Reset() {
	*x = ContactAttributeKeyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactAttributeKeyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactAttributeKeyList) ProtoMessage() {}

func (x *ContactAttributeKeyList) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactAttributeKeyList.ProtoReflect.Descriptor instead.
func (*ContactAttributeKeyList) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{4}
}

func (x *ContactAttributeKeyList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ContactAttributeKeyList) GetList() []*ContactAttributeKey {
	if x != nil {
		return x.List
	}
	return nil
}

type Contact_AttributeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Contact_AttributeList_AttributeItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Contact_AttributeList) Reset() {
	*x = Contact_AttributeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact_AttributeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact_AttributeList) ProtoMessage() {}

func (x *Contact_AttributeList) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact_AttributeList.ProtoReflect.Descriptor instead.
func (*Contact_AttributeList) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Contact_AttributeList) GetList() []*Contact_AttributeList_AttributeItem {
	if x != nil {
		return x.List
	}
	return nil
}

type Contact_AttributeList_AttributeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Contact_AttributeList_AttributeItem) Reset() {
	*x = Contact_AttributeList_AttributeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact_AttributeList_AttributeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact_AttributeList_AttributeItem) ProtoMessage() {}

func (x *Contact_AttributeList_AttributeItem) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact_AttributeList_AttributeItem.ProtoReflect.Descriptor instead.
func (*Contact_AttributeList_AttributeItem) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Contact_AttributeList_AttributeItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Contact_AttributeList_AttributeItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_contact_proto protoreflect.FileDescriptor

var file_contact_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x07, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x74,
	0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x74, 0x6d, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x31, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x31, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x32, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x32, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x86, 0x01, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x35, 0x0a, 0x0d, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x28, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x22, 0x67, 0x0a, 0x15, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5f, 0x0a, 0x17, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xf6, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contact_proto_rawDescOnce sync.Once
	file_contact_proto_rawDescData = file_contact_proto_rawDesc
)

func file_contact_proto_rawDescGZIP() []byte {
	file_contact_proto_rawDescOnce.Do(func() {
		file_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_proto_rawDescData)
	})
	return file_contact_proto_rawDescData
}

var file_contact_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_contact_proto_goTypes = []interface{}{
	(Contact_Gender)(0),                         // 0: proto.Contact.Gender
	(*Contact)(nil),                             // 1: proto.Contact
	(*ContactIdAndCreatedBy)(nil),               // 2: proto.ContactIdAndCreatedBy
	(*AttributeOwnerId)(nil),                    // 3: proto.AttributeOwnerId
	(*ContactAttributeKey)(nil),                 // 4: proto.ContactAttributeKey
	(*ContactAttributeKeyList)(nil),             // 5: proto.ContactAttributeKeyList
	(*Contact_AttributeList)(nil),               // 6: proto.Contact.AttributeList
	(*Contact_AttributeList_AttributeItem)(nil), // 7: proto.Contact.AttributeList.AttributeItem
}
var file_contact_proto_depIdxs = []int32{
	6, // 0: proto.Contact.attributes:type_name -> proto.Contact.AttributeList
	4, // 1: proto.ContactAttributeKeyList.list:type_name -> proto.ContactAttributeKey
	7, // 2: proto.Contact.AttributeList.list:type_name -> proto.Contact.AttributeList.AttributeItem
	1, // 3: proto.ContactService.Create:input_type -> proto.Contact
	1, // 4: proto.ContactService.Update:input_type -> proto.Contact
	2, // 5: proto.ContactService.ReadById:input_type -> proto.ContactIdAndCreatedBy
	3, // 6: proto.ContactService.GetAttributeKeyList:input_type -> proto.AttributeOwnerId
	1, // 7: proto.ContactService.Create:output_type -> proto.Contact
	1, // 8: proto.ContactService.Update:output_type -> proto.Contact
	1, // 9: proto.ContactService.ReadById:output_type -> proto.Contact
	5, // 10: proto.ContactService.GetAttributeKeyList:output_type -> proto.ContactAttributeKeyList
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_contact_proto_init() }
func file_contact_proto_init() {
	if File_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactIdAndCreatedBy); i {
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
		file_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeOwnerId); i {
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
		file_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactAttributeKey); i {
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
		file_contact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactAttributeKeyList); i {
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
		file_contact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact_AttributeList); i {
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
		file_contact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact_AttributeList_AttributeItem); i {
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
			RawDescriptor: file_contact_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contact_proto_goTypes,
		DependencyIndexes: file_contact_proto_depIdxs,
		EnumInfos:         file_contact_proto_enumTypes,
		MessageInfos:      file_contact_proto_msgTypes,
	}.Build()
	File_contact_proto = out.File
	file_contact_proto_rawDesc = nil
	file_contact_proto_goTypes = nil
	file_contact_proto_depIdxs = nil
}
