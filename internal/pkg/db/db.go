package db

import (
	"mb_api/internal/pkg/settings"
	"github.com/jackc/pgx"
	"log"
	"os"
	"sync"
)

var (
	// TODO: add it as environment variable
	host     = "mb_api_db"
	port     = uint16(5432)
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)


var db *pgx.ConnPool = nil
var syncOnce = sync.Once{}

func ConnectToDB() *pgx.ConnPool {
	syncOnce.Do(func() {
		pgxConfig := pgx.ConnConfig{
			Host:     host,
			Port:     port,
			Database: dbname,
			User:     user,
			Password: password,
		}
		pgxConnPoolConfig := pgx.ConnPoolConfig{
			MaxConnections: 1,
			ConnConfig: pgxConfig,
		}
		dbase, err := pgx.NewConnPool(pgxConnPoolConfig)
		if err != nil {
			if settings.UseCaseConf.InHDD {
				log.Fatal("Connection to database was failed")
			}
			db = nil
		} else {
			db = dbase
		}
	})
	return db
}