package handler

import (
	// "encoding/json"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	// "time"
	gmux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/model"
	"github.com/sjoh0704/happysaving/util"
	df "github.com/sjoh0704/happysaving/util/datafactory"
)

func GetCouplesInfo(res http.ResponseWriter, req *http.Request) {

	log.Info("get all couples info")
	couples := []model.Couple{}
	err := df.DbPool.Model(&couples).Select()
	if err != nil {
		log.Error("getting all couples fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", couples, http.StatusAccepted)
}

func GetCoupleInfo(res http.ResponseWriter, req *http.Request) {
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Error("getting couple info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info("get info for couple id: ", id)

	couple := &model.Couple{ID: int64(id)}
	err = df.DbPool.Model(couple).WherePK().Select()

	if err != nil {
		log.Error("getting couple info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info(couple)
	util.SetResponse(res, "success", couple, http.StatusOK)

	// test 
	c := new(model.Couple)
	_ = df.DbPool.Model(c).
		Relation("Sender").
		Where("couples.id = ?", id).
		Select()

	util.SetResponse(res, "success", c, http.StatusOK)

	log.Info(c)
}

func RequestCouple(res http.ResponseWriter, req *http.Request) {
	cr := &model.Couple{}

	err := json.NewDecoder(req.Body).Decode(cr)
	if err != nil {
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	// phase가 공백 문자열일때만 들어올 수 있음
	if cr.Phase != "" {
		util.SetResponse(res, "phase cannot exist", nil, http.StatusBadRequest)
		return
	}

	// 동일 mail을 가진 user가 있는지 check
	count, err := df.DbPool.
		Model(&model.Couple{}).
		Where("sender_id = ? and receiver_id = ?", cr.SenderId, cr.ReceiverId).
		Count()

	log.Info(count)

	if err != nil {
		log.Error("request couple fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	if count >= 1 {
		log.Error("couple request already exists")
		util.SetResponse(res, "couple relation already exists", nil, http.StatusBadRequest)
		return
	}

	log.Info("creating couple relation")

	cr.CreatedAt = time.Now()
	cr.UpdatedAt = time.Now()
	cr.Phase = model.Awaiting

	_, err = df.DbPool.Model(cr).Insert()

	if err != nil {
		log.Error(err)
	}
	util.SetResponse(res, "success", cr, http.StatusCreated)
	log.Info("created couple relation: ", cr.String())
}

func ResponseForRequestCouple(res http.ResponseWriter, req *http.Request) {

	// user := &User{}
	// err := json.NewDecoder(req.Body).Decode(user)
	// if err != nil {
	// 	util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
	// 	return
	// }

	// if user.Name == "" {
	// 	util.SetResponse(res, "name doesn't exist", nil, http.StatusBadRequest)
	// 	return
	// }
	// if user.Mail == "" {
	// 	util.SetResponse(res, "mail doesn't exist", nil, http.StatusBadRequest)
	// 	return
	// }
	// if user.Password == "" {
	// 	util.SetResponse(res, "password doesn't exist", nil, http.StatusBadRequest)
	// 	return
	// }
	// // 동일 mail을 가진 user가 있는지 check
	// count, err := df.DbPool.
	// 	Model(&User{}).
	// 	Where("mail = ?", user.Mail).
	// 	Count()

	// if err != nil {
	// 	log.Error("creating user fails: ", err)
	// 	util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
	// 	return
	// }
	// if count >= 1{
	// 	log.Error("user email already exists")
	// 	util.SetResponse(res, "user already exists", nil, http.StatusBadRequest)
	// 	return
	// }

	// log.Info("creating user")

	// user.CreatedAt = time.Now()
	// user.UpdatedAt = time.Now()
	// hashedPasswd, err := util.HashPassword(user.Password)
	// if err != nil {
	// 	log.Error("creating user fails: ", err)
	// 	util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
	// 	return
	// }
	// user.SetPassword(hashedPasswd)

	// _, err = df.DbPool.Model(user).Insert()

	// if err != nil {
	// 	log.Error(err)
	// }
	// util.SetResponse(res, "success", user, http.StatusCreated)
	// log.Info("created user: ", user.String())
}
