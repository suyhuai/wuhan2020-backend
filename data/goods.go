package data

import (
	"fmt"
	"strings"
	"wuhan2020-backend/models"
)

func SelectGoods(productions, goodsType, status, area string, offset int) ([]*models.GoodsInfo, error) {
	var goodsList []*models.Goods

	session := db.Table("goods").Where("is_delete = 0 AND goods_type = ?", goodsType)
	if productions != "" {
		arr := strings.Split(productions, ",")
		for _, production := range arr {
			session.Where("productions like ?", "%"+production+"%")
		}
	}

	if area != "" {
		session.Where("area=?", area)
	}

	if status != "" {
		session.Where("status=?", status)
	}

	err := session.Limit(2, offset).Asc("id").Find(&goodsList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var result []*models.GoodsInfo

	for _, goods := range goodsList {
		result = append(result, goods.ToGoodsInfo())
	}

	return result, nil
}

func SaveGoods(goodsInfo *models.GoodsInfo) error {

	goods, err := goodsInfo.ToGoods()
	if err != nil {
		return err
	}

	_, err = db.Table("goods").InsertOne(goods)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
