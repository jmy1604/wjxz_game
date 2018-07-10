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
	MSGID_NONE                                         MSGID = 0
	MSGID_C2S_TEST_COMMAND                             MSGID = 1
	MSGID_C2S_HEARTBEAT                                MSGID = 2
	MSGID_S2C_STATE_NOTIFY                             MSGID = 3
	MSGID_C2S_DATA_SYNC_REQUEST                        MSGID = 4
	MSGID_C2S_LOGIN_REQUEST                            MSGID = 10000
	MSGID_S2C_LOGIN_RESPONSE                           MSGID = 10001
	MSGID_S2C_OTHER_PLACE_LOGIN                        MSGID = 10002
	MSGID_C2S_SELECT_SERVER_REQUEST                    MSGID = 10003
	MSGID_S2C_SELECT_SERVER_RESPONSE                   MSGID = 10004
	MSGID_C2S_ENTER_GAME_REQUEST                       MSGID = 10020
	MSGID_S2C_ENTER_GAME_RESPONSE                      MSGID = 10021
	MSGID_S2C_ENTER_GAME_COMPLETE_NOTIFY               MSGID = 10022
	MSGID_C2S_LEAVE_GAME_REQUEST                       MSGID = 10023
	MSGID_S2C_LEAVE_GAME_RESPONSE                      MSGID = 10024
	MSGID_S2C_PLAYER_INFO_RESPONSE                     MSGID = 10025
	MSGID_C2S_ROLES_REQUEST                            MSGID = 10050
	MSGID_S2C_ROLES_RESPONSE                           MSGID = 10051
	MSGID_S2C_ROLES_CHANGE_NOTIFY                      MSGID = 10052
	MSGID_C2S_ROLE_ATTRS_REQUEST                       MSGID = 10053
	MSGID_S2C_ROLE_ATTRS_RESPONSE                      MSGID = 10054
	MSGID_C2S_ROLE_LEVELUP_REQUEST                     MSGID = 10055
	MSGID_S2C_ROLE_LEVELUP_RESPONSE                    MSGID = 10056
	MSGID_C2S_ROLE_RANKUP_REQUEST                      MSGID = 10057
	MSGID_S2C_ROLE_RANKUP_RESPONSE                     MSGID = 10058
	MSGID_C2S_ROLE_DECOMPOSE_REQUEST                   MSGID = 10059
	MSGID_S2C_ROLE_DECOMPOSE_RESPONSE                  MSGID = 10060
	MSGID_C2S_ROLE_FUSION_REQUEST                      MSGID = 10061
	MSGID_S2C_ROLE_FUSION_RESPONSE                     MSGID = 10062
	MSGID_C2S_ROLE_LOCK_REQUEST                        MSGID = 10063
	MSGID_S2C_ROLE_LOCK_RESPONSE                       MSGID = 10064
	MSGID_C2S_ROLE_HANDBOOK_REQUEST                    MSGID = 10065
	MSGID_S2C_ROLE_HANDBOOK_RESPONSE                   MSGID = 10066
	MSGID_C2S_ROLE_LEFTSLOT_OPEN_REQUEST               MSGID = 10067
	MSGID_S2C_ROLE_LEFTSLOT_OPEN_RESPONSE              MSGID = 10068
	MSGID_C2S_ROLE_ONEKEY_EQUIP_REQUEST                MSGID = 10069
	MSGID_S2C_ROLE_ONEKEY_EQUIP_RESPONSE               MSGID = 10070
	MSGID_C2S_ROLE_ONEKEY_UNEQUIP_REQUEST              MSGID = 10071
	MSGID_S2C_ROLE_ONEKEY_UNEQUIP_RESPONSE             MSGID = 10072
	MSGID_C2S_ROLE_LEFTSLOT_RESULT_SAVE_REQUEST        MSGID = 10073
	MSGID_S2C_ROLE_LEFTSLOT_RESULT_SAVE_RESPONSE       MSGID = 10074
	MSGID_C2S_ROLE_LEFTSLOT_RESULT_CANCEL_REQUEST      MSGID = 10075
	MSGID_S2C_ROLE_LEFTSLOT_RESULT_CANCEL_RESPONSE     MSGID = 10076
	MSGID_C2S_BATTLE_RESULT_REQUEST                    MSGID = 10100
	MSGID_S2C_BATTLE_RESULT_RESPONSE                   MSGID = 10101
	MSGID_C2S_BATTLE_RECORD_REQUEST                    MSGID = 10102
	MSGID_S2C_BATTLE_RECORD_RESPONSE                   MSGID = 10103
	MSGID_C2S_BATTLE_RECORD_LIST_REQUEST               MSGID = 10104
	MSGID_S2C_BATTLE_RECORD_LIST_RESPONSE              MSGID = 10105
	MSGID_C2S_BATTLE_RECORD_DELETE_REQUEST             MSGID = 10106
	MSGID_S2C_BATTLE_RECORD_DELETE_RESPONSE            MSGID = 10107
	MSGID_C2S_SET_TEAM_REQUEST                         MSGID = 10200
	MSGID_S2C_SET_TEAM_RESPONSE                        MSGID = 10201
	MSGID_S2C_TEAMS_RESPONSE                           MSGID = 10202
	MSGID_C2S_ITEMS_SYNC_REQUEST                       MSGID = 10300
	MSGID_S2C_ITEMS_SYNC                               MSGID = 10301
	MSGID_S2C_ITEMS_UPDATE                             MSGID = 10302
	MSGID_C2S_ITEM_FUSION_REQUEST                      MSGID = 10303
	MSGID_S2C_ITEM_FUSION_RESPONSE                     MSGID = 10304
	MSGID_C2S_ITEM_SELL_REQUEST                        MSGID = 10305
	MSGID_S2C_ITEM_SELL_RESPONSE                       MSGID = 10306
	MSGID_C2S_ITEM_EQUIP_REQUEST                       MSGID = 10307
	MSGID_S2C_ITEM_EQUIP_RESPONSE                      MSGID = 10308
	MSGID_C2S_ITEM_UNEQUIP_REQUEST                     MSGID = 10309
	MSGID_S2C_ITEM_UNEQUIP_RESPONSE                    MSGID = 10310
	MSGID_C2S_ITEM_UPGRADE_REQUEST                     MSGID = 10311
	MSGID_S2C_ITEM_UPGRADE_RESPONSE                    MSGID = 10312
	MSGID_C2S_ITEM_ONEKEY_UPGRADE_REQUEST              MSGID = 10313
	MSGID_S2C_ITEM_ONEKEY_UPGRADE_RESPONSE             MSGID = 10314
	MSGID_C2S_CAMPAIGN_DATA_REQUEST                    MSGID = 10400
	MSGID_S2C_CAMPAIGN_DATA_RESPONSE                   MSGID = 10401
	MSGID_C2S_CAMPAIGN_HANGUP_INCOME_REQUEST           MSGID = 10402
	MSGID_S2C_CAMPAIGN_HANGUP_INCOME_RESPONSE          MSGID = 10403
	MSGID_C2S_BATTLE_SET_HANGUP_CAMPAIGN_REQUEST       MSGID = 10404
	MSGID_S2C_BATTLE_SET_HANGUP_CAMPAIGN_RESPONSE      MSGID = 10405
	MSGID_C2S_MAIL_SEND_REQUEST                        MSGID = 10500
	MSGID_S2C_MAIL_SEND_RESPONSE                       MSGID = 10501
	MSGID_C2S_MAIL_LIST_REQUEST                        MSGID = 10502
	MSGID_S2C_MAIL_LIST_RESPONSE                       MSGID = 10503
	MSGID_C2S_MAIL_DETAIL_REQUEST                      MSGID = 10504
	MSGID_S2C_MAIL_DETAIL_RESPONSE                     MSGID = 10505
	MSGID_C2S_MAIL_GET_ATTACHED_ITEMS_REQUEST          MSGID = 10506
	MSGID_S2C_MAIL_GET_ATTACHED_ITEMS_RESPONSE         MSGID = 10507
	MSGID_C2S_MAIL_DELETE_REQUEST                      MSGID = 10508
	MSGID_S2C_MAIL_DELETE_RESPONSE                     MSGID = 10509
	MSGID_S2C_MAILS_NEW_NOTIFY                         MSGID = 10510
	MSGID_C2S_TALENT_UP_REQUEST                        MSGID = 10600
	MSGID_S2C_TALENT_UP_RESPONSE                       MSGID = 10601
	MSGID_C2S_TALENT_LIST_REQUEST                      MSGID = 10602
	MSGID_S2C_TALENT_LIST_RESPONSE                     MSGID = 10603
	MSGID_C2S_TOWER_DATA_REQUEST                       MSGID = 10700
	MSGID_S2C_TOWER_DATA_RESPONSE                      MSGID = 10701
	MSGID_C2S_TOWER_RECORDS_INFO_REQUEST               MSGID = 10702
	MSGID_S2C_TOWER_RECORDS_INFO_RESPONSE              MSGID = 10703
	MSGID_C2S_TOWER_RECORD_DATA_REQUEST                MSGID = 10704
	MSGID_S2C_TOWER_RECORD_DATA_RESPONSE               MSGID = 10705
	MSGID_C2S_TOWER_RANKING_LIST_REQUEST               MSGID = 10706
	MSGID_S2C_TOWER_RANKING_LIST_RESPONSE              MSGID = 10707
	MSGID_C2S_DRAW_CARD_REQUEST                        MSGID = 10800
	MSGID_S2C_DRAW_CARD_RESPONSE                       MSGID = 10801
	MSGID_C2S_DRAW_DATA_REQUEST                        MSGID = 10802
	MSGID_S2C_DRAW_DATA_RESPONSE                       MSGID = 10803
	MSGID_C2S_TOUCH_GOLD_REQUEST                       MSGID = 10900
	MSGID_S2C_TOUCH_GOLD_RESPONSE                      MSGID = 10901
	MSGID_C2S_GOLD_HAND_DATA_REQUEST                   MSGID = 10902
	MSGID_S2C_GOLD_HAND_DATA_RESPONSE                  MSGID = 10903
	MSGID_C2S_SHOP_DATA_REQUEST                        MSGID = 11000
	MSGID_S2C_SHOP_DATA_RESPONSE                       MSGID = 11001
	MSGID_C2S_SHOP_BUY_ITEM_REQUEST                    MSGID = 11002
	MSGID_S2C_SHOP_BUY_ITEM_RESPONSE                   MSGID = 11003
	MSGID_C2S_SHOP_REFRESH_REQUEST                     MSGID = 11004
	MSGID_S2C_SHOP_REFRESH_RESPONSE                    MSGID = 11005
	MSGID_S2C_SHOP_AUTO_REFRESH_NOTIFY                 MSGID = 11006
	MSGID_C2S_RANK_LIST_REQUEST                        MSGID = 11100
	MSGID_S2C_RANK_LIST_RESPONSE                       MSGID = 11101
	MSGID_C2S_ARENA_DATA_REQUEST                       MSGID = 11200
	MSGID_S2C_ARENA_DATA_RESPONSE                      MSGID = 11201
	MSGID_C2S_ARENA_PLAYER_DEFENSE_TEAM_REQUEST        MSGID = 11202
	MSGID_S2C_ARENA_PLAYER_DEFENSE_TEAM_RESPONSE       MSGID = 11203
	MSGID_C2S_ARENA_MATCH_PLAYER_REQUEST               MSGID = 11204
	MSGID_S2C_ARENA_MATCH_PLAYER_RESPONSE              MSGID = 11205
	MSGID_C2S_ACTIVE_STAGE_DATA_REQUEST                MSGID = 11300
	MSGID_S2C_ACTIVE_STAGE_DATA_RESPONSE               MSGID = 11301
	MSGID_C2S_ACTIVE_STAGE_CHALLENGE_NUM_BUY_REQUEST   MSGID = 11302
	MSGID_S2C_ACTIVE_STAGE_CHALLENGE_NUM_BUY_RESPONSE  MSGID = 11303
	MSGID_S2C_ACTIVE_STAGE_REFRESH_NOTIFY              MSGID = 11304
	MSGID_C2S_ACTIVE_STAGE_SELECT_ASSIST_ROLE_REQUEST  MSGID = 11305
	MSGID_S2C_ACTIVE_STAGE_SELECT_ASSIST_ROLE_RESPONSE MSGID = 11306
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
	10020: "C2S_ENTER_GAME_REQUEST",
	10021: "S2C_ENTER_GAME_RESPONSE",
	10022: "S2C_ENTER_GAME_COMPLETE_NOTIFY",
	10023: "C2S_LEAVE_GAME_REQUEST",
	10024: "S2C_LEAVE_GAME_RESPONSE",
	10025: "S2C_PLAYER_INFO_RESPONSE",
	10050: "C2S_ROLES_REQUEST",
	10051: "S2C_ROLES_RESPONSE",
	10052: "S2C_ROLES_CHANGE_NOTIFY",
	10053: "C2S_ROLE_ATTRS_REQUEST",
	10054: "S2C_ROLE_ATTRS_RESPONSE",
	10055: "C2S_ROLE_LEVELUP_REQUEST",
	10056: "S2C_ROLE_LEVELUP_RESPONSE",
	10057: "C2S_ROLE_RANKUP_REQUEST",
	10058: "S2C_ROLE_RANKUP_RESPONSE",
	10059: "C2S_ROLE_DECOMPOSE_REQUEST",
	10060: "S2C_ROLE_DECOMPOSE_RESPONSE",
	10061: "C2S_ROLE_FUSION_REQUEST",
	10062: "S2C_ROLE_FUSION_RESPONSE",
	10063: "C2S_ROLE_LOCK_REQUEST",
	10064: "S2C_ROLE_LOCK_RESPONSE",
	10065: "C2S_ROLE_HANDBOOK_REQUEST",
	10066: "S2C_ROLE_HANDBOOK_RESPONSE",
	10067: "C2S_ROLE_LEFTSLOT_OPEN_REQUEST",
	10068: "S2C_ROLE_LEFTSLOT_OPEN_RESPONSE",
	10069: "C2S_ROLE_ONEKEY_EQUIP_REQUEST",
	10070: "S2C_ROLE_ONEKEY_EQUIP_RESPONSE",
	10071: "C2S_ROLE_ONEKEY_UNEQUIP_REQUEST",
	10072: "S2C_ROLE_ONEKEY_UNEQUIP_RESPONSE",
	10073: "C2S_ROLE_LEFTSLOT_RESULT_SAVE_REQUEST",
	10074: "S2C_ROLE_LEFTSLOT_RESULT_SAVE_RESPONSE",
	10075: "C2S_ROLE_LEFTSLOT_RESULT_CANCEL_REQUEST",
	10076: "S2C_ROLE_LEFTSLOT_RESULT_CANCEL_RESPONSE",
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
	10307: "C2S_ITEM_EQUIP_REQUEST",
	10308: "S2C_ITEM_EQUIP_RESPONSE",
	10309: "C2S_ITEM_UNEQUIP_REQUEST",
	10310: "S2C_ITEM_UNEQUIP_RESPONSE",
	10311: "C2S_ITEM_UPGRADE_REQUEST",
	10312: "S2C_ITEM_UPGRADE_RESPONSE",
	10313: "C2S_ITEM_ONEKEY_UPGRADE_REQUEST",
	10314: "S2C_ITEM_ONEKEY_UPGRADE_RESPONSE",
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
	10510: "S2C_MAILS_NEW_NOTIFY",
	10600: "C2S_TALENT_UP_REQUEST",
	10601: "S2C_TALENT_UP_RESPONSE",
	10602: "C2S_TALENT_LIST_REQUEST",
	10603: "S2C_TALENT_LIST_RESPONSE",
	10700: "C2S_TOWER_DATA_REQUEST",
	10701: "S2C_TOWER_DATA_RESPONSE",
	10702: "C2S_TOWER_RECORDS_INFO_REQUEST",
	10703: "S2C_TOWER_RECORDS_INFO_RESPONSE",
	10704: "C2S_TOWER_RECORD_DATA_REQUEST",
	10705: "S2C_TOWER_RECORD_DATA_RESPONSE",
	10706: "C2S_TOWER_RANKING_LIST_REQUEST",
	10707: "S2C_TOWER_RANKING_LIST_RESPONSE",
	10800: "C2S_DRAW_CARD_REQUEST",
	10801: "S2C_DRAW_CARD_RESPONSE",
	10802: "C2S_DRAW_DATA_REQUEST",
	10803: "S2C_DRAW_DATA_RESPONSE",
	10900: "C2S_TOUCH_GOLD_REQUEST",
	10901: "S2C_TOUCH_GOLD_RESPONSE",
	10902: "C2S_GOLD_HAND_DATA_REQUEST",
	10903: "S2C_GOLD_HAND_DATA_RESPONSE",
	11000: "C2S_SHOP_DATA_REQUEST",
	11001: "S2C_SHOP_DATA_RESPONSE",
	11002: "C2S_SHOP_BUY_ITEM_REQUEST",
	11003: "S2C_SHOP_BUY_ITEM_RESPONSE",
	11004: "C2S_SHOP_REFRESH_REQUEST",
	11005: "S2C_SHOP_REFRESH_RESPONSE",
	11006: "S2C_SHOP_AUTO_REFRESH_NOTIFY",
	11100: "C2S_RANK_LIST_REQUEST",
	11101: "S2C_RANK_LIST_RESPONSE",
	11200: "C2S_ARENA_DATA_REQUEST",
	11201: "S2C_ARENA_DATA_RESPONSE",
	11202: "C2S_ARENA_PLAYER_DEFENSE_TEAM_REQUEST",
	11203: "S2C_ARENA_PLAYER_DEFENSE_TEAM_RESPONSE",
	11204: "C2S_ARENA_MATCH_PLAYER_REQUEST",
	11205: "S2C_ARENA_MATCH_PLAYER_RESPONSE",
	11300: "C2S_ACTIVE_STAGE_DATA_REQUEST",
	11301: "S2C_ACTIVE_STAGE_DATA_RESPONSE",
	11302: "C2S_ACTIVE_STAGE_CHALLENGE_NUM_BUY_REQUEST",
	11303: "S2C_ACTIVE_STAGE_CHALLENGE_NUM_BUY_RESPONSE",
	11304: "S2C_ACTIVE_STAGE_REFRESH_NOTIFY",
	11305: "C2S_ACTIVE_STAGE_SELECT_ASSIST_ROLE_REQUEST",
	11306: "S2C_ACTIVE_STAGE_SELECT_ASSIST_ROLE_RESPONSE",
}
var MSGID_value = map[string]int32{
	"NONE":                                         0,
	"C2S_TEST_COMMAND":                             1,
	"C2S_HEARTBEAT":                                2,
	"S2C_STATE_NOTIFY":                             3,
	"C2S_DATA_SYNC_REQUEST":                        4,
	"C2S_LOGIN_REQUEST":                            10000,
	"S2C_LOGIN_RESPONSE":                           10001,
	"S2C_OTHER_PLACE_LOGIN":                        10002,
	"C2S_SELECT_SERVER_REQUEST":                    10003,
	"S2C_SELECT_SERVER_RESPONSE":                   10004,
	"C2S_ENTER_GAME_REQUEST":                       10020,
	"S2C_ENTER_GAME_RESPONSE":                      10021,
	"S2C_ENTER_GAME_COMPLETE_NOTIFY":               10022,
	"C2S_LEAVE_GAME_REQUEST":                       10023,
	"S2C_LEAVE_GAME_RESPONSE":                      10024,
	"S2C_PLAYER_INFO_RESPONSE":                     10025,
	"C2S_ROLES_REQUEST":                            10050,
	"S2C_ROLES_RESPONSE":                           10051,
	"S2C_ROLES_CHANGE_NOTIFY":                      10052,
	"C2S_ROLE_ATTRS_REQUEST":                       10053,
	"S2C_ROLE_ATTRS_RESPONSE":                      10054,
	"C2S_ROLE_LEVELUP_REQUEST":                     10055,
	"S2C_ROLE_LEVELUP_RESPONSE":                    10056,
	"C2S_ROLE_RANKUP_REQUEST":                      10057,
	"S2C_ROLE_RANKUP_RESPONSE":                     10058,
	"C2S_ROLE_DECOMPOSE_REQUEST":                   10059,
	"S2C_ROLE_DECOMPOSE_RESPONSE":                  10060,
	"C2S_ROLE_FUSION_REQUEST":                      10061,
	"S2C_ROLE_FUSION_RESPONSE":                     10062,
	"C2S_ROLE_LOCK_REQUEST":                        10063,
	"S2C_ROLE_LOCK_RESPONSE":                       10064,
	"C2S_ROLE_HANDBOOK_REQUEST":                    10065,
	"S2C_ROLE_HANDBOOK_RESPONSE":                   10066,
	"C2S_ROLE_LEFTSLOT_OPEN_REQUEST":               10067,
	"S2C_ROLE_LEFTSLOT_OPEN_RESPONSE":              10068,
	"C2S_ROLE_ONEKEY_EQUIP_REQUEST":                10069,
	"S2C_ROLE_ONEKEY_EQUIP_RESPONSE":               10070,
	"C2S_ROLE_ONEKEY_UNEQUIP_REQUEST":              10071,
	"S2C_ROLE_ONEKEY_UNEQUIP_RESPONSE":             10072,
	"C2S_ROLE_LEFTSLOT_RESULT_SAVE_REQUEST":        10073,
	"S2C_ROLE_LEFTSLOT_RESULT_SAVE_RESPONSE":       10074,
	"C2S_ROLE_LEFTSLOT_RESULT_CANCEL_REQUEST":      10075,
	"S2C_ROLE_LEFTSLOT_RESULT_CANCEL_RESPONSE":     10076,
	"C2S_BATTLE_RESULT_REQUEST":                    10100,
	"S2C_BATTLE_RESULT_RESPONSE":                   10101,
	"C2S_BATTLE_RECORD_REQUEST":                    10102,
	"S2C_BATTLE_RECORD_RESPONSE":                   10103,
	"C2S_BATTLE_RECORD_LIST_REQUEST":               10104,
	"S2C_BATTLE_RECORD_LIST_RESPONSE":              10105,
	"C2S_BATTLE_RECORD_DELETE_REQUEST":             10106,
	"S2C_BATTLE_RECORD_DELETE_RESPONSE":            10107,
	"C2S_SET_TEAM_REQUEST":                         10200,
	"S2C_SET_TEAM_RESPONSE":                        10201,
	"S2C_TEAMS_RESPONSE":                           10202,
	"C2S_ITEMS_SYNC_REQUEST":                       10300,
	"S2C_ITEMS_SYNC":                               10301,
	"S2C_ITEMS_UPDATE":                             10302,
	"C2S_ITEM_FUSION_REQUEST":                      10303,
	"S2C_ITEM_FUSION_RESPONSE":                     10304,
	"C2S_ITEM_SELL_REQUEST":                        10305,
	"S2C_ITEM_SELL_RESPONSE":                       10306,
	"C2S_ITEM_EQUIP_REQUEST":                       10307,
	"S2C_ITEM_EQUIP_RESPONSE":                      10308,
	"C2S_ITEM_UNEQUIP_REQUEST":                     10309,
	"S2C_ITEM_UNEQUIP_RESPONSE":                    10310,
	"C2S_ITEM_UPGRADE_REQUEST":                     10311,
	"S2C_ITEM_UPGRADE_RESPONSE":                    10312,
	"C2S_ITEM_ONEKEY_UPGRADE_REQUEST":              10313,
	"S2C_ITEM_ONEKEY_UPGRADE_RESPONSE":             10314,
	"C2S_CAMPAIGN_DATA_REQUEST":                    10400,
	"S2C_CAMPAIGN_DATA_RESPONSE":                   10401,
	"C2S_CAMPAIGN_HANGUP_INCOME_REQUEST":           10402,
	"S2C_CAMPAIGN_HANGUP_INCOME_RESPONSE":          10403,
	"C2S_BATTLE_SET_HANGUP_CAMPAIGN_REQUEST":       10404,
	"S2C_BATTLE_SET_HANGUP_CAMPAIGN_RESPONSE":      10405,
	"C2S_MAIL_SEND_REQUEST":                        10500,
	"S2C_MAIL_SEND_RESPONSE":                       10501,
	"C2S_MAIL_LIST_REQUEST":                        10502,
	"S2C_MAIL_LIST_RESPONSE":                       10503,
	"C2S_MAIL_DETAIL_REQUEST":                      10504,
	"S2C_MAIL_DETAIL_RESPONSE":                     10505,
	"C2S_MAIL_GET_ATTACHED_ITEMS_REQUEST":          10506,
	"S2C_MAIL_GET_ATTACHED_ITEMS_RESPONSE":         10507,
	"C2S_MAIL_DELETE_REQUEST":                      10508,
	"S2C_MAIL_DELETE_RESPONSE":                     10509,
	"S2C_MAILS_NEW_NOTIFY":                         10510,
	"C2S_TALENT_UP_REQUEST":                        10600,
	"S2C_TALENT_UP_RESPONSE":                       10601,
	"C2S_TALENT_LIST_REQUEST":                      10602,
	"S2C_TALENT_LIST_RESPONSE":                     10603,
	"C2S_TOWER_DATA_REQUEST":                       10700,
	"S2C_TOWER_DATA_RESPONSE":                      10701,
	"C2S_TOWER_RECORDS_INFO_REQUEST":               10702,
	"S2C_TOWER_RECORDS_INFO_RESPONSE":              10703,
	"C2S_TOWER_RECORD_DATA_REQUEST":                10704,
	"S2C_TOWER_RECORD_DATA_RESPONSE":               10705,
	"C2S_TOWER_RANKING_LIST_REQUEST":               10706,
	"S2C_TOWER_RANKING_LIST_RESPONSE":              10707,
	"C2S_DRAW_CARD_REQUEST":                        10800,
	"S2C_DRAW_CARD_RESPONSE":                       10801,
	"C2S_DRAW_DATA_REQUEST":                        10802,
	"S2C_DRAW_DATA_RESPONSE":                       10803,
	"C2S_TOUCH_GOLD_REQUEST":                       10900,
	"S2C_TOUCH_GOLD_RESPONSE":                      10901,
	"C2S_GOLD_HAND_DATA_REQUEST":                   10902,
	"S2C_GOLD_HAND_DATA_RESPONSE":                  10903,
	"C2S_SHOP_DATA_REQUEST":                        11000,
	"S2C_SHOP_DATA_RESPONSE":                       11001,
	"C2S_SHOP_BUY_ITEM_REQUEST":                    11002,
	"S2C_SHOP_BUY_ITEM_RESPONSE":                   11003,
	"C2S_SHOP_REFRESH_REQUEST":                     11004,
	"S2C_SHOP_REFRESH_RESPONSE":                    11005,
	"S2C_SHOP_AUTO_REFRESH_NOTIFY":                 11006,
	"C2S_RANK_LIST_REQUEST":                        11100,
	"S2C_RANK_LIST_RESPONSE":                       11101,
	"C2S_ARENA_DATA_REQUEST":                       11200,
	"S2C_ARENA_DATA_RESPONSE":                      11201,
	"C2S_ARENA_PLAYER_DEFENSE_TEAM_REQUEST":        11202,
	"S2C_ARENA_PLAYER_DEFENSE_TEAM_RESPONSE":       11203,
	"C2S_ARENA_MATCH_PLAYER_REQUEST":               11204,
	"S2C_ARENA_MATCH_PLAYER_RESPONSE":              11205,
	"C2S_ACTIVE_STAGE_DATA_REQUEST":                11300,
	"S2C_ACTIVE_STAGE_DATA_RESPONSE":               11301,
	"C2S_ACTIVE_STAGE_CHALLENGE_NUM_BUY_REQUEST":   11302,
	"S2C_ACTIVE_STAGE_CHALLENGE_NUM_BUY_RESPONSE":  11303,
	"S2C_ACTIVE_STAGE_REFRESH_NOTIFY":              11304,
	"C2S_ACTIVE_STAGE_SELECT_ASSIST_ROLE_REQUEST":  11305,
	"S2C_ACTIVE_STAGE_SELECT_ASSIST_ROLE_RESPONSE": 11306,
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
	// 1423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x98, 0xd7, 0x8f, 0x1c, 0xc5,
	0x13, 0xc7, 0x7f, 0xc1, 0x20, 0xd4, 0x12, 0x68, 0xdd, 0xf8, 0x6c, 0x9c, 0x73, 0xc0, 0x36, 0x06,
	0xcc, 0x5f, 0xd0, 0x37, 0x5b, 0xb7, 0x3b, 0xba, 0xd9, 0xee, 0xd9, 0xee, 0x9e, 0xdb, 0xbb, 0xa7,
	0x11, 0xc1, 0xb2, 0x2c, 0x61, 0x8c, 0xb0, 0xff, 0x04, 0x72, 0x12, 0xc1, 0xc0, 0x2b, 0xc1, 0x38,
	0xf0, 0x44, 0x78, 0x25, 0x38, 0x62, 0x0c, 0xce, 0x26, 0xd8, 0xc6, 0xbc, 0x13, 0x9e, 0x90, 0x08,
	0x26, 0x0a, 0xf5, 0x4c, 0x55, 0xcf, 0xf4, 0xec, 0x9a, 0xa7, 0x93, 0xe6, 0x5b, 0xdf, 0xcf, 0xd6,
	0x74, 0x55, 0x57, 0xf7, 0x1c, 0x9b, 0x77, 0xef, 0xfd, 0xdb, 0xb6, 0x3c, 0xb0, 0x2b, 0xdf, 0xbe,
	0x65, 0xe7, 0xce, 0xbb, 0xb7, 0x6e, 0xc9, 0xb7, 0xdd, 0xb7, 0xe9, 0xc1, 0x87, 0x76, 0xec, 0xda,
	0xc1, 0xc7, 0xb6, 0xef, 0xdc, 0xba, 0x69, 0x48, 0x5c, 0xff, 0xe3, 0x32, 0x76, 0x5d, 0xcf, 0x74,
	0xe2, 0x36, 0xbf, 0x81, 0xcd, 0x92, 0x4a, 0x42, 0xeb, 0x3f, 0x7c, 0x0e, 0x6b, 0x45, 0x9b, 0x4d,
	0x6e, 0xc1, 0xd8, 0x3c, 0x52, 0xbd, 0x9e, 0x90, 0xed, 0xd6, 0x7f, 0xf9, 0x6c, 0x76, 0xa3, 0x7b,
	0xda, 0x05, 0xa1, 0xed, 0x38, 0x08, 0xdb, 0xfa, 0x9f, 0x0b, 0x34, 0x9b, 0xa3, 0xdc, 0x58, 0x61,
	0x21, 0x97, 0xca, 0xc6, 0x13, 0x33, 0xad, 0xff, 0xf3, 0xf9, 0x6c, 0xcc, 0x05, 0xb6, 0x85, 0x15,
	0xb9, 0x99, 0x91, 0x51, 0xae, 0xa1, 0x9f, 0x81, 0xb1, 0xad, 0x59, 0x7c, 0x2e, 0x9b, 0xed, 0xa4,
	0x44, 0x75, 0x62, 0xe9, 0x1f, 0x3f, 0x27, 0xf9, 0x3c, 0xc6, 0x1d, 0x88, 0x9e, 0x9b, 0x54, 0x49,
	0x03, 0xad, 0xe7, 0x25, 0x5f, 0xc0, 0xc6, 0x9c, 0xa0, 0x6c, 0x17, 0x74, 0x9e, 0x26, 0x22, 0x82,
	0x32, 0xa8, 0xf5, 0x82, 0xe4, 0x4b, 0xd8, 0x7c, 0x07, 0x33, 0x90, 0x40, 0x64, 0x73, 0x03, 0x7a,
	0x0a, 0xb4, 0x87, 0xbe, 0x28, 0xf9, 0x52, 0xb6, 0xa0, 0xc8, 0xae, 0xa1, 0x23, 0x7c, 0xb7, 0xe4,
	0x0b, 0xd9, 0x5c, 0x07, 0x00, 0x69, 0x41, 0xe7, 0x1d, 0xd1, 0x03, 0xef, 0xde, 0x23, 0xf9, 0x22,
	0x36, 0xcf, 0xb9, 0x03, 0x11, 0xad, 0x6f, 0x4a, 0xbe, 0x92, 0x2d, 0x69, 0xa8, 0x91, 0xea, 0xa5,
	0x09, 0x54, 0xeb, 0xb0, 0xd7, 0xf3, 0x13, 0x10, 0x53, 0x10, 0xf2, 0xf7, 0x79, 0x7e, 0x20, 0x22,
	0x7f, 0xbf, 0xe4, 0x8b, 0xd9, 0x2d, 0x4e, 0x4d, 0x13, 0x31, 0x03, 0x3a, 0x8f, 0xe5, 0x84, 0xaa,
	0xe4, 0x03, 0x92, 0xd6, 0x51, 0xab, 0x04, 0x8c, 0x87, 0x1e, 0xf6, 0xeb, 0x48, 0xcf, 0xd1, 0x70,
	0xc4, 0xff, 0x5a, 0x29, 0x44, 0x5d, 0x21, 0x3b, 0x3e, 0xd1, 0xa3, 0x3e, 0x51, 0xa7, 0xe6, 0xc2,
	0x5a, 0x5d, 0x31, 0x8f, 0x05, 0x56, 0x2f, 0x22, 0xf8, 0xe3, 0x22, 0x51, 0x6f, 0x4d, 0x60, 0x0a,
	0x92, 0x2c, 0xf5, 0xe6, 0xe3, 0x45, 0x8d, 0xbc, 0xb9, 0x92, 0xd1, 0xfe, 0x49, 0x01, 0xf7, 0x76,
	0x2d, 0xe4, 0x64, 0xcd, 0x7d, 0xc2, 0xaf, 0x42, 0xa8, 0xa2, 0xf9, 0xd3, 0xa2, 0xc0, 0xde, 0xdc,
	0x06, 0x57, 0x00, 0x65, 0xaa, 0x35, 0xfe, 0x4c, 0xf2, 0x65, 0x6c, 0xa1, 0xf7, 0xd7, 0x03, 0x10,
	0x71, 0x32, 0xfc, 0xfd, 0x89, 0xcc, 0xc4, 0xaa, 0x6a, 0xcb, 0x53, 0xe1, 0xef, 0x7b, 0x15, 0xcd,
	0xa7, 0x8b, 0xe6, 0xac, 0xde, 0x5d, 0x45, 0x93, 0xde, 0x7a, 0xa6, 0x58, 0xd2, 0xea, 0xc5, 0x4b,
	0x0d, 0x8d, 0x67, 0x7d, 0xe7, 0x16, 0x62, 0x57, 0xc8, 0xf6, 0xb8, 0x52, 0x95, 0xf9, 0x9c, 0xef,
	0xdc, 0xa6, 0x8e, 0x80, 0xf3, 0x45, 0xfb, 0xd5, 0x56, 0x7d, 0xc2, 0x9a, 0x44, 0xd9, 0x5c, 0xa5,
	0x50, 0x65, 0x7f, 0x41, 0xf2, 0x55, 0x6c, 0x69, 0x6d, 0xed, 0xc3, 0x20, 0x44, 0x7d, 0x2e, 0xf9,
	0x0a, 0xb6, 0xd8, 0xa3, 0x94, 0x84, 0x49, 0x98, 0xc9, 0xa1, 0x9f, 0xc5, 0x55, 0x1d, 0xbe, 0xf0,
	0xdd, 0x3e, 0x2a, 0x06, 0x41, 0x5f, 0x16, 0x3f, 0xd7, 0x04, 0x65, 0x32, 0x44, 0x7d, 0x25, 0xf9,
	0x6a, 0xb6, 0xac, 0x89, 0xaa, 0xa2, 0x10, 0x76, 0x51, 0xf2, 0xf5, 0x6c, 0xf5, 0xf0, 0x0b, 0x6a,
	0x30, 0x59, 0x62, 0x73, 0xe3, 0x76, 0x0c, 0x21, 0x2f, 0x49, 0xbe, 0x81, 0xad, 0x19, 0x7e, 0xcf,
	0x30, 0x16, 0xc1, 0x97, 0x25, 0xdf, 0xc8, 0xd6, 0x5e, 0x13, 0x1c, 0x09, 0x19, 0x41, 0xe2, 0xd1,
	0x5f, 0x4b, 0x7e, 0x1b, 0x5b, 0x77, 0x4d, 0xb4, 0x8f, 0x46, 0xf8, 0x15, 0x5f, 0xd7, 0x71, 0x61,
	0x6d, 0x02, 0x14, 0x47, 0xb8, 0x9f, 0x7c, 0x5d, 0x9b, 0x3a, 0x02, 0x7e, 0x1e, 0x06, 0x44, 0x4a,
	0xb7, 0x3d, 0xe0, 0x97, 0x61, 0x00, 0xea, 0x08, 0xf8, 0xd5, 0x37, 0x46, 0x18, 0x90, 0xc4, 0xa6,
	0x4a, 0xe3, 0xaa, 0x6f, 0x8c, 0x91, 0x41, 0x88, 0xfa, 0xad, 0xa8, 0xd4, 0x30, 0xaa, 0x0d, 0xc5,
	0x8c, 0x23, 0xd8, 0xef, 0x92, 0xaf, 0x61, 0xcb, 0x87, 0x61, 0x3e, 0x0c, 0x71, 0x7f, 0x48, 0x3e,
	0x9f, 0xcd, 0x29, 0xa7, 0xb5, 0xcd, 0x2d, 0x88, 0x9e, 0x47, 0x5c, 0x54, 0x34, 0xe4, 0x6b, 0x12,
	0xda, 0x2e, 0x29, 0x9a, 0x68, 0xee, 0x79, 0x6d, 0xf0, 0x5c, 0x56, 0x34, 0xb3, 0x62, 0x0b, 0x3d,
	0x13, 0x1e, 0x33, 0xef, 0xa7, 0xfc, 0x66, 0x76, 0x93, 0x73, 0x55, 0x62, 0xeb, 0x83, 0x94, 0x8f,
	0x95, 0xa7, 0x55, 0xf9, 0x30, 0x4b, 0xdb, 0xc2, 0x42, 0xeb, 0xc3, 0x94, 0x46, 0x80, 0x7b, 0xdc,
	0x1c, 0x01, 0x1f, 0xa5, 0x34, 0x02, 0x42, 0x15, 0xb3, 0x38, 0x98, 0xd2, 0x08, 0x28, 0x64, 0x03,
	0x49, 0xd5, 0x3c, 0x87, 0x52, 0x1a, 0x01, 0x75, 0x0d, 0x8d, 0x87, 0xd3, 0x7a, 0xfa, 0x8d, 0xfd,
	0x76, 0x24, 0xa5, 0x91, 0x1b, 0x88, 0x68, 0x3d, 0x9a, 0xd2, 0xc8, 0x2d, 0xd4, 0xe6, 0x0e, 0x3b,
	0x96, 0xd2, 0xc8, 0x6d, 0xc8, 0x34, 0xb1, 0x1b, 0xf6, 0xb4, 0xa3, 0x45, 0xbb, 0xaa, 0xe7, 0xf1,
	0x86, 0xdd, 0xcb, 0x34, 0xb1, 0x53, 0xda, 0xe6, 0x85, 0x4e, 0x1b, 0xb8, 0x41, 0x39, 0x91, 0xd2,
	0x36, 0x1f, 0x1d, 0x45, 0x13, 0x3c, 0xa5, 0x7e, 0x8f, 0x44, 0x2f, 0x15, 0x71, 0x47, 0x96, 0x77,
	0x06, 0xc2, 0xbc, 0xda, 0xa7, 0x7e, 0x6f, 0xea, 0x08, 0x78, 0xad, 0xcf, 0xd7, 0xb2, 0x15, 0x01,
	0xc0, 0x9d, 0x6c, 0x59, 0x9a, 0xc7, 0x32, 0x52, 0xb5, 0xe3, 0xf6, 0xf5, 0x3e, 0x5f, 0xc7, 0x56,
	0x06, 0xa4, 0x66, 0x20, 0x22, 0xdf, 0xe8, 0xbb, 0x71, 0x52, 0xeb, 0x7b, 0xd7, 0x94, 0x18, 0xeb,
	0xbd, 0xfe, 0x96, 0xd0, 0x77, 0xe3, 0xa4, 0xd6, 0xfd, 0xa3, 0x83, 0xe9, 0xd6, 0xd0, 0xa7, 0x6e,
	0xe9, 0x89, 0x38, 0xc9, 0x0d, 0xc8, 0x6a, 0x6b, 0x3f, 0xac, 0xa9, 0x5b, 0xea, 0x1a, 0x1a, 0x1f,
	0xd1, 0x81, 0x31, 0xd8, 0xcd, 0x8f, 0x86, 0xc6, 0x70, 0x13, 0x3f, 0xa6, 0xa9, 0xb9, 0x0b, 0xb1,
	0x0d, 0xd6, 0xfd, 0x21, 0xeb, 0xe3, 0x9a, 0x9a, 0x3b, 0x54, 0xd1, 0xfc, 0x84, 0x76, 0x6b, 0xe6,
	0xcd, 0x1d, 0xb0, 0xee, 0xf4, 0x17, 0x51, 0x17, 0xda, 0xb8, 0x85, 0x08, 0xf4, 0xa4, 0xe6, 0xb7,
	0xb2, 0x55, 0x1e, 0x34, 0x32, 0x12, 0xa1, 0x4f, 0x35, 0x33, 0x0a, 0xa6, 0xc9, 0xd3, 0xcd, 0x8c,
	0xc2, 0x21, 0xf2, 0x8c, 0x76, 0x43, 0x84, 0x64, 0x93, 0x4b, 0x18, 0xd0, 0x1d, 0xe6, 0x59, 0xbf,
	0x44, 0x56, 0x24, 0x20, 0x6d, 0x5e, 0xbb, 0x47, 0x7c, 0xeb, 0x97, 0xa8, 0xae, 0x21, 0xf3, 0x3b,
	0x9f, 0x10, 0x8a, 0xc1, 0xea, 0x7e, 0xef, 0x13, 0x0a, 0x55, 0x34, 0xff, 0xa0, 0x69, 0x1b, 0x5b,
	0x35, 0x00, 0x1d, 0x76, 0xef, 0x49, 0x43, 0xdb, 0x38, 0x10, 0xd1, 0x7a, 0xca, 0xd0, 0xa8, 0x2e,
	0xd5, 0x72, 0x6e, 0x1a, 0xba, 0xe9, 0x95, 0x88, 0xd3, 0x86, 0x46, 0xf5, 0xc8, 0x20, 0x44, 0x9d,
	0x31, 0x74, 0x86, 0xd7, 0xa3, 0xc2, 0x64, 0xce, 0x1a, 0x3a, 0xc3, 0x47, 0xc5, 0x20, 0xe8, 0x5c,
	0x33, 0x27, 0x21, 0x27, 0x63, 0xd9, 0x09, 0x97, 0xe4, 0x7c, 0x33, 0xa7, 0x30, 0x08, 0x51, 0x17,
	0x0c, 0xd5, 0xa3, 0xad, 0xc5, 0x20, 0x8f, 0x44, 0xed, 0x18, 0x7b, 0xdb, 0x52, 0x3d, 0xea, 0x1a,
	0x1a, 0xdf, 0xb1, 0x81, 0x31, 0x78, 0x89, 0x77, 0x43, 0x63, 0x98, 0xfc, 0x7b, 0xb6, 0xaa, 0x45,
	0x16, 0x75, 0xf3, 0x8e, 0x4a, 0xaa, 0x9f, 0xdc, 0x9d, 0x55, 0xb5, 0xa8, 0x89, 0x68, 0x7d, 0x29,
	0xa3, 0x9b, 0x64, 0xf1, 0xdc, 0x5d, 0xb8, 0xc2, 0x1f, 0x7e, 0x39, 0xa3, 0x9b, 0xe4, 0x50, 0x00,
	0x22, 0x5e, 0xc9, 0x28, 0x6d, 0xd3, 0x55, 0x69, 0xe8, 0xbe, 0x9a, 0x51, 0xda, 0x75, 0x8d, 0xce,
	0xd9, 0xcc, 0x7f, 0xc6, 0x38, 0x71, 0x3c, 0x9b, 0x29, 0x67, 0xa6, 0x3f, 0x60, 0x33, 0xff, 0x19,
	0xd3, 0xd0, 0xe9, 0x64, 0xcd, 0x68, 0xa0, 0x17, 0x01, 0x1a, 0x26, 0x34, 0x98, 0xae, 0xf7, 0xff,
	0x99, 0xd1, 0x40, 0x6f, 0xc8, 0x68, 0xff, 0x2b, 0xe3, 0xcb, 0xd9, 0x22, 0xaf, 0x8b, 0xcc, 0x2a,
	0x1f, 0x84, 0x7b, 0xeb, 0x6f, 0xff, 0x6e, 0xae, 0xd6, 0x61, 0x37, 0x5c, 0x99, 0xf2, 0x17, 0xdd,
	0x9a, 0x86, 0xec, 0x6f, 0xa6, 0xa8, 0x24, 0x42, 0x83, 0x14, 0xe1, 0xaa, 0x1c, 0x1c, 0x50, 0x49,
	0x02, 0x11, 0xad, 0x87, 0x06, 0x74, 0x03, 0x2c, 0x55, 0xfc, 0x0e, 0x6a, 0xc3, 0x04, 0x48, 0x03,
	0xe1, 0x05, 0xe2, 0xf0, 0x80, 0x6e, 0x80, 0xff, 0x16, 0x4b, 0x9f, 0x42, 0x03, 0xea, 0xf1, 0x32,
	0xb8, 0x27, 0x6c, 0xd4, 0x25, 0x0b, 0x11, 0x8f, 0x0e, 0xa8, 0xc7, 0x47, 0x06, 0x21, 0xea, 0xd8,
	0x80, 0xf6, 0x9d, 0x88, 0x6c, 0x3c, 0x05, 0xee, 0x33, 0xb8, 0x03, 0xe1, 0x5b, 0xee, 0x99, 0xa6,
	0x7d, 0x37, 0x2a, 0x86, 0x0e, 0x86, 0x69, 0x7e, 0x3b, 0x5b, 0x3f, 0x04, 0x8a, 0xba, 0x22, 0x49,
	0xa0, 0xf8, 0x50, 0xcb, 0x7a, 0x45, 0xe5, 0x89, 0xba, 0x77, 0x9a, 0xdf, 0xc1, 0x36, 0x0c, 0x51,
	0x47, 0x19, 0xf0, 0x27, 0xf6, 0x4d, 0xfb, 0x37, 0xaa, 0x3b, 0x1a, 0x95, 0xde, 0x5f, 0x70, 0x87,
	0x12, 0xc1, 0x0f, 0x68, 0x61, 0x4c, 0x51, 0xde, 0xe2, 0x63, 0x0c, 0x33, 0x39, 0x30, 0xcd, 0xef,
	0x64, 0x1b, 0x87, 0xb8, 0x23, 0x1d, 0x98, 0xca, 0x5b, 0xd3, 0xf7, 0x5c, 0x5f, 0xfc, 0x47, 0xe2,
	0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x24, 0x1b, 0x0e, 0xf8, 0xac, 0x10, 0x00, 0x00,
}
