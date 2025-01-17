package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username          string              `gorm:"type:varchar(100);column:username;not null;unique"`
	Password          string              `gorm:"type:varchar(100);column:password;not null"`
	Email             string              `gorm:"type:varchar(100);column:email;not null;unique"`
	TransactionDetail []TransactionDetail `gorm:"foreignKey:UserID" json:"-"`
	Voucher           []Voucher           `gorm:"foreignKey:UserID" json:"-"`
}

type Brand struct {
	gorm.Model
	name    string
	voucher []Voucher `gorm:"foreignKey:BrandID" json:"-"`
}

type Voucher struct {
	gorm.Model
	name         string
	BrandID      uint
	CostInPoints int
	UserID       uint
}

type Transaction struct {
	gorm.Model
	TotalPoints int
	VoucherID   uint
	Voucher     Voucher `gorm:"foreignKey:VoucherID" json:"-"`
	BrandID     uint
	Brand       Brand `gorm:"foreignKey:BrandID" json:"-"`
}

type TransactionDetail struct {
	gorm.Model
	TransactionID uint
	Transaction   Transaction `gorm:"foreignKey:TransactionID" json:"-"`
	VoucherID     uint
	Voucher       Voucher `gorm:"foreignKey:VoucherID" json:"-"`
	UserID        uint
}
