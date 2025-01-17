package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TotalPoints int
	VoucherID   uint
	Voucher     Voucher `gorm:"foreignKey:VoucherID" json:"-"`
	BrandID     uint
	Brand       Brand `gorm:"foreignKey:BrandID" json:"-"`
}
