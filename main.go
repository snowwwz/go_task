package main

import (
	"github.com/yukinooz/go_task/service/infrastructure"
	"github.com/yukinooz/go_task/service/interfaces"
	"github.com/yukinooz/go_task/service/usecase"
)

func main() {
	//todo conf
	//todo logging

	db := infrastructure.Connect()

	repo := interfaces.NewTaskRepository(db)
	uc := usecase.NewUsecase(repo)
	cont := interfaces.NewController(uc)

	infrastructure.NewHandler(cont)
	infrastructure.Action()
}