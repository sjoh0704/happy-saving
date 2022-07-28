package util

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

type Message struct{
	Payload interface{} `json:"payload"`
	Msg string	`json:"message"`
}


func SetResponse(res http.ResponseWriter, outString string, outJson interface{}, status int) http.ResponseWriter {

	//set Cors
	// res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Max-Age", "3628800")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Type, X-Requested-With, Accept, Authorization, Referer, User-Agent")

	msg := Message{
		Payload: outJson,
		Msg: outString,
	}
	res.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(msg)
	if err != nil { // 500 error 반환 
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	}
	//set StatusCode
	res.WriteHeader(status)
	res.Write(js)
	return res
}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(hashVal, userPw string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
    if err != nil {
        return false
    } else {
        return true
    }
}

// jwt는 header, payload, signature로 이루어져 있음
// payload에는 정보의 한 조각인 claim을 담을 수 있음 
func CreateJWT(Email string) (string, error) {
    mySigningKey := []byte(os.Getenv("SECRET_KEY"))

    aToken := jwt.New(jwt.SigningMethodHS256) 
    claims := aToken.Claims.(jwt.MapClaims)
    claims["Email"] = Email // private claim으로 중복되지 않는 값이 들어가도록 한다. 
    claims["exp"] = time.Now().Add(time.Minute * 20).Unix() // 20분 후에 만료 
	claims["iss"] = "issuer"
	claims["sub"] = "sub title"
	claims["aud"] = "audience"

	// 서명은 다음과 같이 구성
	// HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)
    tk, err := aToken.SignedString(mySigningKey)
    if err != nil {
        return "", err
    }
    return tk, nil
}