package presentation_entry_model

import (
	deme "github.com/husamettinarabaci/SKey/internal/domain/entry/model/entity"
	demo "github.com/husamettinarabaci/SKey/internal/domain/entry/model/object"
)

// TODO: Fix DTO
type Entry struct {
	Id       string `json:"id"`
	Title    string `json:"title" binding:"required"`
	Url      string `json:"url"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Note     string `json:"note"`
}

func (e Entry) String() string {
	return e.Id + " " + e.Title + " " + e.Url + " " + e.UserName + " " + e.Password + " " + e.Note
}

func NewEntry(id string, title string, url string, userName string, password string, note string) Entry {
	return Entry{
		Id:       id,
		Title:    title,
		Url:      url,
		UserName: userName,
		Password: password,
		Note:     note,
	}
}

func (e Entry) IsEmpty() bool {
	if e.Id == "" && e.Title == "" && e.Url == "" && e.UserName == "" && e.Password == "" && e.Note == "" {
		return true
	}
	return false
}

func (e Entry) Equals(entry Entry) bool {
	if e.Id == entry.Id && e.Title == entry.Title && e.Url == entry.Url && e.UserName == entry.UserName && e.Password == entry.Password && e.Note == entry.Note {
		return true
	}
	return false
}

func (e Entry) ToEntity() deme.Entry {
	return deme.Entry{
		Id: demo.NewUId(e.Id),
		Entry: demo.Entry{
			Title:    demo.NewTitle(e.Title),
			Url:      demo.NewUrl(e.Url),
			UserName: demo.NewUserName(e.UserName),
			Password: demo.NewPassword(e.Password),
			Note:     demo.NewNote(e.Note),
		},
	}
}

func FromEntity(entry deme.Entry) Entry {
	return Entry{
		Id:       entry.Id.String(),
		Title:    entry.Entry.Title.String(),
		Url:      entry.Entry.Url.String(),
		UserName: entry.Entry.UserName.String(),
		Password: entry.Entry.Password.String(),
		Note:     entry.Entry.Note.String(),
	}
}
