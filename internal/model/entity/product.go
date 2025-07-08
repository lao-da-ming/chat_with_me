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

type Product struct {
	Name           string         `gorm:"primary_key;column:name;type:VARCHAR;size:100;" json:"name"`
	Code           string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Description    sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Version        int32          `gorm:"column:version;type:INT4;" json:"version"`
	Icon           sql.NullString `gorm:"column:icon;type:VARCHAR;size:255;" json:"icon"`
	ConfigTemplate string         `gorm:"column:config_template;type:JSONB;default:{};" json:"config_template"`
	Metadata       string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Status         int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	ID             sql.NullInt64  `gorm:"column:id;type:INT8;" json:"id"`
	CreatedBy      int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy      sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy      sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt      time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
	Type           int32          `gorm:"column:type;type:INT4;" json:"type"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) BeforeSave(tx *gorm.DB) error {
	return nil
}
