// Code generated by protoc-gen-go.
// source: db_matchsvr.proto
// DO NOT EDIT!

package db

import proto "3p/code.google.com.protobuf/proto"
import json "encoding/json"
import math "math"
import "reflect"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SmallRankRecord struct {
	Rank             *int32  `protobuf:"varint,1,opt" json:"Rank,omitempty"`
	Id               *int32  `protobuf:"varint,2,opt" json:"Id,omitempty"`
	Val              *int32  `protobuf:"varint,3,opt" json:"Val,omitempty"`
	Name             *string `protobuf:"bytes,4,opt" json:"Name,omitempty"`
	TongName         *string `protobuf:"bytes,5,opt" json:"TongName,omitempty"`
	TongIcon         *int32  `protobuf:"varint,6,opt" json:"TongIcon,omitempty"`
	ArenaLvl         *int32  `protobuf:"varint,7,opt" json:"ArenaLvl,omitempty"`
	Camp             *int32  `protobuf:"varint,8,opt" json:"Camp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmallRankRecord) Reset()                { *m = SmallRankRecord{} }
func (m *SmallRankRecord) String() string        { return proto.CompactTextString(m) }
func (*SmallRankRecord) ProtoMessage()           {}
func (*SmallRankRecord) MessageTypeId() uint16   { return 1 }
func (*SmallRankRecord) MessageTypeName() string { return "SmallRankRecord" }

func (m *SmallRankRecord) GetRank() int32 {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return 0
}

func (m *SmallRankRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SmallRankRecord) GetVal() int32 {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return 0
}

func (m *SmallRankRecord) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SmallRankRecord) GetTongName() string {
	if m != nil && m.TongName != nil {
		return *m.TongName
	}
	return ""
}

func (m *SmallRankRecord) GetTongIcon() int32 {
	if m != nil && m.TongIcon != nil {
		return *m.TongIcon
	}
	return 0
}

func (m *SmallRankRecord) GetArenaLvl() int32 {
	if m != nil && m.ArenaLvl != nil {
		return *m.ArenaLvl
	}
	return 0
}

func (m *SmallRankRecord) GetCamp() int32 {
	if m != nil && m.Camp != nil {
		return *m.Camp
	}
	return 0
}

type CustomMatchPlayer struct {
	PlayerId         *int32  `protobuf:"varint,1,opt" json:"PlayerId,omitempty"`
	Score            *int32  `protobuf:"varint,2,opt" json:"Score,omitempty"`
	Name             *string `protobuf:"bytes,3,opt" json:"Name,omitempty"`
	TongIcon         *int32  `protobuf:"varint,4,opt" json:"TongIcon,omitempty"`
	TongName         *string `protobuf:"bytes,5,opt" json:"TongName,omitempty"`
	ArenaLvl         *int32  `protobuf:"varint,6,opt" json:"ArenaLvl,omitempty"`
	Camp             *int32  `protobuf:"varint,7,opt" json:"Camp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CustomMatchPlayer) Reset()                { *m = CustomMatchPlayer{} }
func (m *CustomMatchPlayer) String() string        { return proto.CompactTextString(m) }
func (*CustomMatchPlayer) ProtoMessage()           {}
func (*CustomMatchPlayer) MessageTypeId() uint16   { return 2 }
func (*CustomMatchPlayer) MessageTypeName() string { return "CustomMatchPlayer" }

func (m *CustomMatchPlayer) GetPlayerId() int32 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *CustomMatchPlayer) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *CustomMatchPlayer) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CustomMatchPlayer) GetTongIcon() int32 {
	if m != nil && m.TongIcon != nil {
		return *m.TongIcon
	}
	return 0
}

func (m *CustomMatchPlayer) GetTongName() string {
	if m != nil && m.TongName != nil {
		return *m.TongName
	}
	return ""
}

func (m *CustomMatchPlayer) GetArenaLvl() int32 {
	if m != nil && m.ArenaLvl != nil {
		return *m.ArenaLvl
	}
	return 0
}

func (m *CustomMatchPlayer) GetCamp() int32 {
	if m != nil && m.Camp != nil {
		return *m.Camp
	}
	return 0
}

type CustomMatchBaseInfo struct {
	MasterId         *int32  `protobuf:"varint,1,opt" json:"MasterId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt" json:"RoomType,omitempty"`
	RoomName         *string `protobuf:"bytes,3,opt" json:"RoomName,omitempty"`
	MasterName       *string `protobuf:"bytes,4,opt" json:"MasterName,omitempty"`
	CreateUnix       *int32  `protobuf:"varint,5,opt" json:"CreateUnix,omitempty"`
	MasterTongIcon   *int32  `protobuf:"varint,6,opt" json:"MasterTongIcon,omitempty"`
	MasterTongName   *string `protobuf:"bytes,7,opt" json:"MasterTongName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CustomMatchBaseInfo) Reset()                { *m = CustomMatchBaseInfo{} }
func (m *CustomMatchBaseInfo) String() string        { return proto.CompactTextString(m) }
func (*CustomMatchBaseInfo) ProtoMessage()           {}
func (*CustomMatchBaseInfo) MessageTypeId() uint16   { return 3 }
func (*CustomMatchBaseInfo) MessageTypeName() string { return "CustomMatchBaseInfo" }

func (m *CustomMatchBaseInfo) GetMasterId() int32 {
	if m != nil && m.MasterId != nil {
		return *m.MasterId
	}
	return 0
}

func (m *CustomMatchBaseInfo) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *CustomMatchBaseInfo) GetRoomName() string {
	if m != nil && m.RoomName != nil {
		return *m.RoomName
	}
	return ""
}

func (m *CustomMatchBaseInfo) GetMasterName() string {
	if m != nil && m.MasterName != nil {
		return *m.MasterName
	}
	return ""
}

func (m *CustomMatchBaseInfo) GetCreateUnix() int32 {
	if m != nil && m.CreateUnix != nil {
		return *m.CreateUnix
	}
	return 0
}

func (m *CustomMatchBaseInfo) GetMasterTongIcon() int32 {
	if m != nil && m.MasterTongIcon != nil {
		return *m.MasterTongIcon
	}
	return 0
}

func (m *CustomMatchBaseInfo) GetMasterTongName() string {
	if m != nil && m.MasterTongName != nil {
		return *m.MasterTongName
	}
	return ""
}

type CustomMatchRecords struct {
	InMatchPlayers   []*CustomMatchPlayer `protobuf:"bytes,1,rep" json:"InMatchPlayers,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *CustomMatchRecords) Reset()                { *m = CustomMatchRecords{} }
func (m *CustomMatchRecords) String() string        { return proto.CompactTextString(m) }
func (*CustomMatchRecords) ProtoMessage()           {}
func (*CustomMatchRecords) MessageTypeId() uint16   { return 4 }
func (*CustomMatchRecords) MessageTypeName() string { return "CustomMatchRecords" }

func (m *CustomMatchRecords) GetInMatchPlayers() []*CustomMatchPlayer {
	if m != nil {
		return m.InMatchPlayers
	}
	return nil
}

const ID_SmallRankRecord uint16 = 1
const ID_CustomMatchPlayer uint16 = 2
const ID_CustomMatchBaseInfo uint16 = 3
const ID_CustomMatchRecords uint16 = 4
const ID_ACK uint16 = 0xFFFF

var MessageNames = map[uint16]string{
	1:      "SmallRankRecord",
	2:      "CustomMatchPlayer",
	3:      "CustomMatchBaseInfo",
	4:      "CustomMatchRecords",
	0xFFFF: "ACK",
}
var MessageTypes = map[uint16]reflect.Type{
	1: reflect.TypeOf(SmallRankRecord{}),
	2: reflect.TypeOf(CustomMatchPlayer{}),
	3: reflect.TypeOf(CustomMatchBaseInfo{}),
	4: reflect.TypeOf(CustomMatchRecords{}),
}

func init() {
}
