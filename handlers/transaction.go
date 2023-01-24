package handlers

import (
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go/coreapi"
	"gopkg.in/gomail.v2"
)

// var payment_file = "http://localhost:8000/uploads/"
var total = 90000

// Declare Coreapi Client ...
var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"),
	ClientKey: os.Getenv("CLIENT_KEY"),
}

type handleTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handleTransaction {
	return &handleTransaction{TransactionRepository}
}

func (h *handleTransaction) ShowTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.ShowTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return

		// for i, p := range transaction {
		// 	// transaction[i].Attache = payment_file + p.Attache
		// }
	}
	// for i, p := range transaction {
	// 	// transaction[i].Attache = payment_file + p.Attache
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)

}

func (h *handleTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	date := time.Now()
	// d, _ := time.ParseDuration("720h")
	// days := d.Hours() / 24 // 2 days
	miliTime := date.Unix()
	dueDate, _ := time.Parse("2006-01-02", r.FormValue("dueDate"))
	user_id, _ := strconv.Atoi(r.FormValue("user_id"))
	transaction := models.Transaction{
		ID:        int(miliTime),
		StartDate: date,
		DueDate:   dueDate,
		UserID:    user_id,
		Status:    "Pending",
	}

	addTransaction, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	dataTransaction, err := h.TransactionRepository.GetTransaction(addTransaction.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// var s = snap.Client{}
	// s.New("SB-Mid-server-04rfe58Ldf0IYPUI6aLpZhP7", midtrans.Sandbox)
	// req := &snap.Request{
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  strconv.Itoa(dataTransaction.ID),
	// 		GrossAmt: int64(total),
	// 	},
	// 	CreditCard: &snap.CreditCardDetails{
	// 		Secure: true,
	// 	},
	// 	CustomerDetail: &midtrans.CustomerDetails{
	// 		FName: dataTransaction.User.Fullname,
	// 		Email: dataTransaction.User.Email,
	// 	},
	// }

	// snapResp, _ := s.CreateTransaction(req)
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: dataTransaction}
	json.NewEncoder(w).Encode(response)

}

func (h *handleTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userID := int(UserInfo["id"].(float64))

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	getData, err := h.TransactionRepository.GetTransaction(int(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.FormValue("status") != "" {
		getData.Status = r.FormValue("status")
	}

	data, err := h.TransactionRepository.UpdateTransaction(getData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

// func (h *handleTransaction) Notification(w http.ResponseWriter, r *http.Request) {
// 	var notificationPayload map[string]interface{}

// 	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	transactionStatus := notificationPayload["transaction_status"].(string)
// 	fraudStatus := notificationPayload["fraud_status"].(string)
// 	orderId := notificationPayload["order_id"].(string)
// 	id, _ := strconv.Atoi(orderId)

// 	transaction, _ := h.TransactionRepository.GetTransaction(id)
// 	if transactionStatus == "capture" {
// 		if fraudStatus == "challenge" {
// 			h.TransactionRepository.UpdateTransaction("pending", transaction.ID)
// 		} else if fraudStatus == "accept" {
// 			SendMail("success", transaction)
// 			h.TransactionRepository.UpdateTransaction("success", transaction.ID)
// 		}
// 	} else if transactionStatus == "settlement" {
// 		SendMail("success", transaction)
// 		h.TransactionRepository.UpdateTransaction("success", transaction.ID)
// 	} else if transactionStatus == "deny" {
// 		SendMail("failed", transaction)
// 		h.TransactionRepository.UpdateTransaction("failed", transaction.ID)
// 	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
// 		SendMail("failed", transaction)
// 		h.TransactionRepository.UpdateTransaction("failed", transaction.ID)
// 	} else if transactionStatus == "pending" {

// 		h.TransactionRepository.UpdateTransaction("pending", transaction.ID)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbSound <dumbsound@gmail.com>"
		var CONFIG_AUTH_EMAIL = "afriandifrans@gmail.com"
		var CONFIG_AUTH_PASSWORD = "fapmderihjckvouo"

		var startDate = transaction.StartDate
		var dueDate = transaction.DueDate
		var total_price = strconv.Itoa(total)
		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      <style>
        h1 {
        color: brown;
        }
      </style>
      </head>()
      <body>
      <h2>Subsribe payment :</h2>
      <ul style="list-style-type:none;">
				<li>Price : %s</li>
        <li>Start Date : %s</li>
        <li>Due Date: Rp.%s</li>
        <li>Status : <b>%s</b></li>
      </ul>
      </body>
    </html>`, total_price, startDate, &dueDate, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}
