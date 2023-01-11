package dtotransaction

import (
	"time"
)

type TransactionRequest struct {
	ID        int          `json:"id"`
	StartDate time.Time    `json:"startDate"`
	DueDate   time.Time    `json:"duetDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	Attache   string       `json:"attache"`
	Status    string       `json:"status" `
}

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}
