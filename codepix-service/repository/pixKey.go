package repository

import (
	"codePix/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type PixKeyRepository struct {
	Db *gorm.DB
}

func (r *PixKeyRepository) Register(pixKey *model.PixKey) error {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PixKeyRepository) FindByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey
	// Preloads find all joins
	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? AND key = ?", kind, key)
	if pixKey.Id == "" {
		return nil, fmt.Errorf("there isn't PixKey with: key = %v and kind = %v", key, kind)
	}
	return &pixKey, nil
}

func (r *PixKeyRepository) AddBank(bank *model.Bank) error {
	// how I've already validated data in Model here I just create in db
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PixKeyRepository) FindBank(bankId string) (*model.Bank, error) {
	var bank model.Bank
	r.Db.First(&bank, "id = ?", bankId)
	if bank.Id == "" {
		return nil, fmt.Errorf("there isn't Bank with: id = %v", bankId)
	}
	return &bank, nil
}

func (r *PixKeyRepository) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PixKeyRepository) FindAccount(accountId string) (*model.Account, error) {
	var account model.Account
	r.Db.Preload("Bank").First(&account, "id = ?", accountId)
	if account.Id == "" {
		return nil, fmt.Errorf("there isn't Account with: accountId = %v", accountId)
	}
	return &account, nil
}
