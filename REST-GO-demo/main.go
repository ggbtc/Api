package main

import (
	"REST-GO-demo/services"
	"log"
)

func main() {
	orderBook := services.GetOrderBook("eth-usdt")
	log.Println(orderBook)

	// address := services.GetNewAddress("eth")
	// log.Println(address)

	// info := services.GetUserInfo()
	// log.Println(info)

	// balance := services.GetUserBalance()
	// log.Println(balance)

	// history := services.GetTradeHistory("eth-usdt")
	// log.Println(history)

	// ticker := services.GetTicker("eth-usdt")
	// log.Println(ticker)

	// kline := services.GetKline("eth-usdt", "m")
	// log.Println(kline)

	// createOrder := services.CreateOrder(210, 2.9, "eth-usdt", "buy")
	// log.Println(createOrder)

	// cancel := services.CancelOrder("eth-usdt", 6034143)
	// log.Println(cancel)

	// order := services.GetOrder("eth-usdt", 6034143)
	// log.Println(order)

	// orders := services.GetOrders("eth-usdt")
	// log.Println(orders)

	// trans := services.GetTransByID("eth-usdt", 6034274)
	// log.Println(trans)

	// tradeHistory := services.GetUserTradeHistory("eth-usdt")
	// log.Println(tradeHistory)

}
