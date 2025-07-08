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

type Session struct {
	ID           int64     `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	UserID       int64     `gorm:"column:user_id;type:INT8;" json:"user_id"`
	AppID        int64     `gorm:"column:app_id;type:INT8;" json:"app_id"`
	Token        string    `gorm:"column:token;type:TEXT;" json:"token"`
	RefreshToken string    `gorm:"column:refresh_token;type:TEXT;" json:"refresh_token"`
	Metadata     string    `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	ExpiresAt    time.Time `gorm:"column:expires_at;type:TIMESTAMPTZ;" json:"expires_at"`
	CreatedAt    time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (s *Session) TableName() string {
	return "sessions"
}

func (s *Session) BeforeSave(tx *gorm.DB) error {
	return nil
}
