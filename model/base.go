package model

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `gorm:"column:create_time" json:"create_time"`
	UpdatedAt time.Time  `gorm:"column:update_time" json:"update_time"`
	DeletedAt *time.Time `gorm:"column:delete_time" sql:"index" json:"delete_time"`
}
