package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	gmux "github.com/gorilla/mux"

	"github.com/sjoh0704/happysaving/user"
	"github.com/sjoh0704/happysaving/util"
	datafactory "github.com/sjoh0704/happysaving/util/dataFactory"
)

var (
	mux *gmux.Router
)


func init(){
	util.Init_logging()
	initConnection()
}

func main(){
	port := 8000
	mux = gmux.NewRouter()

	register_multiplexer()

	log.Info("listening port: " + fmt.Sprint(port))
	http.ListenAndServe(":" + fmt.Sprint(port), mux)

}

func register_multiplexer(){
	mux.HandleFunc("/ready", ready)
	mux.HandleFunc("/apis/v1/users", serveUser)
}


func serveUser(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case http.MethodGet:
		user.Get(res, req)
	case http.MethodPost:
		user.Post(res, req)
	case http.MethodPut:
		user.Put(res, req)
	case http.MethodDelete:
		user.Delete(res, req)
	default:
		log.Error("method not acceptable: ", req.Method)
	}
}

func ready(res http.ResponseWriter, req *http.Request){
	log.Info("OK")
	util.SetResponse(res, "OK", nil, 200)
}

func initConnection(){
	datafactory.CreateConnection()
}
