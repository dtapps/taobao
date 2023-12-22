package taobao

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type TbkShopRecommendGetResponse struct {
	TbkShopRecommendGetResponse struct {
		Results struct {
			NTbkShop []struct {
				UserId     int    `json:"user_id"`
				ShopTitle  string `json:"shop_title"`
				ShopType   string `json:"shop_type"`
				SellerNick string `json:"seller_nick"`
				PictUrl    string `json:"pict_url"`
				ShopUrl    string `json:"shop_url"`
			} `json:"n_tbk_shop"`
		} `json:"results"`
	} `json:"tbk_shop_recommend_get_response"`
}

type TbkShopRecommendGetResult struct {
	Result TbkShopRecommendGetResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newTbkShopRecommendGetResult(result TbkShopRecommendGetResponse, body []byte, http gorequest.Response) *TbkShopRecommendGetResult {
	return &TbkShopRecommendGetResult{Result: result, Body: body, Http: http}
}

// TbkShopRecommendGet 淘宝客-公用-店铺关联推荐
// https://open.taobao.com/api.htm?docId=24522&docType=2
func (c *Client) TbkShopRecommendGet(ctx context.Context, notMustParams ...gorequest.Params) (*TbkShopRecommendGetResult, error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.shop.recommend.get", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newTbkShopRecommendGetResult(TbkShopRecommendGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TbkShopRecommendGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkShopRecommendGetResult(response, request.ResponseBody, request), err
}
