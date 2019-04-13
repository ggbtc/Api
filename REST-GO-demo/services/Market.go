package services

import (
	"REST-GO-demo/config"
	"REST-GO-demo/models"
	"REST-GO-demo/utils"
	"encoding/json"
	"strings"
)

const (
	LIMIT int64 = 30 //显示记录条数
)

//获取订单薄
//strSymbol交易对
//limit显示数量
func GetOrderBook(strSymbol string) models.OrderBookResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["limit"] = 30

	strRequestUrl := "/trade/public/getOrderbook"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, "")
	res := models.OrderBookResult{}
	json.Unmarshal([]byte(jsonResult), &res)

	return res
}

//获取用户充值地址
//currency 币种
//需要签名
func GetNewAddress(currency string) models.AddressResult {
	mapParams := make(map[string]interface{})
	mapParams["currency"] = currency
	mapParams["key"] = config.API_KEY //签名

	strRequestUrl := "/user/logined/getNewAddress"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)
	res := models.AddressResult{}
	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//获取用户信息
//需要签名验证
func GetUserInfo() models.UserInfoResult {
	mapParams := make(map[string]interface{})
	mapParams["key"] = config.API_KEY //签名

	strRequestUrl := "/user/logined/getInfo"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)
	res := models.UserInfoResult{}
	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//获取用户余额
//需要签名验证
//注1:以上每一个字段都可能不存在,不存在的字段按照0处理,以上数据包含6位小数点~~
//注2:-trade和-locked均代表锁定余额
//注3: 当余额为空时,字段可能不存在
func GetUserBalance() models.UserBalanceResult {
	mapParams := make(map[string]interface{})
	mapParams["key"] = config.API_KEY //签名
	strRequestUrl := "/user/logined/getBalance"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)
	res := models.UserBalanceResult{}
	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//获取交易历史记录
//strSymbol交易对
//limit显示记录条数
func GetTradeHistory(strSymbol string) models.TradeHistoryResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["limit"] = LIMIT

	strRequestUrl := "/trade/public/getTradeHistory"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, "")
	res := models.TradeHistoryResult{}
	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//获取24小时交易数据
//strSymbol交易对
func GetTicker(strSymbol string) interface{} {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)

	strRequestUrl := "/trade/public/getTicker"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, "")
	var res interface{}

	if strSymbol == "all" {
		res = models.TickerAllPairResult{}
	} else {
		res = models.TickerResult{}
	}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//获取Kline
//strSymbol交易对
//basis单位hour,minute,day,week
//start开始时间(可选)
//end结束时间(可选)
//limit条数(可选)
func GetKline(strSymbol, basis string) models.KlineResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["basis"] = basis
	// mapParams["start"] = startTime
	// mapParams["end"] = endTime
	mapParams["limit"] = LIMIT

	strRequestUrl := "/trade/public/getKline"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, "")

	res := models.KlineResult{}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//用户下单
//price价格
//amount数量
//strSymbol交易对
//tradeType交易类型(sell/buy)
//需要签名
func CreateOrder(price, amount float64, strSymbol, tradeType string) models.CreateOrderResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["price"] = price
	mapParams["amount"] = amount
	mapParams["tradeType"] = tradeType
	mapParams["key"] = config.API_KEY

	strRequestUrl := "/trade/logined/orderCreate"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)

	res := models.CreateOrderResult{}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//取消下单
//strSymbol交易对
//id订单id
//需要签名
func CancelOrder(strSymbol string, id int64) models.CancelOrderResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["id"] = id
	mapParams["key"] = config.API_KEY

	strRequestUrl := "/trade/logined/orderCancel"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)

	res := models.CancelOrderResult{}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//用户查询订单
//strSymbol交易对
//id订单号
func GetOrder(strSymbol string, id int64) models.OrderResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["id"] = id
	mapParams["key"] = config.API_KEY

	strRequestUrl := "/trade/logined/orderGet"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)

	res := models.OrderResult{}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//用户条件查询多条订单信息
//strSymbol交易对
//需要签名
func GetOrders(strSymbol string) models.OrdersResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["limit"] = LIMIT
	mapParams["key"] = config.API_KEY
	// mapParams["time"] = 1443410396717 //查询截止时间
	// mapParams["state"] = "open"	//查询订单状态
	// mapParams["tradeType"] = "buy" //查询买单buy还是卖单sell

	strRequestUrl := "/trade/logined/orderGetUser"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)

	res := models.OrdersResult{}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//用户挂单交易详情
//strSymbol交易对
//id用户挂单id
//需要签名
func GetTransByID(strSymbol string, id int64) models.OrderTransResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["id"] = id
	mapParams["key"] = config.API_KEY

	strRequestUrl := "/trade/logined/transGetByID"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)

	res := models.OrderTransResult{}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}

//获取用户历史交易
//strSymbol交易对
//需要签名
func GetUserTradeHistory(strSymbol string) models.UserTradeHistoryResult {
	mapParams := make(map[string]interface{})
	mapParams["pair"] = strings.ToLower(strSymbol)
	mapParams["key"] = config.API_KEY
	mapParams["limit"] = LIMIT

	strRequestUrl := "/trade/logined/getTradeHistory"
	jsonResult := utils.HttpGetRequest(strRequestUrl, mapParams, config.API_URL, config.SECRET_KEY)

	res := models.UserTradeHistoryResult{}

	json.Unmarshal([]byte(jsonResult), &res)
	return res
}
