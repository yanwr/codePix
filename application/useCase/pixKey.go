package useCase

import "codePix/domain/model"

// PixUseCase it's like Services in REST
type PixUseCase struct {
	// it's like interface implemented
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pixUseCase *PixUseCase) RegisterPixKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := pixUseCase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	err = pixUseCase.PixKeyRepository.Register(pixKey)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (pixUseCase *PixUseCase) FindPixKeyByKind(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pixUseCase.PixKeyRepository.FindByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
