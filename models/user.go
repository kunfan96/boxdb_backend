package models

import "time"

// 用户表
type USER struct {
	ID         string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	Phone      string    `gorm:"type:varchar(255)" json:"phone"`
	Role       int       `gorm:"type:int" json:"role"`
	Status     int       `gorm:"type:int" json:"status"`
	CreateTime time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"`
}

func (USER) TableName() string {
	return "boxdb_users"
}
