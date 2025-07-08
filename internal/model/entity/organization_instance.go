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

type OrganizationInstance struct {
	ID             int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	OrganizationID int64          `gorm:"column:organization_id;type:INT8;" json:"organization_id"`
	CreatedBy      int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy      int64          `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	InstanceCode   sql.NullString `gorm:"column:instance_code;type:VARCHAR;size:255;" json:"instance_code"`
}

func (o *OrganizationInstance) TableName() string {
	return "organization_instances"
}

func (o *OrganizationInstance) BeforeSave(tx *gorm.DB) error {
	return nil
}
