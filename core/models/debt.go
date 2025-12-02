package models

import (
	"gorm.io/gorm"
)

type Debt struct {
	gorm.Model
	CreditorID   uint    `json:"creditorid"`
	Creditor     User	 `gorm:"foreignKey:CreditorID"`
	DebtorID     uint    `json:"debtorid"`
	Debtor		 User	 `gorm:"foreignKey:DebtorID"`
	Amount       uint    `json:"amount"`
}

type ShowDebt struct {
	gorm.Model
	CreditorName string `json:"creditorname"`
	DebtorName   string `json:"debtorname"`
	Amount       uint   `json:"amount"`
}