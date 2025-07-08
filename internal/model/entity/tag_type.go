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

type TagType struct {
	ID          int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name        string         `gorm:"column:name;type:VARCHAR;size:50;" json:"name"`
	Code        string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Metadata    string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Version     int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	BelongsTo   sql.NullInt64  `gorm:"column:belongs_to;type:INT4;" json:"belongs_to"`
	Description sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	CreatedBy   int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy   sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy   sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt   time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (t *TagType) TableName() string {
	return "tag_type"
}

func (t *TagType) BeforeSave(tx *gorm.DB) error {
	return nil
}
