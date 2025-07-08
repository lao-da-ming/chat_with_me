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

type Employee struct {
	ID         int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	UserID     int64          `gorm:"column:user_id;type:INT8;" json:"user_id"`
	Code       string         `gorm:"column:code;type:VARCHAR;size:50;" json:"code"`
	Type       int32          `gorm:"column:type;type:INT4;" json:"type"`
	Status     int32          `gorm:"column:status;type:INT4;default:1;" json:"status"`
	Gender     int32          `gorm:"column:gender;type:INT4;" json:"gender"`
	Birthday   time.Time      `gorm:"column:birthday;type:DATE;" json:"birthday"`
	OfficeMail string         `gorm:"column:office_mail;type:VARCHAR;size:100;" json:"office_mail"`
	LeaderID   sql.NullInt64  `gorm:"column:leader_id;type:INT8;" json:"leader_id"`
	LeaderName sql.NullString `gorm:"column:leader_name;type:VARCHAR;size:50;" json:"leader_name"`
	MentorID   sql.NullInt64  `gorm:"column:mentor_id;type:INT8;" json:"mentor_id"`
	MentorName sql.NullString `gorm:"column:mentor_name;type:VARCHAR;size:50;" json:"mentor_name"`
	IDCard     string         `gorm:"column:id_card;type:VARCHAR;size:20;" json:"id_card"`
	Metadata   string         `gorm:"column:metadata;type:JSONB;default:{};" json:"metadata"`
	Version    int32          `gorm:"column:version;type:INT4;default:1;" json:"version"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	DeletedAt  time.Time      `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
	Name       sql.NullString `gorm:"column:name;type:VARCHAR;size:50;" json:"name"`
	EnName     sql.NullString `gorm:"column:en_name;type:VARCHAR;size:255;" json:"en_name"`
	Phone      sql.NullString `gorm:"column:phone;type:VARCHAR;size:20;" json:"phone"`
	CreatedBy  int64          `gorm:"column:created_by;type:INT8;" json:"created_by"`
	UpdatedBy  sql.NullInt64  `gorm:"column:updated_by;type:INT8;" json:"updated_by"`
	DeletedBy  sql.NullInt64  `gorm:"column:deleted_by;type:INT8;" json:"deleted_by"`
}

func (e *Employee) TableName() string {
	return "employee"
}

func (e *Employee) BeforeSave(tx *gorm.DB) error {
	return nil
}
