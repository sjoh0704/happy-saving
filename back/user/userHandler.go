package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	gmux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/util"
	df "github.com/sjoh0704/happysaving/util/datafactory"
)


func GetUsersInfo(res http.ResponseWriter, req *http.Request) {

	log.Info("get all users info")
	users := []User{}
	err := df.DbPool.Model(&users).Select()
	if err != nil {
		log.Error("getting all users fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	for i:=0; i < len(users); i++{
		users[i].SetPassword("no show")
	}

	util.SetResponse(res, "", users, http.StatusAccepted)
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info("get info for user id: ", id)




	user := &User{Id: int64(id)}
	err = df.DbPool.Model(user).WherePK().Select()

	if err != nil {
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info(user)
	user.SetPassword("no show")
	util.SetResponse(res, "", user, http.StatusOK)
}

func UpdateUserInfo(res http.ResponseWriter, req *http.Request) {
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// user가 있는지 체크 
	existUser := &User{
		Id: int64(id),
	}
	
	err = df.DbPool.Model(existUser).WherePK().Select()
	if err != nil{
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// user가 있으면 업데이트 
	log.Info("update info for user id: ", id)
	
	newUserInfo := &User{
		Id: int64(id),
	}

	err = json.NewDecoder(req.Body).Decode(newUserInfo)
	if newUserInfo.Name != ""{
		existUser.SetName(newUserInfo.Name)
	}
	if newUserInfo.Mail != ""{
		existUser.SetMail(newUserInfo.Mail)
	}
	if newUserInfo.Password != ""{
		existUser.SetPassword(newUserInfo.Password)
	}
	
	existUser.UpdateTime()

	_, err = df.DbPool.Model(existUser).WherePK().Update()

	if err != nil{
		log.Error("updating user fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	util.SetResponse(res, "", existUser, http.StatusOK)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	log.Info("creating user")
	user := &User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil{
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	if user.Name == ""{
		util.SetResponse(res, "name doesn't exist", nil, http.StatusBadRequest)
		return
	}
	if user.Mail == ""{
		util.SetResponse(res, "mail doesn't exist", nil, http.StatusBadRequest)
		return
	}
	if user.Password == ""{
		util.SetResponse(res, "password doesn't exist", nil, http.StatusBadRequest)
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = df.DbPool.Model(user).Insert()

	if err != nil {
		log.Error(err)
	}
	util.SetResponse(res, "", user, http.StatusCreated)
	log.Info("created user: ", user.String())
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	log.Info("hello")
	
}