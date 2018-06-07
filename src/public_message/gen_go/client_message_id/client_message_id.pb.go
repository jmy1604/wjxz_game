// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_message_id.proto

/*
Package msg_client_message_id is a generated protocol buffer package.

It is generated from these files:
	client_message_id.proto

It has these top-level messages:
*/
package msg_client_message_id

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSGID int32

const (
	MSGID_NONE                                    MSGID = 0
	MSGID_C2S_TEST_COMMAND                        MSGID = 1
	MSGID_C2S_HEARTBEAT                           MSGID = 2
	MSGID_S2C_STATE_NOTIFY                        MSGID = 3
	MSGID_C2S_DATA_SYNC_REQUEST                   MSGID = 4
	MSGID_C2S_LOGIN_REQUEST                       MSGID = 10000
	MSGID_S2C_LOGIN_RESPONSE                      MSGID = 10001
	MSGID_S2C_OTHER_PLACE_LOGIN                   MSGID = 10002
	MSGID_C2S_SELECT_SERVER_REQUEST               MSGID = 10003
	MSGID_S2C_SELECT_SERVER_RESPONSE              MSGID = 10004
	MSGID_C2S_ENTER_GAME_REQUEST                  MSGID = 11000
	MSGID_S2C_ENTER_GAME_RESPONSE                 MSGID = 11001
	MSGID_S2C_ENTER_GAME_COMPLETE_NOTIFY          MSGID = 11002
	MSGID_C2S_LEAVE_GAME_REQUEST                  MSGID = 11003
	MSGID_S2C_LEAVE_GAME_RESPONSE                 MSGID = 11004
	MSGID_S2C_PLAYER_INFO_RESPONSE                MSGID = 11005
	MSGID_C2S_ROLES_REQUEST                       MSGID = 11050
	MSGID_S2C_ROLES_RESPONSE                      MSGID = 11051
	MSGID_S2C_ROLES_CHANGE_NOTIFY                 MSGID = 11152
	MSGID_C2S_ROLE_LEVELUP_REQUEST                MSGID = 11153
	MSGID_S2C_ROLE_LEVELUP_RESPONSE               MSGID = 11154
	MSGID_C2S_ROLE_RANKUP_REQUEST                 MSGID = 11155
	MSGID_S2C_ROLE_RANKUP_RESPONSE                MSGID = 11156
	MSGID_C2S_ROLE_DECOMPOSE_REQUEST              MSGID = 11157
	MSGID_S2C_ROLE_DECOMPOSE_RESPONSE             MSGID = 11158
	MSGID_C2S_ROLE_FUSION_REQUEST                 MSGID = 11159
	MSGID_S2C_ROLE_FUSION_RESPONSE                MSGID = 11160
	MSGID_C2S_BATTLE_RESULT_REQUEST               MSGID = 10100
	MSGID_S2C_BATTLE_RESULT_RESPONSE              MSGID = 10101
	MSGID_C2S_BATTLE_RECORD_REQUEST               MSGID = 10102
	MSGID_S2C_BATTLE_RECORD_RESPONSE              MSGID = 10103
	MSGID_C2S_BATTLE_RECORD_LIST_REQUEST          MSGID = 10104
	MSGID_S2C_BATTLE_RECORD_LIST_RESPONSE         MSGID = 10105
	MSGID_C2S_BATTLE_RECORD_DELETE_REQUEST        MSGID = 10106
	MSGID_S2C_BATTLE_RECORD_DELETE_RESPONSE       MSGID = 10107
	MSGID_C2S_SET_TEAM_REQUEST                    MSGID = 10200
	MSGID_S2C_SET_TEAM_RESPONSE                   MSGID = 10201
	MSGID_S2C_TEAMS_RESPONSE                      MSGID = 10202
	MSGID_C2S_ITEMS_SYNC_REQUEST                  MSGID = 10300
	MSGID_S2C_ITEMS_SYNC                          MSGID = 10301
	MSGID_S2C_ITEMS_UPDATE                        MSGID = 10302
	MSGID_C2S_ITEM_FUSION_REQUEST                 MSGID = 10303
	MSGID_S2C_ITEM_FUSION_RESPONSE                MSGID = 10304
	MSGID_C2S_ITEM_SELL_REQUEST                   MSGID = 10305
	MSGID_S2C_ITEM_SELL_RESPONSE                  MSGID = 10306
	MSGID_C2S_CAMPAIGN_DATA_REQUEST               MSGID = 10400
	MSGID_S2C_CAMPAIGN_DATA_RESPONSE              MSGID = 10401
	MSGID_C2S_CAMPAIGN_HANGUP_INCOME_REQUEST      MSGID = 10402
	MSGID_S2C_CAMPAIGN_HANGUP_INCOME_RESPONSE     MSGID = 10403
	MSGID_C2S_BATTLE_SET_HANGUP_CAMPAIGN_REQUEST  MSGID = 10404
	MSGID_S2C_BATTLE_SET_HANGUP_CAMPAIGN_RESPONSE MSGID = 10405
	MSGID_C2S_MAIL_SEND_REQUEST                   MSGID = 10500
	MSGID_S2C_MAIL_SEND_RESPONSE                  MSGID = 10501
	MSGID_C2S_MAIL_LIST_REQUEST                   MSGID = 10502
	MSGID_S2C_MAIL_LIST_RESPONSE                  MSGID = 10503
	MSGID_C2S_MAIL_DETAIL_REQUEST                 MSGID = 10504
	MSGID_S2C_MAIL_DETAIL_RESPONSE                MSGID = 10505
	MSGID_C2S_MAIL_GET_ATTACHED_ITEMS_REQUEST     MSGID = 10506
	MSGID_S2C_MAIL_GET_ATTACHED_ITEMS_RESPONSE    MSGID = 10507
	MSGID_C2S_MAIL_DELETE_REQUEST                 MSGID = 10508
	MSGID_S2C_MAIL_DELETE_RESPONSE                MSGID = 10509
)

var MSGID_name = map[int32]string{
	0:     "NONE",
	1:     "C2S_TEST_COMMAND",
	2:     "C2S_HEARTBEAT",
	3:     "S2C_STATE_NOTIFY",
	4:     "C2S_DATA_SYNC_REQUEST",
	10000: "C2S_LOGIN_REQUEST",
	10001: "S2C_LOGIN_RESPONSE",
	10002: "S2C_OTHER_PLACE_LOGIN",
	10003: "C2S_SELECT_SERVER_REQUEST",
	10004: "S2C_SELECT_SERVER_RESPONSE",
	11000: "C2S_ENTER_GAME_REQUEST",
	11001: "S2C_ENTER_GAME_RESPONSE",
	11002: "S2C_ENTER_GAME_COMPLETE_NOTIFY",
	11003: "C2S_LEAVE_GAME_REQUEST",
	11004: "S2C_LEAVE_GAME_RESPONSE",
	11005: "S2C_PLAYER_INFO_RESPONSE",
	11050: "C2S_ROLES_REQUEST",
	11051: "S2C_ROLES_RESPONSE",
	11152: "S2C_ROLES_CHANGE_NOTIFY",
	11153: "C2S_ROLE_LEVELUP_REQUEST",
	11154: "S2C_ROLE_LEVELUP_RESPONSE",
	11155: "C2S_ROLE_RANKUP_REQUEST",
	11156: "S2C_ROLE_RANKUP_RESPONSE",
	11157: "C2S_ROLE_DECOMPOSE_REQUEST",
	11158: "S2C_ROLE_DECOMPOSE_RESPONSE",
	11159: "C2S_ROLE_FUSION_REQUEST",
	11160: "S2C_ROLE_FUSION_RESPONSE",
	10100: "C2S_BATTLE_RESULT_REQUEST",
	10101: "S2C_BATTLE_RESULT_RESPONSE",
	10102: "C2S_BATTLE_RECORD_REQUEST",
	10103: "S2C_BATTLE_RECORD_RESPONSE",
	10104: "C2S_BATTLE_RECORD_LIST_REQUEST",
	10105: "S2C_BATTLE_RECORD_LIST_RESPONSE",
	10106: "C2S_BATTLE_RECORD_DELETE_REQUEST",
	10107: "S2C_BATTLE_RECORD_DELETE_RESPONSE",
	10200: "C2S_SET_TEAM_REQUEST",
	10201: "S2C_SET_TEAM_RESPONSE",
	10202: "S2C_TEAMS_RESPONSE",
	10300: "C2S_ITEMS_SYNC_REQUEST",
	10301: "S2C_ITEMS_SYNC",
	10302: "S2C_ITEMS_UPDATE",
	10303: "C2S_ITEM_FUSION_REQUEST",
	10304: "S2C_ITEM_FUSION_RESPONSE",
	10305: "C2S_ITEM_SELL_REQUEST",
	10306: "S2C_ITEM_SELL_RESPONSE",
	10400: "C2S_CAMPAIGN_DATA_REQUEST",
	10401: "S2C_CAMPAIGN_DATA_RESPONSE",
	10402: "C2S_CAMPAIGN_HANGUP_INCOME_REQUEST",
	10403: "S2C_CAMPAIGN_HANGUP_INCOME_RESPONSE",
	10404: "C2S_BATTLE_SET_HANGUP_CAMPAIGN_REQUEST",
	10405: "S2C_BATTLE_SET_HANGUP_CAMPAIGN_RESPONSE",
	10500: "C2S_MAIL_SEND_REQUEST",
	10501: "S2C_MAIL_SEND_RESPONSE",
	10502: "C2S_MAIL_LIST_REQUEST",
	10503: "S2C_MAIL_LIST_RESPONSE",
	10504: "C2S_MAIL_DETAIL_REQUEST",
	10505: "S2C_MAIL_DETAIL_RESPONSE",
	10506: "C2S_MAIL_GET_ATTACHED_ITEMS_REQUEST",
	10507: "S2C_MAIL_GET_ATTACHED_ITEMS_RESPONSE",
	10508: "C2S_MAIL_DELETE_REQUEST",
	10509: "S2C_MAIL_DELETE_RESPONSE",
}
var MSGID_value = map[string]int32{
	"NONE":                                    0,
	"C2S_TEST_COMMAND":                        1,
	"C2S_HEARTBEAT":                           2,
	"S2C_STATE_NOTIFY":                        3,
	"C2S_DATA_SYNC_REQUEST":                   4,
	"C2S_LOGIN_REQUEST":                       10000,
	"S2C_LOGIN_RESPONSE":                      10001,
	"S2C_OTHER_PLACE_LOGIN":                   10002,
	"C2S_SELECT_SERVER_REQUEST":               10003,
	"S2C_SELECT_SERVER_RESPONSE":              10004,
	"C2S_ENTER_GAME_REQUEST":                  11000,
	"S2C_ENTER_GAME_RESPONSE":                 11001,
	"S2C_ENTER_GAME_COMPLETE_NOTIFY":          11002,
	"C2S_LEAVE_GAME_REQUEST":                  11003,
	"S2C_LEAVE_GAME_RESPONSE":                 11004,
	"S2C_PLAYER_INFO_RESPONSE":                11005,
	"C2S_ROLES_REQUEST":                       11050,
	"S2C_ROLES_RESPONSE":                      11051,
	"S2C_ROLES_CHANGE_NOTIFY":                 11152,
	"C2S_ROLE_LEVELUP_REQUEST":                11153,
	"S2C_ROLE_LEVELUP_RESPONSE":               11154,
	"C2S_ROLE_RANKUP_REQUEST":                 11155,
	"S2C_ROLE_RANKUP_RESPONSE":                11156,
	"C2S_ROLE_DECOMPOSE_REQUEST":              11157,
	"S2C_ROLE_DECOMPOSE_RESPONSE":             11158,
	"C2S_ROLE_FUSION_REQUEST":                 11159,
	"S2C_ROLE_FUSION_RESPONSE":                11160,
	"C2S_BATTLE_RESULT_REQUEST":               10100,
	"S2C_BATTLE_RESULT_RESPONSE":              10101,
	"C2S_BATTLE_RECORD_REQUEST":               10102,
	"S2C_BATTLE_RECORD_RESPONSE":              10103,
	"C2S_BATTLE_RECORD_LIST_REQUEST":          10104,
	"S2C_BATTLE_RECORD_LIST_RESPONSE":         10105,
	"C2S_BATTLE_RECORD_DELETE_REQUEST":        10106,
	"S2C_BATTLE_RECORD_DELETE_RESPONSE":       10107,
	"C2S_SET_TEAM_REQUEST":                    10200,
	"S2C_SET_TEAM_RESPONSE":                   10201,
	"S2C_TEAMS_RESPONSE":                      10202,
	"C2S_ITEMS_SYNC_REQUEST":                  10300,
	"S2C_ITEMS_SYNC":                          10301,
	"S2C_ITEMS_UPDATE":                        10302,
	"C2S_ITEM_FUSION_REQUEST":                 10303,
	"S2C_ITEM_FUSION_RESPONSE":                10304,
	"C2S_ITEM_SELL_REQUEST":                   10305,
	"S2C_ITEM_SELL_RESPONSE":                  10306,
	"C2S_CAMPAIGN_DATA_REQUEST":               10400,
	"S2C_CAMPAIGN_DATA_RESPONSE":              10401,
	"C2S_CAMPAIGN_HANGUP_INCOME_REQUEST":      10402,
	"S2C_CAMPAIGN_HANGUP_INCOME_RESPONSE":     10403,
	"C2S_BATTLE_SET_HANGUP_CAMPAIGN_REQUEST":  10404,
	"S2C_BATTLE_SET_HANGUP_CAMPAIGN_RESPONSE": 10405,
	"C2S_MAIL_SEND_REQUEST":                   10500,
	"S2C_MAIL_SEND_RESPONSE":                  10501,
	"C2S_MAIL_LIST_REQUEST":                   10502,
	"S2C_MAIL_LIST_RESPONSE":                  10503,
	"C2S_MAIL_DETAIL_REQUEST":                 10504,
	"S2C_MAIL_DETAIL_RESPONSE":                10505,
	"C2S_MAIL_GET_ATTACHED_ITEMS_REQUEST":     10506,
	"S2C_MAIL_GET_ATTACHED_ITEMS_RESPONSE":    10507,
	"C2S_MAIL_DELETE_REQUEST":                 10508,
	"S2C_MAIL_DELETE_RESPONSE":                10509,
}

func (x MSGID) String() string {
	return proto.EnumName(MSGID_name, int32(x))
}
func (MSGID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("msg.client_message_id.MSGID", MSGID_name, MSGID_value)
}

func init() { proto.RegisterFile("client_message_id.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x95, 0xd9, 0x4e, 0x1c, 0x39,
	0x14, 0x86, 0x67, 0x61, 0x46, 0x23, 0x4b, 0x83, 0x0a, 0x0f, 0x0d, 0x03, 0xcc, 0x00, 0x33, 0x30,
	0x30, 0x59, 0xc4, 0x05, 0x79, 0x02, 0x53, 0x75, 0xe8, 0x2e, 0xa5, 0xca, 0xae, 0xb6, 0x5d, 0xdd,
	0xe2, 0xca, 0xca, 0x82, 0x10, 0x52, 0x08, 0x51, 0xe0, 0x11, 0xb2, 0x2f, 0x12, 0x4b, 0xb6, 0xcb,
	0xac, 0x37, 0xc9, 0x2b, 0x64, 0x7f, 0x8a, 0x90, 0xe7, 0xc8, 0x42, 0x20, 0x91, 0x22, 0x77, 0xdb,
	0xee, 0x72, 0x35, 0x57, 0x2d, 0x9d, 0xff, 0xfc, 0x9f, 0x4f, 0x9f, 0x73, 0xec, 0x42, 0x83, 0xa7,
	0xce, 0x2c, 0x2f, 0x9e, 0x5d, 0x57, 0x2b, 0x8b, 0x6b, 0x6b, 0x27, 0x96, 0x16, 0xd5, 0xf2, 0xe9,
	0x99, 0x73, 0xe7, 0x57, 0xd7, 0x57, 0x71, 0x65, 0x65, 0x6d, 0x69, 0xa6, 0x4b, 0x3c, 0xbc, 0xd3,
	0x8b, 0x7e, 0x49, 0x45, 0x35, 0x8e, 0xf0, 0x6f, 0xa8, 0x87, 0x32, 0x0a, 0xc1, 0x0f, 0xb8, 0x1f,
	0x05, 0xe1, 0xac, 0x50, 0x12, 0x84, 0x54, 0x21, 0x4b, 0x53, 0x42, 0xa3, 0xe0, 0x47, 0xdc, 0x87,
	0x7e, 0xd7, 0xd1, 0x1a, 0x10, 0x2e, 0xe7, 0x80, 0xc8, 0xe0, 0x27, 0x9d, 0x28, 0x66, 0x43, 0x25,
	0x24, 0x91, 0xa0, 0x28, 0x93, 0xf1, 0xfc, 0x42, 0xf0, 0x33, 0x1e, 0x42, 0x15, 0x9d, 0x18, 0x11,
	0x49, 0x94, 0x58, 0xa0, 0xa1, 0xe2, 0x50, 0xcf, 0x41, 0xc8, 0xa0, 0x07, 0x0f, 0xa0, 0x3e, 0x2d,
	0x25, 0xac, 0x1a, 0x53, 0x17, 0xde, 0xa0, 0x78, 0x10, 0x61, 0x0d, 0xb2, 0x71, 0x91, 0x31, 0x2a,
	0x20, 0xd8, 0xa4, 0x78, 0x18, 0x55, 0xb4, 0xc0, 0x64, 0x0d, 0xb8, 0xca, 0x12, 0x12, 0x42, 0x3b,
	0x29, 0xd8, 0xa2, 0x78, 0x14, 0x0d, 0x69, 0x98, 0x80, 0x04, 0x42, 0xa9, 0x04, 0xf0, 0x06, 0x70,
	0x07, 0xdd, 0xa6, 0x78, 0x0c, 0x0d, 0xb7, 0xaa, 0x2b, 0xe9, 0x06, 0x7e, 0x8b, 0xe2, 0x11, 0x34,
	0xa0, 0x01, 0x40, 0x25, 0x70, 0x55, 0x25, 0x29, 0x38, 0xf7, 0x6e, 0x8e, 0xff, 0x42, 0x83, 0xda,
	0xed, 0x89, 0xc6, 0xfa, 0x25, 0xc7, 0x13, 0x68, 0xb4, 0xa4, 0x86, 0x2c, 0xcd, 0x12, 0xe8, 0xf4,
	0x61, 0x2f, 0xb7, 0xfc, 0x04, 0x48, 0x03, 0x7c, 0xfe, 0xbe, 0xe3, 0x7b, 0xa2, 0xe1, 0x7f, 0xcd,
	0xf1, 0xdf, 0xe8, 0x4f, 0xad, 0x66, 0x09, 0x59, 0x00, 0xae, 0x62, 0x3a, 0xcf, 0x3a, 0xf2, 0xb7,
	0xdc, 0xf6, 0x91, 0xb3, 0x04, 0x84, 0x83, 0x3e, 0x6d, 0xd8, 0x3e, 0xda, 0xb8, 0x31, 0x3c, 0x6b,
	0xd8, 0xd3, 0xda, 0x42, 0x58, 0x23, 0xb4, 0xea, 0x0a, 0xdd, 0x68, 0xea, 0xd3, 0x2c, 0x4e, 0x25,
	0xd0, 0x80, 0x24, 0xcf, 0x1c, 0x75, 0xb3, 0xa9, 0x1b, 0x6d, 0xcd, 0x05, 0xd9, 0xc0, 0xb7, 0x9a,
	0x1a, 0xee, 0xec, 0x9c, 0xd0, 0xe3, 0x05, 0xf7, 0x76, 0xd3, 0xfe, 0x15, 0x5f, 0xb5, 0x43, 0x68,
	0xea, 0x29, 0x39, 0x73, 0x04, 0xba, 0x8b, 0x4c, 0x74, 0x1a, 0x75, 0xbb, 0x89, 0xc7, 0xd1, 0x88,
	0xf3, 0x17, 0x13, 0x0c, 0xe2, 0x8e, 0x7f, 0xfe, 0x7c, 0x2e, 0x62, 0xd6, 0xd9, 0xad, 0xbb, 0xfe,
	0xf9, 0x4e, 0x35, 0xe6, 0x7b, 0x4d, 0xbb, 0x45, 0x73, 0x44, 0xca, 0xa4, 0x85, 0xcd, 0x13, 0xe9,
	0xec, 0x1f, 0xdc, 0x16, 0x95, 0x75, 0x03, 0xf8, 0x48, 0xbb, 0x00, 0x21, 0xe3, 0x91, 0x03, 0x7c,
	0xea, 0x06, 0x18, 0xdd, 0x00, 0x3e, 0x53, 0xbd, 0x4b, 0xdd, 0x80, 0x24, 0x16, 0x9d, 0x32, 0x76,
	0x29, 0x9e, 0x44, 0x63, 0xdd, 0x14, 0x93, 0x64, 0xd7, 0x92, 0xe2, 0xff, 0xd0, 0x78, 0x37, 0x2a,
	0x82, 0xd6, 0x5e, 0x5a, 0xd8, 0x1e, 0xc5, 0x53, 0xe8, 0x9f, 0x6e, 0x98, 0x4b, 0x33, 0xb8, 0x7d,
	0x8a, 0x87, 0x50, 0x7f, 0xfb, 0x86, 0x49, 0x25, 0x81, 0xa4, 0x0e, 0xf1, 0x8e, 0xd9, 0x8b, 0x59,
	0x90, 0x8c, 0x6d, 0x87, 0xd9, 0x2d, 0xd4, 0xf1, 0xc2, 0x16, 0xbe, 0x67, 0xf6, 0x42, 0xc4, 0x12,
	0x52, 0xe1, 0x3f, 0x0d, 0xcf, 0x33, 0xfc, 0x07, 0xea, 0xd5, 0xae, 0x8e, 0x18, 0xbc, 0xc8, 0x70,
	0xa5, 0xfd, 0xc2, 0xb4, 0x83, 0x79, 0x16, 0x11, 0x09, 0xc1, 0xcb, 0xcc, 0x4e, 0x5c, 0x87, 0xcb,
	0x13, 0x7f, 0x95, 0xd9, 0x89, 0xfb, 0xaa, 0xa9, 0xe2, 0x75, 0xa6, 0x4b, 0x77, 0x66, 0x01, 0x49,
	0xe2, 0xac, 0x6f, 0x32, 0x5d, 0xa1, 0xb3, 0x1a, 0xcd, 0x18, 0xdf, 0x66, 0x76, 0xd2, 0x21, 0x49,
	0x33, 0x12, 0x57, 0x69, 0xfb, 0x85, 0xb3, 0xe6, 0xfb, 0x75, 0x3b, 0xe9, 0xb2, 0x6e, 0x00, 0x0f,
	0xea, 0x78, 0x1a, 0xfd, 0xeb, 0x01, 0xf4, 0x3d, 0xcc, 0x33, 0x15, 0xd3, 0x90, 0x15, 0x1e, 0x87,
	0x87, 0x75, 0xfc, 0x3f, 0x9a, 0xf0, 0x48, 0xe5, 0x44, 0x83, 0x7c, 0x54, 0xc7, 0x47, 0xd0, 0x54,
	0x61, 0xe2, 0x7a, 0x1c, 0x26, 0xd7, 0x79, 0x2d, 0xf6, 0x71, 0x1d, 0x1f, 0x45, 0xd3, 0x85, 0xb9,
	0x1f, 0x9c, 0x6c, 0xd0, 0x4f, 0xea, 0xb6, 0x4f, 0x29, 0x89, 0x13, 0x25, 0x80, 0x76, 0x96, 0xfa,
	0x02, 0xb7, 0x7d, 0x2a, 0x6a, 0xc6, 0x78, 0x91, 0x7b, 0x46, 0x6f, 0x8f, 0x2f, 0xf9, 0x46, 0x7f,
	0x7d, 0x2f, 0x73, 0x3b, 0xd6, 0x96, 0x18, 0x81, 0xd4, 0x3f, 0xd6, 0x7a, 0x85, 0xdb, 0xb1, 0xfa,
	0xaa, 0x31, 0x5f, 0xe5, 0xba, 0x67, 0xce, 0x5c, 0x05, 0xa9, 0x88, 0x94, 0x24, 0xac, 0x41, 0x64,
	0x96, 0xc7, 0x82, 0xae, 0x71, 0x7c, 0x08, 0x4d, 0x3a, 0xd0, 0x81, 0x99, 0x06, 0x7a, 0xbd, 0x5c,
	0x91, 0x77, 0x8f, 0x6e, 0x94, 0x2b, 0xf2, 0xaf, 0xcf, 0x4d, 0x7e, 0xf2, 0xd7, 0xd6, 0x97, 0xf7,
	0xd8, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x0d, 0x10, 0x09, 0x94, 0x07, 0x00, 0x00,
}
