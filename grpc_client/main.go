package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_client/delivery"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_client/model"
	"google.golang.org/grpc"
	"os"
	"os/exec"
)

func main() {
	delivery.Home()
}

func init(){
	grpc.EnableTracing = true

	//log := glg.FileWriter("log/app.log", 0666)
	//glg.Get().
	//	SetMode(glg.BOTH).
	//	AddLevelWriter(glg.LOG, log).
	//	AddLevelWriter(glg.ERR, log).
	//	AddLevelWriter(glg.WARN, log).
	//	AddLevelWriter(glg.DEBG, log).
	//	AddLevelWriter(glg.INFO, log)
	//
	//
	//_ = glg.Log("Loading configuration file....")
	//err := godotenv.Load(".env")
	//if err != nil {
	//	_ = glg.Error("Error loading configuration file!")
	//	os.Exit(1)
	//}

	model.ClearScreen = map[string]func(){}
	model.ClearScreen["darwin"] = func() {
		cmd:= exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	model.ClearScreen["linux"] = model.ClearScreen["darwin"]
	model.ClearScreen["windows"] = func() {
		cmd:= exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading configuration file!")
		os.Exit(1)
	}
}
