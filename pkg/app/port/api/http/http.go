package app_port_api_htpp

import (
	"context"

	asde "github.com/husamettinarabaci/SKey/pkg/app/shared/dto/entry"
)

type APIHttp interface {
	GetAllEntry(context.Context) ([]asde.Entry, error)
	GetEntryById(context.Context, string) (asde.Entry, error)
}
