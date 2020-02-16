package models

import (
	"time"
)

type Goods struct {
	Id          int       `xorm:"not null pk default nextval('goods_id_seq'::regclass) autoincr unique INTEGER"`
	Productions string    `xorm:"not null TEXT"`
	Amount      string    `xorm:"not null TEXT"`
	Price       string    `xorm:"default ''::text TEXT"`
	Company     string    `xorm:"default ''::text TEXT"`
	Person      string    `xorm:"not null TEXT"`
	Tel         string    `xorm:"not null TEXT"`
	Area        string    `xorm:"not null TEXT"`
	Remark      string    `xorm:"default ''::text TEXT"`
	GoodsType   string    `xorm:"not null TEXT"`
	Status      string    `xorm:"default 'unverified'::text TEXT"`
	IsDelete    int       `xorm:"not null default 0 INTEGER"`
	CreateTime  time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (goods *Goods) ToGoodsInfo() *GoodsInfo {
	return &GoodsInfo{
		Amount:      goods.Amount,
		Area:        goods.Area,
		Company:     goods.Company,
		Person:      goods.Person,
		Price:       goods.Price,
		Productions: goods.Productions,
		Remark:      goods.Remark,
		Tel:         goods.Tel,
		GoodsType:   goods.GoodsType,
		Status:      goods.Status,
		Time:        goods.CreateTime.String(),
	}
}
