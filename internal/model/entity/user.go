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
	ID           int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	AppID        int64          `gorm:"column:app_id;type:INT8;" json:"app_id"`
	Username     string         `gorm:"column:username;type:VARCHAR;size:50;" json:"username"`
	PasswordHash string         `gorm:"column:password_hash;type:VARCHAR;size:255;" json:"password_hash"`
	Name         string         `gorm:"column:name;type:VARCHAR;size:50;" json:"name"`
	EnName       string         `gorm:"column:en_name;type:VARCHAR;size:255;" json:"en_name"`
	Code         string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Email        sql.NullString `gorm:"column:email;type:VARCHAR;size:100;" json:"email"`
	Phone        sql.NullString `gorm:"column:phone;type:VARCHAR;size:20;" json:"phone"`
	UserType     int32          `gorm:"column:user_type;type:INT4;" json:"user_type"`
	Status       int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	Version      int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	DeletedAt    time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	LastLoginAt  time.Time      `gorm:"column:last_login_at;type:TIMESTAMPTZ;" json:"last_login_at"`
	Mfa          sql.NullInt64  `gorm:"column:mfa;type:INT4;default:1;" json:"mfa"`
	MfaType      sql.NullInt64  `gorm:"column:mfa_type;type:INT4;" json:"mfa_type"`
	Salt         sql.NullString `gorm:"column:salt;type:BYTEA;" json:"salt"`
	CreatedBy    int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy    sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy    sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	return nil
}
