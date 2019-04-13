package services

import (
	"fmt"
	"testing"
)

//测试获取合约信息接口
func Test_GetOrderBook(t *testing.T) {
	orderBook := GetOrderBook("eth-usdt")
	fmt.Println("获取订单薄: ", orderBook)
}

//测试获取用户充值地址
func Test_GetNewAddress(t *testing.T) {
	address := GetNewAddress("eth")
	fmt.Println("获取用户充值地址: ", address)
}

//测试获取用户信息
func Test_GetUserInfo(t *testing.T) {
	userInfo := GetUserInfo()
	fmt.Println("获取用户信息: ", userInfo)
}

//测试获取用户余额
func Test_GetUserBalance(t *testing.T) {
	balance := GetUserBalance()
	fmt.Println("获取用户余额：", balance)
}

//测试获取交易历史记录
func Test_GetTradeHistory(t *testing.T) {
	tradeHistory := GetTradeHistory("eth-usdt")
	fmt.Println("获取交易历史记录：", tradeHistory)
}

//测试获取24小时交易数据
func Test_GetTicker(t *testing.T) {
	ticker := GetTicker("eth-usdt")
	fmt.Println("获取24小时交易数据：", ticker)
}

//测试获取Kline
func Test_GetKline(t *testing.T) {
	kline := GetKline("eth-usdt", "month")
	fmt.Println("获取Kline：", kline)
}

//测试创建订单
func Test_CreateOrder(t *testing.T) {
	create := CreateOrder(200.16, 1.96, "eth-usdt", "buy")
	fmt.Println("创建订单：", create)
}

//测试取消订单
func Test_CancelOrder(t *testing.T) {
	cancel := CancelOrder("eth-usdt", 1234564)
	fmt.Println("取消订单：", cancel)
}

//测试用户查询订单
func Test_GetOrder(t *testing.T) {
	order := GetOrder("eth-usdt", 1234564)
	fmt.Println("用户查询订单：", order)
}

//测试用户查询多条订单
func Test_GetOrders(t *testing.T) {
	orders := GetOrders("eth-usdt")
	fmt.Println("用户查询多条订单：", orders)
}

//测试用户挂单交易详情
func Test_GetTransByID(t *testing.T) {
	trans := GetTransByID("eth-usdt", 1245645)
	fmt.Println("获取用户挂单交易详情：", trans)
}

//测试获取用户历史交易
func Test_GetUserTradeHistory(t *testing.T) {
	history := GetUserTradeHistory("eth-usdt")
	fmt.Println("获取用户历史交易：", history)
}
