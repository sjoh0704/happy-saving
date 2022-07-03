package util

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init_logging(){

	lum := &lumberjack.Logger{
        Filename:   "log/server.log",
        MaxSize:    500,
        MaxBackups: 3, 
        MaxAge:     28,
        Compress:   true,
  	}
//  textformatter나 jsonformatter를 사용할 수 있음 
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceQuote:    true,
	})

 // stdout 및 lumberjack으로 Output 설정 
	log.SetOutput(lum)
	log.SetOutput(os.Stdout)

 // 지정된 모듈에 대한 로깅 수준을 설정 -> DebugLevel 이상 부터 로깅.
 	log.SetLevel(log.DebugLevel)	
}