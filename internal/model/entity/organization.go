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

type Organization struct {
	ID                      int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	CompanyID               int64          `gorm:"column:company_id;type:INT8;" json:"company_id"`
	Name                    string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Code                    string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Description             sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	OrganizationTypeID      int64          `gorm:"column:organization_type_id;type:INT8;" json:"organization_type_id"`
	ParentID                sql.NullInt64  `gorm:"column:parent_id;type:INT8;" json:"parent_id"`
	Path                    string         `gorm:"column:path;type:VARCHAR;" json:"path"`
	Metadata                string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	SortOrder               int32          `gorm:"column:sort_order;type:INT4;default:0;" json:"sort_order"`
	Version                 int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status                  int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt               time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt               time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	IsReference             bool           `gorm:"column:is_reference;type:BOOL;default:false;" json:"is_reference"`
	ReferenceType           sql.NullInt64  `gorm:"column:reference_type;type:INT4;" json:"reference_type"`
	RefOrgID                sql.NullInt64  `gorm:"column:ref_org_id;type:INT8;" json:"ref_org_id"`
	SyncConfig              string         `gorm:"column:sync_config;type:JSONB;default:{'sync_identity': false, 'sync_removals': true, 'sync_additions': true};" json:"sync_config"`
	ShortName               sql.NullString `gorm:"column:short_name;type:VARCHAR;size:100;" json:"short_name"`
	OrganizationStructureID sql.NullInt64  `gorm:"column:organization_structure_id;type:INT8;" json:"organization_structure_id"`
	Icon                    sql.NullString `gorm:"column:icon;type:VARCHAR;" json:"icon"`
	Exclusions              string         `gorm:"column:exclusions;type:JSONB;default:{};" json:"exclusions"`
	Associations            string         `gorm:"column:associations;type:JSONB;default:{};" json:"associations"`
	IsDisplay               bool           `gorm:"column:is_display;type:BOOL;" json:"is_display"`
	CreatedBy               int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy               sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy               sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt               time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (o *Organization) TableName() string {
	return "organizations"
}

func (o *Organization) BeforeSave(tx *gorm.DB) error {
	return nil
}
