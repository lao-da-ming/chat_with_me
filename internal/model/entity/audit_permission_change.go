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

type AuditPermissionChange struct {
	ID           int64          `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	OperationID  int64          `gorm:"column:operation_id;type:INT8;" json:"operation_id"`
	PermissionID sql.NullInt64  `gorm:"column:permission_id;type:INT8;" json:"permission_id"`
	ResourceID   int64          `gorm:"column:resource_id;type:INT8;" json:"resource_id"`
	ResourceCode sql.NullString `gorm:"column:resource_code;type:VARCHAR;size:100;" json:"resource_code"`
	ResourceName sql.NullString `gorm:"column:resource_name;type:VARCHAR;size:100;" json:"resource_name"`
	Action       sql.NullString `gorm:"column:action;type:VARCHAR;size:50;" json:"action"`
	AssigneeType int32          `gorm:"column:assignee_type;type:INT2;" json:"assignee_type"`
	AssigneeID   int64          `gorm:"column:assignee_id;type:INT8;" json:"assignee_id"`
	GrantType    int32          `gorm:"column:grant_type;type:INT2;" json:"grant_type"`
	Constraints  sql.NullString `gorm:"column:constraints;type:JSONB;" json:"constraints"`
	ChangeType   string         `gorm:"column:change_type;type:VARCHAR;size:100;" json:"change_type"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
}

func (a *AuditPermissionChange) TableName() string {
	return "audit_permission_changes"
}

func (a *AuditPermissionChange) BeforeSave(tx *gorm.DB) error {
	return nil
}
