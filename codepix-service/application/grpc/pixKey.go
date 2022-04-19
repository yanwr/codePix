package grpc

import (
	"codePix/application/grpc/pb"
	"codePix/application/useCase"
	"codePix/domain/model"
	"context"
)

// PixKeyGrpcServiceController It's like Controller in REST
type PixKeyGrpcServiceController struct {
	PixUseCase useCase.PixUseCase
	pb.UnimplementedPixServiceControllerServer
}

// RegisterPixKey Here I'll "implements" interface created in file "grpc/pb/pixKey_grpc.pb.go"       all objects has the same arguments of the pixKey.proto (my contract gRPC)
// PixKeyRegistration is the body of Request in PixKey format, like a PixKeyDTO. And PixKeyResult would be a Response.
func (p *PixKeyGrpcServiceController) RegisterPixKey(ctx context.Context, pixKeyRegistration *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterPixKey(pixKeyRegistration.Key, pixKeyRegistration.Kind, pixKeyRegistration.AccountId)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: model.NOT_CREATED,
			Error:  err.Error(),
		}, err
	}
	return &pb.PixKeyCreatedResult{
		Id:     key.Id,
		Status: model.CREATED,
	}, nil
}

func (p *PixKeyGrpcServiceController) FindPixKey(ctx context.Context, inPixKey *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindPixKeyByKind(inPixKey.Key, inPixKey.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, err
	}
	return &pb.PixKeyInfo{
		Id:   pixKey.Id,
		Kind: pixKey.Kind,
		Key:  pixKey.Key,
		Account: &pb.Account{
			AccountId:     pixKey.AccountId,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankId,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

func NewPixGrpcServiceController(u useCase.PixUseCase) *PixKeyGrpcServiceController {
	return &PixKeyGrpcServiceController{PixUseCase: u}
}
