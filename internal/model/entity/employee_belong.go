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

type EmployeeBelong struct {
	ID           int64     `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	EmployeeID   int64     `gorm:"column:employee_id;type:INT8;" json:"employee_id"`
	BelongTo     int64     `gorm:"column:belong_to;type:INT8;" json:"belong_to"`
	BelongToType int32     `gorm:"column:belong_to_type;type:INT4;" json:"belong_to_type"`
	IsMain       bool      `gorm:"column:is_main;type:BOOL;default:false;" json:"is_main"`
	Version      int32     `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt    time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (e *EmployeeBelong) TableName() string {
	return "employee_belongs"
}

func (e *EmployeeBelong) BeforeSave(tx *gorm.DB) error {
	return nil
}
