package models

type UserInfoResult struct {
	Message string `json:"message"`
	Result  struct {
		User             string        `json:"User"`
		NickName         string        `json:"NickName"`
		RealName         string        `json:"RealName"`
		UserVerify       string        `json:"UserVerify"`
		Referrer         string        `json:"Referrer"`
		Pwd              string        `json:"Pwd"`
		RegIP            string        `json:"RegIP"`
		Time             string        `json:"Time"`
		Email            string        `json:"Email"`
		Phone            string        `json:"Phone"`
		IsLocked         bool          `json:"IsLocked"`
		IsBalanceLocked  bool          `json:"IsBalanceLocked"`
		IsWithdrawLocked bool          `json:"IsWithdrawLocked"`
		Viplv            int64         `json:"Viplv"`
		Group            string        `json:"Group"`
		Balance          interface{}   `json:"Balance"`
		Address          interface{}   `json:"Address"`
		WrAddress        []interface{} `json:"WrAddress"`
		Favourite        string        `json:"Favourite"`
		GoogleKey        string        `json:"GoogleKey"`
		GoogleVerify     bool          `json:"GoogleVerify"`
		PayPwd           bool          `json:"PayPwd"`
		FeeScale         interface{}   `json:"FeeScale"`
		Oauth            interface{}   `json:"Oauth"`
	} `json:"result"`
}
