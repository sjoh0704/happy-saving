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

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	gmux "github.com/gorilla/mux"
// 	log "github.com/sirupsen/logrus"
// 	"github.com/sjoh0704/happysaving/util"
// 	df "github.com/sjoh0704/happysaving/util/datafactory"
// )

// 모든 post 정보 가져오기
func GetPosts(res http.ResponseWriter, req *http.Request) {

	log.Info("get all posts info")
	posts := []model.Post{}
	err := df.DbPool.Model(&posts).Select()
	if err != nil {
		log.Error("getting all posts fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", posts, http.StatusOK)
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Error("getting post info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info("get info for post id: ", id)

	post := &model.Post{ID: int64(id)}
	err = df.DbPool.Model(post).WherePK().Select()

	if err != nil {
		log.Error("getting post info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	log.Info(post)
	util.SetResponse(res, "success", post, http.StatusOK)
}

func UpdatePost(res http.ResponseWriter, req *http.Request) {
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("getting post fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// post가 있는지 체크
	existPost := &model.Post{
		ID: int64(id),
	}

	err = df.DbPool.Model(existPost).WherePK().Select()
	if err != nil {
		log.Error("getting post info fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// post가 있으면 업데이트
	log.Info("update info for post id: ", id)

	newPost := &model.Post{
		ID: int64(id),
	}

	err = json.NewDecoder(req.Body).Decode(newPost)
	// TODO title 변경은 추후
	// if newPost.Title != "" {
	// 	existPost.SetName(newPost.GetTitle())
	// }
	if newPost.Content != "" {
		existPost.SetContent(newPost.GetContent())
	}
	if newPost.ImageURL != "" {
		existPost.SetImageURL(newPost.GetImageURL())
	}

	existPost.UpdateTime()

	_, err = df.DbPool.Model(existPost).WherePK().Update()
	if err != nil {
		log.Error("updating post fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	util.SetResponse(res, "success", existPost, http.StatusOK)
}

// post 생성
func CreatePost(res http.ResponseWriter, req *http.Request) {

	post := &model.Post{}
	err := json.NewDecoder(req.Body).Decode(post)
	if err != nil {
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	if post.AuthorID == 0 {
		util.SetResponse(res, "author_id doesn't exist", nil, http.StatusBadRequest)
		return
	}
	if post.CoupleID == 0 {
		util.SetResponse(res, "couple_id doesn't exist", nil, http.StatusBadRequest)
		return
	}
	if post.Content == "" {
		util.SetResponse(res, "content doesn't exist", nil, http.StatusBadRequest)
		return
	}
	// 한 couple id 내에서 동일 title을 가진 post가 있는지 check
	count, err := df.DbPool.
		Model(&model.Post{}).
		Where("couple_id = ?", post.CoupleID).
		Where("title = ?", post.Title).
		Count()

	if err != nil {
		log.Error("creating post fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	if count >= 1 {
		log.Error("post already exists")
		util.SetResponse(res, "post already exists", nil, http.StatusBadRequest)
		return
	}

	log.Info("creating post")

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err = df.DbPool.Model(post).Insert()

	if err != nil {
		log.Error(err)
	}
	util.SetResponse(res, "success", post, http.StatusCreated)
	log.Info("created post: ", post.String())
}

func DeletePost(res http.ResponseWriter, req *http.Request) {
	log.Info("Deleting post")
	vars := gmux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("deleting post fails: ", err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}
	post := &model.Post{ID: int64(id)}
	_, err = df.DbPool.Model(post).WherePK().Delete()

	if err != nil {
		log.Error("deleting post fails")
		util.SetResponse(res, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	util.SetResponse(res, "success", nil, http.StatusOK)
}
