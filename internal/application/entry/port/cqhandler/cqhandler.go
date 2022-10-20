package application_entry_port_cqhandler

import (
	"context"

	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
)

type CQHandler interface {
	GetAllEntry(ctx context.Context) ([]deme.Entry, error)
	GetEntryById(ctx context.Context, uid string) (deme.Entry, error)
	CreateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error)
	UpdateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error)
	DeleteEntry(ctx context.Context, uid string) error
}
