package database

import (
	"Training/go-crud-with-oracle/infrastructure/persistance/repository"
	"Training/go-crud-with-oracle/technical_service/config"
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"time"
)

type Database struct {
	DB *sql.DB
	UserRepo *repository.UserRepository
}

func InitDB() (*Database, error) {
	url := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%v/%s"`, config.C.Database.DBUser, config.C.Database.DBPass, config.C.Database.DBHost, config.C.Database.DBPort, config.C.Database.DBName)

	db, err := sql.Open("godror", url)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	setDB := &Database{
		DB: db,
		UserRepo: repository.BuildUserRepository(db),
	}

	return setDB, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}

func (d *Database) Ping() error {
	return d.DB.Ping()
}