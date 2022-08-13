package taobao

import (
	"context"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
)

func (c *Client) request(ctx context.Context, params map[string]interface{}) (gorequest.Response, error) {

	// 签名
	c.Sign(params)

	// 创建请求
	client := c.client

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.PgsqlDb != nil {
		go c.log.GormMiddlewareCustom(ctx, gostring.ToString(params["method"]), request, Version)
	}
	if c.config.MongoDb != nil {
		go c.log.MongoMiddlewareCustom(ctx, gostring.ToString(params["method"]), request, Version)
	}

	return request, err
}
