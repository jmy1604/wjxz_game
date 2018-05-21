package table_config

import (
	"encoding/xml"
	"io/ioutil"
	"libs/log"
)

type XmlCampaignItem struct {
	Id                  int32  `xml:"CampaignID,attr"`
	StageId             int32  `xml:"StageID,attr"`
	UnlockMap           int32  `xml:"UnlockMap,attr"`
	Difficulty          int32  `xml:"Difficulty,attr"`
	ChapterMap          int32  `xml:"ChapterMap,attr"`
	StaticRewardSec     int32  `xml:"StaticRewardSec,attr"`
	StaticRewardItemStr string `xml:"StaticRewardItem,attr"`
	StaticRewardItem    []int32
	RandomDropSec       int32  `xml:"RandomDropSec,attr"`
	RandomDropIDListStr string `xml:"RandomDropIDList,attr"`
	RandomDropIDList    []int32
}

type XmlCampaignConfig struct {
	Items []XmlCampaignItem `xml:"item"`
}

type CampaignTableMgr struct {
	Map                  map[int32]*XmlCampaignItem
	Array                []*XmlCampaignItem
	Chapter2Campaigns    map[int32][]int32
	Difficulty2Campaigns map[int32][]int32
}

func (this *CampaignTableMgr) Init() bool {
	if !this.Load() {
		log.Error("CampaignTableMgr Init load failed !")
		return false
	}
	return true
}

func (this *CampaignTableMgr) Load() bool {
	data, err := ioutil.ReadFile("../game_data/Campaign.xml")
	if nil != err {
		log.Error("CampaignTableMgr read file err[%s] !", err.Error())
		return false
	}

	tmp_cfg := &XmlCampaignConfig{}
	err = xml.Unmarshal(data, tmp_cfg)
	if nil != err {
		log.Error("CampaignTableMgr xml Unmarshal failed error [%s] !", err.Error())
		return false
	}

	if this.Map == nil {
		this.Map = make(map[int32]*XmlCampaignItem)
	}
	if this.Array == nil {
		this.Array = make([]*XmlCampaignItem, 0)
	}
	if this.Chapter2Campaigns == nil {
		this.Chapter2Campaigns = make(map[int32][]int32)
	}
	if this.Difficulty2Campaigns == nil {
		this.Difficulty2Campaigns = make(map[int32][]int32)
	}

	tmp_len := int32(len(tmp_cfg.Items))

	var tmp_item *XmlCampaignItem
	for idx := int32(0); idx < tmp_len; idx++ {
		tmp_item = &tmp_cfg.Items[idx]
		tmp_item.StaticRewardItem = parse_xml_str_arr2(tmp_item.StaticRewardItemStr, ",")
		if tmp_item.StaticRewardItem == nil {
			tmp_item.StaticRewardItem = make([]int32, 0)
		}
		tmp_item.RandomDropIDList = parse_xml_str_arr2(tmp_item.RandomDropIDListStr, ",")
		if tmp_item.RandomDropIDList == nil {
			tmp_item.RandomDropIDList = make([]int32, 0)
		}

		this.Map[tmp_item.Id] = tmp_item
		this.Array = append(this.Array, tmp_item)

		c2c := this.Chapter2Campaigns[tmp_item.ChapterMap]
		if c2c == nil {
			c2c = []int32{tmp_item.Id}
			this.Chapter2Campaigns[tmp_item.ChapterMap] = c2c
		} else {
			this.Chapter2Campaigns[tmp_item.ChapterMap] = append(c2c, tmp_item.Id)
		}

		d2c := this.Difficulty2Campaigns[tmp_item.Difficulty]
		if d2c == nil {
			d2c = []int32{tmp_item.Id}
			this.Difficulty2Campaigns[tmp_item.ChapterMap] = append(d2c, tmp_item.Id)
		}
	}

	return true
}

func (this *CampaignTableMgr) Get(id int32) *XmlCampaignItem {
	return this.Map[id]
}

func (this *CampaignTableMgr) GetChapterCampaign(chapter_id int32) []int32 {
	return this.Chapter2Campaigns[chapter_id]
}

func (this *CampaignTableMgr) GetDifficultyCampaign(difficulty int32) []int32 {
	return this.Difficulty2Campaigns[difficulty]
}