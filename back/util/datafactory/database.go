package datafactory

import (
	"fmt"
    "github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

var (
	DbPool *pg.DB
)


func ConnectDB() {
	db_user := "postgres"
	db_password := "1q2w3e4r"
	hostname := "localhost"
	db_port := "5432"
	db_name := "postgres"

	DbPool = pg.Connect(&pg.Options{
		User: db_user,
		Password: db_password,
		Database: db_name,
		Addr: fmt.Sprintf("%s:%s", hostname, db_port),
	})

	if DbPool == nil{
		log.Error("unable to connect to DB")
		panic("panic")
	}else{
		log.Info("db connection success")
	}
}

func CloseDB(){
	err := DbPool.Close()
	if err != nil {
		log.Error("Error while closing the connection", err)
		panic(err)
	}
	log.Info("db connection closed")
}