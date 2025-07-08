package entity

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type User struct {
	ID   int32          `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	Name sql.NullString `gorm:"column:name;type:VARCHAR;size:255;" json:"name"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	return nil
}
