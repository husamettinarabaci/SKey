package application_entry_adapter_cqhandler

import (
	"context"

	aes "github.com/husamettinarabaci/SKey/internal/application/entry/service"
	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
)

type CommandHandler struct {
	service aes.CoreService
}

func NewCommandHandler(s aes.CoreService) CommandHandler {
	return CommandHandler{
		service: s,
	}
}

func (h CommandHandler) CreateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return h.service.CreateEntry(ctx, entry)
}
func (h CommandHandler) UpdateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return h.service.UpdateEntry(ctx, entry)
}
func (h CommandHandler) DeleteEntry(ctx context.Context, uid string) error {

	return h.service.DeleteEntry(ctx, uid)
}
