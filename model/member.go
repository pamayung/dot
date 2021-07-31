package model

import "time"

type Member struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	AccountID uint32    `gorm:"size:11;not null;unique" json:"id"`
	Status    int       `gorm:"size:1;not null;" json:"first_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
