package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"
	"google.golang.org/grpc"
)

type ClientRepo struct {
	Host string
}

func NewClientRepo(host string) model.UserHandlerClient {
	return &ClientRepo{
		Host:host,
	}
}

func (_c *ClientRepo) AddUser(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*model.Response, error) {
	client, err := _c.getClient()
	if err != nil {
		return nil, err
	}
	//_ = glg.Log(">> AddUser Request: ", in.String())
	return client.AddUser(ctx, in)
}

func (_c *ClientRepo) GetUser(ctx context.Context, in *model.UserId, opts ...grpc.CallOption) (*model.User, error) {
	client, err := _c.getClient()
	if err != nil {
		return nil, err
	}
	//_ = glg.Log(">> GetUser Request: ", in.String())
	return client.GetUser(ctx, in)
}

func (_c *ClientRepo) ListUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*model.UserList, error) {
	client, err := _c.getClient()
	if err != nil {
		return nil, err
	}
	//_ = glg.Log(">> ListUser Request: ", in.String())
	return client.ListUser(ctx, in)
}

func (_c *ClientRepo) getClient() (model.UserHandlerClient, error) {
	conn, err := grpc.Dial(_c.Host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := model.NewUserHandlerClient(conn)
	return client, nil
}
