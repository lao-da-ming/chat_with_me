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

type OrganizationManager struct {
	ID             int64         `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	OrganizationID int64         `gorm:"column:organization_id;type:INT8;" json:"organization_id"`
	ManageID       int64         `gorm:"column:manage_id;type:INT8;" json:"manage_id"`
	ManageType     int32         `gorm:"column:manage_type;type:INT4;" json:"manage_type"`
	AssignType     int32         `gorm:"column:assign_type;type:INT4;" json:"assign_type"`
	CreatedAt      time.Time     `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	ExpireAt       time.Time     `gorm:"column:expire_at;type:TIMESTAMPTZ;" json:"expire_at"`
	EndAt          time.Time     `gorm:"column:end_at;type:TIMESTAMPTZ;" json:"end_at"`
	StartAt        time.Time     `gorm:"column:start_at;type:TIMESTAMPTZ;" json:"start_at"`
	CreatedBy      int64         `gorm:"column:created_by;type:INT8;default:4;" json:"created_by"`
	UpdatedBy      sql.NullInt64 `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
}

func (o *OrganizationManager) TableName() string {
	return "organization_managers"
}

func (o *OrganizationManager) BeforeSave(tx *gorm.DB) error {
	return nil
}
