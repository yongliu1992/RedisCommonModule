package RedisCommonModule 
import (
	"github.com/go-sql-driver/mysql"
	"time"
)
type ResourceAssessment2 struct {
	ID               int            `gorm:"column:id;primary_key" json:"id" db:"id"`
	Country          string         `gorm:"column:country" json:"country" db:"country"`
	AccessDepartment string         `gorm:"column:access_department" json:"access_department" db:"access_department"`
	AdviserUID       int            `gorm:"column:adviser_uid" json:"adviser_uid" db:"adviser_uid"`
	Name             string         `gorm:"column:name" json:"name" db:"name"`
	SourceID         int            `gorm:"column:source_id" json:"source_id" db:"source_id"`
	SendTime         mysql.NullTime `gorm:"column:send_time" json:"send_time" db:"send_time"`
	Channel          int            `gorm:"column:channel" json:"channel" db:"channel"`
	GrUserID         string         `gorm:"column:gr_user_id" json:"gr_user_id" db:"gr_user_id"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"-" db:"updated_at"`
	DeletedAt        mysql.NullTime `gorm:"column:deleted_at" json:"-" db:"deleted_at"`
	City             string         `gorm:"column:city" json:"city" db:"city"`
	Status           int            `gorm:"column:status" json:"status" db:"status"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"-" db:"created_at"`
	Phone            string         `gorm:"column:phone" json:"phone" db:"phone"`
	DataType         int            `gorm:"column:data_type" json:"data_type" db:"data_type"`
	Need             string         `gorm:"column:need" json:"need" db:"need"`
	RequestParam     string         `gorm:"column:request_param" json:"request_param" db:"request_param"`
	AccessCity       string         `gorm:"column:access_city" json:"access_city" db:"access_city"`
	AdviserName      string         `gorm:"column:adviser_name" json:"adviser_name" db:"adviser_name"`
	Category         int            `gorm:"column:category" json:"category" db:"category"`
	WxPushStatus     int            `gorm:"column:wx_push_status" json:"wx_push_status" db:"wx_push_status"`
	JobNumber        string         `gorm:"column:job_number"  db:"job_number" json:"job_number"`
	PhoneEncrypted   string         `gorm:"column:phone_encrypted" json:"phone_encrypted" db:"phone_encrypted"`
	PhoneMask        string         `gorm:"column:phone_mask" json:"phone_mask" db:"phone_mask"`
	ReferPageUrl     string         `json:"refer_page_url"` //来源网址
}
