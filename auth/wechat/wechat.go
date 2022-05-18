/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-18 22:55:45
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-18 23:01:01
 */
package wechat

import (
	"fmt"

	"github.com/medivhzhan/weapp/v2"
)

type Service struct {
	AppID     string
	AppSecret string
}

func (s *Service) Resolve(code string) (string, error) {
	resp, err := weapp.Login(s.AppID, s.AppSecret, code)
	if err != nil {
		return "", fmt.Errorf("weapp.Login: %v", err)
	}
	if resp.GetResponseError() != nil {
		return "", fmt.Errorf("weapp.Login: %v", resp.GetResponseError())
	}

	return resp.OpenID, nil
}
