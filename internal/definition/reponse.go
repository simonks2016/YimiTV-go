package definition

type Response[data any] struct {
	Code           int               `json:"code"`
	Content        data              `json:"content"`
	Message        map[string]string `json:"message"`
	RedirectUrl    string            `json:"redirect_url"`
	RedirectParams string            `json:"redirect_params"`
	IsUseCaching   bool              `json:"is_use_caching"`
}
