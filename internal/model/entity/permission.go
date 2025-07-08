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

type Permission struct {
	ID           int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	ObjectID     int64          `gorm:"column:object_id;type:INT8;" json:"object_id"`
	Action       string         `gorm:"column:action;type:VARCHAR;size:255;" json:"action"`
	AssigneeType int32          `gorm:"column:assignee_type;type:INT4;" json:"assignee_type"`
	AssigneeID   int64          `gorm:"column:assignee_id;type:INT8;" json:"assignee_id"`
	GrantType    int32          `gorm:"column:grant_type;type:INT4;" json:"grant_type"`
	AssignedAt   time.Time      `gorm:"column:assigned_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"assigned_at"`
	Constraints  string         `gorm:"column:constraints;type:JSONB;default:{};" json:"constraints"`
	Description  sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Version      int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status       int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	CreatedBy    int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy    sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy    sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt    time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
	Type         int32          `gorm:"column:type;type:INT4;" json:"type"`
}

func (p *Permission) TableName() string {
	return "permissions"
}

func (p *Permission) BeforeSave(tx *gorm.DB) error {
	return nil
}
