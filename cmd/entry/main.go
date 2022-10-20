package main

import (
	"github.com/gin-gonic/gin"
	aeac "github.com/husamettinarabaci/SKey/internal/application/entry/adapter/cqhandler"
	pesp "github.com/husamettinarabaci/SKey/pkg/presentation/entry/service/restapi"
)

func main() {

	//TODO: IoC Container

	c := gin.Default()

	handler := aeac.NewCQHandler()
	pesp.NewRestAPI(c, handler)

	c.Run(":8080")

}

//TODO: Read data on Config
