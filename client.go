package taobao

import (
	"fmt"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
	"regexp"
	"strconv"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppKey           string             // 应用Key
	AppSecret        string             // 密钥
	AdzoneId         int64              // mm_xxx_xxx_xxx的第三位
	ApiGormClientFun golog.ApiClientFun // 日志配置
	Debug            bool               // 日志开关
	ZapLog           *golog.ZapLog      // 日志服务
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	config        struct {
		appKey    string // 应用Key
		appSecret string // 密钥
		adzoneId  int64  // mm_xxx_xxx_xxx的第三位
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.zapLog = config.ZapLog

	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret
	c.config.adzoneId = config.AdzoneId

	c.requestClient = gorequest.NewHttp()
	c.requestClient.Uri = apiUrl

	apiGormClient := config.ApiGormClientFun()
	if apiGormClient != nil {
		c.log.client = apiGormClient
		c.log.status = true
	}

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

func (c *Client) ZkFinalPriceParseInt64(ZkFinalPrice string) int64 {
	parseInt, err := strconv.ParseInt(ZkFinalPrice, 10, 64)
	if err != nil {
		re := regexp.MustCompile("[0-9]+")
		SalesTipMap := re.FindAllString(ZkFinalPrice, -1)
		if len(SalesTipMap) == 2 {
			return gostring.ToInt64(fmt.Sprintf("%s%s", SalesTipMap[0], SalesTipMap[1])) * 10
		} else {
			return gostring.ToInt64(SalesTipMap[0]) * 100
		}
	} else {
		return parseInt * 100
	}
}

func (c *Client) CommissionRateParseInt64(CommissionRate string) int64 {
	parseInt, err := strconv.ParseInt(CommissionRate, 10, 64)
	if err != nil {
		re := regexp.MustCompile("[0-9]+")
		SalesTipMap := re.FindAllString(CommissionRate, -1)
		if len(SalesTipMap) == 2 {
			return gostring.ToInt64(fmt.Sprintf("%s%s", SalesTipMap[0], SalesTipMap[1]))
		} else {
			return gostring.ToInt64(SalesTipMap[0])
		}
	} else {
		return parseInt
	}
}

func (c *Client) CouponAmountToInt64(CouponAmount int64) int64 {
	return CouponAmount * 100
}

func (c *Client) CommissionIntegralToInt64(GoodsPrice, CouponProportion int64) int64 {
	return (GoodsPrice * CouponProportion) / 100
}
