package table_config

import (
	"encoding/xml"
	"io/ioutil"
	"libs/log"
)

type XmlStatusItem struct {
	Id               int32  `xml:"BuffID,attr"`
	Type             int32  `xml:"Type,attr"`
	Effect           string `xml:"Effect,attr"`
	ResistCountMax   int32  `xml:"ResistCountMax,attr"`
	MutexType        int32  `xml:"MutexType,attr"`
	ResistMutexType  string `xml:"ResistMutexType,attr"`
	CancelMutexType  string `xml:"CancelMutexType,attr"`
	ResistMutexID    string `xml:"ResistMutexID,attr"`
	CancelMutexID    string `xml:"CancelMutexID,attr"`
	ResistMutexTypes []int32
	CancelMutexTypes []int32
	ResistMutexIDs   []int32
	CancelMutexIDs   []int32
}

type XmlStatusConfig struct {
	Items []XmlStatusItem `xml:"item"`
}

type StatusTableMgr struct {
	Map   map[int32]*XmlStatusItem
	Array []*XmlStatusItem
}

func (this *StatusTableMgr) Init() bool {
	if !this.Load() {
		log.Error("StatusTableMgr Init load failed !")
		return false
	}
	return true
}

func (this *StatusTableMgr) Load() bool {
	data, err := ioutil.ReadFile("../game_data/status.xml")
	if nil != err {
		log.Error("StatusTableMgr read file err[%s] !", err.Error())
		return false
	}

	tmp_cfg := &XmlStatusConfig{}
	err = xml.Unmarshal(data, tmp_cfg)
	if nil != err {
		log.Error("StatusTableMgr xml Unmarshal failed error [%s] !", err.Error())
		return false
	}

	if this.Map == nil {
		this.Map = make(map[int32]*XmlStatusItem)
	}
	if this.Array == nil {
		this.Array = make([]*XmlStatusItem, 0)
	}

	tmp_len := int32(len(tmp_cfg.Items))

	var tmp_item *XmlStatusItem
	for idx := int32(0); idx < tmp_len; idx++ {
		tmp_item = &tmp_cfg.Items[idx]

		tmp_item.ResistMutexTypes = parse_xml_str_arr(tmp_item.ResistMutexType, ",")
		if tmp_item.ResistMutexTypes == nil {
			tmp_item.ResistMutexTypes = make([]int32, 0)
		}
		tmp_item.CancelMutexTypes = parse_xml_str_arr(tmp_item.CancelMutexType, ",")
		if tmp_item.CancelMutexTypes == nil {
			tmp_item.CancelMutexTypes = make([]int32, 0)
		}
		tmp_item.ResistMutexIDs = parse_xml_str_arr(tmp_item.ResistMutexID, ",")
		if tmp_item.ResistMutexIDs == nil {
			tmp_item.ResistMutexIDs = make([]int32, 0)
		}
		tmp_item.CancelMutexIDs = parse_xml_str_arr(tmp_item.CancelMutexID, ",")
		if tmp_item.CancelMutexIDs == nil {
			tmp_item.CancelMutexIDs = make([]int32, 0)
		}

		this.Map[tmp_item.Id] = tmp_item
		this.Array = append(this.Array, tmp_item)
	}

	return true
}

func (this *StatusTableMgr) Get(id int32) *XmlStatusItem {
	return this.Map[id]
}
