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




