package model

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	TransactionID uint
	Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"-"`
	VoucherID     uint
	Voucher       Voucher `gorm:"foreignKey:VoucherID" json:"-"`
	UserID        uint
}
