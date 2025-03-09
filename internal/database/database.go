package database

import (
	"app/internal/dal"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

const Path = "database.db"

func initialize() error {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err = gorm.Open(sqlite.Open("file:"+Path+"?_fk=1"), &gorm.Config{
		Logger:               newLogger,
		FullSaveAssociations: false,
	})
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
