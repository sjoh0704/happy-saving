package user

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/util/datafactory"
)

func Get(http.ResponseWriter, *http.Request) {
	log.Info("hello")

	user := &User{
		Name: "test",
		Mail: "test@test.com",
		Password: "1234",
	}

	_, err := datafactory.DbPool.Model(user).Insert()

	if err != nil {
		log.Error("fail", err)
	}
}

func Put(http.ResponseWriter, *http.Request) {
	log.Info("hello")
		
}

func Post(http.ResponseWriter, *http.Request) {
	log.Info("hello")
	// db.Dbpool.Exec(context.TODO(), "INSERT INTO test (name, cluster, member_id, member_name, attribute, role, status, createdTime, updatedTime) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")

}

func Delete(http.ResponseWriter, *http.Request) {
	log.Info("hello")
		
}