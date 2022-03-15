// Package wechat 是跟微信官方的服务器交互的包，不是跟微信小程序前端交互的
package wechat

import (
	"fmt"
	weapp "github.com/medivhzhan/weapp/v2"
)

type Service struct {
	AppID		string
	AppSecret	string
}

func (s *Service) Resolve(code string) (string, error) {
	resp, err := weapp.Login(s.AppID, s.AppSecret, code)
	if err != nil {
		return "", err
	}
	if err = resp.GetResponseError(); err != nil {
		return "", fmt.Errorf("weapp response err: %v", err)
	}

	return resp.OpenID, nil
}



