package initialize

import "go.uber.org/zap"

// Replace global variable in zap with `logger`
func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
