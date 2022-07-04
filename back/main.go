package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	gmux "github.com/gorilla/mux"

	"github.com/sjoh0704/happysaving/user"
	"github.com/sjoh0704/happysaving/util"
)

var (
	mux *gmux.Router
)


func init(){
	util.Init_logging()
}

func main(){
	log.Info("main.go 실행")
	mux = gmux.NewRouter()

	register_multiplexer()

	http.ListenAndServe(":8000", mux)
}

func register_multiplexer(){
	mux.HandleFunc("/ready", ready)
	mux.HandleFunc("/apis/v1/users", serveUser)
}


func serveUser(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case http.MethodGet:
		user.Get(res, req)
	default:
		log.Error("method not acceptable")
	}
}

func ready(res http.ResponseWriter, req *http.Request){
	log.Info("OK")
}