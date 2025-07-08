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

type UserAttribute struct {
	ID        int64         `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	UserID    int64         `gorm:"column:user_id;type:INT8;" json:"user_id"`
	AttrKey   string        `gorm:"column:attr_key;type:VARCHAR;size:50;" json:"attr_key"`
	AttrValue string        `gorm:"column:attr_value;type:JSONB;" json:"attr_value"`
	AttrType  int32         `gorm:"column:attr_type;type:INT4;" json:"attr_type"`
	Version   int32         `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt time.Time     `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	CreatedBy int64         `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy sql.NullInt64 `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy sql.NullInt64 `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt time.Time     `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (u *UserAttribute) TableName() string {
	return "user_attributes"
}

func (u *UserAttribute) BeforeSave(tx *gorm.DB) error {
	return nil
}
