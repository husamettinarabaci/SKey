package app_shared_dto_entry

import (
	demo "github.com/husamettinarabaci/SKey/pkg/domain/entry/model/object"
)

type Entry struct {
	Id    demo.UId
	Entry demo.Entry
}

func (e Entry) String() string {
	return e.Id.String() + " " + e.Entry.String()
}

func NewEntry(id demo.UId, entry demo.Entry) Entry {
	return Entry{
		Id:    id,
		Entry: entry,
	}
}

func (e Entry) IsEmpty() bool {
	return e.Id.IsEmpty()
}

func (e Entry) Equals(entry Entry) bool {
	return e.Id.Equals(entry.Id)
}
