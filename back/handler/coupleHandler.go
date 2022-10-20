package handler

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
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

	log.Info("GetCouplesInfo")
	couples := []model.Couple{}
	err := df.DbPool.Model(&couples).Select()
	if err != nil {
		log.Error("getting all couples fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", couples, http.StatusOK)
}

func GetCoupleInfo(res http.ResponseWriter, req *http.Request) {
	log.Info("GetCoupleInfo")
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Error("getting couple info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info("get info for couple id: ", id)

	couple := &model.Couple{ID: int64(id)}
	err = df.DbPool.Model(couple).
		Relation("Receiver").
		Relation("Sender").
		WherePK().
		Select()

	if err != nil {
		log.Error("getting couple info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	log.Info(couple)
	util.SetResponse(res, "success", couple, http.StatusOK)
}

func RequestCouple(res http.ResponseWriter, req *http.Request) {
	log.Info("RequestCouple")
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
	if cr.SenderId == 0 || cr.ReceiverId == 0 || cr.SenderId == cr.ReceiverId {
		util.SetResponse(res, "sender id or receiver id has problem", nil, http.StatusBadRequest)
		return
	}

	// couple이 이미 있는지를 체크
	count, err := df.DbPool.
		Model(&model.Couple{}).
		Where("sender_id = ? and receiver_id = ?", cr.SenderId, cr.ReceiverId).
		Count()

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

	// 있는 User인지 체크
	user1 := &model.User{}
	err = df.DbPool.Model(user1).Where("id = ?", cr.SenderId).Select()
	if err != nil {
		log.Error("user doesn't exist")
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	user2 := &model.User{}
	err = df.DbPool.Model(user2).Where("id = ?", cr.ReceiverId).Select()
	if err != nil {
		log.Error("user doesn't exist")
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// user 성별 체크
	if user1.Gender == user2.Gender {
		log.Error("users have same genders. cannot create relation")
		util.SetResponse(res, "users have same genders. cannot create relation", nil, http.StatusBadRequest)
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

// c := &model.Couple{}
// err = df.DbPool.Model(c).
// 	Relation("Sender").
// 	Relation("Receiver").
// 	Where("couple.id = ?", id).
// 	Select()
// if err != nil {
// 	log.Error(err)
// }

// util.SetResponse(res, "success", c, http.StatusOK)

// log.Info(c)

func ResponseForRequestCouple(res http.ResponseWriter, req *http.Request) {

	log.Info("ResponseForRequestCouple")
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Error(err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	// 제대로된 요청인지 확인
	cr := &model.Couple{}
	err = json.NewDecoder(req.Body).Decode(cr)
	if err != nil {
		log.Error(err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	// 클라이언트는 리시버 id로 보내야 함
	// if cr.ReceiverId == 0 {
	// 	log.Error("doesn't send receiver id")
	// 	util.SetResponse(res, "doesn't send receiver id", nil, http.StatusBadRequest)
	// 	return
	// }
	// phase 체크
	if !(cr.Phase == model.Approved || cr.Phase == model.Denyed || cr.Phase == model.Awaiting) {
		log.Error("phase only can be approved, denyed, awaiting")
		util.SetResponse(res, "phase only can be approved, denyed, awaiting", nil, http.StatusBadRequest)
		return
	}

	// couple이 있는지 확인
	existCouple := &model.Couple{}
	err = df.DbPool.Model(existCouple).
		Where("id = ?", id).
		Select()

	if err != nil {
		log.Error(err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// 이미 처리된 Phase라면 패스
	if existCouple.Phase == model.Approved || existCouple.Phase == model.Denyed {
		log.Info("phase already is processed")
		util.SetResponse(res, "phase already is processed", nil, http.StatusBadRequest)
		return
	}

	// 보낸 요청이 정확한 올바른 요청인지 체크
	// if !(existCouple.SenderId == cr.SenderId && existCouple.ReceiverId == cr.ReceiverId) {
	// 	log.Error("receiver or sender id is not matched with request")
	// 	util.SetResponse(res, "receiver or sender id is not matched with request", nil, http.StatusBadRequest)
	// 	return
	// }

	// 승인 상태일때
	if cr.Phase == model.Approved {
		existCouple.Phase = model.Approved
		existCouple.UpdatedAt = time.Now()
		_, err = df.DbPool.Model(existCouple).WherePK().Update()
		if err != nil {
			log.Error(err)
			util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
			return
		}
		log.Info("phase is changed to Approved")
		util.SetResponse(res, "phase is changed to Approved", existCouple, http.StatusOK)
		return
	} else if cr.Phase == model.Denyed {
		// 거절일 때
		existCouple.Phase = model.Denyed
		existCouple.UpdatedAt = time.Now()
		_, err = df.DbPool.Model(existCouple).WherePK().Update()
		if err != nil {
			log.Error(err)
			util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
			return
		}
		log.Info("phase is changed to Denyed")
		util.SetResponse(res, "phase is changed to Denyed", existCouple, http.StatusOK)
		return
	}

	util.SetResponse(res, "cannot process phase", nil, http.StatusBadRequest)
	return
}

// sender id을 통해서 가져오기
func GetCoupleInfoByUserId(res http.ResponseWriter, req *http.Request) {
	log.Info("Getting couple info by sender id")
	userId, err := strconv.Atoi(req.FormValue("userid"))
	if err != nil {
		log.Error("cannot parse user id")
		util.SetResponse(res, "cannot parse user id", nil, http.StatusInternalServerError)
		return
	}
	if userId == 0 {
		log.Error("user id doesn't exist")
		util.SetResponse(res, "user id doesn't exist", nil, http.StatusBadRequest)
		return
	}

	log.Info("Getting couple info by sender id: ", userId)

	couple := &model.Couple{}

	count, err := df.DbPool.
		Model(couple).
		Relation("Receiver").
		Relation("Sender").
		Where("sender_id = ?", userId).
		Where("phase = ? ", "awaiting").
		Order("created_at DESC").
		Limit(1).
		SelectAndCount()

	// 커플이 존재하지 않는 경우
	if count == 0 {
		log.Info(fmt.Sprintf("couple by userid: %d doesn't exist", userId))
		util.SetResponse(res, fmt.Sprintf("couple by userid: %d doesn't exist", userId), nil, http.StatusNoContent)
		return
	}

	if err != nil {
		log.Error("getting couple info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", couple, http.StatusOK)
}

// user id을 통해서 모든 송신자 couple request 가져오기
func GetAllCouplesReqByUserId(res http.ResponseWriter, req *http.Request) {
	log.Info("Getting all couples request by user id")
	userId, err := strconv.Atoi(req.FormValue("userid"))
	if err != nil {
		log.Error("cannot parse user id")
		util.SetResponse(res, "cannot parse user id", nil, http.StatusInternalServerError)
		return
	}
	if userId == 0 {
		log.Error("user id doesn't exist")
		util.SetResponse(res, "user id doesn't exist", nil, http.StatusBadRequest)
		return
	}

	log.Info("Getting all couple req by receiver id: ", userId)

	var couples []model.Couple

	count, err := df.DbPool.
		Model(&couples).
		Relation("Sender").
		Where("receiver_id = ?", userId).
		Where("phase = ? ", "awaiting").
		Order("created_at DESC").
		SelectAndCount()

	// 커플이 존재하지 않는 경우
	if count == 0 {
		log.Info(fmt.Sprintf("couple by receiverid: %d doesn't exist", userId))
		util.SetResponse(res, fmt.Sprintf("couple by receiverid: %d doesn't exist", userId), nil, http.StatusNoContent)
		return
	}

	if err != nil {
		log.Error("getting all couples request fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", couples, http.StatusOK)
}

// sender id을 통해서 phase와 일치하는 요청 가져오기
func GetCoupleInfoBySenderId(res http.ResponseWriter, req *http.Request) {
	log.Info("Getting approved couple info by sender id")
	userId, err := strconv.Atoi(req.FormValue("userid"))
	if err != nil {
		log.Error("cannot parse user id")
		util.SetResponse(res, "cannot parse user id", nil, http.StatusInternalServerError)
		return
	}
	if userId == 0 {
		log.Error("user id doesn't exist")
		util.SetResponse(res, "user id doesn't exist", nil, http.StatusBadRequest)
		return
	}

	phase := req.FormValue("phase")

	if phase == "" {
		log.Error("user id doesn't exist")
		util.SetResponse(res, "user id doesn't exist", nil, http.StatusBadRequest)
		return
	}

	// if phase != "approved" || phase != "awaiting" || phase != "denyed" {
	// 	log.Error("phase doesn't exist")
	// 	util.SetResponse(res, "phase doesn't exist", nil, http.StatusBadRequest)
	// 	return
	// }

	log.Info("Getting approved couple info by sender id: ", userId)

	couple := &model.Couple{}

	count, err := df.DbPool.
		Model(couple).
		Relation("Receiver").
		Relation("Sender").
		Where("sender_id = ?", userId).
		Where("phase = ? ", phase).
		Order("created_at DESC").
		Limit(1).
		SelectAndCount()

	// 커플이 존재하지 않는 경우
	if count == 0 {
		log.Info(fmt.Sprintf("couple by userid: %d doesn't exist", userId))
		util.SetResponse(res, fmt.Sprintf("couple by userid: %d doesn't exist", userId), nil, http.StatusNoContent)
		return
	}

	if err != nil {
		log.Error("getting couple info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", couple, http.StatusOK)
}
