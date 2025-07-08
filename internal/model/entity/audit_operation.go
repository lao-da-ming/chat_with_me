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

type AuditOperation struct {
	ID            int64          `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	OperationType int32          `gorm:"column:operation_type;type:INT4;" json:"operation_type"`
	OperationTime time.Time      `gorm:"column:operation_time;type:TIMESTAMPTZ;" json:"operation_time"`
	OperatorID    int64          `gorm:"column:operator_id;type:INT8;" json:"operator_id"`
	OperatorName  string         `gorm:"column:operator_name;type:VARCHAR;size:100;" json:"operator_name"`
	TargetType    int32          `gorm:"column:target_type;type:INT2;" json:"target_type"`
	TargetID      int64          `gorm:"column:target_id;type:INT8;" json:"target_id"`
	TargetName    string         `gorm:"column:target_name;type:VARCHAR;size:100;" json:"target_name"`
	ResourceID    sql.NullInt64  `gorm:"column:resource_id;type:INT8;" json:"resource_id"`
	ResourceCode  sql.NullString `gorm:"column:resource_code;type:VARCHAR;size:100;" json:"resource_code"`
	ResourceName  sql.NullString `gorm:"column:resource_name;type:VARCHAR;size:100;" json:"resource_name"`
	Action        sql.NullString `gorm:"column:action;type:VARCHAR;size:50;" json:"action"`
	Reason        sql.NullString `gorm:"column:reason;type:TEXT;" json:"reason"`
	Details       sql.NullString `gorm:"column:details;type:JSONB;" json:"details"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
}

func (a *AuditOperation) TableName() string {
	return "audit_operations"
}

func (a *AuditOperation) BeforeSave(tx *gorm.DB) error {
	return nil
}
