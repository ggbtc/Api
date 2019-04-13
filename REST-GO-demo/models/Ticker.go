package models

type SignalPairTickerRsp struct {
	Chg    string `json:"chg"`  //增加或减少的比例
	High   string `json:"high"` //24小时最高价
	Low    string `json:"low"`  //24小时最低价
	Name   string `json:"name"` //交易对
	ID     string `json:"id"`   //排序
	Last   string `json:"last"` //上一次成交价格
	Turn   string `json:"turn"` //24小时成交额(既-后面的币种的成交数量)
	Vol    string `json:"vol"`  //24小时成交量(既-前面的币种的成交数量)
	Closed bool   `json:"closed"`
}
type TickerResult struct {
	Message string              `json:"message"`
	Result  SignalPairTickerRsp `json:"result"`
}

type TickerAllPairResult struct {
	Message string `json:"message"`
	Result  struct {
		AnkrBtc SignalPairTickerRsp `json:"ankr-btc"`
		ApkUsdt SignalPairTickerRsp `json:"apk-usdt"`
		EthUsdt SignalPairTickerRsp `json:"eth-usdt"`
	} `json:"result"`
}
