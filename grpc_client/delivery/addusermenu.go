package delivery

import (
	"bufio"
	"fmt"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_client/usecase"
	"strings"
)

func ShowAddUserMenu(reader bufio.Reader){
	fmt.Println("MENU: Add User")
	fmt.Println("===========================================")
	fmt.Print("Entry User Id: ")
	uid, _ := reader.ReadString('\n')
	uid = strings.Trim(uid,"\n")
	fmt.Print("Entry name of user: ")
	name, _ := reader.ReadString('\n')
	name = strings.Trim(name, "\n")
	fmt.Print("Entry user password: ")
	pwd, _ := reader.ReadString('\n')
	pwd = strings.Trim(pwd, "\n")
	fmt.Print("Entry user gender[L/P]: ")
	gender, _ := reader.ReadString('\n')
	gender = strings.Trim(gender, "\n")

	ucase := usecase.NewRequestUsecase()
	fmt.Println(ucase.AddUser(uid,name,pwd,gender))

}
