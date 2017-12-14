package coupons

import (
	"github.com/extrame/goblet"
	"module"
	"time"
	"util"
)

func TempData() {
	types := []string{
		"加油券",
		"充值券",
		"非油线下券",
		"非油商城券",
		"工时券",
		"洗车券",
	}

	coupons := []module.Coupon{}

	for i := 0; i < 60; i++ {
		k := i / 10
		coupon := module.Coupon{
			CardId:       util.RandomAlphabetic(8),
			CardCode:     types[k],
			FromUserName: "1",
			Money:        100,
			CardType:     types[k],
			GetTime:      time.Now(),
			Status:       i % 2,
		}

		coupons = append(coupons, coupon)

	}

	goblet.DB.Insert(&coupons)

}
