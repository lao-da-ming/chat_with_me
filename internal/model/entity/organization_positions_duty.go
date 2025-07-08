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

type OrganizationPositionsDuty struct {
	ID                     int64         `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	OrganizationPositionID int64         `gorm:"column:organization_position_id;type:INT8;" json:"organization_position_id"`
	OperatorID             int64         `gorm:"column:operator_id;type:INT8;" json:"operator_id"`
	Version                int32         `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt              time.Time     `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt              time.Time     `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	DutyID                 sql.NullInt64 `gorm:"column:duty_id;type:INT8;" json:"duty_id"`
	CreatedBy              int64         `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy              sql.NullInt64 `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
}

func (o *OrganizationPositionsDuty) TableName() string {
	return "organization_positions_duties"
}

func (o *OrganizationPositionsDuty) BeforeSave(tx *gorm.DB) error {
	return nil
}
