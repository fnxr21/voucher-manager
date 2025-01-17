package model

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	name    string
	voucher []Voucher `gorm:"foreignKey:BrandID" json:"-"`
}
