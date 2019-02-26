package usecase

import "github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"

type ServiceUsecase interface {
	SaveUser(user model.User) (*model.Response, error)
	GetUserById(id string) (*model.User, error)
	GetUsers() (*model.UserList, error)
}
