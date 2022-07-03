package user

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/util"
)

func Test(){
	fmt.Println("user")
}

func Hello (http.ResponseWriter, *http.Request) {
	util.Init_logging()
	log.Info("hello")
		
}