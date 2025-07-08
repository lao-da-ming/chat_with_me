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

type EmployeePositionHistory struct {
	ID            int64          `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	EmployeeID    int64          `gorm:"column:employee_id;type:INT8;" json:"employee_id"`
	PositionID    int64          `gorm:"column:position_id;type:INT8;" json:"position_id"`
	OperationType int32          `gorm:"column:operation_type;type:INT2;" json:"operation_type"`
	OperationTime time.Time      `gorm:"column:operation_time;type:TIMESTAMP;" json:"operation_time"`
	OperatorID    int64          `gorm:"column:operator_id;type:INT8;" json:"operator_id"`
	OperatorName  string         `gorm:"column:operator_name;type:VARCHAR;size:100;" json:"operator_name"`
	Reason        sql.NullString `gorm:"column:reason;type:TEXT;" json:"reason"`
	Details       string         `gorm:"column:details;type:JSONB;default:{};" json:"details"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
}

func (e *EmployeePositionHistory) TableName() string {
	return "employee_position_histories"
}

func (e *EmployeePositionHistory) BeforeSave(tx *gorm.DB) error {
	return nil
}
