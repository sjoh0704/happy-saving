package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	gmux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/model"
	"github.com/sjoh0704/happysaving/util"
	df "github.com/sjoh0704/happysaving/util/datafactory"
)

func GetUsersInfo(res http.ResponseWriter, req *http.Request) {

	log.Info("get all users info")
	users := []model.User{}
	err := df.DbPool.Model(&users).Select()
	if err != nil {
		log.Error("getting all users fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", users, http.StatusAccepted)
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info("get info for user id: ", id)

	user := &model.User{ID: int64(id)}
	err = df.DbPool.Model(user).WherePK().Select()

	if err != nil {
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info(user)
	util.SetResponse(res, "success", user, http.StatusOK)
}

func UpdateUserInfo(res http.ResponseWriter, req *http.Request) {
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// user가 있는지 체크
	existUser := &model.User{
		ID: int64(id),
	}

	err = df.DbPool.Model(existUser).WherePK().Select()
	if err != nil {
		log.Error("getting user info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// user가 있으면 업데이트
	log.Info("update info for user id: ", id)

	newUserInfo := &model.User{
		ID: int64(id),
	}

	err = json.NewDecoder(req.Body).Decode(newUserInfo)
	if newUserInfo.Name != "" {
		existUser.SetName(newUserInfo.Name)
	}
	if newUserInfo.Mail != "" {
		existUser.SetMail(newUserInfo.Mail)
	}
	if newUserInfo.Password != "" {
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

	if err != nil {
		log.Error("updating user fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	util.SetResponse(res, "success", existUser, http.StatusOK)
}

// 사용자 회원가입
func CreateUser(res http.ResponseWriter, req *http.Request) {

	user := &model.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		util.SetResponse(res, "name doesn't exist", nil, http.StatusBadRequest)
		return
	}
	if user.Mail == "" {
		util.SetResponse(res, "mail doesn't exist", nil, http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		util.SetResponse(res, "password doesn't exist", nil, http.StatusBadRequest)
		return
	}

	if user.Gender != model.Female && user.Gender != model.Male {
		util.SetResponse(res, "gender info is not correct", nil, http.StatusBadRequest)
		return
	}

	// 동일 mail을 가진 user가 있는지 check
	count, err := df.DbPool.
		Model(&model.User{}).
		Where("mail = ?", user.Mail).
		Count()
	
	if err != nil {
		log.Error("creating user fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	if count >= 1{
		log.Error("user email already exists")
		util.SetResponse(res, "user already exists", nil, http.StatusBadRequest)
		return
	}


	log.Info("creating user")

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	hashedPasswd, err := util.HashPassword(user.Password)
	if err != nil {
		log.Error("creating user fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	user.SetPassword(hashedPasswd)

	_, err = df.DbPool.Model(user).Insert()

	if err != nil {
		log.Error(err)
	}
	util.SetResponse(res, "success", user, http.StatusCreated)
	log.Info("created user: ", user.String())
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	log.Info("Deleting user")
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("deleting user fails")
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	user := &model.User{ID: int64(id)}
	_, err = df.DbPool.Model(user).WherePK().Delete()

	if err != nil {
		log.Error("deleting user fails")
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", nil, http.StatusAccepted)
}
