package domain_entry_model_object

import (
	"testing"

	"github.com/google/uuid"
	demo "github.com/husamettinarabaci/SKey/internal/domain/entry/model/object"
	"github.com/stretchr/testify/assert"
)

func TestEntry_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		id    demo.UId
		entry demo.Entry
	}
	sampeUId := demo.NewUId(uuid.NewString())
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestEntry_String",
			args: args{
				id: sampeUId,
				entry: demo.NewEntry(
					demo.NewTitle("test_title"),
					demo.NewUrl("test_url"),
					demo.NewUserName("test_userName"),
					demo.NewPassword("test_password"),
					demo.NewNote("test_note")),
			},
			want: sampeUId.String() + " test_title test_url test_userName test_password test_note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.id, tt.args.entry).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestEntry_NewEntry(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		id    demo.UId
		entry demo.Entry
	}
	sampeUId := demo.NewUId(uuid.NewString())
	tests := []struct {
		name string
		args args
		want Entry
	}{
		{
			name: "TestEntry_NewEntry",
			args: args{
				id: sampeUId,
				entry: demo.Entry{
					Title:    demo.NewTitle("test_title"),
					Url:      demo.NewUrl("test_url"),
					UserName: demo.NewUserName("test_userName"),
					Password: demo.NewPassword("test_password"),
					Note:     demo.NewNote("test_note"),
				},
			},
			want: Entry{
				Id: sampeUId,
				Entry: demo.Entry{
					Title:    demo.NewTitle("test_title"),
					Url:      demo.NewUrl("test_url"),
					UserName: demo.NewUserName("test_userName"),
					Password: demo.NewPassword("test_password"),
					Note:     demo.NewNote("test_note"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.id, tt.args.entry)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestEntry_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		id    demo.UId
		entry demo.Entry
	}
	sampeUId := demo.NewUId(uuid.NewString())
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestEntry_IsEmpty_Full",
			args: args{
				id: sampeUId,
				entry: demo.Entry{
					Title:    demo.NewTitle("test_title"),
					Url:      demo.NewUrl("test_url"),
					UserName: demo.NewUserName("test_userName"),
					Password: demo.NewPassword("test_password"),
					Note:     demo.NewNote("test_note"),
				},
			},
			want: false,
		},
		{
			name: "TestEntry_IsEmpty_One",
			args: args{
				id: "",
				entry: demo.Entry{
					Title:    demo.NewTitle("test_title"),
					Url:      demo.NewUrl("test_url"),
					UserName: demo.NewUserName("test_userName"),
					Password: demo.NewPassword("test_password"),
					Note:     demo.NewNote("test_note"),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.id, tt.args.entry).IsEmpty()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestEntry_Equals(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		entry1 Entry
		entry2 Entry
	}
	sampeUId1 := demo.NewUId(uuid.NewString())
	sampeUId2 := demo.NewUId(uuid.NewString())
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestEntry_Equals_True",
			args: args{
				entry1: Entry{
					Id: sampeUId1,
					Entry: demo.Entry{
						Title:    demo.NewTitle("test_title"),
						Url:      demo.NewUrl("test_url"),
						UserName: demo.NewUserName("test_userName"),
						Password: demo.NewPassword("test_password"),
						Note:     demo.NewNote("test_note"),
					},
				},
				entry2: Entry{
					Id: sampeUId1,
					Entry: demo.Entry{
						Title:    demo.NewTitle("test_title"),
						Url:      demo.NewUrl("test_url"),
						UserName: demo.NewUserName("test_userName"),
						Password: demo.NewPassword("test_password"),
						Note:     demo.NewNote("test_note"),
					},
				},
			},
			want: true,
		},
		{
			name: "TestEntry_Equals_False",
			args: args{
				entry1: Entry{
					Id: sampeUId1,
					Entry: demo.Entry{
						Title:    demo.NewTitle("test_title"),
						Url:      demo.NewUrl("test_url"),
						UserName: demo.NewUserName("test_userName"),
						Password: demo.NewPassword("test_password"),
						Note:     demo.NewNote("test_note"),
					},
				},
				entry2: Entry{
					Id: sampeUId2,
					Entry: demo.Entry{
						Title:    demo.NewTitle("test_title"),
						Url:      demo.NewUrl("test_url"),
						UserName: demo.NewUserName("test_userName"),
						Password: demo.NewPassword("test_password"),
						Note:     demo.NewNote("test_note"),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewEntry(tt.args.entry1.Id, tt.args.entry1.Entry)
			got2 := NewEntry(tt.args.entry2.Id, tt.args.entry2.Entry)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}
