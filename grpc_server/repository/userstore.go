package repository

import (
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_server/model/store"
)

type UserStore struct{

}

func (_u *UserStore) IsExistUser(id string) (bool) {
	for _, user := range store.UserMemStorage.List {
		if user.Id == id {
			return true
		}
	}
	return false
}

func (_u *UserStore) GetUserById(id string) (*model.User, error) {

	for _, user := range store.UserMemStorage.List {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, model.ErrUserNotFound
}

func (_u *UserStore) GetUsers() (*model.UserList, error) {
	return store.UserMemStorage, nil
}

func (_u *UserStore) SaveUser(user model.User) error {
	store.UserMemStorage.List = append(store.UserMemStorage.List, &user)
	return nil
}

func NewUserStore() UserStoreRepo {
	return &UserStore{}
}


