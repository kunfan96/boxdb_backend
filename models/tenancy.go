package models

// 租户相关表
import (
	"time"
)

// 租户表
type Tenancy struct {
	ID           string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Name         string    `gorm:"type:varchar(255)" json:"name"`
	Status       int       `gorm:"type:int" json:"status"`
	CreateUserID string    `gorm:"type:varchar(255)" json:"create_user_id"`
	CreateTime   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"`
}

// 租户-用户关联表
type TenancyUser struct {
	ID         string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	UserID     string    `gorm:"type:varchar(255)" json:"user_id"`
	TenancyID  string    `gorm:"type:varchar(255)" json:"tenancy_id"`
	Status     int       `gorm:"type:int" json:"status"`
	CreateTime time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"`
}

func (Tenancy) TableName() string {
	return "boxdb_tenancy"
}

func (TenancyUser) TableName() string {
	return "boxdb_tenancy_user"
}
