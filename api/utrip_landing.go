package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lj-211/wechat_work/internal"
	log "github.com/sirupsen/logrus"
)

var gUtripApiLogger = log.WithFields(log.Fields{
	"module": "api-utrip",
})

// 1. 处理企业微信落地跳转
// 2. 拼接跳转优行商旅地址
func Landing(c *gin.Context) {
	// 1. 获取code
	code := c.Query("code")
	fmt.Println("code: ", code)

	// 2. 获取用户信息
	info, err := internal.GetUserInfo(code)

	if err != nil {
		gUtripApiLogger.Infof("获取用户信息失败 %s", err.Error())
		c.String(200, err.Error())
		return
	}

	// 3. 拼接url
	user_phone := info.Mobile
	user_name := info.Name

	url := internal.PackUtripUrl(user_phone, user_name)

	gUtripApiLogger.Debugf("Utrip 返回跳转地址 %s", url)

	// 4. 回复302跳转
	c.Redirect(http.StatusMovedPermanently, url)
}
