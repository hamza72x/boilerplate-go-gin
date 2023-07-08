package accounts

import (
	"gorm.io/gorm"
)

type AccountType string

const (
	ADMIN AccountType = "admin"
	USER  AccountType = "user"
)

type Account struct {
	gorm.Model
	AccountType AccountType `gorm:"column:account_type;index" json:"account_type"`
	Name        string      `gorm:"column:name" json:"name"`
}

func (Account) TableName() string {
	return "accounts"
}
