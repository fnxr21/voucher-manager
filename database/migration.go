package database

import (
	"fmt"

	"github.com/fnxr21/voucher-manager/internal/model"
	"github.com/fnxr21/voucher-manager/pkg/mysql"
)

func RunMigration() {
	var err error
	// main migration
	err = mysql.DB.AutoMigrate(
		&model.User{},
		&model.Brand{},
		&model.Voucher{},
		&model.Transaction{},
		&model.TransactionDetail{},
	)

	if err != nil {
		fmt.Println(err)
		panic("DB Migration Failed ")
	}

	fmt.Println("All Migration Success")
}
