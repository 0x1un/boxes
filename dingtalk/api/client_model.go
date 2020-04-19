package api

import (
	"net/http"
	"net/url"

	"github.com/0x1un/boxes/dingtalk/misc"
)

// 错误响应内容
type ErrResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Response 响应的内容
type Response struct {
	StatusCode int
	Text       []byte
	URL        string
}

type AccessTokenResponse struct {
	ErrResponse
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
	Created     int64
}

// api主结构, 所有的api都围绕此结构体
type DingTalkClient struct {
	Client      *http.Client
	Parameters  url.Values
	Data        misc.Data
	ProcessReq  CreateProcessInstanceReq
	APPKEY      string
	APPSECRET   string
	BaseURI     string
	AccessToken string
}
