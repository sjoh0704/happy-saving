package user

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/util"
	df "github.com/sjoh0704/happysaving/util/datafactory"
)


func GetUsersInfo(res http.ResponseWriter, req *http.Request) {

	users := []User{}
	err := df.DbPool.Model(&users).Select()
	if err != nil {
		log.Error("getting all users fails ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	for i:=0; i < len(users); i++{
		users[i].SetPassword("no show")
	}

	util.SetResponse(res, "", users, http.StatusAccepted)
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
	log.Info("hello")

	user := &User{
		Name: "test",
		Mail: "test@test.com",
		Password: "1234",
	}

	_, err := df.DbPool.Model(user).Insert()

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