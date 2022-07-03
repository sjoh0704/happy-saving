package main

import (
	"fmt"
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

	fmt.Println("hi")
	user.Test()
}

func register_multiplexer(){

	mux.HandleFunc("/", user.Hello)
}
