package main

import (
	"github.com/lj-211/wechat_work/config"
	"github.com/lj-211/wechat_work/internal"
	log "github.com/sirupsen/logrus"
)

func main() {

	initLog()

	init_logger := log.WithFields(log.Fields{
		"module": "init",
	})

	err := config.LoadConfig()
	if err != nil {
		init_logger.Fatal("加载配置失败")
	}

	err = spawnTokenRefresher()
	if err != nil {
		init_logger.Fatal("首次获取企业微信token失败")
	}

	err = internal.SetUtripKey(config.GConfig.Utrip.UtripCorpCode,
		config.GConfig.Utrip.UtripKey)
	if err != nil {
		init_logger.Fatal("设置utrip关键配置失败")
	}

	startGin()
}

func initLog() {
	log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.SetLevel(log.DebugLevel)
}
