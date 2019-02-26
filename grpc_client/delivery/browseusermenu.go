package delivery

import (
	"bufio"
	"fmt"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_client/usecase"
	"strings"
)

func ShowBrowseUserMenu(reader bufio.Reader){
	fmt.Println("MENU: Browse User")
	fmt.Println("===========================================")
	fmt.Print("Entry user id: ")
	uid,_ := reader.ReadString('\n')
	uid = strings.Trim(uid, "\n")
	fmt.Println()
	ucase := usecase.NewRequestUsecase()
	fmt.Println(ucase.GetUserById(uid))
}
