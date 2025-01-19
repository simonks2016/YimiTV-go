package YimiTV_go

import "YimiTV-go/pkg"

func NewClient(appId, appSecret string, isDebug bool) *pkg.Client {

	cli := pkg.NewClient(appId, appSecret)
	if isDebug {
		cli.SetHost("http://127.0.0.1:8080")
	} else {
		cli.SetHost("https://apiv1.emicdn.com")
	}
	return cli
}

//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
