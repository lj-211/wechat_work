package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lj-211/wechat_work/api"
	"github.com/lj-211/wechat_work/config"
	"github.com/lj-211/wechat_work/internal"
	log "github.com/sirupsen/logrus"
)

var wechatLogger = log.WithFields(log.Fields{
	"module": "wechat",
})

func spawnTokenRefresher() error {
	is_ok := make(chan bool)

	go func() {
		var once sync.Once
		time_delta := 30 * time.Second

		try := 3

		for {
			err := internal.RefreshToken(config.GConfig.WechatWork.CorpId,
				config.GConfig.WechatWork.CorpSecret)
			if err != nil {
				wechatLogger.Errorf("初始化token失败 %s", err.Error())
			}

			try = try - 1

			if internal.IsTokenReady() {
				once.Do(func() {
					is_ok <- true
				})
				//break
			} else if try == 0 {
				is_ok <- false
				return
			}
		}

		time.Sleep(time_delta)

		for {
			_ = internal.RefreshToken(config.GConfig.WechatWork.CorpId,
				config.GConfig.WechatWork.CorpSecret)

			time.Sleep(time_delta)
		}
	}()

	for ok := true; ok; ok = false {
		result := <-is_ok
		if !result {
			return fmt.Errorf("初始化企业微信Token失败")
		}
		close(is_ok)
	}

	return nil
}

func startGin() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	wechat_work := r.Group("/wechat_work")
	{
		wechat_work.GET("/landing", api.Landing)
		wechat_work.POST("/echo", api.Echo)
	}

	err := r.Run()
	if err != nil {
		wechatLogger.Errorf("gin error %s", err.Error())
	}
}

func GetEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("获取环境变量失败")
	}

	return val, nil
}
