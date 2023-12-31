package models

import (
	"fmt"

	"gorm.io/gorm"
)

type TypeCurrency string

const (
	usd TypeCurrency = "USD"
	brl TypeCurrency = "BRL"
	eur TypeCurrency = "EUR"
)

type Account struct {
	gorm.Model
	Id       string       `json:"id" gorm:"primaryKey, autoIncrement"`
	Owner    string       `json:"owner" gorm:"omitempty"`
	Balance  int64        `json:"balance" gorm:"omitempty"`
	Email    string       `json:"email" gorm:"omitempty"`
	Currency TypeCurrency `json:"currency" gorm:"omitempty"`
	Entries  []Entries    `json:"entries" gorm:"foreignKey:account_id;references:id"`
}

func PreparteData() {
	fmt.Println(usd, brl, eur)
}
