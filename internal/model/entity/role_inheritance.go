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

type RoleInheritance struct {
	ID            int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT8;" json:"id"`
	RoleID        int64     `gorm:"column:role_id;type:INT8;" json:"role_id"`
	InheritRoleID int64     `gorm:"column:inherit_role_id;type:INT8;" json:"inherit_role_id"`
	Version       int32     `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt     time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

func (r *RoleInheritance) TableName() string {
	return "role_inheritances"
}

func (r *RoleInheritance) BeforeSave(tx *gorm.DB) error {
	return nil
}
