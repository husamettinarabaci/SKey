package application_entry_port_cqhandler

import (
	"context"

	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
)

type QueryHandler interface {
	GetAllEntry(ctx context.Context) ([]deme.Entry, error)
	GetEntryById(ctx context.Context, uid string) (deme.Entry, error)
}
