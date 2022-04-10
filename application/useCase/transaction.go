package useCase

import "codePix/domain/model"

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepository
	PixKeyRepository      model.PixKeyRepository
}

func (transactionUseCase *TransactionUseCase) RegisterTransaction(accountIdFrom string, amount float64, pixKeyIdTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	accountFrom, err := transactionUseCase.PixKeyRepository.FindAccount(accountIdFrom)
	if err != nil {
		return nil, err
	}

	pixKeyTo, err := transactionUseCase.PixKeyRepository.FindByKind(pixKeyIdTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(accountFrom, amount, pixKeyTo, description)
	if err != nil {
		return nil, err
	}

	err = transactionUseCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (transactionUseCase *TransactionUseCase) CompleteTransaction(transactionId string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	err = transaction.Complete()
	if err != nil {
		return nil, err
	}
	return SaveTransaction(transactionUseCase, transaction)
}

func (transactionUseCase *TransactionUseCase) ConfirmTransaction(transactionId string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	err = transaction.Confirm()
	if err != nil {
		return nil, err
	}
	return SaveTransaction(transactionUseCase, transaction)
}

func (transactionUseCase *TransactionUseCase) CancelTransaction(transactionId string, cancelDescription string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	err = transaction.Cancel(cancelDescription)
	if err != nil {
		return nil, err
	}
	return SaveTransaction(transactionUseCase, transaction)
}

func SaveTransaction(transactionUseCase *TransactionUseCase, transaction *model.Transaction) (*model.Transaction, error) {
	err := transactionUseCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
