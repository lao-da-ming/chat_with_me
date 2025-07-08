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

type Application struct {
	ID          int64          `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	Name        string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Description sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Code        string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	AppType     int32          `gorm:"column:app_type;type:INT4;" json:"app_type"`
	AppKey      string         `gorm:"column:app_key;type:VARCHAR;size:100;" json:"app_key"`
	AppSecret   string         `gorm:"column:app_secret;type:VARCHAR;size:255;" json:"app_secret"`
	Metadata    string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Version     int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status      int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (a *Application) TableName() string {
	return "applications"
}

func (a *Application) BeforeSave(tx *gorm.DB) error {
	return nil
}
