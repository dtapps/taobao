package taobao

import (
	"fmt"
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
	"regexp"
	"strconv"
)

type ConfigClient struct {
	AppKey     string           // 应用Key
	AppSecret  string           // 密钥
	AdzoneId   int64            // mm_xxx_xxx_xxx的第三位
	GormClient *dorm.GormClient // 日志数据库
	LogClient  *golog.ZapLog    // 日志驱动
	LogDebug   bool             // 日志开关
}

type Client struct {
	requestClient *gorequest.App   // 请求服务
	logClient     *golog.ApiClient // 日志服务
	config        *ConfigClient    // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.requestClient = gorequest.NewHttp()
	c.requestClient.Uri = apiUrl

	if c.config.GormClient.Db != nil {
		c.logClient, err = golog.NewApiClient(&golog.ApiClientConfig{
			GormClient: c.config.GormClient,
			TableName:  logTable,
			LogClient:  c.config.LogClient,
			LogDebug:   c.config.LogDebug,
		})
		if err != nil {
			return nil, err
		}
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
