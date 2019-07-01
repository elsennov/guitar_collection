package console

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
)

func StartMigrations() {
	db, err := sql.Open("mysql", "root:12345elsen@/guitar_collection?charset=utf8&parseTime=True&loc=Local&multiStatements=true")
	if err != nil {
		log.Println("Error sql.Open")
		panic(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Error mysql.WithInstance")
		panic(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"guitar_collection",
		driver,
	)
	if err != nil {
		log.Println("Error migrate.NewWithDatabaseInstance")
		panic(err)
	}

	if migration != nil {
		err := migration.Steps(1)
		if err != nil {
			panic(err)
		}
	}
}

func RollbackLastMigrations() {
	db, err := sql.Open("mysql", "root:12345elsen@/guitar_collection?charset=utf8&parseTime=True&loc=Local&multiStatements=true")
	if err != nil {
		log.Println("Error sql.Open")
		panic(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Error mysql.WithInstance")
		panic(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"guitar_collection",
		driver,
	)
	if err != nil {
		log.Println("Error migrate.NewWithDatabaseInstance")
		panic(err)
	}

	if migration != nil {
		err := migration.Steps(-1)
		if err != nil {
			panic(err)
		}
	}
}
