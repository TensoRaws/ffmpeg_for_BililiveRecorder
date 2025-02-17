package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Getconf() (string, string, int, string, int) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("workdir", "")
	viper.SetDefault("outputdir", "")
	viper.SetDefault("ffmpegtype", 10)
	viper.SetDefault("ffmpegcom", "")
	viper.SetDefault("rmtime", -1)
	viper.SafeWriteConfig()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("read config failed: %v", err)
	}
	if viper.Get("workdir") == "" || viper.Get("outputdir") == "" {
		logrus.Fatal("启动失败请完善配置文件")
	}
	return viper.GetString("workdir"), viper.GetString("outputdir"), viper.GetInt("ffmpegtype"), viper.GetString("ffmpegcom"), viper.GetInt("rmtime")

}
