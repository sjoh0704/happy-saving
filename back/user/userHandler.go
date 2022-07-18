package user

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/util"
	"github.com/sjoh0704/happysaving/util/datafactory"
)

func GetUsersInfo(res http.ResponseWriter, req *http.Request) {
	log.Info("hello")

	user := &User{
		Name: "test",
		Mail: "test@test.com",
		Password: "1234",
	}

	_, err := datafactory.DbPool.Model(user).Insert()

	if err != nil {
		log.Error("fail", err)
	}
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
	log.Info("hello")

	user := &User{
		Name: "test",
		Mail: "test@test.com",
		Password: "1234",
	}

	_, err := datafactory.DbPool.Model(user).Insert()

	if err != nil {
		log.Error("fail", err)
	}
}

func UpdateUserInfo(res http.ResponseWriter, req *http.Request) {
	log.Info("hello")
		
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

	_, err = datafactory.DbPool.Model(user).Insert()

	if err != nil {
		log.Error(err)
	}
	util.SetResponse(res, "", user, http.StatusCreated)
	log.Info("created user: ", user.Name)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	log.Info("hello")
	
}