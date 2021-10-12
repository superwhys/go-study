package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func basic() {
	log.WithFields(log.Fields{"animal": "dog"}).Info("a dog has appear")
}

func basic2() {
	// 创建一个新的logger实例。可以创建任意多个。
	var logger = log.New()

	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	// 默认的输出是os.Stderr
	logger.Out = os.Stdout

	log1 := logger.WithFields(log.Fields{
		"animal": "dog",
		"size":   10,
	})
	log1.Info("this is basic2 info")
	log1.Warn("this is basic2 warn")
}

func basic3() {
	var logger = log.New()
	logger.Out = os.Stdout
	logger.Info("this is basic3")
}

func logFormat() {
	var logger = log.New()
	//logger.Formatter = &log.JSONFormatter{
	//	TimestampFormat:   time.RFC1123,
	//	DisableTimestamp:  false,
	//	DisableHTMLEscape: false,
	//	DataKey:           "",
	//	FieldMap:          nil,
	//	CallerPrettyfier:  nil,
	//	PrettyPrint:       false,
	//}
	logger.Formatter = &log.TextFormatter{}

	logger.Info("this is logFormat")
}
func main() {
	basic()
	basic2()
	basic3()
	logFormat()
}
