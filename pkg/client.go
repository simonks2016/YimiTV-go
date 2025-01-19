package pkg

import (
	"YimiTV-go/internal/Route"
	"YimiTV-go/internal/client"
	"YimiTV-go/internal/definition"
	"github.com/segmentio/ksuid"
	"time"
)

type Client struct {
	appId  string
	appKey string
	host   string
}

func (self *Client) SetHost(host string) {
	self.host = host
}

func NewClient(appId, appKey string) *Client {
	return &Client{
		appId:  appId,
		appKey: appKey,
		host:   "https://apiv1.emicdn.com",
	}
}

func (this *Client) ValidateToken(token string) (*ValidTokenResult, error) {

	var p = definition.ValidateTokenParameter{
		Token:     token,
		PageToken: ksuid.New().String(),
		TimeUnix:  time.Now().Unix(),
	}

	//发送post请求
	post, err := client.Post[ValidTokenResult](
		this.AppId,
		this.appKey,
		Route.GetRoute(this.host, Route.ValidateToken),
		&p)
	if err != nil {
		return nil, err
	}
	return post, nil
}
