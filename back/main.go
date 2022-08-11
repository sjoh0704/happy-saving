package main

import (
	"fmt"
	"net/http"

	"github.com/go-pg/pg/orm"
	gmux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/auth"
	"github.com/sjoh0704/happysaving/handler"
	"github.com/sjoh0704/happysaving/model"
	"github.com/sjoh0704/happysaving/util"

	"github.com/joho/godotenv"
	// md "github.com/sjoh0704/happysaving/middleware"
	"github.com/sjoh0704/happysaving/util/datafactory"
)

var (
	mux *gmux.Router
)

var apiVersion string = "/apis/v1"

func init(){
	util.Init_logging()
	initDbConnection()
	godotenv.Load(".env") // .env에서 필요한 변수 가져오기 
}

func main(){
	defer datafactory.CloseDB()
	port := 8000
	mux = gmux.NewRouter()

	register_multiplexer()
	// mux.Use(md.TokenAuthMiddleware)
	log.Info("listening port: " + fmt.Sprint(port))
	http.ListenAndServe(":" + fmt.Sprint(port), mux)

}

func register_multiplexer(){
	mux.HandleFunc("/ready", ready).Methods("GET")
	mux.HandleFunc("/auth", auth.Auth).Methods("POST")
	serveUser()
	serveCouple()
}


func serveUser(){
	mux.HandleFunc(apiVersion + "/users", handler.CreateUser).Methods("POST")
	mux.HandleFunc(apiVersion + "/users", handler.GetUsersInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", handler.GetUserInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", handler.UpdateUserInfo).Methods("POST")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")
}


func serveCouple(){
	mux.HandleFunc(apiVersion + "/couples", handler.GetCouplesInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/couples/{id:[0-9]+}", handler.GetCoupleInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/couples", handler.RequestCouple).Methods("POST")
	mux.HandleFunc(apiVersion + "/couples/{id:[0-9]+}", handler.ResponseForRequestCouple).Methods("PUT")
}


func servePost(){
	mux.HandleFunc(apiVersion + "/users", handler.CreateUser).Methods("POST")
	mux.HandleFunc(apiVersion + "/users", handler.GetUsersInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", handler.GetUserInfo).Methods("GET")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", handler.UpdateUserInfo).Methods("POST")
	mux.HandleFunc(apiVersion + "/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")
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
		(*model.User)(nil),
		(*model.Post)(nil),
		(*model.Couple)(nil),
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