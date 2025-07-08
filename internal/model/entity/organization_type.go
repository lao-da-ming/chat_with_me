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

type OrganizationType struct {
	ID              int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name            string         `gorm:"column:name;type:VARCHAR;size:50;" json:"name"`
	Description     sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	Code            string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Level           int32          `gorm:"column:level;type:INT4;" json:"level"`
	AttributeSchema string         `gorm:"column:attribute_schema;type:JSONB;default:{};" json:"attribute_schema"`
	Rules           string         `gorm:"column:rules;type:JSONB;default:{};" json:"rules"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	Version         int32          `gorm:"column:version;type:INT4;" json:"version"`
	IsLawEntity     bool           `gorm:"column:is_law_entity;type:BOOL;default:false;" json:"is_law_entity"`
	Order           sql.NullInt64  `gorm:"column:order;type:INT4;" json:"order"`
	IsCompany       sql.NullBool   `gorm:"column:is_company;type:BOOL;" json:"is_company"`
	Icon            sql.NullString `gorm:"column:icon;type:VARCHAR;size:255;" json:"icon"`
	CreatedBy       int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy       sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy       sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt       time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (o *OrganizationType) TableName() string {
	return "organization_types"
}

func (o *OrganizationType) BeforeSave(tx *gorm.DB) error {
	return nil
}
