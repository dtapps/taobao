package taobao

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppKey    string // 应用Key
	AppSecret string // 密钥
	AdzoneId  int64  // mm_xxx_xxx_xxx的第三位
}

// Client 实例
type Client struct {
	config struct {
		appKey    string // 应用Key
		appSecret string // 密钥
		adzoneId  int64  // mm_xxx_xxx_xxx的第三位
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret
	c.config.adzoneId = config.AdzoneId

	return c, nil
}

type ErrResp struct {
	ErrorResponse struct {
		Code      int    `json:"code"`
		Msg       string `json:"msg"`
		SubCode   string `json:"sub_code"`
		SubMsg    string `json:"sub_msg"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}
