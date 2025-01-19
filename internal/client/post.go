package client

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/simonks2016/YimiTV-go/internal/auth"
	"github.com/simonks2016/YimiTV-go/internal/definition"
)

func Post[data any](appId, appKey, url string, param any) (*data, error) {

	var response definition.Response[data]
	post, err := grequests.Post(url, &grequests.RequestOptions{
		Headers: map[string]string{
			"app_id":    appId,
			"signature": auth.Signature(appId, appKey, param),
		},
		JSON: param,
	})
	if err != nil {
		return nil, err
	}

	if post.Ok {
		if err = post.JSON(&response); err != nil {
			return nil, err
		} else {
			if response.Code == 200 {
				return &response.Content, nil
			} else {
				return nil, errors.New(response.Message["error_message"])
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("failed to connect api server,code:%d", post.StatusCode))
}
