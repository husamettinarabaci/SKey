package infrastructure_entry_adapter_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
	iem "github.com/husamettinarabaci/SKey/pkg/infrastructure/entry/model"
)

type Repository struct {
	filePath string
}

func NewRepository(filePath string) Repository {
	return Repository{
		filePath: filePath,
	}
}

func (h Repository) GetAllEntry(ctx context.Context) ([]deme.Entry, error) {
	fmt.Println(h.filePath)
	return nil, nil
}
func (h Repository) GetEntryById(ctx context.Context, uid string) (deme.Entry, error) {

	return iem.Entry{}.ToEntity(), nil
}
func (h Repository) CreateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	fmt.Println(h.filePath)
	var iEntry iem.Entry = iem.FromEntity(entry)
	file, _ := json.MarshalIndent(iEntry, "", " ")
	err := ioutil.WriteFile(h.filePath, file, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return deme.Entry{}, nil
}
func (h Repository) UpdateEntry(ctx context.Context, entry deme.Entry) (deme.Entry, error) {

	return deme.Entry{}, nil
}
func (h Repository) DeleteEntry(ctx context.Context, uid string) error {

	return nil
}
