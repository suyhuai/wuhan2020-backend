package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"wuhan2020-backend/common"
	"wuhan2020-backend/data"
	"wuhan2020-backend/models"
)

type Response struct {
	Success bool
	Data    interface{}
	Err     string
}

func Home(ctx iris.Context) {
	ctx.View("index.html")
}

func GetSupply(ctx iris.Context) {
	area := ctx.URLParam("area")
	productions := ctx.URLParam("productions")
	status := ctx.URLParam("status")
	offset, err := ctx.URLParamInt("offset")
	if err != nil {
		fmt.Println(err)
		Res(ctx, nil, err)
		return
	}

	fmt.Println(fmt.Sprintf("%s", area))
	SelectGoods, err := data.SelectGoods(productions, common.GOODS_TYPE_SUPPLY, status, area, offset)
	if err != nil {
		fmt.Println(err)
		Res(ctx, nil, err)
		return
	}

	Res(ctx, SelectGoods, nil)
}

func GetNeeds(ctx iris.Context) {
	area := ctx.URLParam("area")
	productions := ctx.URLParam("productions")
	status := ctx.URLParam("status")
	offset, err := ctx.URLParamInt("offset")
	if err != nil {
		fmt.Println(err)
		Res(ctx, nil, err)
		return
	}

	fmt.Println(fmt.Sprintf("%s", area))
	SelectGoods, err := data.SelectGoods(productions, common.GOODS_TYPE_NEEDS, status, area, offset)
	if err != nil {
		fmt.Println(err)
		Res(ctx, nil, err)
		return
	}

	Res(ctx, SelectGoods, nil)
}

func PostSupply(ctx iris.Context) {
	var supplyInfo = new(models.SupplyInfo)
	err := ctx.ReadJSON(&supplyInfo)
	if err != nil {
		Res(ctx, nil, err)
		return
	}

	goodsInfo, err := supplyInfo.ToGoodsInfo()
	if err != nil {
		Res(ctx, nil, err)
		return
	}

	err = data.SaveGoods(goodsInfo)
	if err != nil {
		Res(ctx, nil, err)
		return
	}

	Res(ctx, nil, nil)
}

func PostNeeds(ctx iris.Context) {
	var needsInfo = new(models.NeedsInfo)
	err := ctx.ReadJSON(&needsInfo)
	if err != nil {
		Res(ctx, nil, err)
		return
	}

	goodsInfo, err := needsInfo.ToGoodsInfo()
	if err != nil {
		Res(ctx, nil, err)
		return
	}

	err = data.SaveGoods(goodsInfo)
	if err != nil {
		Res(ctx, nil, err)
		return
	}

	Res(ctx, nil, nil)
}

func Res(ctx iris.Context, data interface{}, err error) {
	var errMsg string
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		errMsg = err.Error()
	}
	ctx.JSON(Response{
		Success: true,
		Data:    data,
		Err:     errMsg,
	})
}
