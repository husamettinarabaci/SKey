package application_entry_adapter_cqhandler

import (
	"context"

	aes "github.com/husamettinarabaci/SKey/internal/application/entry/service"
	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
)

type QueryHandler struct {
	service aes.CoreService
}

func NewQueryHandler(s aes.CoreService) QueryHandler {
	return QueryHandler{
		service: s,
	}
}

func (h QueryHandler) GetAllEntry(ctx context.Context) ([]deme.Entry, error) {

	return h.service.GetAllEntry(ctx)
}
func (h QueryHandler) GetEntryById(ctx context.Context, uid string) (deme.Entry, error) {

	return h.service.GetEntryById(ctx, uid)
}
