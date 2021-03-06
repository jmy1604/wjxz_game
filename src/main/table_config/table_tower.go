package table_config

import (
	"encoding/xml"
	"io/ioutil"
	"libs/log"
)

type XmlTowerItem struct {
	Id      int32 `xml:"TowerID,attr"`
	StageId int32 `xml:"StageID,attr"`
	Next    *XmlTowerItem
}

type XmlTowerConfig struct {
	Items []XmlTowerItem `xml:"item"`
}

type TowerTableMgr struct {
	Map   map[int32]*XmlTowerItem
	Array []*XmlTowerItem
}

func (this *TowerTableMgr) Init() bool {
	if !this.Load() {
		log.Error("TowerTableMgr Init load failed !")
		return false
	}
	return true
}

func (this *TowerTableMgr) Load() bool {
	data, err := ioutil.ReadFile("../game_data/Tower.xml")
	if nil != err {
		log.Error("TowerTableMgr read file err[%s] !", err.Error())
		return false
	}

	tmp_cfg := &XmlTowerConfig{}
	err = xml.Unmarshal(data, tmp_cfg)
	if nil != err {
		log.Error("TowerTableMgr xml Unmarshal failed error [%s] !", err.Error())
		return false
	}

	if this.Map == nil {
		this.Map = make(map[int32]*XmlTowerItem)
	}
	if this.Array == nil {
		this.Array = make([]*XmlTowerItem, 0)
	}
	tmp_len := int32(len(tmp_cfg.Items))

	var tmp_item, prev *XmlTowerItem
	for idx := int32(0); idx < tmp_len; idx++ {
		tmp_item = &tmp_cfg.Items[idx]
		if prev != nil {
			prev.Next = tmp_item
		}

		prev = tmp_item

		this.Map[tmp_item.Id] = tmp_item
		this.Array = append(this.Array, tmp_item)
	}

	return true
}

func (this *TowerTableMgr) Get(tower_id int32) *XmlTowerItem {
	return this.Map[tower_id]
}
