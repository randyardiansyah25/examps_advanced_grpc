package repository

import "github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"

type UserStoreRepo interface {
	SaveUser(user model.User) error
	GetUserById(id string) (*model.User, error)
	GetUsers() (*model.UserList, error)
	IsExistUser(id string) (bool)
}
