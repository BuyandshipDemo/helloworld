package entity

import "gorm.io/gorm"

type Helloworld struct {
	gorm.Model
	Key string `column:"key"`
	Val string `column:"val"`
}

func (u *Helloworld) TableName() string {
	return "helloworld"
}