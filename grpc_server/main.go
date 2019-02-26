package main

import (
	"github.com/joho/godotenv"
	"github.com/kpango/glg"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_common/model"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_server/delivery"
	"github.com/randyardiansyah25/examps_advanced_grpc/grpc_server/model/store"
	"os"
)

func main() {
	delivery.Start()
}

func init(){

	log := glg.FileWriter("log/app.log", 0666)
	glg.Get().
		SetMode(glg.BOTH).
		AddLevelWriter(glg.LOG, log).
		AddLevelWriter(glg.ERR, log).
		AddLevelWriter(glg.WARN, log).
		AddLevelWriter(glg.DEBG, log).
		AddLevelWriter(glg.INFO, log)


	_ = glg.Log("Loading configuration file....")
	err := godotenv.Load(".env")
	if err != nil {
		_ = glg.Error("Error loading configuration file!")
		os.Exit(1)
	}

	store.UserMemStorage = &model.UserList{
		List: make([]*model.User, 0),
	}
}
