package models

type UserTradeHistoryResult struct {
	Message string `json:"message"`
	Result  []struct {
		P  float64 `json:"P"`  //价格
		A  float64 `json:"A"`  //数量
		T  string  `json:"T"`  //时间
		Tp string  `json:"Tp"` //买卖类型
		B  int64   `json:"B"`  //买单ID
		S  int64   `json:"S"`  //卖单ID
		ID string  `json:"ID"` //订单号
	} `json:"result"`
}
