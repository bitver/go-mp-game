// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: internal/protos/events.proto

package protos

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

type EventType int32

const (
	EventType_init        EventType = 0
	EventType_connect     EventType = 1
	EventType_exit        EventType = 2
	EventType_move        EventType = 3
	EventType_empty       EventType = 4
	EventType_state       EventType = 5
	EventType_state_patch EventType = 6
	EventType_cast        EventType = 7
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "init",
		1: "connect",
		2: "exit",
		3: "move",
		4: "empty",
		5: "state",
		6: "state_patch",
		7: "cast",
	}
	EventType_value = map[string]int32{
		"init":        0,
		"connect":     1,
		"exit":        2,
		"move":        3,
		"empty":       4,
		"state":       5,
		"state_patch": 6,
		"cast":        7,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_events_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_internal_protos_events_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=protos.EventType" json:"type,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Event_Connect
	//	*Event_Move
	//	*Event_State
	//	*Event_StatePatch
	//	*Event_Cast
	Data     isEvent_Data `protobuf_oneof:"data"`
	PlayerId uint32       `protobuf:"varint,7,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_init
}

func (m *Event) GetData() isEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Event) GetConnect() *EventConnect {
	if x, ok := x.GetData().(*Event_Connect); ok {
		return x.Connect
	}
	return nil
}

func (x *Event) GetMove() *EventMove {
	if x, ok := x.GetData().(*Event_Move); ok {
		return x.Move
	}
	return nil
}

func (x *Event) GetState() *GameState {
	if x, ok := x.GetData().(*Event_State); ok {
		return x.State
	}
	return nil
}

func (x *Event) GetStatePatch() *GameStatePatch {
	if x, ok := x.GetData().(*Event_StatePatch); ok {
		return x.StatePatch
	}
	return nil
}

func (x *Event) GetCast() *EventCast {
	if x, ok := x.GetData().(*Event_Cast); ok {
		return x.Cast
	}
	return nil
}

func (x *Event) GetPlayerId() uint32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

type isEvent_Data interface {
	isEvent_Data()
}

type Event_Connect struct {
	Connect *EventConnect `protobuf:"bytes,2,opt,name=connect,proto3,oneof"`
}

type Event_Move struct {
	Move *EventMove `protobuf:"bytes,3,opt,name=move,proto3,oneof"`
}

type Event_State struct {
	State *GameState `protobuf:"bytes,4,opt,name=state,proto3,oneof"`
}

type Event_StatePatch struct {
	StatePatch *GameStatePatch `protobuf:"bytes,5,opt,name=state_patch,json=statePatch,proto3,oneof"`
}

type Event_Cast struct {
	Cast *EventCast `protobuf:"bytes,6,opt,name=cast,proto3,oneof"`
}

func (*Event_Connect) isEvent_Data() {}

func (*Event_Move) isEvent_Data() {}

func (*Event_State) isEvent_Data() {}

func (*Event_StatePatch) isEvent_Data() {}

func (*Event_Cast) isEvent_Data() {}

type EventConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit *Unit `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *EventConnect) Reset() {
	*x = EventConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventConnect) ProtoMessage() {}

func (x *EventConnect) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventConnect.ProtoReflect.Descriptor instead.
func (*EventConnect) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{1}
}

func (x *EventConnect) GetUnit() *Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

type EventCast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityId uint32 `protobuf:"varint,1,opt,name=ability_id,json=abilityId,proto3" json:"ability_id,omitempty"`
}

func (x *EventCast) Reset() {
	*x = EventCast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCast) ProtoMessage() {}

func (x *EventCast) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventCast.ProtoReflect.Descriptor instead.
func (*EventCast) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{2}
}

func (x *EventCast) GetAbilityId() uint32 {
	if x != nil {
		return x.AbilityId
	}
	return 0
}

type EventMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction *Vector2 `protobuf:"bytes,1,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *EventMove) Reset() {
	*x = EventMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMove) ProtoMessage() {}

func (x *EventMove) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMove.ProtoReflect.Descriptor instead.
func (*EventMove) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{3}
}

func (x *EventMove) GetDirection() *Vector2 {
	if x != nil {
		return x.Direction
	}
	return nil
}

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities map[uint32]*NetworkEntity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{4}
}

func (x *GameState) GetEntities() map[uint32]*NetworkEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type GameStatePatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities        map[uint32]*PatchNetworkEntity `protobuf:"bytes,7,rep,name=entities,proto3" json:"entities,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedEntities map[uint32]*NetworkEntity      `protobuf:"bytes,8,rep,name=created_entities,json=createdEntities,proto3" json:"created_entities,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DeletedEntities map[uint32]*Empty              `protobuf:"bytes,9,rep,name=deleted_entities,json=deletedEntities,proto3" json:"deleted_entities,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GameStatePatch) Reset() {
	*x = GameStatePatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatePatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatePatch) ProtoMessage() {}

func (x *GameStatePatch) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatePatch.ProtoReflect.Descriptor instead.
func (*GameStatePatch) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{5}
}

func (x *GameStatePatch) GetEntities() map[uint32]*PatchNetworkEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *GameStatePatch) GetCreatedEntities() map[uint32]*NetworkEntity {
	if x != nil {
		return x.CreatedEntities
	}
	return nil
}

func (x *GameStatePatch) GetDeletedEntities() map[uint32]*Empty {
	if x != nil {
		return x.DeletedEntities
	}
	return nil
}

var File_internal_protos_events_proto protoreflect.FileDescriptor

var file_internal_protos_events_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x6f, 0x76,
	0x65, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x63, 0x61, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x42, 0x06, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x6e, 0x69,
	0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x2a, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65,
	0x12, 0x2d, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x32, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x9c, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x52, 0x0a, 0x0d, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x89,
	0x04, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x40, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x10, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x1a, 0x57, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x59, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x67, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x65, 0x78, 0x69, 0x74, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x10, 0x04, 0x12, 0x09, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x63, 0x61, 0x73,
	0x74, 0x10, 0x07, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protos_events_proto_rawDescOnce sync.Once
	file_internal_protos_events_proto_rawDescData = file_internal_protos_events_proto_rawDesc
)

func file_internal_protos_events_proto_rawDescGZIP() []byte {
	file_internal_protos_events_proto_rawDescOnce.Do(func() {
		file_internal_protos_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protos_events_proto_rawDescData)
	})
	return file_internal_protos_events_proto_rawDescData
}

var file_internal_protos_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_protos_events_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_protos_events_proto_goTypes = []any{
	(EventType)(0),             // 0: protos.EventType
	(*Event)(nil),              // 1: protos.Event
	(*EventConnect)(nil),       // 2: protos.EventConnect
	(*EventCast)(nil),          // 3: protos.EventCast
	(*EventMove)(nil),          // 4: protos.EventMove
	(*GameState)(nil),          // 5: protos.GameState
	(*GameStatePatch)(nil),     // 6: protos.GameStatePatch
	nil,                        // 7: protos.GameState.EntitiesEntry
	nil,                        // 8: protos.GameStatePatch.EntitiesEntry
	nil,                        // 9: protos.GameStatePatch.CreatedEntitiesEntry
	nil,                        // 10: protos.GameStatePatch.DeletedEntitiesEntry
	(*Unit)(nil),               // 11: protos.Unit
	(*Vector2)(nil),            // 12: protos.Vector2
	(*NetworkEntity)(nil),      // 13: protos.NetworkEntity
	(*PatchNetworkEntity)(nil), // 14: protos.PatchNetworkEntity
	(*Empty)(nil),              // 15: protos.Empty
}
var file_internal_protos_events_proto_depIdxs = []int32{
	0,  // 0: protos.Event.type:type_name -> protos.EventType
	2,  // 1: protos.Event.connect:type_name -> protos.EventConnect
	4,  // 2: protos.Event.move:type_name -> protos.EventMove
	5,  // 3: protos.Event.state:type_name -> protos.GameState
	6,  // 4: protos.Event.state_patch:type_name -> protos.GameStatePatch
	3,  // 5: protos.Event.cast:type_name -> protos.EventCast
	11, // 6: protos.EventConnect.unit:type_name -> protos.Unit
	12, // 7: protos.EventMove.direction:type_name -> protos.Vector2
	7,  // 8: protos.GameState.entities:type_name -> protos.GameState.EntitiesEntry
	8,  // 9: protos.GameStatePatch.entities:type_name -> protos.GameStatePatch.EntitiesEntry
	9,  // 10: protos.GameStatePatch.created_entities:type_name -> protos.GameStatePatch.CreatedEntitiesEntry
	10, // 11: protos.GameStatePatch.deleted_entities:type_name -> protos.GameStatePatch.DeletedEntitiesEntry
	13, // 12: protos.GameState.EntitiesEntry.value:type_name -> protos.NetworkEntity
	14, // 13: protos.GameStatePatch.EntitiesEntry.value:type_name -> protos.PatchNetworkEntity
	13, // 14: protos.GameStatePatch.CreatedEntitiesEntry.value:type_name -> protos.NetworkEntity
	15, // 15: protos.GameStatePatch.DeletedEntitiesEntry.value:type_name -> protos.Empty
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_internal_protos_events_proto_init() }
func file_internal_protos_events_proto_init() {
	if File_internal_protos_events_proto != nil {
		return
	}
	file_internal_protos_unit_proto_init()
	file_internal_protos_utils_proto_init()
	file_internal_protos_network_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_protos_events_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
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
		file_internal_protos_events_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EventConnect); i {
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
		file_internal_protos_events_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EventCast); i {
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
		file_internal_protos_events_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EventMove); i {
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
		file_internal_protos_events_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GameState); i {
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
		file_internal_protos_events_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GameStatePatch); i {
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
	file_internal_protos_events_proto_msgTypes[0].OneofWrappers = []any{
		(*Event_Connect)(nil),
		(*Event_Move)(nil),
		(*Event_State)(nil),
		(*Event_StatePatch)(nil),
		(*Event_Cast)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_protos_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_protos_events_proto_goTypes,
		DependencyIndexes: file_internal_protos_events_proto_depIdxs,
		EnumInfos:         file_internal_protos_events_proto_enumTypes,
		MessageInfos:      file_internal_protos_events_proto_msgTypes,
	}.Build()
	File_internal_protos_events_proto = out.File
	file_internal_protos_events_proto_rawDesc = nil
	file_internal_protos_events_proto_goTypes = nil
	file_internal_protos_events_proto_depIdxs = nil
}
