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

type Menu struct {
	ID          int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name        string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Description sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Code        string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Path        string         `gorm:"column:path;type:VARCHAR;" json:"path"`
	Metadata    string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Version     int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status      int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	AppID       int64          `gorm:"column:app_id;type:INT8;" json:"app_id"`
	Icon        string         `gorm:"column:icon;type:VARCHAR;size:255;" json:"icon"`
	ProductCode sql.NullString `gorm:"column:product_code;type:VARCHAR;size:50;" json:"product_code"`
	Route       sql.NullString `gorm:"column:route;type:VARCHAR;size:255;" json:"route"`
	EntryPoint  sql.NullString `gorm:"column:entry_point;type:VARCHAR;size:255;" json:"entry_point"`
	CreatedBy   int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy   sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy   sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt   time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
	ParentID    sql.NullInt64  `gorm:"column:parent_id;type:INT8;" json:"parent_id"`
	Type        int32          `gorm:"column:type;type:INT4;" json:"type"`
}

func (m *Menu) TableName() string {
	return "menus"
}

func (m *Menu) BeforeSave(tx *gorm.DB) error {
	return nil
}
