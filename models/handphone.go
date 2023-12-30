package models

import "gorm.io/gorm"

type Handphone struct {
	ID          uint   `gorm:"primary key;autoIncrement" json:"id"`
	NoHandphone string `json:"no_handphone"`
	Provider    string `json:"provider"`
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Handphone{})
	return err
}
