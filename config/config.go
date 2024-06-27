package config

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

var ConfigInfo config

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("./config") // path从调用Init的地方开始算起
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			klog.Fatalf("config file not found: %v", err)
		} else {
			klog.Fatalf("config error: %v", err)
		}
	}
	if err := viper.Unmarshal(&ConfigInfo); err != nil {
		klog.Fatalf("config decode error: %v", err)
	}

}
