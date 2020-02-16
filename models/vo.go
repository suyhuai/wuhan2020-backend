package models

import (
	"fmt"
	"wuhan2020-backend/common"
)

type SupplyInfo struct {
	Amount      string `json:"amount,omitempty"`
	Area        string `json:"area,omitempty"`
	Company     string `json:"company,omitempty"`
	Person      string `json:"person,omitempty"`
	Price       string `json:"price,omitempty"`
	Productions string `json:"productions,omitempty"`
	Remark      string `json:"remark,omitempty"`
	Tel         string `json:"tel,omitempty"`
	GoodsType   string `json:"goodsType,omitempty"`
}

type NeedsInfo struct {
	Amount      string `json:"amount,omitempty"`
	Area        string `json:"area,omitempty"`
	Person      string `json:"person,omitempty"`
	Productions string `json:"productions,omitempty"`
	Remark      string `json:"remark,omitempty"`
	Tel         string `json:"tel,omitempty"`
	GoodsType   string `json:"goodsType,omitempty"`
}

func (supply *SupplyInfo) ToGoodsInfo() (*GoodsInfo, error) {
	if supply.GoodsType != common.GOODS_TYPE_SUPPLY {
		return nil, fmt.Errorf("wrong goods type, expected %s", common.GOODS_TYPE_SUPPLY)
	}

	return &GoodsInfo{
		Amount:      supply.Amount,
		Area:        supply.Area,
		Company:     supply.Company,
		Person:      supply.Person,
		Price:       supply.Price,
		Productions: supply.Productions,
		Remark:      supply.Remark,
		Tel:         supply.Tel,
		GoodsType:   supply.GoodsType,
		Status:      common.GOODS_STATUS_UNVERIFIED,
	}, nil
}

func (needs *NeedsInfo) ToGoodsInfo() (*GoodsInfo, error) {
	if needs.GoodsType != common.GOODS_TYPE_NEEDS {
		return nil, fmt.Errorf("wrong goods type, expected %s", common.GOODS_TYPE_NEEDS)
	}

	return &GoodsInfo{
		Amount:      needs.Amount,
		Area:        needs.Area,
		Company:     "",
		Person:      needs.Person,
		Price:       "",
		Productions: needs.Productions,
		Remark:      needs.Remark,
		Tel:         needs.Tel,
		GoodsType:   needs.GoodsType,
		Status:      common.GOODS_STATUS_UNVERIFIED,
	}, nil
}
