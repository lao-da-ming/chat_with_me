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

type AuditRoleAssignmentChange struct {
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	OperationID int64     `gorm:"column:operation_id;type:INT8;" json:"operation_id"`
	UserID      int64     `gorm:"column:user_id;type:INT8;" json:"user_id"`
	RoleID      int64     `gorm:"column:role_id;type:INT8;" json:"role_id"`
	ChangeType  string    `gorm:"column:change_type;type:VARCHAR;size:10;" json:"change_type"`
	CreatedAt   time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
}

func (a *AuditRoleAssignmentChange) TableName() string {
	return "audit_role_assignment_changes"
}

func (a *AuditRoleAssignmentChange) BeforeSave(tx *gorm.DB) error {
	return nil
}
