package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	aeac "github.com/husamettinarabaci/SKey/internal/application/entry/adapter/cqhandler"
	aepc "github.com/husamettinarabaci/SKey/internal/application/entry/port/cqhandler"
	aepr "github.com/husamettinarabaci/SKey/internal/application/entry/port/repository"
	iear "github.com/husamettinarabaci/SKey/pkg/infrastructure/entry/adapter/repository"
	pesp "github.com/husamettinarabaci/SKey/pkg/presentation/entry/service/restapi"
)

func main() {
	var err error
	c := gin.Default()

	cont := container.New()

	err = cont.Singleton(func() aepr.Repository {
		return iear.NewRepository("/home/usrv/Documents/test.json")
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(repository aepr.Repository) aepc.CQHandler {
		return aeac.NewCQHandler(repository)
	})
	if err != nil {
		panic(err)
	}

	var repository aepr.Repository
	err = cont.Resolve(&repository)
	if err != nil {
		panic(err)
	}

	var cqHandler aepc.CQHandler
	err = cont.Resolve(&cqHandler)
	if err != nil {
		panic(err)
	}

	pesp.NewRestAPI(c, cqHandler)
	c.Run(":8080")

}

//TODO: Read data on Config
