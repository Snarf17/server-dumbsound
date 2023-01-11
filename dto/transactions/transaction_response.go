package dtotransaction

import "time"

type TransactionResponse struct {
	StartDate time.Time    `json:"startDate"`
	DueDate   time.Time    `json:"duetDate"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	Attache   string       `json:"attache"`
	Status    string       `json:"status" `
}
