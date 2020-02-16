package models

import "time"

type GoodsInfo struct {
	Amount      string `json:"amount,omitempty"`
	Area        string `json:"area,omitempty"`
	Company     string `json:"company,omitempty"`
	Person      string `json:"person,omitempty"`
	Price       string `json:"price,omitempty"`
	Productions string `json:"productions,omitempty"`
	Remark      string `json:"remark,omitempty"`
	Tel         string `json:"tel,omitempty"`
	GoodsType   string `json:"type,omitempty"`
	Status      string `json:"status,omitempty"`
	Time        string `json:"time,omitempty"`
}

func (info *GoodsInfo) ToGoods() (*Goods, error) {
	return &Goods{
		Productions: info.Productions,
		Amount:      info.Amount,
		Price:       info.Price,
		Company:     info.Company,
		Person:      info.Person,
		Tel:         info.Tel,
		Area:        info.Area,
		Remark:      info.Remark,
		GoodsType:   info.GoodsType,
		Status:      info.Status,
		IsDelete:    0,
		CreateTime:  time.Now(),
	}, nil
}
