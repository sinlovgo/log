## 简介

The log package refers to  [lexkong/log](https://github.com/lexkong/log), and some convenience changes
The function is exactly the same, because the original author is no longer maintained and does not support `go mod`

The log packages commonly used when develop by Go:

+ log
+ glog
+ logrus

`log` and` glog` are relatively simple and cannot meet production-level program development.
`logrus` is powerful, but not support rotate. You need to rotate the log file by an external program yourself. 
This log package summarizes the requirements commonly used in enterprise development, and integrates these functions in a log package. 
After testing, the performance of the log package can fully meet the needs of enterprise-level production.

## 使用方法

Before using the log package, you need to initialize the log package. 
The initialization functions are:`InitWithConfig()`, `InitWithFile()`

A simple example:

```go
package main

import (
	"fmt"

	"github.com/sinlovgo/log"
	"github.com/sinlovgo/log/lager"
)

func main() {
	_ := log.InitWithFile("log.yaml")

	for i := 0; i < 1; i++ {
		log.Infof("Hi %s, system is starting up ...", "paas-bot")
		log.Info("check-info", lager.Data{
			"info": "something",
		})

		log.Debug("check-info", lager.Data{
			"info": "something",
		})

		log.Warn("failed-to-do-somthing", lager.Data{
			"info": "something",
		})

		err := fmt.Errorf("This is an error")
		log.Error("failed-to-do-somthing", err)

		log.Info("shutting-down")
	}
}
```

log.yaml file content：

```yaml
writers: file,stdout
logger_level: DEBUG
logger_file: logs/log.log
log_format_text: false
rollingPolicy: size # size, daily
log_rotate_date: 1
log_rotate_size: 1
log_backup_count: 7
```

## Log parameters

+ `writers`: file,stdout。`file` will let `logger_file` to file，`stdout` will show at std, most of time use bose
+ `logger_level`: log level: DEBUG, INFO, WARN, ERROR, FATAL
+ `logger_file`: log file setting
+ `log_format_text`: format `true` will format json, `false` will show abs
+ `rollingPolicy`: rotate policy, can choose as: daily, size. `daily` store as daily，`size` will save as max
+ `log_rotate_date`: rotate date, coordinate `rollingPolicy: daily`
+ `log_rotate_size`: rotate size，coordinate `rollingPolicy: size`
+ `log_backup_count`: backup max count, log system will compress the log file when log reaches rotate set, this set is max file count
