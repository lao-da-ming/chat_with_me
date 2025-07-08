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

type AppInstance struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	AppID     int64     `gorm:"column:app_id;type:INT8;" json:"app_id"`
	ServiceID int64     `gorm:"column:service_id;type:INT8;" json:"service_id"`
	Metadata  string    `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Version   int32     `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status    int32     `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (a *AppInstance) TableName() string {
	return "app_instances"
}

func (a *AppInstance) BeforeSave(tx *gorm.DB) error {
	return nil
}
