package factory

import (
	"codePix/application/useCase"
	"codePix/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) useCase.TransactionUseCase {
	pixRepository := repository.PixKeyRepository{Db: database}
	transactionRepository := repository.TransactionRepository{Db: database}

	return useCase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixKeyRepository:      &pixRepository,
	}
}
