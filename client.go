package YimiTV_go

import "github.com/simonks2016/YimiTV-go/pkg"

func NewClient(appId, appSecret string, isDebug bool) *pkg.Client {

	cli := pkg.NewClient(appId, appSecret)
	if isDebug {
		cli.SetHost("http://127.0.0.1:8080")
	} else {
		cli.SetHost("https://apiv1.emicdn.com")
	}
	return cli
}
