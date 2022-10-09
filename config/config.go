package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

var GConfig = struct {
	WechatWork struct {
		CorpId     string `default:"ayla" yaml:"corp_id"`
		CorpSecret string `yaml:"corp_secret"`
	} `yaml:"wechat_work"`
	Utrip struct {
		UtripCorpCode string `yaml:"utrip_corp_code"`
		UtripKey      string `yaml:"utrip_key"`
	}
}{}

func LoadConfig() error {
	config_path := "./config/wechat_work_config.yml"

	if env_path := os.Getenv("CONFIGOR_CONFIG_PATH"); env_path != "" {
		config_path = env_path
	}

	err := configor.New(&configor.Config{Debug: true,
		Verbose:              true, //ENVPrefix:            "WEB",
		ErrorOnUnmatchedKeys: true}).Load(&GConfig, config_path)

	if err != nil {
		return fmt.Errorf("加载配置失败 %w", err)
	}

	return nil
}

func GetEnvironment() string {
	return os.Getenv("CONFIGOR_ENV")
}
