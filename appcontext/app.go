package appcontext

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type appContext struct {
	db *sqlx.DB
}

var context *appContext

func Init() {
	db := initDb()
	context = &appContext{
		db: db,
	}
}

func initDb() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:12345elsen@/guitar_collection?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("Database is nil")
	} else {
		if err = db.Ping(); err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(3)
	}

	return db
}

func GetDB() *sqlx.DB {
	return context.db
}
