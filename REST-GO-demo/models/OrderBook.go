package models

type OD struct {
	P float64 `json:"P"` //价格
	A float64 `json:"A"` //数量
}

type OrderBookResult struct {
	Message string `json:"message"`
	Result  struct {
		Sell []OD    `json:"Sell"` //卖单
		Buy  []OD    `JSON:"Buy"`  //买单
		Time int64   `json:"Time"`
		Last float64 `json:"Last"`
	} `json:"result"`
}
