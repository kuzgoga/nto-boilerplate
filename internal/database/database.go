package database

import (
	"app/internal/dal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

const Path = "database.db"

func initialize() error {
	var err error
	db, err = gorm.Open(sqlite.Open("file:"+Path+"?_fk=1"), &gorm.Config{})
	if err != nil {
		return err
	}
	if res := db.Exec(`PRAGMA foreign_keys = ON`); res.Error != nil {
		return res.Error
	}

	if err := limitConnectionPool(); err != nil {
		return err
	}

	dal.SetDefault(db)

	log.Println("Initialized database")
	return nil
}

func limitConnectionPool() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func Shutdown() error {
	db, err := db.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	once = sync.Once{}
	return nil
}

func GetInstance() *gorm.DB {
	once.Do(func() {
		err := initialize()
		if err != nil {
			panic(err)
		}
	})
	return db
}
