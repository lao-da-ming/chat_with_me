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

type AuditOperationChange struct {
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	OperationID int64     `gorm:"column:operation_id;type:INT8;" json:"operation_id"`
	BeforeState string    `gorm:"column:before_state;type:JSONB;" json:"before_state"`
	AfterState  string    `gorm:"column:after_state;type:JSONB;" json:"after_state"`
	CreatedAt   time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
}

func (a *AuditOperationChange) TableName() string {
	return "audit_operation_changes"
}

func (a *AuditOperationChange) BeforeSave(tx *gorm.DB) error {
	return nil
}
