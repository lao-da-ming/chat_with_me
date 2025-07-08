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

type PositionRole struct {
	ID         int64         `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	PositionID int64         `gorm:"column:position_id;type:INT8;" json:"position_id"`
	Version    int32         `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt  time.Time     `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt  time.Time     `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	OperatorID int64         `gorm:"column:operator_id;type:INT8;" json:"operator_id"`
	RoleID     sql.NullInt64 `gorm:"column:role_id;type:INT8;" json:"role_id"`
	CreatedBy  int64         `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy  sql.NullInt64 `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
}

func (p *PositionRole) TableName() string {
	return "position_roles"
}

func (p *PositionRole) BeforeSave(tx *gorm.DB) error {
	return nil
}
