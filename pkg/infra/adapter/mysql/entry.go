package infra_adapter_mysql

import (
	"context"

	deme "github.com/husamettinarabaci/SKey/pkg/domain/entry/model/entity"
	demo "github.com/husamettinarabaci/SKey/pkg/domain/entry/model/object"
)

type EntryMySQLRepository struct {
}

func (e *EntryMySQLRepository) Store(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return deme.Entry{}, nil
}
func (e *EntryMySQLRepository) FindById(ctx context.Context, uid demo.UId) (deme.Entry, error) {

	return deme.Entry{}, nil
}
func (e *EntryMySQLRepository) FindAll(ctx context.Context) ([]deme.Entry, error) {

	return []deme.Entry{}, nil
}
func (e *EntryMySQLRepository) Update(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return deme.Entry{}, nil
}
func (e *EntryMySQLRepository) Delete(ctx context.Context, uid demo.UId) error {

	return nil
}
