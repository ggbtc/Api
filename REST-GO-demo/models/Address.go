package models

type AddressResult struct {
	Message string `json:"message"`
	Result  struct {
		Address  string `json:"Address"`
		Currency string `json:"Currency"`
		User     string `json:"User"`
	} `json:"result"`
}
