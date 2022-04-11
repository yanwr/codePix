package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(transactionId string) (*Transaction, error)
}

const (
	TRANSACTION_PENDING   string = "PENDING"
	TRANSACTION_COMPLETED string = "COMPLETED"
	TRANSACTION_CANCELED  string = "CANCELED"
	TRANSACTION_CONFIRMED string = "CONFIRMED"
)

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Id                string    `json:"id" gorm:"type:uuid;primary_key" valid:"required"`
	AccountFrom       *Account  `valid:"-"`
	AccountFromId     string    `json:"accountFromId" gorm:"column:accountFromId;type:uuid;not null" valid:"notnull"`
	Amount            float64   `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey   `valid:"-"`
	PixKeyToId        string    `json:"pixKeyToId" gorm:"column:pixKeyToId;type:uuid;not null" valid:"notnull"`
	Status            string    `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string    `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string    `json:"cancelDescription" gorm:"type:varchar(255)" valid:"-"`
	CreatedAt         time.Time `json:"createdAt" valid:"required"`
	UpdatedAt         time.Time `json:"updatedAt" valid:"required"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}
	err = transaction.isSameAccountValid()
	if err != nil {
		return err
	}
	err = transaction.isStatusValid()
	if err != nil {
		return err
	}
	err = transaction.isAmountValid()
	if err != nil {
		return err
	}
	return nil
}

func (transaction *Transaction) isAmountValid() error {
	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than 0 in Transaction")
	}
	return nil
}

func (transaction *Transaction) isStatusValid() error {
	if transaction.Status != TRANSACTION_PENDING && transaction.Status != TRANSACTION_COMPLETED && transaction.Status != TRANSACTION_CANCELED && transaction.Status != TRANSACTION_CONFIRMED {
		return errors.New("invalid status of Transaction")
	}
	return nil
}

func (transaction *Transaction) isSameAccountValid() error {
	if transaction.PixKeyTo.AccountId == transaction.AccountFrom.Id {
		return errors.New("the source and destination account cannot be the same in Transaction")
	}
	return nil
}

func (transaction *Transaction) Complete() error {
	changeStatusAndUpdate(transaction, TRANSACTION_COMPLETED)
	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Confirm() error {
	changeStatusAndUpdate(transaction, TRANSACTION_CONFIRMED)
	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Cancel(cancelDescription string) error {
	changeStatusAndUpdate(transaction, TRANSACTION_CANCELED)
	transaction.CancelDescription = cancelDescription
	err := transaction.isValid()
	return err
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		Id:                uuid.NewV4().String(),
		AccountFrom:       accountFrom,
		AccountFromId:     accountFrom.Id,
		Amount:            amount,
		PixKeyTo:          pixKeyTo,
		PixKeyToId:        pixKeyTo.Id,
		Status:            TRANSACTION_PENDING,
		Description:       description,
		CancelDescription: "",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func changeStatusAndUpdate(transaction *Transaction, newStatus string) {
	transaction.Status = newStatus
	transaction.UpdatedAt = time.Now()
}
