package usecase

import (
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_server/repository"
)

type Service struct{
	repo repository.UserStoreRepo
}

func (_s *Service) SaveUser(user model.User) (*model.Response, error) {
	var resp model.Response
	if _s.repo.IsExistUser(user.Id) {
		resp.ResponseCode = "1111"
		resp.ResponseMsg = "Duplicate user!"
	}else {
		err := _s.repo.SaveUser(user)
		if err != nil {
			resp.ResponseCode = "1111"
			resp.ResponseMsg = err.Error()
		}

		resp.ResponseCode = "0000"
		resp.ResponseMsg = "Success!"
	}
	return &resp, nil
}

func (_s *Service) GetUserById(id string) (*model.User, error) {
	return _s.repo.GetUserById(id)
}

func (_s *Service) GetUsers() (*model.UserList, error) {
	return _s.repo.GetUsers()
}

func NewService() ServiceUsecase {
	return &Service{
		repo: repository.NewUserStore(),
	}
}
