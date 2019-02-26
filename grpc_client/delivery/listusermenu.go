package delivery

import (
	"fmt"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_client/usecase"
)

func ShowListUser(){
	fmt.Println("MENU: List User")
	fmt.Println("==================================================================")
	ucase := usecase.NewRequestUsecase()
	fmt.Println(ucase.GetUsers())
}
