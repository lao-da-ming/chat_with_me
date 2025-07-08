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

type Service struct {
	ID           int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	AccessScope  int32          `gorm:"column:access_scope;type:INT4;" json:"access_scope"`
	Name         string         `gorm:"column:name;type:VARCHAR;size:100;" json:"name"`
	Code         string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Metadata     string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	LastReportAt time.Time      `gorm:"column:last_report_at;type:TIMESTAMPTZ;" json:"last_report_at"`
	ProductCode  sql.NullString `gorm:"column:product_code;type:VARCHAR;size:255;" json:"product_code"`
}

func (s *Service) TableName() string {
	return "services"
}

func (s *Service) BeforeSave(tx *gorm.DB) error {
	return nil
}
