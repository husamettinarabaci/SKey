package application_entry_adapter_cqhandler

import (
	"context"

	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
)

type CQHandler struct {
}

func NewCQHandler() CQHandler {
	return CQHandler{}
}

func (h CQHandler) GetAllEntry(ctx context.Context) ([]deme.Entry, error) {

	return nil, nil
}
func (h CQHandler) GetEntryById(ctx context.Context, uid string) (deme.Entry, error) {

	return deme.Entry{}, nil
}
func (h CQHandler) CreateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return deme.Entry{}, nil
}
func (h CQHandler) UpdateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return deme.Entry{}, nil
}
func (h CQHandler) DeleteEntry(ctx context.Context, uid string) error {

	return nil
}
