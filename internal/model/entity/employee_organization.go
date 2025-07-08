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

type EmployeeOrganization struct {
	ID             int64         `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	EmployeeID     int64         `gorm:"column:employee_id;type:INT8;" json:"employee_id"`
	OrganizationID int64         `gorm:"column:organization_id;type:INT8;" json:"organization_id"`
	IsMain         bool          `gorm:"column:is_main;type:BOOL;default:false;" json:"is_main"`
	SyncSourceID   sql.NullInt64 `gorm:"column:sync_source_id;type:INT8;" json:"sync_source_id"`
	CreatedAt      time.Time     `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	IsSynced       sql.NullBool  `gorm:"column:is_synced;type:BOOL;" json:"is_synced"`
	SourceType     sql.NullInt64 `gorm:"column:source_type;type:INT4;default:1;" json:"source_type"`
	ResourceType   sql.NullInt64 `gorm:"column:resource_type;type:INT4;" json:"resource_type"`
	Source         int32         `gorm:"column:source;type:INT4;default:0;" json:"source"`
	IsExclusion    bool          `gorm:"column:is_exclusion;type:BOOL;default:false;" json:"is_exclusion"`
	CreatedBy      int64         `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy      sql.NullInt64 `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
}

func (e *EmployeeOrganization) TableName() string {
	return "employee_organizations"
}

func (e *EmployeeOrganization) BeforeSave(tx *gorm.DB) error {
	return nil
}
