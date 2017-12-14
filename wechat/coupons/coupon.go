package coupons

import (
	"encoding/json"
	"github.com/chanxuehong/wechat.v2/mp/card"
	"log"
	"util/wechat/core"
)

func InitCoupon(appid, secret string) {

	//Create 创建卡券
	var (
		c     card.Card
		clt   = core.NewClient(appid, secret).Clt
		query = &card.BatchGetQuery{
			0,
			100,
			nil,
		}
	)

	temp, err := card.BatchGet(clt, query)
	log.Println(temp, err)

	if temp != nil {
		return
	}

	if err := json.Unmarshal([]byte(testCoupon), &c); err != nil {
		log.Println(err.Error())
		return
	}

	cardid, err := card.Create(clt, &c)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(cardid)

}
