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

type PositionLevel struct {
	ID             int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name           string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Description    sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Level          int32          `gorm:"column:level;type:INT4;" json:"level"`
	PositionTypeID int64          `gorm:"column:position_type_id;type:INT8;" json:"position_type_id"`
	Status         int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	Version        sql.NullInt64  `gorm:"column:version;type:INT4;" json:"version"`
	CreatedBy      int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy      sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy      sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt      time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (p *PositionLevel) TableName() string {
	return "position_levels"
}

func (p *PositionLevel) BeforeSave(tx *gorm.DB) error {
	return nil
}
