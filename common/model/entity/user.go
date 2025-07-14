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
	ID      int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name    sql.NullString `gorm:"column:name;type:VARCHAR;size:255;" json:"name"`
	Attr    sql.NullString `gorm:"column:attr;type:JSONB;" json:"attr"`
	Path    sql.NullString `gorm:"column:path;type:VARCHAR;" json:"path"`
	Version int32          `gorm:"column:version;type:INT4;default:0;" json:"version"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	return nil
}
