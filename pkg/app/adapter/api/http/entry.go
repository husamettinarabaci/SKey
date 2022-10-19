package app_adapter_api_http

import (
	"context"

	asde "github.com/husamettinarabaci/SKey/pkg/app/shared/dto/entry"
)

func (api *APIHttp) GetAllEntry(ctx context.Context) ([]asde.Entry, error) {

	return nil, nil
}

func (api *APIHttp) GetEntryById(ctx context.Context, uid string) (asde.Entry, error) {

	return asde.Entry{}, nil
}
