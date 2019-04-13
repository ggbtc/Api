package models

type CreateOrderResult struct {
	Message string `json:"message"`
	Result  int64  `json:"result"`
}

type CancelOrderResult struct {
	Message string `json:"message"`
	Result  string `json:"result"`
}

type OrderRsp struct {
	User         string  `json:"user"`         //用户
	ID           int64   `json:"id"`           //订单号
	Pair         string  `json:"pair"`         //交易对
	TradeType    string  `json:"tradetype"`    //买卖类型
	OrderPrice   float64 `json:"orderprice"`   //下单价格
	OrderAmount  float64 `json:"orderamount"`  //下单数量
	Time         int64   `json:"time"`         //下单时间
	IsMarket     bool    `json:"ismarket"`     //是否市场
	FeeScale     float64 `json:"feescale"`     //手续费比例
	RemainAmount float64 `json:"remainamount"` //未成交剩余数量
	FeeAmount    float64 `json:"feeamount"`    //手续费数量
	FillAmount   float64 `json:"fillamount"`   //当为卖单: 这里代表用户卖掉eth,获取usdt的数量/当为买单 : 这里代表用户未使用eth的数量
	AvgPrice     float64 `json:"avgprice"`     //成交平均价格
	State        string  `json:"state"`        //订单状态(open\closed\canceled)
	LastUpdate   string  `json:"lastupdate"`   //最后一次更新时间
}

//用户单个订单信息
type OrderResult struct {
	Message string   `json:"message"`
	Result  OrderRsp `json:"result"`
}

//用户所有订单信息
type OrdersResult struct {
	Message string     `json:"message"`
	Result  []OrderRsp `json:"result"`
}

//挂单交易详情
type OrderTransResult struct {
	Message string `json:"message"`
	Result  []struct {
		P float64 `json:"P"` //价格
		A float64 `json:"A"` //数量
		T string  `json:"T"` //时间
	} `json:"result"`
}
