package pkg

type CustomerSimple struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Icon       string `json:"icon"`
	Background string `json:"background"`
	BrandName  string `json:"brand_name"`
	FansAmount int64  `json:"fans_amount"`
}

type ValidTokenResult struct {
	Result   bool            `json:"result"`
	UserInfo *CustomerSimple `json:"user_info"`
}
