package models

import "time"

type Transaction struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time    `json:"startDate"`
	DueDate   time.Time    `json:"duetDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	Attache   string       `json:"attache" gorm:"type:varchar(255)"`
	Status    string       `json:"status" gorm:"type:varchar(255)"`
}

type TransactionResponse struct {
	ID        int          `json:"id"`
	StartDate time.Time    `json:"startDate"`
	DueDate   time.Time    `json:"duetDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	Attache   string       `json:"attache"`
	Status    string       `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
