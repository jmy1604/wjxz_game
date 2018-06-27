package main

import (
	"libs/log"
	"main/table_config"
	"math/rand"
	"net/http"
	"public_message/gen_go/client_message"
	"public_message/gen_go/client_message_id"

	_ "time"

	"github.com/golang/protobuf/proto"
)

const (
	ROLE_MAX_COUNT = 300
)

func (this *dbPlayerRoleColumn) BuildMsg() (roles []*msg_client_message.Role) {
	this.m_row.m_lock.UnSafeRLock("dbPlayerRoleColumn.BuildMsg")
	defer this.m_row.m_lock.UnSafeRUnlock()

	for _, v := range this.m_data {
		is_lock := false
		if v.IsLock > 0 {
			is_lock = true
		}
		role := &msg_client_message.Role{
			Id:      v.Id,
			TableId: v.TableId,
			Rank:    v.Rank,
			Level:   v.Level,
			IsLock:  is_lock,
			Equips:  v.Equip,
		}
		roles = append(roles, role)
	}
	return
}

func (this *dbPlayerRoleColumn) BuildSomeMsg(ids []int32) (roles []*msg_client_message.Role) {
	this.m_row.m_lock.UnSafeRLock("dbPlayerRoleColumn.BuildOneMsg")
	defer this.m_row.m_lock.UnSafeRUnlock()

	for i := 0; i < len(ids); i++ {
		v, o := this.m_data[ids[i]]
		if !o {
			return
		}

		is_lock := false
		if v.IsLock > 0 {
			is_lock = true
		}
		role := &msg_client_message.Role{
			Id:      v.Id,
			TableId: v.TableId,
			Rank:    v.Rank,
			Level:   v.Level,
			IsLock:  is_lock,
			Equips:  v.Equip,
		}
		roles = append(roles, role)
	}
	return
}

func (this *Player) new_role(role_id int32, rank int32, level int32) int32 {
	if this.db.Roles.NumAll() >= ROLE_MAX_COUNT {
		item_info := &msg_client_message.ItemInfo{ItemCfgId: role_id, ItemNum: 1}
		SendMail(nil, this.Id, MAIL_TYPE_SYSTEM, "", "", []*msg_client_message.ItemInfo{item_info})
		return -1
	}

	card := card_table_mgr.GetRankCard(role_id, rank)
	if card == nil {
		log.Error("Cant get role card by id[%v] rank[%v]", role_id, rank)
		return 0
	}
	var role dbPlayerRoleData
	role.TableId = role_id
	role.Id = this.db.Global.IncbyCurrentRoleId(1)
	role.Rank = rank
	role.Level = level
	this.db.Roles.Add(&role)

	this.roles_id_change_info.id_add(role.Id)

	handbook := this.db.RoleHandbook.GetRole()
	if handbook == nil {
		this.db.RoleHandbook.SetRole([]int32{role_id})
		if !this.is_handbook_adds {
			this.is_handbook_adds = true
		}
	} else {
		found := false
		for i := 0; i < len(handbook); i++ {
			if handbook[i] == role_id {
				if !this.is_handbook_adds {
					this.is_handbook_adds = true
				}
				found = true
				break
			}
		}
		if !found {
			handbook = append(handbook, role_id)
			this.db.RoleHandbook.SetRole(handbook)
		}
	}

	log.Debug("Player[%v] create new role[%v] table_id[%v]", this.Id, role.Id, role_id)

	return role.Id
}

func (this *Player) has_role(id int32) bool {
	all := this.db.Roles.GetAllIndex()
	for i := 0; i < len(all); i++ {
		table_id, o := this.db.Roles.GetTableId(all[i])
		if o && table_id == id {
			return true
		}
	}
	return false
}

func (this *Player) rand_role() int32 {
	if card_table_mgr.Array == nil {
		return 0
	}

	c := len(card_table_mgr.Array)
	r := rand.Intn(c)
	cr := r
	table_id := int32(0)
	for {
		table_id = card_table_mgr.Array[r%c].Id
		if !this.has_role(table_id) {
			break
		}
		r += 1
		if r-cr >= c {
			// 允许重复
			//table_id = 0
			break
		}
	}

	id := int32(0)
	if table_id > 0 {
		id = this.db.Global.IncbyCurrentRoleId(1)
		this.db.Roles.Add(&dbPlayerRoleData{
			Id:      id,
			TableId: table_id,
			Rank:    1,
			Level:   1,
		})

		this.roles_id_change_info.id_add(id)
		log.Debug("Player[%v] rand role[%v]", this.Id, table_id)
	}

	return id
}

func (this *Player) delete_role(role_id int32) bool {
	if !this.db.Roles.HasIndex(role_id) {
		return false
	}
	this.db.Roles.Remove(role_id)
	if this.team_member_mgr != nil {
		m := this.team_member_mgr[role_id]
		if m != nil {
			delete(this.team_member_mgr, role_id)
			team_member_pool.Put(m)
		}
	}
	this.roles_id_change_info.id_remove(role_id)
	return true
}

func (this *Player) check_and_send_roles_change() {
	if this.roles_id_change_info.is_changed() {
		var msg msg_client_message.S2CRolesChangeNotify
		if this.roles_id_change_info.add != nil {
			roles := this.db.Roles.BuildSomeMsg(this.roles_id_change_info.add)
			if roles != nil {
				msg.Adds = roles
			}
		}
		if this.roles_id_change_info.remove != nil {
			msg.Removes = this.roles_id_change_info.remove
		}
		if this.roles_id_change_info.update != nil {
			roles := this.db.Roles.BuildSomeMsg(this.roles_id_change_info.update)
			if roles != nil {
				msg.Updates = roles
			}
		}
		this.roles_id_change_info.reset()
		this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLES_CHANGE_NOTIFY), &msg)
	}
}

func (this *Player) add_init_roles() {
	var team []int32
	init_roles := global_config_mgr.GetGlobalConfig().InitRoles
	for i := 0; i < len(init_roles)/3; i++ {
		iid := this.new_role(init_roles[3*i], init_roles[3*i+1], init_roles[3*i+2])
		if team == nil {
			team = []int32{iid}
		} else if len(team) < BATTLE_TEAM_MEMBER_MAX_NUM {
			team = append(team, iid)
		}
	}
	this.db.BattleTeam.SetAttackMembers(team)
	this.db.BattleTeam.SetDefenseMembers(team)
	this.db.BattleTeam.SetCampaignMembers(team)
}

func (this *Player) send_roles() {
	msg := &msg_client_message.S2CRolesResponse{}
	msg.Roles = this.db.Roles.BuildMsg()
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLES_RESPONSE), msg)
}

func (this *Player) get_team_member_by_role(role_id int32, team *BattleTeam, pos int32) (m *TeamMember) {
	var table_id, rank, level int32
	var o bool
	table_id, o = this.db.Roles.GetTableId(role_id)
	if !o {
		log.Error("Cant get table id by role id[%v]", role_id)
		return
	}
	rank, o = this.db.Roles.GetRank(role_id)
	if !o {
		log.Error("Cant get rank by role id[%v]", role_id)
		return
	}
	level, o = this.db.Roles.GetLevel(role_id)
	if !o {
		log.Error("Cant get level by role id[%v]", role_id)
		return
	}
	role_card := card_table_mgr.GetRankCard(table_id, rank)
	if role_card == nil {
		log.Error("Cant get card by role_id[%v] and rank[%v]", table_id, rank)
		return
	}

	if this.team_member_mgr == nil {
		this.team_member_mgr = make(map[int32]*TeamMember)
	}
	m = this.team_member_mgr[role_id]
	if m == nil {
		m = team_member_pool.Get()
		this.team_member_mgr[role_id] = m
	}
	if team == nil {
		m.init_attrs_equips_skills(level, role_card, nil)
		this.role_update_suit_attr_power(role_id, true, true)
	} else {
		m.init_all(team, role_id, level, role_card, pos, nil)
	}
	return
}

func (this *Player) send_role_attrs(role_id int32) int32 {
	if !this.db.Roles.HasIndex(role_id) {
		log.Error("Player[%v] no role[%v], send attrs failed", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}

	m := this.get_team_member_by_role(role_id, nil, -1)
	if m == nil {
		log.Error("Player[%v] get team member with role[%v] failed, cant send role attrs", this.Id, role_id)
		return -1
	}

	power := this.roles_power[role_id]
	response := &msg_client_message.S2CRoleAttrsResponse{
		RoleId: role_id,
		Attrs:  m.attrs,
		Power:  power,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_ATTRS_RESPONSE), response)
	log.Debug("Player[%v] send role[%v] attrs: %v  power: %v", this.Id, role_id, m.attrs, power)
	return 1
}

func (this *Player) lock_role(role_id int32, is_lock bool) int32 {
	if !this.db.Roles.HasIndex(role_id) {
		log.Error("Player[%v] not found role[%v], lock failed", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}
	if is_lock {
		this.db.Roles.SetIsLock(role_id, 1)
	} else {
		this.db.Roles.SetIsLock(role_id, 0)
	}

	this.roles_id_change_info.id_update(role_id)

	response := &msg_client_message.S2CRoleLockResponse{
		RoleId: role_id,
		IsLock: is_lock,
	}

	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_LOCK_RESPONSE), response)
	return 1
}

func (this *Player) _levelup_role(role_id, lvl int32) int32 {
	if len(levelup_table_mgr.Array) <= int(lvl) {
		log.Error("Player[%v] is already max level[%v]", this.Id, lvl)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_LEVEL_IS_MAX)
	}

	levelup_data := levelup_table_mgr.Get(lvl)
	if levelup_data == nil {
		log.Error("cant found level[%v] data", lvl)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_LEVEL_DATA_NOT_FOUND)
	}

	if levelup_data.CardLevelUpRes != nil {
		for i := 0; i < len(levelup_data.CardLevelUpRes)/2; i++ {
			resource_id := levelup_data.CardLevelUpRes[2*i]
			resource_num := levelup_data.CardLevelUpRes[2*i+1]
			now_num := this.get_resource(resource_id)
			if now_num < resource_num {
				log.Error("Player[%v] levelup role[%v] cost resource[%v] not enough, need[%v] now[%v]", this.Id, role_id, resource_id, resource_num, now_num)
				return int32(msg_client_message.E_ERR_PLAYER_ITEM_NUM_NOT_ENOUGH)
			}
			if this.tmp_cache_items == nil || len(this.tmp_cache_items) > 0 {
				this.tmp_cache_items = make(map[int32]int32)
			}
			num := this.tmp_cache_items[resource_id]
			if num == 0 {
				this.tmp_cache_items[resource_id] = resource_num
			} else {
				this.tmp_cache_items[resource_id] = num + resource_num
			}
		}
	}
	return 1
}

func (this *Player) levelup_role(role_id, up_num int32) int32 {
	if up_num == 0 {
		up_num = 1
	}

	lvl, o := this.db.Roles.GetLevel(role_id)
	if !o {
		log.Error("Player[%v] not have role[%v]", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}

	res := int32(0)
	for i := int32(1); i <= up_num; i++ {
		res = this._levelup_role(role_id, lvl+i)
		if res < 0 {
			return res
		}
	}

	for id, num := range this.tmp_cache_items {
		this.add_resource(id, -num)
	}

	this.db.Roles.SetLevel(role_id, lvl+up_num)
	this.tmp_cache_items = nil
	this.roles_id_change_info.id_update(role_id)

	response := &msg_client_message.S2CRoleLevelUpResponse{
		RoleId: role_id,
		Level:  lvl + up_num,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_LEVELUP_RESPONSE), response)

	log.Debug("Player[%v] role[%v] up to level[%v]", this.Id, role_id, lvl+up_num)

	return lvl
}

func (this *Player) rankup_role(role_id int32) int32 {
	rank, o := this.db.Roles.GetRank(role_id)
	if !o {
		log.Error("Player[%v] not have role[%v]", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}

	table_id, _ := this.db.Roles.GetTableId(role_id)
	cards := card_table_mgr.GetCards(table_id)
	if len(cards) <= int(rank) {
		log.Error("Player[%v] is already max rank[%v]", this.Id, rank)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_RANK_IS_MAX)
	}

	card_data := card_table_mgr.GetRankCard(table_id, rank)
	if card_data == nil {
		log.Error("Cant found card[%v,%v] data", table_id, rank)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_TABLE_ID_NOT_FOUND)
	}

	rank_data := rankup_table_mgr.Get(rank)
	if rank_data == nil {
		log.Error("Cant found rankup[%v] data", rank)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_RANKUP_DATA_NOT_FOUND)
	}
	var cost_resources []int32
	if card_data.Type == 1 {
		cost_resources = rank_data.Type1RankUpRes
	} else if card_data.Type == 2 {
		cost_resources = rank_data.Type2RankUpRes
	} else if card_data.Type == 3 {
		cost_resources = rank_data.Type3RankUpRes
	} else {
		log.Error("Card[%v,%v] type[%v] invalid", table_id, rank, card_data.Type)
		return -1
	}

	for i := 0; i < len(cost_resources)/2; i++ {
		resource_id := cost_resources[2*i]
		resource_num := cost_resources[2*i+1]
		rn := this.get_resource(resource_id)
		if rn < resource_num {
			log.Error("Player[%v] rank[%] up failed, resource[%v] num[%v] not enough", this.Id, rank, resource_id, rn)
			return int32(msg_client_message.E_ERR_PLAYER_ITEM_NUM_NOT_ENOUGH)
		}
	}

	for i := 0; i < len(cost_resources)/2; i++ {
		resource_id := cost_resources[2*i]
		resource_num := cost_resources[2*i+1]
		this.add_resource(resource_id, -resource_num)
	}

	rank += 1
	this.db.Roles.SetRank(role_id, rank)
	this.roles_id_change_info.id_update(role_id)

	response := &msg_client_message.S2CRoleRankUpResponse{
		RoleId: role_id,
		Rank:   rank,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_RANKUP_RESPONSE), response)

	log.Debug("Player[%v] role[%v] up rank[%v]", this.Id, role_id, rank)

	return rank
}

func get_decompose_rank_res(table_id, rank int32) []int32 {
	rank_data := rankup_table_mgr.Get(rank)
	if rank_data == nil {
		log.Error("Cant get rankup[%v] data", rank)
		return nil
	}
	var resources []int32
	card_data := card_table_mgr.GetRankCard(table_id, rank)
	if card_data == nil {
		log.Error("Cant found card[%v,%v] data", table_id, rank)
		return nil
	}
	if card_data.Type == 1 {
		resources = rank_data.Type1DecomposeRes
	} else if card_data.Type == 2 {
		resources = rank_data.Type2DecomposeRes
	} else if card_data.Type == 3 {
		resources = rank_data.Type3DecomposeRes
	} else {
		log.Error("Card[%v,%v] type[%v] invalid", table_id, rank, card_data.Type)
		return nil
	}

	return resources
}

func (this *Player) team_has_role(team_id int32, role_id int32) bool {
	var members []int32
	if team_id == BATTLE_ATTACK_TEAM {
		members = this.db.BattleTeam.GetAttackMembers()
	} else if team_id == BATTLE_DEFENSE_TEAM {
		members = this.db.BattleTeam.GetDefenseMembers()
	}
	if members != nil {
		for _, m := range members {
			if role_id == m {
				return true
			}
		}
	}
	return false
}

func (this *Player) decompose_role(role_id int32) int32 {
	level, o := this.db.Roles.GetLevel(role_id)
	if !o {
		log.Error("Player[%v] not have role[%v]", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}

	is_lock, _ := this.db.Roles.GetIsLock(role_id)
	if is_lock > 0 {
		log.Error("Player[%v] role[%v] is locked, cant decompose", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_IS_LOCKED)
	}

	/*if this.team_member_mgr[role_id] != nil {
		log.Error("Player[%v] team has role[%v], cant decompose", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_IN_TEAM_CANT_DECOMPOSE)
	}*/

	if this.team_has_role(BATTLE_ATTACK_TEAM, role_id) {
		log.Error("Player[%v] attack team has role[%v], cant decompose", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_IN_TEAM_CANT_DECOMPOSE)
	}

	if this.team_has_role(BATTLE_DEFENSE_TEAM, role_id) {
		log.Error("Player[%v] defense team has role[%v], cant decompose", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_IN_TEAM_CANT_DECOMPOSE)
	}

	rank, _ := this.db.Roles.GetRank(role_id)
	table_id, _ := this.db.Roles.GetTableId(role_id)

	card_data := card_table_mgr.GetRankCard(table_id, rank)
	if card_data == nil {
		log.Error("Not found card data by table_id[%v] and rank[%v]", table_id, rank)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_TABLE_ID_NOT_FOUND)
	}

	for i := 0; i < len(card_data.DecomposeRes)/2; i++ {
		item_id := card_data.DecomposeRes[2*i]
		item_num := card_data.DecomposeRes[2*i+1]
		this.add_resource(item_id, item_num)
		if this.tmp_cache_items == nil {
			this.tmp_cache_items = make(map[int32]int32)
		}
		this.tmp_cache_items[item_id] += item_num
	}

	levelup_data := levelup_table_mgr.Get(level)
	if levelup_data == nil {
		log.Error("Not found levelup[%v] data", level)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_LEVEL_DATA_NOT_FOUND)
	}
	if levelup_data.CardDecomposeRes != nil {
		for i := 0; i < len(levelup_data.CardDecomposeRes)/2; i++ {
			item_id := levelup_data.CardDecomposeRes[2*i]
			item_num := levelup_data.CardDecomposeRes[2*i+1]
			this.add_resource(item_id, item_num)
			if this.tmp_cache_items == nil {
				this.tmp_cache_items = make(map[int32]int32)
			}
			this.tmp_cache_items[item_id] += item_num
		}
	}

	rank_res := get_decompose_rank_res(table_id, rank)
	if rank_res != nil {
		for i := 0; i < len(rank_res)/2; i++ {
			this.add_resource(rank_res[2*i], rank_res[2*i+1])
			if this.tmp_cache_items == nil {
				this.tmp_cache_items = make(map[int32]int32)
			}
			this.tmp_cache_items[rank_res[2*i]] += rank_res[2*i+1]
		}
	}

	this.delete_role(role_id)
	role := this.team_member_mgr[role_id]
	if role != nil {
		team_member_pool.Put(role)
		delete(this.team_member_mgr, role_id)
	}

	response := &msg_client_message.S2CRoleDecomposeResponse{
		RoleId: role_id,
	}
	if this.tmp_cache_items != nil {
		for k, v := range this.tmp_cache_items {
			response.GetItems = append(response.GetItems, &msg_client_message.ItemInfo{
				ItemCfgId: k,
				ItemNum:   v,
			})
		}
		this.tmp_cache_items = nil
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_DECOMPOSE_RESPONSE), response)

	log.Debug("Player[%v] decompose role[%v]", this.Id, role_id)

	return 1
}

func (this *Player) check_fusion_role_cond(cost_role_ids []int32, cost_cond *table_config.FusionCostCond) int32 {
	for i := 0; i < len(cost_role_ids); i++ {
		if !this.db.Roles.HasIndex(cost_role_ids[i]) {
			log.Error("Player[%v] fusion role need role[%v] not found", this.Id, cost_role_ids[i])
			return int32(msg_client_message.E_ERR_PLAYER_FUSION_NEED_ROLE_NOT_FOUND)
		}

		is_lock, _ := this.db.Roles.GetIsLock(cost_role_ids[i])
		if is_lock > 0 {
			log.Error("Player[%v] role[%v] is locked, fusion check role cond failed", this.Id, cost_role_ids[i])
			return int32(msg_client_message.E_ERR_PLAYER_ROLE_IS_LOCKED)
		}

		table_id, _ := this.db.Roles.GetTableId(cost_role_ids[i])
		if cost_cond.CostId > 0 && table_id != cost_cond.CostId {
			log.Error("Player[%v] fusion cost role[%v] invalid", this.Id, cost_role_ids[i])
			return int32(msg_client_message.E_ERR_PLAYER_FUSION_ROLE_INVALID)
		} else {
			rank, _ := this.db.Roles.GetRank(cost_role_ids[i])
			card := card_table_mgr.GetRankCard(table_id, rank)
			if card == nil {
				log.Error("Player[%v] fusion role[%v] not found card[%v] with rank[%v]", this.Id, cost_role_ids[i], table_id, rank)
				return int32(msg_client_message.E_ERR_PLAYER_FUSION_ROLE_INVALID)
			}
			if cost_cond.CostCamp > 0 && card.Camp != cost_cond.CostCamp {
				log.Error("Player[%v] fusion role[%v] camp[%v] invalid", this.Id, cost_role_ids[i], card.Camp)
				return int32(msg_client_message.E_ERR_PLAYER_FUSION_ROLE_INVALID)
			}
			if cost_cond.CostType > 0 && card.Type != cost_cond.CostType {
				log.Error("Player[%v] fusion role[%v] type[%v] invalid", this.Id, cost_role_ids[i], card.Type)
				return int32(msg_client_message.E_ERR_PLAYER_FUSION_ROLE_INVALID)
			}
			if cost_cond.CostStar > 0 && card.Rarity != cost_cond.CostStar {
				log.Error("Player[%v] fusion role[%v] star[%v] invalid", this.Id, cost_role_ids[i], card.Type)
				return int32(msg_client_message.E_ERR_PLAYER_FUSION_ROLE_INVALID)
			}
		}
	}
	return 1
}

// 返还升级升阶消耗的资源
func (this *Player) _return_role_resource(role_id int32) (items []*msg_client_message.ItemInfo) {
	lvl, _ := this.db.Roles.GetLevel(role_id)
	rank, _ := this.db.Roles.GetRank(role_id)

	for i := int32(1); i < lvl; i++ {
		levelup_data := levelup_table_mgr.Get(i)
		if levelup_data == nil {
			return
		}
		d := levelup_data.CardDecomposeRes
		if d != nil {
			for j := 0; j < len(d)/2; j++ {
				items = append(items, &msg_client_message.ItemInfo{
					ItemCfgId: d[2*j],
					ItemNum:   d[2*j+1],
				})
			}
		}
	}

	for i := int32(1); i < rank; i++ {
		rankup_data := rankup_table_mgr.Get(i)
		if rankup_data == nil {
			return
		}
		dd := [][]int32{rankup_data.Type1DecomposeRes, rankup_data.Type2DecomposeRes, rankup_data.Type3DecomposeRes}
		for _, d := range dd {
			if d == nil {
				continue
			}
			for j := 0; j < len(d)/2; j++ {
				items = append(items, &msg_client_message.ItemInfo{
					ItemCfgId: d[2*j],
					ItemNum:   d[2*j+1],
				})
			}

		}
	}

	if items != nil {
		for _, item := range items {
			this.add_resource(item.GetItemCfgId(), item.GetItemNum())
		}
	}

	return
}

func (this *Player) fusion_role(fusion_id, main_role_id int32, cost_role_ids [][]int32) int32 {
	fusion := fusion_table_mgr.Get(fusion_id)
	if fusion == nil {
		log.Error("Fusion[%v] table data not found", fusion_id)
		return int32(msg_client_message.E_ERR_PLAYER_FUSION_ROLE_TABLE_DATA_NOT_FOUND)
	}

	// 资源是否足够
	for i := 0; i < len(fusion.ResCondition)/2; i++ {
		res_id := fusion.ResCondition[2*i]
		res_num := fusion.ResCondition[2*i+1]
		rn := this.get_resource(res_id)
		if rn < res_num {
			log.Error("Player[%v] fusion[%v] resource[%v] num[%v] not enough, need %v", this.Id, fusion_id, res_id, rn, res_num)
			return int32(msg_client_message.E_ERR_PLAYER_FUSION_NEED_RESOURCE_NOT_ENOUGH)
		}
	}

	// 固定配方
	if fusion.FusionType == 1 {
		if !this.db.Roles.HasIndex(main_role_id) {
			log.Error("Player[%v] fusion[%v] not found main role[%v]", this.Id, fusion_id, main_role_id)
			return int32(msg_client_message.E_ERR_PLAYER_FUSION_MAIN_ROLE_NOT_FOUND)
		}

		is_lock, _ := this.db.Roles.GetIsLock(main_role_id)
		if is_lock > 0 {
			log.Error("Player[%v] role[%v] is locked, cant fusion", this.Id, main_role_id)
			return int32(msg_client_message.E_ERR_PLAYER_ROLE_IS_LOCKED)
		}

		main_card_id, _ := this.db.Roles.GetTableId(main_role_id)
		if main_card_id != fusion.MainCardID {
			log.Error("Player[%v] fusion[%v] main card id[%v] is invalid", this.Id, fusion_id, main_card_id)
			return int32(msg_client_message.E_ERR_PLAYER_FUSION_MAIN_CARD_INVALID)
		}

		main_role_level, _ := this.db.Roles.GetLevel(main_role_id)
		if main_role_level < fusion.MainCardLevelCond {
			log.Error("Player[%v] fusion[%v] main card id[%v] level[%v] not enough, need level[%v]", this.Id, fusion_id, main_card_id, main_role_level, fusion.MainCardLevelCond)
			return int32(msg_client_message.E_ERR_PLAYER_FUSION_MAIN_CARD_INVALID)
		}
	} else {
		if this.db.Roles.NumAll() >= global_config_mgr.GetGlobalConfig().MaxRoleCount {
			log.Error("Player[%v] role inventory is full", this.Id)
			return int32(msg_client_message.E_ERR_PLAYER_ROLE_INVENTORY_NOT_ENOUGH_SPACE)
		}
	}

	for i := 0; i < len(cost_role_ids); i++ {
		if i >= len(fusion.CostConds) {
			break
		}
		cn := int32(0)
		if cost_role_ids[i] != nil {
			cn = int32(len(cost_role_ids[i]))
		}
		if fusion.CostConds[i].CostNum > cn {
			log.Error("Player[%v] fusion[%v] cost num %v not enough, need %v", this.Id, fusion_id, cn, fusion.CostConds[i].CostNum)
			return int32(msg_client_message.E_ERR_PLAYER_FUSION_ROLE_MATERIAL_NOT_ENOUGH)
		}
		res := this.check_fusion_role_cond(cost_role_ids[i], fusion.CostConds[i])
		if res < 0 {
			return res
		}
	}

	var item *msg_client_message.ItemInfo
	var o bool
	if o, item = this.drop_item_by_id(fusion.ResultDropID, true, false); !o {
		log.Error("Player[%v] fusion[%v] drop new card failed", this.Id, fusion_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_FUSION_FAILED)
	}

	new_role_id := int32(0)
	if fusion.FusionType == 1 {
		this.db.Roles.SetTableId(main_role_id, item.ItemCfgId)
		this.roles_id_change_info.id_update(main_role_id)
	} else {
		new_role_id = this.new_role(item.ItemCfgId, 1, 1)
	}

	var get_items []*msg_client_message.ItemInfo
	for i := 0; i < len(cost_role_ids); i++ {
		for j := 0; j < len(cost_role_ids[i]); j++ {
			items := this._return_role_resource(cost_role_ids[i][j])
			get_items = append(get_items, items...)
			this.delete_role(cost_role_ids[i][j])
		}
	}

	for i := 0; i < len(fusion.ResCondition)/2; i++ {
		res_id := fusion.ResCondition[2*i]
		res_num := fusion.ResCondition[2*i+1]
		this.add_resource(res_id, -res_num)
	}

	this.check_and_send_roles_change()

	response := &msg_client_message.S2CRoleFusionResponse{
		NewCardId: item.ItemCfgId,
		RoleId:    new_role_id,
		GetItems:  get_items,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_FUSION_RESPONSE), response)

	log.Debug("Player[%v] fusion[%v] main_card[%v] get new role[%v] new card[%v], cost cards[%v]", this.Id, fusion_id, main_role_id, new_role_id, item.ItemCfgId, cost_role_ids)

	return 1
}

func (this *Player) get_role_handbook() int32 {
	response := &msg_client_message.S2CRoleHandbookResponse{
		Roles: this.db.RoleHandbook.GetRole(),
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_HANDBOOK_RESPONSE), response)
	return 1
}

func (this *Player) role_open_left_slot(role_id int32) int32 {
	open, ok := this.db.Roles.GetLeftSlotIsOpen(role_id)
	if !ok {
		log.Error("Player[%v] not found role[%v]", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}
	if open > 0 {
		log.Warn("Player[%v] role[%v] left slot already opened", this.Id, role_id)
	}

	this.db.Roles.SetLeftSlotIsOpen(role_id, 1)

	this.roles_id_change_info.id_update(role_id)

	response := &msg_client_message.S2CRoleLeftSlotOpenResponse{
		RoleId: role_id,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_LEFTSLOT_OPEN_RESPONSE), response)

	return 1
}

func (this *Player) role_one_key_equip(role_id int32, equips []int32) int32 {
	role_equips, o := this.db.Roles.GetEquip(role_id)
	if !o {
		log.Error("Player[%v] no role[%v], one key equip failed", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}

	if equips == nil {
		equips = make([]int32, EQUIP_TYPE_MAX)
		all_item := this.db.Items.GetAllIndex()
		for _, item_id := range all_item {
			item := item_table_mgr.Get(item_id)
			if item == nil || item.EquipType < 1 || item.EquipType >= EQUIP_TYPE_LEFT_SLOT {
				continue
			}
			eq := item_table_mgr.Get(equips[item.EquipType])
			if equips[item.EquipType] == 0 || (eq != nil && eq.BattlePower < item.BattlePower) {
				equips[item.EquipType] = item_id
			}

			if role_equips != nil && item.EquipType < int32(len(role_equips)) {
				if role_equips[item.EquipType] <= 0 {
					continue
				}
				e := item_table_mgr.Get(role_equips[item.EquipType])
				if e == nil {
					log.Warn("Player[%v] role[%v] equip type %v item %v not found table data", this.Id, role_id, item.EquipType, role_equips[item.EquipType])
					continue
				}
				// 已装备的大于背包中的，不替换
				if eq != nil && e.BattlePower >= eq.BattlePower {
					equips[item.EquipType] = role_equips[item.EquipType]
				}
			}
		}

		for i := 0; i < len(equips); i++ {
			if i >= EQUIP_TYPE_LEFT_SLOT {
				break
			}
			if equips[i] > 0 {
				if role_equips != nil && i < len(role_equips) && role_equips[i] > 0 {
					if equips[i] != role_equips[i] {
						this.del_item(equips[i], 1)
						this.add_item(role_equips[i], 1)
					}
				} else {
					this.del_item(equips[i], 1)
				}
			}
		}
	} else {
		for _, equip_id := range equips {
			if !this.db.Items.HasIndex(equip_id) {
				log.Error("Player[%v] no item[%v], role[%v] one key equip failed", this.Id, equip_id, role_id)
				return int32(msg_client_message.E_ERR_PLAYER_ITEM_NOT_FOUND)
			}
		}
		if role_equips != nil {
			for i := 0; i < len(role_equips); i++ {
				this.add_item(role_equips[i], 1)
			}
		}
		for _, equip_id := range equips {
			this.del_item(equip_id, 1)
		}
	}

	this.db.Roles.SetEquip(role_id, equips)

	this.roles_id_change_info.id_update(role_id)

	response := &msg_client_message.S2CRoleOneKeyEquipResponse{
		RoleId: role_id,
		Equips: equips,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_ONEKEY_EQUIP_RESPONSE), response)

	log.Debug("Player[%v] role[%v] one key equips[%v]", this.Id, role_id, equips)

	return 1
}

func (this *Player) role_one_key_unequip(role_id int32) int32 {
	equips, o := this.db.Roles.GetEquip(role_id)
	if !o {
		log.Error("Player[%v] not found role[%v], one key equip failed", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}

	if equips != nil {
		for i := 0; i < len(equips); i++ {
			if i == EQUIP_TYPE_LEFT_SLOT || equips[i] == 0 {
				continue
			}
			this.add_item(equips[i], 1)
			equips[i] = 0
		}
		this.db.Roles.SetEquip(role_id, equips)
	}

	this.roles_id_change_info.id_update(role_id)

	response := &msg_client_message.S2CRoleOneKeyUnequipResponse{
		RoleId: role_id,
		Equips: equips,
	}
	this.Send(uint16(msg_client_message_id.MSGID_S2C_ROLE_ONEKEY_UNEQUIP_RESPONSE), response)
	return 1
}

func (this *Player) set_power(role_id, pow int32) {
	if this.roles_power == nil {
		return
	}
	this.roles_power[role_id] = pow
}

func (this *Player) role_update_suit_attr_power(role_id int32, get_suit_attr, get_power bool) int32 {
	equips, o := this.db.Roles.GetEquip(role_id)
	if !o {
		log.Error("Player[%v] not found role[%v], update suits failed", this.Id, role_id)
		return int32(msg_client_message.E_ERR_PLAYER_ROLE_NOT_FOUND)
	}

	if equips == nil {
		return -1
	}

	power := int32(0)
	suits := make(map[*table_config.XmlSuitItem]int32)
	for _, e := range equips {
		if e <= 0 {
			continue
		}
		equip := item_table_mgr.Get(e)
		if equip == nil {
			log.Warn("Player[%v] role[%v] equip[%v] table data not found", this.Id, role_id, e)
			continue
		}

		if get_power {
			power += equip.BattlePower
		}

		if equip.SuitId <= 0 {
			continue
		}

		suit_data := suit_table_mgr.Get(equip.SuitId)
		if suit_data == nil {
			log.Warn("Suit id[%v] not found", equip.SuitId)
			continue
		}

		sn := suits[suit_data]
		if sn == 0 {
			suits[suit_data] = 1
		} else {
			suits[suit_data] = sn + 1
		}
	}

	var mem *TeamMember
	if get_suit_attr {
		mem = this.team_member_mgr[role_id]
	}

	for s, n := range suits {
		attrs := s.SuitAddAttrs[n]
		if mem != nil && attrs != nil {
			mem.add_attrs(attrs)
		}
		if get_power {
			pow := s.SuitPowers[n]
			if pow > 0 {
				power += pow
			}
		}
	}

	if get_power {
		this.set_power(role_id, power)
	}

	return 1
}

func C2SRoleAttrsHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleAttrsRequest
	err := proto.Unmarshal(msg_data, &req)
	if nil != err {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.send_role_attrs(req.GetRoleId())
}

func C2SRoleLockHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleLockRequest
	err := proto.Unmarshal(msg_data, &req)
	if nil != err {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.lock_role(req.GetRoleId(), req.GetIsLock())
}

func C2SRoleLevelUpHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleLevelUpRequest
	err := proto.Unmarshal(msg_data, &req)
	if nil != err {
		log.Error("Unmarshal msg failed err(%s) !", err.Error())
		return -1
	}
	return p.levelup_role(req.GetRoleId(), req.GetUpNum())
}

func C2SRoleRankUpHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleRankUpRequest
	err := proto.Unmarshal(msg_data, &req)
	if nil != err {
		log.Error("Unmarshal msg failed err(%s) !", err.Error())
		return -1
	}
	return p.rankup_role(req.GetRoleId())
}

func C2SRoleDecomposeHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleDecomposeRequest
	err := proto.Unmarshal(msg_data, &req)
	if nil != err {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.decompose_role(req.GetRoleId())
}

func C2SRoleFusionHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleFusionRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.fusion_role(req.GetFusionId(), req.GetMainCardId(), [][]int32{req.GetCost1CardIds(), req.GetCost2CardIds(), req.GetCost3CardIds()})
}

func C2SRoleHandbookHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleHandbookRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}

	return p.get_role_handbook()
}

func C2SRoleLeftSlotOpenHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleLeftSlotOpenRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.role_open_left_slot(req.GetRoleId())
}

func C2SRoleOneKeyEquipHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleOneKeyEquipRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.role_one_key_equip(req.GetRoleId(), req.GetEquips())
}

func C2SRoleOneKeyUnequipHandler(w http.ResponseWriter, r *http.Request, p *Player, msg_data []byte) int32 {
	var req msg_client_message.C2SRoleOnekeyUnequipRequest
	err := proto.Unmarshal(msg_data, &req)
	if err != nil {
		log.Error("Unmarshal msg failed err(%s)!", err.Error())
		return -1
	}
	return p.role_one_key_unequip(req.GetRoleId())
}
