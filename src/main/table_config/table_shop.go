package table_config

import (
	"encoding/xml"
	"io/ioutil"
	"libs/log"
)

type XmlShopItem struct {
	Id              int32  `xml:"ID,attr"`
	Type            int32  `xml:"ShopType,attr"`
	ShopMaxSlot     int32  `xml:"ShopMaxSlot,attr"`
	AutoRefreshTime string `xml:"AutoRefreshTime,attr"`
	FreeRefreshTime int32  `xml:"FreeRefreshTime,attr"`
	RefreshResStr   string `xml:"RefreshRes,attr"`
	RefreshRes      []int32
}

type XmlShopConfig struct {
	Items []*XmlShopItem `xml:"item"`
}

type ShopTableManager struct {
	shops_map   map[int32]*XmlShopItem
	shops_array []*XmlShopItem
}

func (this *ShopTableManager) Init() bool {
	data, err := ioutil.ReadFile("../game_data/Shop.xml")
	if nil != err {
		log.Error("ShopTableManager Load read file err[%s] !", err.Error())
		return false
	}

	tmp_cfg := &XmlShopConfig{}
	err = xml.Unmarshal(data, tmp_cfg)
	if nil != err {
		log.Error("ShopTableManager Load xml Unmarshal failed error [%s] !", err.Error())
		return false
	}

	this.shops_map = make(map[int32]*XmlShopItem)
	this.shops_array = []*XmlShopItem{}
	for i := 0; i < len(tmp_cfg.Items); i++ {
		c := tmp_cfg.Items[i]
		c.RefreshRes = parse_xml_str_arr2(c.RefreshResStr, ",")
		if c.RefreshRes == nil || len(c.RefreshRes)%2 != 0 {
			return false
		}
		this.shops_map[c.Id] = c
		this.shops_array = append(this.shops_array, c)
	}

	log.Info("Shop table load items count(%v)", len(tmp_cfg.Items))

	return true
}

func (this *ShopTableManager) Get(shop_id int32) *XmlShopItem {
	return this.shops_map[shop_id]
}
