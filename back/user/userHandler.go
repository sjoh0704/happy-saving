package user

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Get(http.ResponseWriter, *http.Request) {
	log.Info("hello")
		
}

func Post(http.ResponseWriter, *http.Request) {
	log.Info("hello")
		
}

func Delete(http.ResponseWriter, *http.Request) {
	log.Info("hello")
		
}