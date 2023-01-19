package dtotransaction

import "time"

type TransactionResponse struct {
	StartDate time.Time    `json:"startDate"`
	DueDate   time.Time    `json:"dueDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	Status    string       `json:"status" `
}
