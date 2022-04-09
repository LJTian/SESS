package initialize

import "go.uber.org/zap"

func InitLoger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
