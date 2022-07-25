package main

import (
	"fmt"
	"net/http"
	"github.com/go-pg/pg/orm"
	log "github.com/sirupsen/logrus"
	gmux "github.com/gorilla/mux"
	"github.com/sjoh0704/happysaving/user"
	"github.com/sjoh0704/happysaving/util"
	"github.com/sjoh0704/happysaving/util/datafactory"
)

var (
	mux *gmux.Router
)

var apiVersion string = "/apis/v1"

func init(){
	util.Init_logging()
	initDbConnection()

}

func main(){
	defer datafactory.CloseDB()
	port := 8000
	mux = gmux.NewRouter()

	register_multiplexer()
	log.Info("listening port: " + fmt.Sprint(port))
	http.ListenAndServe(":" + fmt.Sprint(port), mux)

}

func register_multiplexer(){
	mux.HandleFunc("/ready", ready)
	serveUser()
}


func serveUser(){
	mux.HandleFunc(apiVersion + "/users", user.CreateUser).Methods("POST")
	mux.HandleFunc(apiVersion + "/users", user.GetUsersInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", user.GetUserInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", user.UpdateUserInfo).Methods("POST")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", user.DeleteUser).Methods("DELETE")
}

func ready(res http.ResponseWriter, req *http.Request){
	log.Info("OK")
	util.SetResponse(res, "OK", nil, http.StatusAccepted)
}

func initDbConnection(){
	
	datafactory.ConnectDB()

	if err := CreateSchema(); err != nil{
		log.Error("cannot create schema", err)
		panic(err)
	}
}

func CreateSchema() error {
	models := []interface{}{
		(*user.User)(nil),
		// (*Story)(nil),
	}
	for _, model := range models {
		err := datafactory.DbPool.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}