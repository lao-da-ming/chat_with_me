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

type OrganizationStructure struct {
	ID          int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name        string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Description sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	DomainID    int64          `gorm:"column:domain_id;type:INT8;" json:"domain_id"`
	Order       int32          `gorm:"column:order;type:INT4;default:0;" json:"order"`
	Version     int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	CreatedBy   int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy   int64          `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	Code        sql.NullString `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	DeletedBy   sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt   time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (o *OrganizationStructure) TableName() string {
	return "organization_structures"
}

func (o *OrganizationStructure) BeforeSave(tx *gorm.DB) error {
	return nil
}
