package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kpango/glg"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_server/usecase"
)

func NewServerDelivery() model.UserHandlerServer {
	return &ServerHandler{
		service: usecase.NewService(),
	}
}

type ServerHandler struct {
	service usecase.ServiceUsecase
}

func (_s *ServerHandler) AddUser(c context.Context, u *model.User) (*model.Response, error) {
	_ = glg.Log(">> AddUser command: ", u.String())
	return _s.service.SaveUser(*u)
}

func (_s *ServerHandler) GetUser(c context.Context, u *model.UserId) (*model.User, error) {
	_ = glg.Log(">> GetUser command: ", u.String())
	return _s.service.GetUserById(u.Id)
}

func (_s *ServerHandler) ListUser(c context.Context, e *empty.Empty) (*model.UserList, error) {
	_ = glg.Log(">> ListUser command: ()")
	return _s.service.GetUsers()
}








