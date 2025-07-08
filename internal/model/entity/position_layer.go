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

type PositionLayer struct {
	ID          int64          `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	Name        string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Description sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Code        string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	LayerLevel  int32          `gorm:"column:layer_level;type:INT4;" json:"layer_level"`
	Status      int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	Version     int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedBy   int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedBy   sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;" json:"updated_at"`
	DeletedBy   sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt   time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (p *PositionLayer) TableName() string {
	return "position_layers"
}

func (p *PositionLayer) BeforeSave(tx *gorm.DB) error {
	return nil
}
