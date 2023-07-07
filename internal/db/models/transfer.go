package models

import (
	"gorm.io/gorm"
)

type Transfers struct {
	Id            string    `json:"id" gorm:"primaryKey, autoIncrement"`
	ToAccountId   []Account `json:"to_account_id" gorm:"many2many:to_account_id, foreignKey:Id"`
	FromAccountId []Account `json:"from_account_id" gorm:"many2many:from_account_id, foreignKey:Id"`
	Amount        int64     `json:"amount"`
	gorm.Model
}
