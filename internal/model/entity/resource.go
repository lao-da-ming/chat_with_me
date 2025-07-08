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

type Resource struct {
	ID           int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name         string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Description  sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Code         string         `gorm:"column:code;type:VARCHAR;size:100;" json:"code"`
	ResourceType int32          `gorm:"column:resource_type;type:INT4;" json:"resource_type"`
	ResourcePath string         `gorm:"column:resource_path;type:VARCHAR;" json:"resource_path"`
	AccessScope  int32          `gorm:"column:access_scope;type:INT4;" json:"access_scope"`
	ServiceName  string         `gorm:"column:service_name;type:VARCHAR;size:255;" json:"service_name"`
	Metadata     string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	OwnerType    int32          `gorm:"column:owner_type;type:INT4;" json:"owner_type"`
	OwnerID      int64          `gorm:"column:owner_id;type:INT8;" json:"owner_id"`
	Version      int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status       int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	AppID        sql.NullInt64  `gorm:"column:app_id;type:INT8;default:0;" json:"app_id"`
	ProductCode  string         `gorm:"column:product_code;type:VARCHAR;size:255;default:0;" json:"product_code"`
	CreatedBy    int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy    sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy    sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt    time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (r *Resource) TableName() string {
	return "resources"
}

func (r *Resource) BeforeSave(tx *gorm.DB) error {
	return nil
}
