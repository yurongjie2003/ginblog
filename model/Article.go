package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid         int      `gorm:"type:int;not null;" json:"cid"`
	Category    Category `gorm:"foreignKey:Cid"`
	Description string   `gorm:"type:varchar(200)" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Img         string   `gorm:"type:varchar(100)" json:"img"`
}
