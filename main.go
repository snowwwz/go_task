package main

import (
	"github.com/yukinooz/go_task/service/infrastructure"
	"github.com/yukinooz/go_task/service/interfaces"
	"github.com/yukinooz/go_task/service/usecase"
)

func main() {
	//todo conf
	//todo logging

	repo := interfaces.NewTaskRepository(infrastructure.Connect())
	uc := usecase.NewUsecase(repo)
	cont := interfaces.NewController(uc)

	infrastructure.NewHandler(cont)
	infrastructure.Run()
}
