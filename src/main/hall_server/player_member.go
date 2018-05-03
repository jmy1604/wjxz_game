package main

import (
	"libs/log"
	"main/table_config"
	"math/rand"
	_ "net/http"
	_ "public_message/gen_go/client_message"
	_ "public_message/gen_go/client_message_id"
	_ "time"

	_ "github.com/golang/protobuf/proto"
)

// 基础属性
const (
	ATTR_HP_MAX             = 1  // 最大血量
	ATTR_HP                 = 2  // 当前血量
	ATTR_MP                 = 3  // 气势
	ATTR_ATTACK             = 4  // 攻击
	ATTR_DEFENSE            = 5  // 防御
	ATTR_DODGE_COUNT        = 6  // 闪避次数
	ATTR_INJURED_MAX        = 7  // 受伤上限
	ATTR_SHIELD             = 8  // 护盾
	ATTR_CRITICAL           = 9  // 暴击率
	ATTR_CRITICAL_MULTI     = 10 // 暴击伤害倍率
	ATTR_ANTI_CRITICAL      = 11 // 抗暴率
	ATTR_BLOCK_RATE         = 12 // 格挡率
	ATTR_BLOCK_DEFENSE_RATE = 13 // 格挡减伤率
	ATTR_BREAK_BLOCK_RATE   = 14 // 破格率

	ATTR_TOTAL_DAMAGE_ADD      = 15 // 总增伤
	ATTR_CLOSE_DAMAGE_ADD      = 16 // 近战增伤
	ATTR_REMOTE_DAMAGE_ADD     = 17 // 远程增伤
	ATTR_NORMAL_DAMAGE_ADD     = 18 // 普攻增伤
	ATTR_RAGE_DAMAGE_ADD       = 19 // 怒气增伤
	ATTR_TOTAL_DAMAGE_SUB      = 20 // 总减伤
	ATTR_CLOSE_DAMAGE_SUB      = 21 // 近战减伤
	ATTR_REMOTE_DAMAGE_SUB     = 22 // 远程减伤
	ATTR_NORMAL_DAMAGE_SUB     = 23 // 普攻减伤
	ATTR_RAGE_DAMAGE_SUB       = 24 // 怒气减伤
	ATTR_CLOSE_VAMPIRE         = 25 // 近战吸血
	ATTR_REMOTE_VAMPIRE        = 26 // 远程吸血
	ATTR_CURE_RATE_CORRECT     = 27 // 治疗率修正
	ATTR_CURED_RATE_CORRECT    = 28 // 被治疗率修正
	ATTR_CLOSE_REFLECT         = 29 // 近战反击系数
	ATTR_REMOTE_REFLECT        = 30 // 远程反击系数
	ATTR_ARMOR_ADD             = 31 // 护甲增益
	ATTR_BREAK_ARMOR           = 32 // 破甲
	ATTR_POISON_INJURED_RESIST = 33 // 毒气受伤抗性
	ATTR_BURN_INJURED_RESIST   = 34 // 点燃受伤抗性
	ATTR_BLEED_INJURED_RESIST  = 35 // 流血受伤抗性
	ATTR_HP_PERCENT_BONUS      = 36 // 血量百分比
	ATTR_ATTACK_PERCENT_BONUS  = 37 // 攻击百分比
	ATTR_DEFENSE_PERCENT_BONUS = 38 // 防御百分比
	ATTR_DAMAGE_PERCENT_BONUS  = 39 // 伤害百分比
	ATTR_COUNT_MAX             = 40
)

// 战斗结束类型
const (
	BATTLE_END_BY_ALL_DEAD   = 1 // 一方全死
	BATTLE_END_BY_ROUND_OVER = 2 // 回合用完
)

// 最大回合数
const (
	BATTLE_ROUND_MAX_NUM = 30
)

const (
	BATTLE_TEAM_MEMBER_INIT_ENERGY       = 3 // 初始能量
	BATTLE_TEAM_MEMBER_MAX_ENERGY        = 6 // 最大能量
	BATTLE_TEAM_MEMBER_ADD_ENERGY        = 2 // 能量增加量
	BATTLE_TEAM_MEMBER_MAX_NUM           = 9 // 最大人数
	BATTLE_FORMATION_LINE_NUM            = 3 // 阵型列数
	BATTLE_FORMATION_ONE_LINE_MEMBER_NUM = 3 // 每列人数
)

// 阵容类型
const (
	BATTLE_ATTACK_TEAM  = 1
	BATTLE_DEFENSE_TEAM = 2
	BATTLE_STAGE_TEAM   = 99
)

type MemberPassiveTriggerData struct {
	skill      *table_config.XmlSkillItem
	battle_num int32
	round_num  int32
}

type DelaySkill struct {
	trigger_event int32
	skill         *table_config.XmlSkillItem
	user          *TeamMember
	target_team   *BattleTeam
	trigger_pos   []int32
}

type TeamMember struct {
	team                    *BattleTeam
	pos                     int32
	id                      int32
	level                   int32
	card                    *table_config.XmlCardItem
	hp                      int32
	energy                  int32
	attack                  int32
	defense                 int32
	act_num                 int32                                 // 行动次数
	attrs                   []int32                               // 属性
	bufflist_arr            []*BuffList                           // BUFF
	passive_triggers        map[int32][]*MemberPassiveTriggerData // 被动技触发事件
	temp_normal_skill       int32                                 // 临时普通攻击
	temp_super_skill        int32                                 // 临时怒气攻击
	use_temp_skill          bool                                  // 是否使用临时技能
	temp_changed_attrs      []int32                               // 临时改变的属性
	temp_changed_attrs_used int32                                 // 临时改变属性计算状态 0 忽略 1 已初始化 2 已计算
	delay_skills            []*DelaySkill                         // 延迟的技能效果
	passive_skills          map[int32]int32                       // 被动技
}

func (this *TeamMember) add_attrs(attrs []int32) {
	for i := 0; i < len(attrs)/2; i++ {
		attr := attrs[2*i]
		this.add_attr(attr, attrs[2*i+1])
	}
}

func (this *TeamMember) add_skill_attr(skill_id int32) {
	skill := skill_table_mgr.Get(skill_id)
	if skill == nil {
		return
	}
	this.add_attrs(skill.SkillAttr)
}

func (this *TeamMember) init_passive_data(skills []int32) {
	if skills == nil {
		return
	}
	for i := 0; i < len(skills); i++ {
		if !this.add_passive_trigger(skills[i]) {
			log.Warn("Team[%v] member[%v] add passive skill[%v] failed", this.team.side, this.pos, skills[i])
		} else {
			log.Debug("Team[%v] member[%v] add passive skill[%v]", this.team.side, this.pos, skills[i])
		}
	}
}

func (this *TeamMember) init_passive_round_num() bool {
	if this.passive_triggers == nil {
		return false
	}
	for _, d := range this.passive_triggers {
		for i := 0; i < len(d); i++ {
			if d[i] != nil && d[i].skill.TriggerRoundMax > 0 {
				d[i].round_num = d[i].skill.TriggerRoundMax
			}
		}
	}
	return true
}

func (this *TeamMember) add_passive_trigger(skill_id int32) bool {
	skill := skill_table_mgr.Get(skill_id)
	if skill == nil || skill.Type != SKILL_TYPE_PASSIVE {
		return false
	}

	if this.passive_skills == nil {
		this.passive_skills = make(map[int32]int32)
	}
	if _, o := this.passive_skills[skill_id]; o {
		return false
	}

	if this.passive_triggers == nil {
		this.passive_triggers = make(map[int32][]*MemberPassiveTriggerData)
	}

	d := passive_trigger_data_pool.Get()
	d.skill = skill
	d.battle_num = skill.TriggerBattleMax
	d.round_num = skill.TriggerRoundMax
	if d.battle_num == 0 {
		d.battle_num = -1
	}
	if d.round_num == 0 {
		d.round_num = -1
	}
	datas := this.passive_triggers[skill.SkillTriggerType]
	if datas == nil {
		this.passive_triggers[skill.SkillTriggerType] = []*MemberPassiveTriggerData{d}
	} else {
		this.passive_triggers[skill.SkillTriggerType] = append(datas, d)
	}

	this.passive_skills[skill_id] = skill_id

	return true
}

func (this *TeamMember) delete_passive_trigger(skill_id int32) bool {
	skill := skill_table_mgr.Get(skill_id)
	if skill == nil || skill.Type != SKILL_TYPE_PASSIVE {
		return false
	}

	if this.passive_skills == nil {
		this.passive_skills = make(map[int32]int32)
	}
	if _, o := this.passive_skills[skill_id]; !o {
		return false
	}

	if this.passive_triggers == nil {
		return false
	}

	triggers := this.passive_triggers[skill.SkillTriggerType]
	if triggers == nil {
		return false
	}

	l := len(triggers)
	i := l - 1
	for ; i >= 0; i-- {
		if triggers[i] == nil {
			continue
		}
		if triggers[i].skill.Id == skill_id {
			passive_trigger_data_pool.Put(triggers[i])
			triggers[i] = nil
			break
		}
	}

	if i >= 0 {
		for n := i; n < l-1; n++ {
			triggers[n] = triggers[n+1]
		}
		if l > 1 {
			this.passive_triggers[skill.SkillTriggerType] = triggers[:l-1]
		} else {
			delete(this.passive_triggers, skill.SkillTriggerType)
		}
	}

	delete(this.passive_skills, skill_id)

	return true
}

func (this *TeamMember) can_passive_trigger(trigger_event int32, skill_id int32) (trigger bool) {
	d, o := this.passive_triggers[trigger_event]
	if !o || d == nil {
		return
	}

	for i := 0; i < len(d); i++ {
		if d[i] == nil {
			continue
		}
		if d[i].skill.Id != skill_id {
			continue
		}
		if d[i].battle_num != 0 && d[i].round_num != 0 {
			trigger = true
		}
		break
	}

	return
}

func (this *TeamMember) used_passive_trigger_count(trigger_event int32, skill_id int32) {
	d, o := this.passive_triggers[trigger_event]
	if !o || d == nil {
		return
	}

	for i := 0; i < len(d); i++ {
		if d[i] != nil && d[i].skill.Id == skill_id {
			if d[i].battle_num > 0 {
				d[i].battle_num -= 1
				log.Debug("Team[%v] member[%v] 减少一次技能[%v]战斗触发事件[%v]次数", this.team.side, this.pos, skill_id, trigger_event)
			}
			if d[i].round_num > 0 {
				d[i].round_num -= 1
				log.Debug("Team[%v] member[%v] 减少一次技能[%v]回合触发事件[%v]次数", this.team.side, this.pos, skill_id, trigger_event)
			}
			if d[i].battle_num == 0 || d[i].round_num == 0 {
				//passive_trigger_data_pool.Put(d[i])
			}
			break
		}
	}
}

func (this *TeamMember) has_trigger_event(trigger_events []int32) bool {
	n := 0
	for i := 0; i < len(trigger_events); i++ {
		d, o := this.passive_triggers[trigger_events[i]]
		if !o || d == nil {
			break
		}

		for j := 0; j < len(d); j++ {
			if d[i] == nil {
				continue
			}
			if d[i].battle_num != 0 && d[i].round_num != 0 {
				n += 1
			}
			break
		}
	}
	if n != len(trigger_events) {
		return false
	}
	return true
}

func (this *TeamMember) init_equip(equip_id int32) {
	d := item_table_mgr.Get(equip_id)
	if d == nil {
		return
	}
	this.init_passive_data(d.EquipSkill)
	if d.EquipSkill != nil {
		for i := 0; i < len(d.EquipSkill); i++ {
			this.add_skill_attr(d.EquipSkill[i])
		}
	}
	this.add_attrs(d.EquipAttr)
	log.Debug("@@@@@@@@@@@@@@############## team[%v] member[%v] init equip [%v] skill[%v]", this.team.side, this.pos, equip_id, d.EquipSkill)
}

func (this *TeamMember) init_equips() {
	equips, o := this.team.player.db.Roles.GetEquip(this.id)
	if !o {
		return
	}
	if equips == nil || len(equips) == 0 {
		return
	}
	for i := 0; i < len(equips); i++ {
		this.init_equip(equips[i])
	}
}

func (this *TeamMember) init(team *BattleTeam, id int32, level int32, role_card *table_config.XmlCardItem, pos int32) {
	if this.attrs == nil {
		this.attrs = make([]int32, ATTR_COUNT_MAX)
	} else {
		for i := 0; i < len(this.attrs); i++ {
			this.attrs[i] = 0
		}
	}

	if this.bufflist_arr != nil {
		for i := 0; i < len(this.bufflist_arr); i++ {
			this.bufflist_arr[i].clear()
			this.bufflist_arr[i].owner = this
		}
	}

	this.passive_skills = make(map[int32]int32)

	this.team = team
	this.id = id
	this.pos = pos
	this.level = level
	this.card = role_card
	this.energy = BATTLE_TEAM_MEMBER_INIT_ENERGY
	this.act_num = 0

	// 技能增加属性
	if role_card.NormalSkillID > 0 {
		this.add_skill_attr(role_card.NormalSkillID)
	}
	if role_card.SuperSkillID > 0 {
		this.add_skill_attr(role_card.SuperSkillID)
	}
	for i := 0; i < len(role_card.PassiveSkillIds); i++ {
		this.add_skill_attr(role_card.PassiveSkillIds[i])
	}

	this.init_passive_data(role_card.PassiveSkillIds)
	if this.team.team_type == BATTLE_STAGE_TEAM {
		// id表示怪物装备
		this.init_equip(id)
	} else {
		this.init_equips()
	}

	this.hp = (role_card.BaseHP + (level-1)*role_card.GrowthHP/100) * (10000 + this.attrs[ATTR_HP_PERCENT_BONUS]) / 10000
	this.attack = (role_card.BaseAttack + (level-1)*role_card.GrowthAttack/100) * (10000 + this.attrs[ATTR_ATTACK_PERCENT_BONUS]) / 10000
	this.defense = (role_card.BaseDefence + (level-1)*role_card.GrowthDefence/100) * (10000 + this.attrs[ATTR_DEFENSE_PERCENT_BONUS]) / 10000
	this.attrs[ATTR_HP_MAX] = this.hp
	this.attrs[ATTR_HP] = this.hp
	this.attrs[ATTR_ATTACK] = this.attack
	this.attrs[ATTR_DEFENSE] = this.defense
}

func (this *TeamMember) init_with_summon(user *TeamMember, team *BattleTeam, id int32, level int32, role_card *table_config.XmlCardItem, pos int32) {
	this.init(team, id, level, role_card, pos)
	for i := 0; i < len(user.attrs); i++ {
		this.attrs[i] = user.attrs[i]
	}
	// 技能增加属性
	if role_card.NormalSkillID > 0 {
		this.add_skill_attr(role_card.NormalSkillID)
	}
	if role_card.SuperSkillID > 0 {
		this.add_skill_attr(role_card.SuperSkillID)
	}
	for i := 0; i < len(role_card.PassiveSkillIds); i++ {
		this.add_skill_attr(role_card.PassiveSkillIds[i])
	}
}

func (this *TeamMember) add_attr(attr int32, value int32) {
	if attr == ATTR_HP {
		this.add_hp(value)
	} else if attr == ATTR_HP_MAX {
		this.add_max_hp(value)
	} else {
		this.attrs[attr] += value
	}
}

func (this *TeamMember) add_hp(hp int32) {
	if hp > 0 {
		if this.attrs[ATTR_HP]+hp > this.attrs[ATTR_HP_MAX] {
			this.attrs[ATTR_HP] = this.attrs[ATTR_HP_MAX]
		} else {
			this.attrs[ATTR_HP] += hp
		}
	} else if hp < 0 {
		if this.attrs[ATTR_HP]+hp < 0 {
			this.attrs[ATTR_HP] = 0
		} else {
			this.attrs[ATTR_HP] += hp
		}
	}
	this.hp = this.attrs[ATTR_HP]
	if hp != 0 && this.hp == 0 {
		log.Debug("+++++++++++++++++++++++++++ team[%v] mem[%v] 将死", this.team.side, this.pos)
	}
}

func (this *TeamMember) add_max_hp(add int32) {
	if add < 0 {
		if this.attrs[ATTR_HP_MAX]+add < this.attrs[ATTR_HP] {
			this.attrs[ATTR_HP] = this.attrs[ATTR_HP_MAX] + add
		}
	}
	this.attrs[ATTR_HP_MAX] += add
}

func (this *TeamMember) round_start() {
	this.act_num += 1
	this.init_passive_round_num()
}

func (this *TeamMember) round_end() {
	for i := 0; i < len(this.bufflist_arr); i++ {
		buffs := this.bufflist_arr[i]
		buffs.on_round_end()
	}

	for _, v := range this.passive_triggers {
		if v == nil {
			continue
		}
		for i := 0; i < len(v); i++ {
			if v[i].skill.TriggerRoundMax > 0 {
				v[i].round_num = v[i].skill.TriggerRoundMax
			}
		}
	}

	this.energy += BATTLE_TEAM_MEMBER_ADD_ENERGY
}

func (this *TeamMember) get_use_skill() (skill_id int32) {
	if this.act_num <= 0 {
		return
	}

	// 能量满用绝杀
	if this.energy >= BATTLE_TEAM_MEMBER_MAX_ENERGY {
		skill_id = this.card.SuperSkillID
	} else {
		skill_id = this.card.NormalSkillID
	}
	return
}

func (this *TeamMember) used_skill() {
	if this.energy >= BATTLE_TEAM_MEMBER_MAX_ENERGY {
		this.energy -= BATTLE_TEAM_MEMBER_MAX_ENERGY
	}
	if this.act_num > 0 {
		this.act_num -= 1
	}
}

func (this *TeamMember) add_buff(attacker *TeamMember, skill_effect []int32) (buff_id int32) {
	b := buff_table_mgr.Get(skill_effect[1])
	if b == nil {
		return
	}

	if this.bufflist_arr == nil {
		this.bufflist_arr = make([]*BuffList, BUFF_EFFECT_TYPE_COUNT)
		for i := 0; i < BUFF_EFFECT_TYPE_COUNT; i++ {
			this.bufflist_arr[i] = &BuffList{}
			this.bufflist_arr[i].owner = this
		}
	}

	// 互斥
	for i := 0; i < len(this.bufflist_arr); i++ {
		h := this.bufflist_arr[i]
		if h != nil && h.check_buff_mutex(b) {
			return
		}
	}

	if rand.Int31n(10000) >= skill_effect[2] {
		return
	}

	buff_id = this.bufflist_arr[b.Effect[0]].add_buff(attacker, b, skill_effect)
	return buff_id
}

func (this *TeamMember) has_buff(buff_id int32) bool {
	if this.bufflist_arr != nil {
		for i := 0; i < len(this.bufflist_arr); i++ {
			bufflist := this.bufflist_arr[i]
			buff := bufflist.head
			for buff != nil {
				if buff.buff.Id == buff_id {
					return true
				}
			}
		}
	}
	return false
}

func (this *TeamMember) remove_buff_effect(buff *Buff) {
	if buff.buff == nil || buff.buff.Effect == nil {
		return
	}

	if len(buff.buff.Effect) >= 2 {
		effect_type := buff.buff.Effect[0]
		if effect_type == BUFF_EFFECT_TYPE_MODIFY_ATTR {
			this.add_attr(buff.buff.Effect[1], -buff.param)
		} else if effect_type == BUFF_EFFECT_TYPE_TRIGGER_SKILL {
			this.delete_passive_trigger(buff.buff.Effect[1])
		}
	}
}

func (this *TeamMember) is_disable_normal_attack() bool {
	if this.bufflist_arr == nil {
		return false
	}
	disable := false
	bufflist := this.bufflist_arr[BUFF_EFFECT_TYPE_DISABLE_NORMAL_ATTACK]
	if bufflist.head != nil {
		disable = true
	} else {
		bufflist = this.bufflist_arr[BUFF_EFFECT_TYPE_DISABLE_ACTION]
		if bufflist.head != nil {
			disable = true
		}
	}
	return disable
}

func (this *TeamMember) is_disable_super_attack() bool {
	if this.bufflist_arr == nil {
		return false
	}
	disable := false
	bufflist := this.bufflist_arr[BUFF_EFFECT_TYPE_DISABLE_SUPER_ATTACK]
	if bufflist.head != nil {
		disable = true
	} else {
		bufflist = this.bufflist_arr[BUFF_EFFECT_TYPE_DISABLE_ACTION]
		if bufflist.head != nil {
			disable = true
		}
	}
	return disable
}

func (this *TeamMember) is_disable_attack() bool {
	if this.bufflist_arr == nil {
		return false
	}
	disable := false
	bufflist := this.bufflist_arr[BUFF_EFFECT_TYPE_DISABLE_ACTION]
	if bufflist.head != nil {
		disable = true
	} else {
		bufflist = this.bufflist_arr[BUFF_EFFECT_TYPE_DISABLE_NORMAL_ATTACK]
		if bufflist.head != nil {
			disable = true
		} else {
			bufflist = this.bufflist_arr[BUFF_EFFECT_TYPE_DISABLE_SUPER_ATTACK]
			if bufflist.head != nil {
				disable = true
			}
		}
	}
	return disable
}

func (this *TeamMember) is_dead() bool {
	if this.hp < 0 {
		return true
	}
	return false
}

func (this *TeamMember) is_will_dead() bool {
	if this.hp == 0 {
		return true
	}
	return false
}

func (this *TeamMember) set_dead() {
	this.hp = -1
	log.Debug("+++++++++++++++++++++++++ team[%v] mem[%v] 死了", this.team.side, this.pos)
}

func (this *TeamMember) has_delay_skills() bool {
	if this.delay_skills == nil {
		return false
	}
	return true
}

func (this *TeamMember) push_delay_skill(trigger_event int32, skill *table_config.XmlSkillItem, user *TeamMember, target_team *BattleTeam, trigger_pos []int32) {
	ds := delay_skill_pool.Get()
	ds.trigger_event = trigger_event
	ds.skill = skill
	ds.user = user
	ds.target_team = target_team
	ds.trigger_pos = trigger_pos
	this.delay_skills = append(this.delay_skills, ds)
}

func (this *TeamMember) delay_skills_effect(target_team *BattleTeam) {
	if this.delay_skills == nil {
		return
	}

	for i := 0; i < len(this.delay_skills); i++ {
		ds := this.delay_skills[i]
		if ds == nil {
			continue
		}

		one_passive_skill_effect(ds.trigger_event, ds.skill, ds.user, ds.target_team, ds.trigger_pos, true)
	}
}

func (this *TeamMember) clear_delay_skills() {
	if this.delay_skills == nil {
		return
	}
	for i := 0; i < len(this.delay_skills); i++ {
		delay_skill_pool.Put(this.delay_skills[i])
	}
	this.delay_skills = nil
}

func (this *TeamMember) has_delay_trigger_event_skill(trigger_event int32) bool {
	if this.delay_skills == nil {
		return false
	}

	for i := 0; i < len(this.delay_skills); i++ {
		if this.delay_skills[i] == nil {
			continue
		}
		if this.delay_skills[i].trigger_event == trigger_event {
			return true
		}
	}
	return false
}

func (this *TeamMember) on_will_dead(attacker *TeamMember) {
	if passive_skill_effect_with_self_pos(EVENT_BEFORE_TARGET_DEAD, attacker, this.team, this.pos, nil, nil, true) {
		log.Debug("Team[%v] member[%v] 触发了死亡前被动技能", attacker.team.side, attacker.pos)
	}
}

func (this *TeamMember) on_after_will_dead(attacker *TeamMember) {
	passive_skill_effect_with_self_pos(EVENT_AFTER_TARGET_DEAD, attacker, this.team, this.pos, attacker.team, nil, true)
	log.Debug("+++++++++++++ Team[%v] member[%v] 触发死亡后触发器", this.team.side, this.pos)
}

func (this *TeamMember) on_dead(attacker *TeamMember) {
	// 队友死亡触发
	for pos := int32(0); pos < BATTLE_TEAM_MEMBER_MAX_NUM; pos++ {
		team_mem := this.team.members[pos]
		if team_mem == nil || team_mem.is_dead() {
			continue
		}
		if pos != this.pos {
			passive_skill_effect_with_self_pos(EVENT_AFTER_TEAMMATE_DEAD, attacker, this.team, pos, this.team, []int32{this.pos}, true)
		}
	}
	// 相对于敌方死亡时触发
	for pos := int32(0); pos < BATTLE_TEAM_MEMBER_MAX_NUM; pos++ {
		team_mem := attacker.team.members[pos]
		if team_mem == nil || team_mem.is_dead() {
			continue
		}
		passive_skill_effect_with_self_pos(EVENT_AFTER_ENEMY_DEAD, attacker, attacker.team, pos, this.team, []int32{this.pos}, true)
	}
}

func (this *TeamMember) on_battle_finish() {
	if this.passive_triggers != nil {
		for _, d := range this.passive_triggers {
			if d == nil {
				continue
			}
			for i := 0; i < len(d); i++ {
				if d[i] != nil {
					passive_trigger_data_pool.Put(d[i])
				}
			}
		}
		this.passive_triggers = nil
	}

	if this.passive_skills != nil {
		this.passive_skills = nil
	}
}