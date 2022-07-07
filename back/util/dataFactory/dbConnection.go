package datafactory

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	connUrl string
	ctx context.Context
	Dbpool *pgxpool.Pool
	DBPassWordPath string
)

func CreateConnection(){
	db_driver := "postgres"
	db_user := "postgres"
	db_password := "1q2w3e4r"
	hostname := "localhost"
	db_port := "5432"
	db_name := "postgres"
	connUrl = fmt.Sprintf("%s://%s:%s@%s:%s/%s", db_driver, db_user, db_password, hostname, db_port, db_name)

	var err error
	Dbpool, err = pgxpool.Connect(context.Background(), connUrl)
	if err != nil {
		log.Error("unable to connect to DB", err)
		panic(err)
	}
	log.Info("db connection success")
 
}