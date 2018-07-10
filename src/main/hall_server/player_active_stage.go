package main

import (
	"libs/log"
	"libs/utils"
	"main/table_config"
	_ "math/rand"
	"net/http"
	"public_message/gen_go/client_message"
	"public_message/gen_go/client_message_id"
	_ "sync"
	"time"

	"github.com/golang/protobuf/proto"
)

func (this *Player) _send_active_stage_data() {
	last_refresh := this.db.ActiveStage.GetLastRefreshTime()
	response := &msg_client_message.S2CActiveStageDataResponse{
		CanChallengeNum:            this.db.ActiveStage.GetCanChallengeNum(),
		MaxChallengeNum:            global_config.ActiveStageChallengeNumOfDay,
		RemainSeconds4ChallengeNum: utils.GetRemainSeconds2NextDayTime(last_refresh, global_config.ActiveStageRefreshTime),
		ChallengeNumPrice:          global_config.ActiveStageChallengeNumPrice,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ACTIVE_STAGE_DATA_RESPONSE), response)
	log.Debug("Player[%v] active stage data: %v", this.Id, response)
}

func (this *Player) check_active_stage_refresh() bool {
	// 固定时间点自动刷新
	if global_config.ActiveStageRefreshTime == "" {
		return false
	}

	now_time := int32(time.Now().Unix())
	last_refresh := this.db.ActiveStage.GetLastRefreshTime()

	if last_refresh > 0 && !utils.CheckDayTimeArrival(last_refresh, global_config.ActiveStageRefreshTime) {
		return false
	}

	this.db.ActiveStage.SetCanChallengeNum(global_config.ActiveStageChallengeNumOfDay)
	this.db.ActiveStage.SetLastRefreshTime(now_time)

	this._send_active_stage_data()

	notify := &msg_client_message.S2CActiveStageRefreshNotify{}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ACTIVE_STAGE_REFRESH_NOTIFY), notify)

	log.Debug("Player[%v] active stage refreshed", this.Id)
	return true
}

func (this *Player) send_active_stage_data() int32 {
	if this.check_active_stage_refresh() {
		return 1
	}
	this._send_active_stage_data()
	return 1
}

func (this *Player) active_stage_select_friend_role(friend_id int32, role_id int32, pos int32) int32 {
	return 1
}

func (this *Player) fight_active_stage(active_stage_id int32) int32 {
	var active_stage *table_config.XmlActiveStageItem
	active_stage = active_stage_table_mgr.Get(active_stage_id)
	if active_stage == nil {
		log.Error("Active stage %v table data not found", active_stage_id)
		return -1
	}

	stage_id := active_stage.StageId
	stage := stage_table_mgr.Get(stage_id)
	if stage == nil {
		log.Error("Active stage[%v] stage[%v] not found", active_stage_id, stage_id)
		return -1
	}

	if this.db.ActiveStage.GetCanChallengeNum() <= 0 {
		log.Error("Player[%v] active stage challenge num used out", this.Id)
		return -1
	}

	is_win, my_team, target_team, enter_reports, rounds, _ := this.FightInStage(4, stage)
	this.db.ActiveStage.IncbyCanChallengeNum(-1)
	member_damages := this.active_stage_team.common_data.members_damage
	member_cures := this.active_stage_team.common_data.members_cure
	response := &msg_client_message.S2CBattleResultResponse{
		IsWin:               is_win,
		MyTeam:              my_team,
		TargetTeam:          target_team,
		EnterReports:        enter_reports,
		Rounds:              rounds,
		MyMemberDamages:     member_damages[this.active_stage_team.side],
		TargetMemberDamages: member_damages[this.target_stage_team.side],
		MyMemberCures:       member_cures[this.active_stage_team.side],
		TargetMemberCures:   member_cures[this.target_stage_team.side],
		BattleType:          4,
		BattleParam:         active_stage_id,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_BATTLE_RESULT_RESPONSE), response)

	if is_win {
		this.send_stage_reward(stage, 4)
	}

	Output_S2CBattleResult(this, response)

	return 1
}

func C2SActiveStageDataHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SActiveStageDataRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.send_active_stage_data()
}

func C2SActiveStageSelectFriendRoleHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SActiveStageSelectAssistRoleRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.active_stage_select_friend_role(req.GetFriendId(), req.GetRoleId(), req.GetPos())
}
