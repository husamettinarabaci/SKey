package presentation_entry_service_cli

import (
	"github.com/gin-gonic/gin"
	aepc "github.com/husamettinarabaci/SKey/internal/application/entry/port/cqhandler"
	pem "github.com/husamettinarabaci/SKey/pkg/presentation/entry/model"
)

type Cli struct {
	commandHandler aepc.CommandHandler
	queryHandler   aepc.QueryHandler
}

func NewRestAPI(ch aepc.CommandHandler, qh aepc.QueryHandler) *Cli {
	api := &Cli{
		commandHandler: ch,
		queryHandler:   qh,
	}
	if qh != nil {
	}
	if ch != nil {
	}
	return api
}

func (api *Cli) GetAllEntry(c *gin.Context) {
	entries, err := api.queryHandler.GetAllEntry(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var entriesResponse []pem.Entry
		for _, entry := range entries {
			entriesResponse = append(entriesResponse, pem.FromEntity(entry))
		}

		c.JSON(200, gin.H{
			"entries": entriesResponse,
		})
		return
	}
}

func (api *Cli) GetEntryById(c *gin.Context) {
	id := c.Param("id")
	entry, err := api.queryHandler.GetEntryById(c, id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var entryResponse pem.Entry = pem.FromEntity(entry)
		c.JSON(200, gin.H{
			"entry": entryResponse,
		})
		return
	}
}

func (api *Cli) CreateEntry(c *gin.Context) {
	var inEntry pem.Entry
	err := c.BindJSON(&inEntry)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	entry, err := api.commandHandler.CreateEntry(c, inEntry.ToEntity())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var entryResponse pem.Entry = pem.FromEntity(entry)
		c.JSON(200, gin.H{
			"entry": entryResponse,
		})
		return
	}
}

func (api *Cli) UpdateEntry(c *gin.Context) {
	var inEntry pem.Entry
	err := c.BindJSON(&inEntry)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	entry, err := api.commandHandler.UpdateEntry(c, inEntry.ToEntity())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var entryResponse pem.Entry = pem.FromEntity(entry)
		c.JSON(200, gin.H{
			"entry": entryResponse,
		})
		return
	}
}

func (api *Cli) DeleteEntry(c *gin.Context) {
	id := c.Param("id")
	err := api.commandHandler.DeleteEntry(c, id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{})
		return
	}
}
