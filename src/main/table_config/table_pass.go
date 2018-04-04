package table_config

import (
	"encoding/xml"
	"io/ioutil"
	"libs/log"
)

type XmlPassItem struct {
	Id               int32  `xml:"Id,attr"`
	MonsterList      string `xml:"MonsterList,attr"`
	MaxRound         int32  `xml:"MaxRound,attr"`
	TimeUpWin        int32  `xml:"TimeUpWin,attr"`
	PlayerCardMax    int32  `xml:"PlayerCardMax,attr"`
	FriendSupportMax int32  `xml:"FriendSupportMax,attr"`
	NpcSupportList   string `xml:"NpcSupportList,attr"`
}

type XmlPassConfig struct {
	Items []XmlPassItem `xml:"item"`
}

type PassTableMgr struct {
	Map   map[int32]*XmlPassItem
	Array []*XmlPassItem
}

func (this *PassTableMgr) Init() bool {
	if !this.Load() {
		log.Error("PassTableMgr Init load failed !")
		return false
	}
	return true
}

func (this *PassTableMgr) Load() bool {
	data, err := ioutil.ReadFile("../game_data/pass.xml")
	if nil != err {
		log.Error("PassTableMgr read file err[%s] !", err.Error())
		return false
	}

	tmp_cfg := &XmlPassConfig{}
	err = xml.Unmarshal(data, tmp_cfg)
	if nil != err {
		log.Error("PassTableMgr xml Unmarshal failed error [%s] !", err.Error())
		return false
	}

	if this.Map == nil {
		this.Map = make(map[int32]*XmlPassItem)
	}
	if this.Array == nil {
		this.Array = make([]*XmlPassItem, 0)
	}

	tmp_len := int32(len(tmp_cfg.Items))

	var tmp_item *XmlPassItem
	for idx := int32(0); idx < tmp_len; idx++ {
		tmp_item = &tmp_cfg.Items[idx]

		this.Map[tmp_item.Id] = tmp_item
		this.Array = append(this.Array, tmp_item)
	}

	return true
}

func (this *PassTableMgr) Get(id int32) *XmlPassItem {
	return this.Map[id]
}