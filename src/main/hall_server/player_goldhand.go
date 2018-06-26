package main

import (
	"libs/log"
	"net/http"
	"public_message/gen_go/client_message"
	"public_message/gen_go/client_message_id"
	_ "sync"
	"time"

	_ "3p/code.google.com.protobuf/proto"
	"github.com/golang/protobuf/proto"
)

func (this *Player) send_gold_hand() int32 {
	lvl := this.db.Info.GetLvl()
	gold_hand_data := goldhand_table_mgr.Get(lvl)
	if gold_hand_data == nil {
		log.Error("Goldhand data with level %v not found", lvl)
		return int32(msg_client_message.E_ERR_PLAYER_GOLDHAND_DATA_NOT_FOUND)
	}

	last_refresh := this.db.GoldHand.GetLastRefreshTime()
	now_time := int32(time.Now().Unix())
	remain_seconds := gold_hand_data.RefreshCD - (now_time - last_refresh)
	if remain_seconds < 0 {
		remain_seconds = 0
	}
	response := &msg_client_message.S2CGoldHandDataResponse{
		RemainRefreshSeconds: remain_seconds,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_GOLD_HAND_DATA_RESPONSE), response)
	return 1
}

func (this *Player) touch_gold(t int32) int32 {
	lvl := this.db.Info.GetLvl()
	gold_hand := goldhand_table_mgr.Get(lvl)
	if gold_hand == nil {
		log.Error("Goldhand data with level %v not found", lvl)
		return int32(msg_client_message.E_ERR_PLAYER_GOLDHAND_DATA_NOT_FOUND)
	}
	last_refresh := this.db.GoldHand.GetLastRefreshTime()
	now_time := int32(time.Now().Unix())
	if now_time-last_refresh < gold_hand.RefreshCD {
		log.Error("Player[%v] gold hand is cooling down", this.Id)
		return int32(msg_client_message.E_ERR_PLAYER_GOLDHAND_REFRESH_IS_COOLINGDOWN)
	}

	var gold, diamond int32
	if t == 1 {
		gold = gold_hand.GoldReward1
		diamond = gold_hand.GemCost1
	} else if t == 2 {
		gold = gold_hand.GoldReward2
		diamond = gold_hand.GemCost2
	} else if t == 3 {
		gold = gold_hand.GoldReward3
		diamond = gold_hand.GemCost3
	} else {
		log.Error("Gold Hand type[%v] invalid")
		return -1
	}

	if this.get_diamond() < diamond {
		log.Error("Player[%v] diamond not enough, cant touch gold %v", this.Id, t)
		return int32(msg_client_message.E_ERR_PLAYER_DIAMOND_NOT_ENOUGH)
	}

	this.add_gold(gold)
	this.add_diamond(-diamond)

	this.db.GoldHand.SetLastRefreshTime(now_time)

	response := &msg_client_message.S2CTouchGoldResponse{
		Type:                     t,
		GetGold:                  gold,
		CostDiamond:              diamond,
		NextRefreshRemainSeconds: gold_hand.RefreshCD,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_TOUCH_GOLD_RESPONSE), response)

	return 1
}

func C2SGoldHandDataHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SGoldHandDataRequest
	err := proto.Unmarshal(msg_data, &req)
	if nil != err {
		log.Error("Unmarshal msg failed err(%s)", err.Error())
		return -1
	}
	return p.send_gold_hand()
}

func C2STouchGoldHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2STouchGoldRequest
	err := proto.Unmarshal(msg_data, &req)
	if nil != err {
		log.Error("Unmarshal msg failed err(%s)", err.Error())
		return -1
	}
	return p.touch_gold(req.GetType())
}