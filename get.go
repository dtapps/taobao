package taobao

import "go.dtapp.net/golog"

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetAdzoneId() int64 {
	return c.config.adzoneId
}

func (c *Client) GetLogGorm() *golog.ApiClient {
	return c.log.logGormClient
}

func (c *Client) GetLogMongo() *golog.ApiClient {
	return c.log.logMongoClient
}
