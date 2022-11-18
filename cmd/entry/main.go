package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	aeac "github.com/husamettinarabaci/SKey/internal/application/entry/adapter/cqhandler"
	aepc "github.com/husamettinarabaci/SKey/internal/application/entry/port/cqhandler"
	aes "github.com/husamettinarabaci/SKey/internal/application/entry/service"
	des "github.com/husamettinarabaci/SKey/internal/domain/entry/service"
	pesp "github.com/husamettinarabaci/SKey/pkg/presentation/entry/service/restapi"
	"gopkg.in/yaml.v2"
)

const configPath = "configs/entry.yml"

type Config struct {
	Debug bool `yaml:"debug"`
	Cli   struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"cli"`
	Restapi struct {
		Enabled        bool   `yaml:"enabled"`
		Host           string `yaml:"host"`
		Port           string `yaml:"port"`
		QueryHandler   bool   `yaml:"queryHandler"`
		CommandHandler bool   `yaml:"commandHandler"`
	} `yaml:"restapi"`
}

var AppConfig *Config

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ReadConfig()
	var err error
	c := gin.Default()

	cont := container.New()

	err = cont.Singleton(func() des.CoreService {
		return des.NewCoreService()
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s des.CoreService) aes.CoreService {
		return aes.NewCoreService(s)
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s aes.CoreService) aepc.CommandHandler {
		return aeac.NewCommandHandler(s)
	})
	if err != nil {
		panic(err)
	}

	err = cont.Singleton(func(s aes.CoreService) aepc.QueryHandler {
		return aeac.NewQueryHandler(s)
	})
	if err != nil {
		panic(err)
	}

	if AppConfig.Restapi.Enabled {
		var commandHandler aepc.CommandHandler
		if AppConfig.Restapi.CommandHandler {
			err = cont.Resolve(&commandHandler)
			if err != nil {
				panic(err)
			}
		}

		var queryHandler aepc.QueryHandler
		if AppConfig.Restapi.QueryHandler {
			err = cont.Resolve(&queryHandler)
			if err != nil {
				panic(err)
			}
		}
		pesp.NewRestAPI(c, commandHandler, queryHandler)
		c.Run(AppConfig.Restapi.Host + ":" + AppConfig.Restapi.Port)
	}
}

//TODO: Read data on Config
