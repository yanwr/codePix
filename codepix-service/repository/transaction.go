package repository

import (
	"codePix/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type TransactionRepository struct {
	Db *gorm.DB
}

func (r *TransactionRepository) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionRepository) Save(transaction *model.Transaction) error {
	// update
	err := r.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionRepository) Find(transactionId string) (*model.Transaction, error) {
	var transaction model.Transaction
	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", transactionId)
	if transaction.Id == "" {
		return nil, fmt.Errorf("there isn't Transanction with: id = %v", transactionId)
	}
	return &transaction, nil
}
