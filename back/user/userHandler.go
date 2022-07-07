package user

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
	db "github.com/sjoh0704/happysaving/util/dataFactory"
)

func Get(http.ResponseWriter, *http.Request) {
	log.Info("hello")

		
}

func Put(http.ResponseWriter, *http.Request) {
	log.Info("hello")
		
}

func Post(http.ResponseWriter, *http.Request) {
	log.Info("hello")
	db.Dbpool.Exec(context.TODO(), "INSERT INTO test (name, cluster, member_id, member_name, attribute, role, status, createdTime, updatedTime) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")

}

func Delete(http.ResponseWriter, *http.Request) {
	log.Info("hello")
		
}