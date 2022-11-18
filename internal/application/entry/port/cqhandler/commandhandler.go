package application_entry_port_cqhandler

import (
	"context"

	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
)

type CommandHandler interface {
	CreateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error)
	UpdateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error)
	DeleteEntry(ctx context.Context, uid string) error
}
