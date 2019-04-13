package models

type KlineResult struct {
	Message string `json:"message"`
	Result  []struct {
		T int64   `json:"T"` //时间戳
		S float64 `json:"S"` //开始价格
		E float64 `json:"E"` //结束价格
		H float64 `json:"H"` //最高价格
		L float64 `json:"L"` //最低价格
		A float64 `json:"A"` //数量
	} `json:"result"`
}
