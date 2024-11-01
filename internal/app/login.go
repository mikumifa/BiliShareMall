package app

import (
	"github.com/mikumifa/BiliShareMall/internal/http"
	"github.com/rs/zerolog/log"
)

type LoginInfo struct {
	Key      string `json:"key"`
	LoginUrl string `json:"login_url"`
}

func (a *App) GetLoginKeyAndUrl() LoginInfo {
	key, loginUrl := http.GetLoginKeyAndUrl()
	loginInfo := LoginInfo{
		Key:      key,
		LoginUrl: loginUrl,
	}
	return loginInfo
}

// VerifyLoginResponse 封装登录验证响应的结构体
type VerifyLoginResponse struct {
	CookieStr string `json:"cookies"`
}

func (a *App) VerifyLogin(loginKey string) VerifyLoginResponse {
	str, err := http.VerifyLogin(loginKey)
	if err != nil {
		log.Error().Err(err).Msg("VerifyLogin error")
		return VerifyLoginResponse{}
	}
	ret := VerifyLoginResponse{
		CookieStr: str,
	}
	return ret
}
