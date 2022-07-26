package auth

import (
	"encoding/json"
	"net/http"
	"strconv"
	gmux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/util"
	df "github.com/sjoh0704/happysaving/util/datafactory"
	"github.com/sjoh0704/happysaving/user"
)

func Auth(res http.ResponseWriter, req *http.Request) {
	
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// user가 있는지 체크 
	existUser := &user.User{
		ID: int64(id),
	}
	
	err = df.DbPool.Model(existUser).WherePK().Select()
	if err != nil{
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// user가 있으면 업데이트 
	log.Info("update info for user id: ", id)
	
	newUserInfo := &user.User{
		ID: int64(id),
	}

	err = json.NewDecoder(req.Body).Decode(newUserInfo)
	if newUserInfo.Name != ""{
		existUser.SetName(newUserInfo.Name)
	}
	if newUserInfo.Mail != ""{
		existUser.SetMail(newUserInfo.Mail)
	}
	if newUserInfo.Password != ""{
		hashedPasswd, err := util.HashPassword(newUserInfo.Password)
		if err != nil {
			log.Error("getting user info fails: ", err)
			util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
			return
		}
		existUser.SetPassword(hashedPasswd)
	}
	
	existUser.UpdateTime()

	_, err = df.DbPool.Model(existUser).WherePK().Update()

	if err != nil{
		log.Error("updating user fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	util.SetResponse(res, "success", existUser, http.StatusOK)
}