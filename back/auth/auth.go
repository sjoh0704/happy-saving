package auth

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/user"
	"github.com/sjoh0704/happysaving/util"
	df "github.com/sjoh0704/happysaving/util/datafactory"
)

// 인증
func Auth(res http.ResponseWriter, req *http.Request) {

	authUser := &user.User{}
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

	userCheck := &user.User{}

	// user mail이 있는지 check
	userExist, err := df.DbPool.
		Model(userCheck).
		Where("mail = ?", authUser.Mail).
		SelectAndCount()

	if err != nil {
		log.Error("auth error: ", err.Error())
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	} else if userExist == 0 { // user가 없다면
		util.SetResponse(res, "email or password is not correct", nil, http.StatusBadRequest)
		return
	}

	if util.CheckPasswordHash(userCheck.Password, authUser.Password) { // login 성공 
		log.Info("user login success: ", userCheck)
		util.SetResponse(res, "login success", nil, http.StatusOK)
	
	} else { // login 실패 
		log.Info("user login fails: ", userCheck)
		util.SetResponse(res, "email or password is not correct", nil, http.StatusBadRequest)

	}

}
