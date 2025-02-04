package main

import (
	"github.com/apache/yunikorn-core/pkg/log"
	"github.com/apache/yunikorn-release/soak/framework"
	"go.uber.org/zap"
)

const (
	ConfigFileName = "conf.yaml"
)

var logger *zap.Logger = log.Log(log.Test)

func main() {
	conf, err := framework.InitConfig(ConfigFileName)
	if err != nil {
		logger.Fatal("failed to parse config", zap.Error(err))
	}
	logger.Info("config successully loaded", zap.Any("conf", conf))
}
