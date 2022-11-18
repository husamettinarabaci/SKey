package domain_entry_service

import (
	"context"
	"fmt"

	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
)

type CoreService struct {
}

func NewCoreService() CoreService {
	return CoreService{}
}

func (c CoreService) GetAllEntry(ctx context.Context) ([]deme.Entry, error) {
	fmt.Println("GetAllEntry")
	return []deme.Entry{}, nil
}
func (c CoreService) GetEntryById(ctx context.Context, uid string) (deme.Entry, error) {
	fmt.Println("GetEntryById")
	fmt.Println(uid)
	return deme.Entry{}, nil
}

func (c CoreService) CreateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {
	fmt.Println("CreateEntry")
	fmt.Println(entry)
	return deme.Entry{}, nil
}
func (c CoreService) UpdateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {
	fmt.Println("UpdateEntry")
	fmt.Println(entry)
	return deme.Entry{}, nil
}
func (c CoreService) DeleteEntry(ctx context.Context, uid string) error {
	fmt.Println("DeleteEntry")
	fmt.Println(uid)
	return nil
}
