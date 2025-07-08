package data

import (
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

func NewData() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:123456@10.60.33.25:5432/postgres"), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	sdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sdb.SetConnMaxLifetime(30 * time.Minute)
	sdb.SetMaxIdleConns(50)
	sdb.SetMaxOpenConns(100)
	return db
}
