package presentation_entry_service_restapi

import (
	"github.com/gin-gonic/gin"
	aepc "github.com/husamettinarabaci/SKey/internal/application/entry/port/cqhandler"
	pem "github.com/husamettinarabaci/SKey/pkg/presentation/entry/model"
)

type RestAPI struct {
	ginEngine *gin.Engine
	handler   aepc.CQHandler
}

func NewRestAPI(c *gin.Engine, h aepc.CQHandler) *RestAPI {
	api := &RestAPI{
		ginEngine: c,
		handler:   h,
	}
	api.ginEngine.GET("/api/v1/entries", api.GetAllEntry)
	api.ginEngine.GET("/api/v1/entry/:id", api.GetEntryById)
	api.ginEngine.POST("/api/v1/entry", api.CreateEntry)
	api.ginEngine.PATCH("/api/v1/entry", api.UpdateEntry)
	api.ginEngine.DELETE("/api/v1/entry/:id", api.DeleteEntry)
	return api
}

func (api *RestAPI) GetAllEntry(c *gin.Context) {
	entries, err := api.handler.GetAllEntry(c)
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

func (api *RestAPI) GetEntryById(c *gin.Context) {
	id := c.Param("id")
	entry, err := api.handler.GetEntryById(c, id)
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

func (api *RestAPI) CreateEntry(c *gin.Context) {
	var inEntry pem.Entry
	err := c.BindJSON(&inEntry)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	entry, err := api.handler.CreateEntry(c, inEntry.ToEntity())
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

func (api *RestAPI) UpdateEntry(c *gin.Context) {
	var inEntry pem.Entry
	err := c.BindJSON(&inEntry)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	entry, err := api.handler.UpdateEntry(c, inEntry.ToEntity())
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

func (api *RestAPI) DeleteEntry(c *gin.Context) {
	id := c.Param("id")
	err := api.handler.DeleteEntry(c, id)
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
