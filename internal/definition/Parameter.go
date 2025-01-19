package definition

type ValidateTokenParameter struct {
	Token     string `json:"token"`
	PageToken string `json:"page_token"`
	TimeUnix  int64  `json:"time_unix"`
}
