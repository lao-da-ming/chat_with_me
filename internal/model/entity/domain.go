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

type Domain struct {
	ID           int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name         string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Description  sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	BelongTo     int64          `gorm:"column:belong_to;type:INT8;" json:"belong_to"`
	BelongToType int32          `gorm:"column:belong_to_type;type:INT4;" json:"belong_to_type"`
	Metadata     string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Version      int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	CreatedBy    int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy    sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy    sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt    time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (d *Domain) TableName() string {
	return "domains"
}

func (d *Domain) BeforeSave(tx *gorm.DB) error {
	return nil
}
