package models

type UserBalanceResult struct {
	Message string `json:"message"`
	Result  struct {
		Eth        float64 `json:"eth"`                  //eth可用余额
		EthTrade   float64 `json:"eth-trade,omitempty"`  //eth锁定余额
		EthLocked  float64 `json:"eth-locked,omitempty"` //eth锁定余额
		Usdt       float64 `json:"usdt"`
		UsdtTrade  float64 `json:"usdt-trade"`
		UsdtLocked float64 `json:"usdt-locked"`
		Btc        float64 `json:"btc"`        //btc可用余额
		BtcTrade   float64 `json:"btc-trade"`  //btc锁定余额(交易中的锁定余额)
		Hkdt       float64 `json:"hkdt"`       //hkdt可用余额
		HkdtTrade  float64 `json:"hkdt-trade"` //hkdt锁定余额(交易中的锁定余额)
	} `json:"result"`
}
