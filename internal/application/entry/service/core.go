package application_entry_service

import (
	"context"

	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
	des "github.com/husamettinarabaci/SKey/internal/domain/entry/service"
)

type CoreService struct {
	service des.CoreService
}

func NewCoreService(s des.CoreService) CoreService {
	return CoreService{
		service: s,
	}
}

func (c CoreService) GetAllEntry(ctx context.Context) ([]deme.Entry, error) {

	return c.service.GetAllEntry(ctx)
}
func (c CoreService) GetEntryById(ctx context.Context, uid string) (deme.Entry, error) {

	return c.service.GetEntryById(ctx, uid)
}

func (c CoreService) CreateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return c.service.CreateEntry(ctx, entry)
}
func (c CoreService) UpdateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return c.service.UpdateEntry(ctx, entry)
}
func (c CoreService) DeleteEntry(ctx context.Context, uid string) error {

	return c.service.DeleteEntry(ctx, uid)
}
