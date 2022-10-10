package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/lj-211/wechat_work/config"
	"github.com/lj-211/wechat_work/internal/util"
)

var gUtripCorpCode = ""
var gUtripKey = ""

// http://ubtrip.eatuo.com:9081/#/singleLogin?&user=apitest@ssharing.com&usertype=2&name=测试&corpcode=TestCorp&sign=MD5(user+corpcode+key+2018-01-01)&type=flight&sdate=1502985600&scity=SHA&ecity=CAN
// 参数
//		user = 手机
//		usertype = 3
//		name = 企业微信名称
//		corpdor = TestCorp
//		md5
//		type = "home"
//
// 正式
//		http://www.ssharing.net:9081/#/singleLogin
func PackUtripUrl(user_phone string, user_name string) string {
	env := config.GetEnvironment()

	base_url := getUtripBaseUrl(env)

	base_url += ("&user=" + user_phone)
	base_url += "&usertype=3"
	base_url += ("&name=" + user_name)
	base_url += ("&corpcode=" + gUtripCorpCode)
	base_url += ("&sign=" + utripSign(user_phone, gUtripCorpCode, gUtripKey))
	base_url += ("&type=home")

	return base_url
}

func getUtripBaseUrl(env string) string {
	base_url := ""
	switch env {
	case "production":
		base_url = "http://www.ssharing.net:9081/#/singleLogin?"
	default:
		base_url = "http://ubtrip.eatuo.com:9081/#/singleLogin?"
	}

	return base_url
}

func utripSign(user string, corp_code string, key string) string {
	// MD5（user + corpcode+ Key + 当前日期（ yyyy-mm-dd））
	sign_str := user + corp_code + key + time.Now().Format("2006-01-02")

	h := md5.New()
	h.Write([]byte(sign_str))

	ret := hex.EncodeToString(h.Sum(nil))

	return ret
}

func SetUtripKey(corp_code string, key string) error {

	if corp_code == "" {
		return fmt.Errorf("无效的utrip corp code %w", util.ParamError)
	}

	if key == "" {
		return fmt.Errorf("无效的utrip key %w", util.ParamError)
	}

	gUtripCorpCode = corp_code
	gUtripKey = key

	return nil
}
