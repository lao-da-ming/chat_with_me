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

type OrganizationPositionsRole struct {
	ID                     int64     `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	OrganizationPositionID int64     `gorm:"column:organization_position_id;type:INT8;" json:"organization_position_id"`
	OperatorID             int64     `gorm:"column:operator_id;type:INT8;" json:"operator_id"`
	Version                int32     `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt              time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (o *OrganizationPositionsRole) TableName() string {
	return "organization_positions_roles"
}

func (o *OrganizationPositionsRole) BeforeSave(tx *gorm.DB) error {
	return nil
}
