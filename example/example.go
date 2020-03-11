package main

import (
	"fmt"

	"github.com/sinlovgo/log"
	"github.com/sinlovgo/log/lager"
)

func main() {
	logPath := "log.yaml"
	err := log.InitWithFile(logPath)
	if err != nil {
		fmt.Printf("init log error at path %v\n", logPath)
	}

	for i := 0; i < 3; i++ {
		log.Infof("Hi %s, system is starting up ...", "paas-bot")
		log.Info("check-info", lager.Data{
			"info": "something",
		})

		log.Debug("check-info", lager.Data{
			"id":   string(i),
			"info": "something",
		})

		log.Warn("failed-to-do-somthing", lager.Data{
			"id":   string(i),
			"info": "something",
		})

		err := fmt.Errorf("This is an error")
		log.Error("failed-to-do-somthing", err)

		log.Info("shutting-down")
	}
}
