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

type Company struct {
	ID                  int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name                string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	EnName              string         `gorm:"column:en_name;type:VARCHAR;size:255;" json:"en_name"`
	ShortName           string         `gorm:"column:short_name;type:VARCHAR;size:50;" json:"short_name"`
	Description         sql.NullString `gorm:"column:description;type:TEXT;" json:"description"`
	LegalRepresentative string         `gorm:"column:legal_representative;type:VARCHAR;size:50;" json:"legal_representative"`
	OrganizationID      int64          `gorm:"column:organization_id;type:INT8;" json:"organization_id"`
	Code                string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Country             int32          `gorm:"column:country;type:INT4;" json:"country"`
	Province            int32          `gorm:"column:province;type:INT4;" json:"province"`
	City                int32          `gorm:"column:city;type:INT4;" json:"city"`
	District            int32          `gorm:"column:district;type:INT4;" json:"district"`
	Area                int32          `gorm:"column:area;type:INT4;" json:"area"`
	Address             string         `gorm:"column:address;type:VARCHAR;size:255;" json:"address"`
	Phone               string         `gorm:"column:phone;type:VARCHAR;size:20;" json:"phone"`
	Email               string         `gorm:"column:email;type:VARCHAR;size:100;" json:"email"`
	Website             string         `gorm:"column:website;type:VARCHAR;size:255;" json:"website"`
	Logo                string         `gorm:"column:logo;type:VARCHAR;size:255;" json:"logo"`
	IsBusinessCompany   bool           `gorm:"column:is_business_company;type:BOOL;default:false;" json:"is_business_company"`
	IsMainOffice        bool           `gorm:"column:is_main_office;type:BOOL;default:false;" json:"is_main_office"`
	CompanyType         int32          `gorm:"column:company_type;type:INT4;" json:"company_type"`
	BusinessLicense     string         `gorm:"column:business_license;type:VARCHAR;size:255;" json:"business_license"`
	TaxID               string         `gorm:"column:tax_id;type:VARCHAR;size:255;" json:"tax_id"`
	Metadata            string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Version             int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	Status              int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	IsMainland          bool           `gorm:"column:is_mainland;type:BOOL;default:false;" json:"is_mainland"`
	Contact             sql.NullString `gorm:"column:contact;type:VARCHAR;size:50;" json:"contact"`
	CreatedBy           int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy           sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy           sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
	DeletedAt           time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

func (c *Company) TableName() string {
	return "companies"
}

func (c *Company) BeforeSave(tx *gorm.DB) error {
	return nil
}
