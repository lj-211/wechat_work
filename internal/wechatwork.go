package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/lj-211/common/net/chttp"
	"github.com/lj-211/wechat_work/internal/util"
	log "github.com/sirupsen/logrus"
)

var gAccessToken string = ""
var gTokenMutex sync.RWMutex

var gWechatLogger = log.WithFields(log.Fields{
	"module": "wechat",
})

type UserBaseInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserId  string `json:"UserId"`
}

type UserInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserId  string `json:"userid"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}

func IsTokenReady() bool {
	gTokenMutex.RLock()
	defer gTokenMutex.RUnlock()

	return gAccessToken != ""
}

// 文档地址
// https://work.weixin.qq.com/api/doc/90000/90135/91039
func RefreshToken(corp_id string, corp_secret string) error {
	gTokenMutex.Lock()

	defer gTokenMutex.Unlock()

	url := joinGetTokenUrl(corp_id, corp_secret)

	raw_data, req_err := chttp.GetUrl(context.Background(), url)

	if req_err != nil {
		return fmt.Errorf("internal:RefreshToken: request wechat : %w", req_err)
	}

	gWechatLogger.Debugf("Token return %s", string(raw_data))

	return nil
}

func joinGetTokenUrl(corp_id string, corp_secret string) string {
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	url = fmt.Sprintf(url, corp_id, corp_secret)

	return url
}

func ReadToken(action func()) {
	gTokenMutex.RLock()
	defer gTokenMutex.RUnlock()

	action()
}

// 文档地址
// https://work.weixin.qq.com/api/doc/90000/90135/91023
func GetUserInfo(code string) (UserInfo, error) {

	base_info, err := getUserBaseInfo(code)

	if err != nil {
		return UserInfo{},
			fmt.Errorf("获取用户基础信息失败 %w", err)
	}

	gWechatLogger.Debugf("用户基础身份: %v", base_info)

	return getUserInfo(base_info.UserId)
}

func getUserBaseInfo(code string) (UserBaseInfo, error) {
	if code == "" {
		return UserBaseInfo{},
			fmt.Errorf("没有提供正确的code(%s) %w", code, util.ParamError)
	}

	base_info_url := joinGetUserBaseInfoUrl(code)

	base_raw_data, err := chttp.GetUrl(context.Background(), base_info_url)

	if err != nil {
		return UserBaseInfo{},
			fmt.Errorf("请求用户基础信息失败 %w", err)
	}

	base_info := UserBaseInfo{}
	_ = json.Unmarshal(base_raw_data, &base_info)

	if base_info.Errcode != 0 {
		return UserBaseInfo{},
			fmt.Errorf("微信返回用户基础信息失败 %s", base_info.Errmsg)
	}
	return base_info, nil
}

func joinGetUserBaseInfoUrl(code string) string {
	base_info_url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s",
		gAccessToken, code)

	return base_info_url
}

func getUserInfo(user_id string) (UserInfo, error) {

	if user_id == "" {
		return UserInfo{}, fmt.Errorf("没有提供正确的用户id(%s) %w", user_id, util.ParamError)
	}

	info_url := joinGetUserInfoUrl(user_id)

	info_raw_data, err := chttp.GetUrl(context.Background(), info_url)

	if err != nil {
		return UserInfo{},
			fmt.Errorf("请求用户信息失败 %w", err)
	}

	user_info := UserInfo{}
	_ = json.Unmarshal(info_raw_data, &user_info)

	fmt.Println("用户信息: ", user_info)

	if user_info.Errcode != 0 {
		return UserInfo{},
			fmt.Errorf("微信返回用户信息失败 %s", user_info.Errmsg)
	}

	return user_info, nil
}

func joinGetUserInfoUrl(user_id string) string {
	info_url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s",
		gAccessToken, user_id)

	return info_url
}
