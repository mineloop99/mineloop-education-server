// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: account_informationpb/account_information.proto

package account_informationpb

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

type AccountInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName        string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserBirthday    string `protobuf:"bytes,2,opt,name=user_birthday,json=userBirthday,proto3" json:"user_birthday,omitempty"`
	UserSex         string `protobuf:"bytes,3,opt,name=user_sex,json=userSex,proto3" json:"user_sex,omitempty"`
	UserPhoneNumber string `protobuf:"bytes,4,opt,name=user_phone_number,json=userPhoneNumber,proto3" json:"user_phone_number,omitempty"`
	UserEmail       string `protobuf:"bytes,5,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
}

func (x *AccountInformation) Reset() {
	*x = AccountInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_informationpb_account_information_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInformation) ProtoMessage() {}

func (x *AccountInformation) ProtoReflect() protoreflect.Message {
	mi := &file_account_informationpb_account_information_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInformation.ProtoReflect.Descriptor instead.
func (*AccountInformation) Descriptor() ([]byte, []int) {
	return file_account_informationpb_account_information_proto_rawDescGZIP(), []int{0}
}

func (x *AccountInformation) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AccountInformation) GetUserBirthday() string {
	if x != nil {
		return x.UserBirthday
	}
	return ""
}

func (x *AccountInformation) GetUserSex() string {
	if x != nil {
		return x.UserSex
	}
	return ""
}

func (x *AccountInformation) GetUserPhoneNumber() string {
	if x != nil {
		return x.UserPhoneNumber
	}
	return ""
}

func (x *AccountInformation) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

type WelcomeMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WelcomeMessageRequest) Reset() {
	*x = WelcomeMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_informationpb_account_information_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomeMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeMessageRequest) ProtoMessage() {}

func (x *WelcomeMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_informationpb_account_information_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeMessageRequest.ProtoReflect.Descriptor instead.
func (*WelcomeMessageRequest) Descriptor() ([]byte, []int) {
	return file_account_informationpb_account_information_proto_rawDescGZIP(), []int{1}
}

type WelcomeMessageRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WelcomeMessage string `protobuf:"bytes,1,opt,name=welcome_message,json=welcomeMessage,proto3" json:"welcome_message,omitempty"`
}

func (x *WelcomeMessageRespone) Reset() {
	*x = WelcomeMessageRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_informationpb_account_information_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomeMessageRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeMessageRespone) ProtoMessage() {}

func (x *WelcomeMessageRespone) ProtoReflect() protoreflect.Message {
	mi := &file_account_informationpb_account_information_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeMessageRespone.ProtoReflect.Descriptor instead.
func (*WelcomeMessageRespone) Descriptor() ([]byte, []int) {
	return file_account_informationpb_account_information_proto_rawDescGZIP(), []int{2}
}

func (x *WelcomeMessageRespone) GetWelcomeMessage() string {
	if x != nil {
		return x.WelcomeMessage
	}
	return ""
}

type FetchUserInformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchUserInformationRequest) Reset() {
	*x = FetchUserInformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_informationpb_account_information_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchUserInformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserInformationRequest) ProtoMessage() {}

func (x *FetchUserInformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_informationpb_account_information_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserInformationRequest.ProtoReflect.Descriptor instead.
func (*FetchUserInformationRequest) Descriptor() ([]byte, []int) {
	return file_account_informationpb_account_information_proto_rawDescGZIP(), []int{3}
}

type FetchUserInformationRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountInformation *AccountInformation `protobuf:"bytes,1,opt,name=account_information,json=accountInformation,proto3" json:"account_information,omitempty"`
}

func (x *FetchUserInformationRespone) Reset() {
	*x = FetchUserInformationRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_informationpb_account_information_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchUserInformationRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchUserInformationRespone) ProtoMessage() {}

func (x *FetchUserInformationRespone) ProtoReflect() protoreflect.Message {
	mi := &file_account_informationpb_account_information_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchUserInformationRespone.ProtoReflect.Descriptor instead.
func (*FetchUserInformationRespone) Descriptor() ([]byte, []int) {
	return file_account_informationpb_account_information_proto_rawDescGZIP(), []int{4}
}

func (x *FetchUserInformationRespone) GetAccountInformation() *AccountInformation {
	if x != nil {
		return x.AccountInformation
	}
	return nil
}

type EditUserInformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountInformation *AccountInformation `protobuf:"bytes,1,opt,name=account_information,json=accountInformation,proto3" json:"account_information,omitempty"`
}

func (x *EditUserInformationRequest) Reset() {
	*x = EditUserInformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_informationpb_account_information_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserInformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserInformationRequest) ProtoMessage() {}

func (x *EditUserInformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_informationpb_account_information_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserInformationRequest.ProtoReflect.Descriptor instead.
func (*EditUserInformationRequest) Descriptor() ([]byte, []int) {
	return file_account_informationpb_account_information_proto_rawDescGZIP(), []int{5}
}

func (x *EditUserInformationRequest) GetAccountInformation() *AccountInformation {
	if x != nil {
		return x.AccountInformation
	}
	return nil
}

type EditUserInformationRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditUserInformationRespone) Reset() {
	*x = EditUserInformationRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_informationpb_account_information_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserInformationRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserInformationRespone) ProtoMessage() {}

func (x *EditUserInformationRespone) ProtoReflect() protoreflect.Message {
	mi := &file_account_informationpb_account_information_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserInformationRespone.ProtoReflect.Descriptor instead.
func (*EditUserInformationRespone) Descriptor() ([]byte, []int) {
	return file_account_informationpb_account_information_proto_rawDescGZIP(), []int{6}
}

var File_account_informationpb_account_information_proto protoreflect.FileDescriptor

var file_account_informationpb_account_information_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x17, 0x0a, 0x15, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40,
	0x0a, 0x15, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x65, 0x6c, 0x63, 0x6f,
	0x6d, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x1d, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x77, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x58,
	0x0a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x1a, 0x45, 0x64, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x1c, 0x0a, 0x1a, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x32, 0xfa,
	0x02, 0x0a, 0x19, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0e,
	0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x7a, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x65, 0x12, 0x77, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_informationpb_account_information_proto_rawDescOnce sync.Once
	file_account_informationpb_account_information_proto_rawDescData = file_account_informationpb_account_information_proto_rawDesc
)

func file_account_informationpb_account_information_proto_rawDescGZIP() []byte {
	file_account_informationpb_account_information_proto_rawDescOnce.Do(func() {
		file_account_informationpb_account_information_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_informationpb_account_information_proto_rawDescData)
	})
	return file_account_informationpb_account_information_proto_rawDescData
}

var file_account_informationpb_account_information_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_account_informationpb_account_information_proto_goTypes = []interface{}{
	(*AccountInformation)(nil),          // 0: account_information.AccountInformation
	(*WelcomeMessageRequest)(nil),       // 1: account_information.WelcomeMessageRequest
	(*WelcomeMessageRespone)(nil),       // 2: account_information.WelcomeMessageRespone
	(*FetchUserInformationRequest)(nil), // 3: account_information.FetchUserInformationRequest
	(*FetchUserInformationRespone)(nil), // 4: account_information.FetchUserInformationRespone
	(*EditUserInformationRequest)(nil),  // 5: account_information.EditUserInformationRequest
	(*EditUserInformationRespone)(nil),  // 6: account_information.EditUserInformationRespone
}
var file_account_informationpb_account_information_proto_depIdxs = []int32{
	0, // 0: account_information.FetchUserInformationRespone.account_information:type_name -> account_information.AccountInformation
	0, // 1: account_information.EditUserInformationRequest.account_information:type_name -> account_information.AccountInformation
	1, // 2: account_information.AccountInformationService.WelcomeMessage:input_type -> account_information.WelcomeMessageRequest
	3, // 3: account_information.AccountInformationService.FetchUserInformation:input_type -> account_information.FetchUserInformationRequest
	5, // 4: account_information.AccountInformationService.EditUserInformation:input_type -> account_information.EditUserInformationRequest
	2, // 5: account_information.AccountInformationService.WelcomeMessage:output_type -> account_information.WelcomeMessageRespone
	4, // 6: account_information.AccountInformationService.FetchUserInformation:output_type -> account_information.FetchUserInformationRespone
	6, // 7: account_information.AccountInformationService.EditUserInformation:output_type -> account_information.EditUserInformationRespone
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_informationpb_account_information_proto_init() }
func file_account_informationpb_account_information_proto_init() {
	if File_account_informationpb_account_information_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_informationpb_account_information_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInformation); i {
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
		file_account_informationpb_account_information_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomeMessageRequest); i {
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
		file_account_informationpb_account_information_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomeMessageRespone); i {
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
		file_account_informationpb_account_information_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchUserInformationRequest); i {
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
		file_account_informationpb_account_information_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchUserInformationRespone); i {
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
		file_account_informationpb_account_information_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserInformationRequest); i {
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
		file_account_informationpb_account_information_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserInformationRespone); i {
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
			RawDescriptor: file_account_informationpb_account_information_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_informationpb_account_information_proto_goTypes,
		DependencyIndexes: file_account_informationpb_account_information_proto_depIdxs,
		MessageInfos:      file_account_informationpb_account_information_proto_msgTypes,
	}.Build()
	File_account_informationpb_account_information_proto = out.File
	file_account_informationpb_account_information_proto_rawDesc = nil
	file_account_informationpb_account_information_proto_goTypes = nil
	file_account_informationpb_account_information_proto_depIdxs = nil
}
