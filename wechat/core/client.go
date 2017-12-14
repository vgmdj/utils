package core

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
)

type Client struct {
	Clt *core.Client
}

func NewClient(appid, secret string) *Client {
	default_server := core.NewDefaultAccessTokenServer(appid, secret, nil)
	client := new(core.Client)
	client.AccessTokenServer = default_server

	return &Client{
		Clt: client,
	}

}
