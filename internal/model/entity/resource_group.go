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

type ResourceGroup struct {
	ID          int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	AppID       sql.NullInt64  `gorm:"column:app_id;type:INT8;" json:"app_id"`
	ProductCode string         `gorm:"column:product_code;type:VARCHAR;size:255;" json:"product_code"`
	ServiceName string         `gorm:"column:service_name;type:VARCHAR;size:255;" json:"service_name"`
	Name        string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Code        string         `gorm:"column:code;type:VARCHAR;size:100;" json:"code"`
	Description sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Version     int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedBy   int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedBy   sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;" json:"updated_at"`
	DeletedBy   sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt   time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (r *ResourceGroup) TableName() string {
	return "resource_groups"
}

func (r *ResourceGroup) BeforeSave(tx *gorm.DB) error {
	return nil
}
