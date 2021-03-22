package main

import (
	"../logger"
	"fmt"
	"time"
)

func initLogger(name, logPath, logName string, level string) (err error) {
	m := make(map[string]string, 16)
	m["log_path"] = logPath
	m["log_name"] = logName
	m["level"] = level
	m["log_split_type"] = "size"
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}

	logger.Debug("init logger success")
	return
}

func Run() {
	for {
		logger.Debug("1111")
		time.Sleep(time.Second)
	}
}

func main() {
	err := initLogger("file", "D:/logs/", "address", "debug")
	if err != nil {
		fmt.Println(err)
		return
	}
	Run()
	return
}
