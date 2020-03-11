package main

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/sinlovgo/log"
	"github.com/sinlovgo/log/lager"
)

func main() {

	logPath := "log.yaml"

	viper.SetConfigType("yaml")
	viper.SetConfigFile(logPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("InitWithFile err: %v", err)
		return
	}

	logFormat := viper.GetBool("log.log_format_text")
	//fmt.Printf("logFormat: %v\n", logFormat)
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  logFormat,
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	err := log.InitWithConfig(&passLagerCfg)
	if err != nil {
		fmt.Printf("init log error at path %v, err is %v\n", logPath, err)
	}

	//logPath := "log.json"
	//err := log.InitWithFile(logPath, "json")
	//if err != nil {
	//	fmt.Printf("init log error at path %v, err is %v\n", logPath, err)
	//	return
	//}

	for i := 0; i < 2; i++ {
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
