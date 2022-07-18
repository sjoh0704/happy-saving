package util

import (
	"encoding/json"
	"net/http"
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


func SetResponse(res http.ResponseWriter, outString string, outJson interface{}, status int) http.ResponseWriter {

	//set Cors
	// res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Max-Age", "3628800")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Type, X-Requested-With, Accept, Authorization, Referer, User-Agent")

	//set Out
	if outJson != nil {
		res.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(outJson)
		if err != nil { // 500 error 반환 
			http.Error(res, err.Error(), http.StatusInternalServerError)
			log.Error(err)
		}
		//set StatusCode
		res.WriteHeader(status)
		res.Write(js)
		return res

	} else {
		//set StatusCode
		res.WriteHeader(status)
		res.Write([]byte(outString))
		return res
	}
}

