package delivery

import (
	"bufio"
	"fmt"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_client/model"
	"os"
	"runtime"
	"strings"
)

func Home() {
	reader := bufio.NewReader(os.Stdin)
	for {
		clrFunc, ok := model.ClearScreen[runtime.GOOS]
		if ok {
			clrFunc()
		}

		fmt.Println("MENU")
		fmt.Println("1. Add User")
		fmt.Println("2. Browse User")
		fmt.Println("3. List User")
		fmt.Println("4. Exit")
		fmt.Print("Choise menu [1/2/3/4]: ")
		ops, _ := reader.ReadString('\n')
		ops = strings.Trim(ops,"\n")
		if ops == "1" {
			clrFunc()
			ShowAddUserMenu(*reader)
		}else if ops == "2" {
			clrFunc()
			ShowBrowseUserMenu(*reader)
		}else if ops == "3" {
			clrFunc()
			ShowListUser()
		}else if ops == "4" {
			fmt.Println("Good bye....")
			os.Exit(1)
		}else {
			fmt.Println("Invalid option menu!!")
			fmt.Println()
		}

		fmt.Println("press enter key to continue...")
		var garbage string
		_, _ = fmt.Scanln(&garbage)
		_ = garbage
	}
}



