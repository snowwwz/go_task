package main

import (
	"github.com/BurntSushi/toml"
	"github.com/yukinooz/go_task/config"
	"github.com/yukinooz/go_task/service/infrastructure"
	"github.com/yukinooz/go_task/service/interfaces"
	"github.com/yukinooz/go_task/service/usecase"
)

func main() {
	var conf config.Config
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		panic(err)
	}

	repo := interfaces.NewTaskRepository(infrastructure.Connect(conf.DB))
	uc := usecase.NewUsecase(repo)
	cont := interfaces.NewController(uc)

	infrastructure.NewHandler(cont)
	infrastructure.Run()
}
