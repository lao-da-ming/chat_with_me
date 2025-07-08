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

type EmployeePosition struct {
	ID         int64         `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	EmployeeID int64         `gorm:"column:employee_id;type:INT8;" json:"employee_id"`
	PositionID int64         `gorm:"column:position_id;type:INT8;" json:"position_id"`
	IsMain     bool          `gorm:"column:is_main;type:BOOL;default:false;" json:"is_main"`
	CreatedAt  time.Time     `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt  time.Time     `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	Status     int32         `gorm:"column:status;type:INT4;default:1;" json:"status"`
	StartDate  time.Time     `gorm:"column:start_date;type:DATE;default:CURRENT_DATE;" json:"start_date"`
	EndDate    time.Time     `gorm:"column:end_date;type:DATE;" json:"end_date"`
	CreatedBy  int64         `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy  sql.NullInt64 `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
}

func (e *EmployeePosition) TableName() string {
	return "employee_positions"
}

func (e *EmployeePosition) BeforeSave(tx *gorm.DB) error {
	return nil
}
