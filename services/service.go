package services

import (
	"context"
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/domain/models/engine"
)

type Service interface {
	Call(context context.Context, request *adapters.Request) engine.Response
}

type apiService struct {
}

func (a apiService) Call(context context.Context, request *adapters.Request) engine.Response {
	//TODO implement me
	panic("implement me")
}

func NewApiService() Service {
	return &apiService{}
}

type socketService struct {
	address string
	service string
}

func (s socketService) Call(context context.Context, request *adapters.Request) engine.Response {
	//TODO implement me
	panic("implement me")
}

func NewSocketService(address, service string) Service {
	return &socketService{
		address: address,
		service: service,
	}
}
