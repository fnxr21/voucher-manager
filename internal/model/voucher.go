package model

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model
	name         string
	BrandID      uint
	CostInPoints int
	UserID       uint
}
