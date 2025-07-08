package data

import (
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var ProviderSet = wire.NewSet(NewData)

func NewData() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(""), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	db = db.Debug()
	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sdb.SetConnMaxLifetime(30 * time.Minute)
	sdb.SetMaxIdleConns(50)
	sdb.SetMaxOpenConns(100)
	return db, nil
}
