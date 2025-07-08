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

type Instance struct {
	Name         string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Code         string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	EntryPoint   string         `gorm:"column:entry_point;type:VARCHAR;size:255;" json:"entry_point"`
	Icon         sql.NullString `gorm:"column:icon;type:VARCHAR;size:255;" json:"icon"`
	APIBaseURL   string         `gorm:"column:api_base_url;type:VARCHAR;size:255;" json:"api_base_url"`
	Config       string         `gorm:"column:config;type:JSONB;default:{};" json:"config"`
	Version      int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status       int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	Metadata     sql.NullString `gorm:"column:metadata;type:JSONB;" json:"metadata"`
	ProductCode  sql.NullString `gorm:"column:product_code;type:VARCHAR;size:255;" json:"product_code"`
	CreatedBy    int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy    sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy    sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt    time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
	ID           int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	InstanceType int32          `gorm:"column:instance_type;type:INT4;" json:"instance_type"`
	Release      sql.NullString `gorm:"column:release;type:VARCHAR;size:50;" json:"release"`
}

func (i *Instance) TableName() string {
	return "instances"
}

func (i *Instance) BeforeSave(tx *gorm.DB) error {
	return nil
}
