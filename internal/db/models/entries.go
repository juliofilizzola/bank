package models

import (
	"gorm.io/gorm"
)

type Entries struct {
	gorm.Model
	Id        string `json:"id" gorm:"primaryKey, autoIncrement"`
	AccountId string `json:"account_id"`
	Amount    int64  `json:"amount"`
}
