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

type OrganizationPosition struct {
	ID         int64         `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	PositionID int64         `gorm:"column:position_id;type:INT8;" json:"position_id"`
	OrgID      int64         `gorm:"column:org_id;type:INT8;" json:"org_id"`
	Version    int32         `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt  time.Time     `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt  time.Time     `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	CreatedBy  int64         `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy  sql.NullInt64 `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
}

func (o *OrganizationPosition) TableName() string {
	return "organization_positions"
}

func (o *OrganizationPosition) BeforeSave(tx *gorm.DB) error {
	return nil
}
