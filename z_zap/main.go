package main

import (
	"go.uber.org/zap"
	"net/http"
)

/*
	Zap提供了两种类型的日志记录器—Sugared Logger和Logger。
		在性能很好但不是很关键的上下文中，使用SugaredLogger
		在每一微秒和每一次内存分配都很重要的上下文中，使用Logger
 */

var loggerProduction *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLoggerProduction() {
	loggerProduction, _ = zap.NewProduction()
}

func InitLoggerSugar() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGetProduction(url string) {
	resp, err := http.Get(url)
	if err != nil {
		loggerProduction.Error("Error fetching url",
			zap.String("url", url),
			zap.Error(err))
	} else {
		loggerProduction.Info("success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func simpleHttpGetSugar(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	InitLoggerProduction()
	InitLoggerSugar()
	defer loggerProduction.Sync()
	defer sugarLogger.Sync()
	simpleHttpGetProduction("https://www.baidu.com")
	simpleHttpGetSugar("https://www.baidu.com")
}
