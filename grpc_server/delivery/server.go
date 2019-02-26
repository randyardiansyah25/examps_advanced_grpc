package delivery

import (
	"fmt"
	"github.com/kpango/glg"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_server/delivery/handler"
	"google.golang.org/grpc"
	"net"
	"os"
)

func Start(){
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("LISTENER_PORT")))
	if err != nil {
		_ = glg.Error(err.Error())
		os.Exit(1)
	}

	srv := grpc.NewServer()
	serverHandler := handler.NewServerDelivery()

	model.RegisterUserHandlerServer(srv, serverHandler)

	_ = glg.Info("Listening at : ", os.Getenv("LISTENER_PORT"))

	err = srv.Serve(l)
	if err != nil {
		_ = glg.Error(err.Error())
		os.Exit(1)
	}

}
