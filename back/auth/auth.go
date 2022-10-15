package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/model"
	"github.com/sjoh0704/happysaving/util"
	df "github.com/sjoh0704/happysaving/util/datafactory"
)

// 로그인(인증)
func Auth(res http.ResponseWriter, req *http.Request) {

	authUser := &model.User{}
	err := json.NewDecoder(req.Body).Decode(authUser)
	if err != nil {
		log.Error("auth error: ", err.Error())
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	if authUser.Mail == "" {
		util.SetResponse(res, "mail doesn't exist", nil, http.StatusBadRequest)
		return
	}
	if authUser.Password == "" {
		util.SetResponse(res, "password doesn't exist", nil, http.StatusBadRequest)
		return
	}

	user := &model.User{}

	// user mail이 있는지 check
	userCount, err := df.DbPool.
		Model(user).
		Where("mail = ?", authUser.Mail).
		SelectAndCount()

	if err != nil {
		log.Error("auth error: ", err.Error())
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	} else if userCount == 0 { // user가 없다면
		util.SetResponse(res, "email or password is not correct", nil, http.StatusBadRequest)
		return
	}

	if !util.CheckPasswordHash(user.Password, authUser.Password) { // login 실패
		log.Info("user login fails: ", user)
		util.SetResponse(res, "email or password is not correct", nil, http.StatusBadRequest)
		return

	}
	accessToken, err := CreateJWT(authUser.Mail)
	if err != nil {
		log.Info("user login fails: ", user)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	http.SetCookie(res, &http.Cookie{
		Name:     "access-token",
		Value:    accessToken,
		HttpOnly: true,
		// Expires:  time.Now().Add(time.Hour * 24),
		Expires:  time.Now().Add(time.Second * 3600),
	})

	log.Info("user login success: ", user)
	util.SetResponse(res, "login success", user, http.StatusOK)
}


// jwt는 header, payload, signature로 이루어져 있음
// payload에는 정보의 한 조각인 claim을 담을 수 있음 
func CreateJWT(Email string) (string, error) {
    mySigningKey := []byte(os.Getenv("SECRET_KEY"))

    aToken := jwt.New(jwt.SigningMethodHS256) 
    claims := aToken.Claims.(jwt.MapClaims)
    claims["Email"] = Email // private claim으로 중복되지 않는 값이 들어가도록 한다. 
    claims["exp"] = time.Now().Add(time.Minute * 100).Unix() // 20분 후에 만료 
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

func VerifiyJWTToken(token string)(bool, error){
	if token == ""{
		return false, jwt.ErrInvalidKey
	}

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return false, err
	}
	if email, ok := claims["Email"]; ok {
		log.Info("Authenticated user: ", email)	
	}else{
		return false, nil
	}
	
	return true, nil
}