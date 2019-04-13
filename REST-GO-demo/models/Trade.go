package models

type TradeRsp struct {
	User         string  `json:"user"`         //用户名
	ID           float64 `json:"id"`           //订单id
	Pair         string  `json:"pair"`         //交易对
	Orderprice   float64 `json:"orderprice""`  //订单价格
	OrderAmount  float64 `json:"orderamount"`  //订单数量
	RemainAmount float64 `json:"remainamount"` //剩余数量
	FillAmount   float64 `json:"fillamount"`   //当为卖单: 这里代表用户卖掉eth,获取usdt的数量；当为买单 : 这里代表用户未使用eth的数量
	State        string  `json:"state"`        //订单状态
	TradeType    string  `json:"tradetype"`    //交易类型sell/buy
}

type TradeResult struct {
	Message string   `json:"message"`
	Result  TradeRsp `json:"result"`
}
