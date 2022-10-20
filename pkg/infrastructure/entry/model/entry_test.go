package infrastructure_entry_model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntry_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		entry Entry
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestEntry_String",
			args: args{
				entry: NewEntry(
					"test_id",
					"test_title",
					"test_url",
					"test_userName",
					"test_password",
					"test_note"),
			},
			want: "test_id test_title test_url test_userName test_password test_note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.entry.Id, tt.args.entry.Title, tt.args.entry.Url, tt.args.entry.UserName, tt.args.entry.Password, tt.args.entry.Note).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestEntry_NewEntry(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		entry Entry
	}
	tests := []struct {
		name string
		args args
		want Entry
	}{
		{
			name: "TestEntry_NewEntry",
			args: args{
				entry: NewEntry(
					"test_id",
					"test_title",
					"test_url",
					"test_userName",
					"test_password",
					"test_note"),
			},
			want: Entry{
				Id:       "test_id",
				Title:    "test_title",
				Url:      "test_url",
				UserName: "test_userName",
				Password: "test_password",
				Note:     "test_note",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.entry.Id, tt.args.entry.Title, tt.args.entry.Url, tt.args.entry.UserName, tt.args.entry.Password, tt.args.entry.Note)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestEntry_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		entry Entry
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestEntry_IsEmpty_False",
			args: args{
				entry: NewEntry(
					"test_id",
					"test_title",
					"test_url",
					"test_userName",
					"test_password",
					"test_note"),
			},
			want: false,
		},
		{
			name: "TestEntry_IsEmpty_True",
			args: args{
				entry: NewEntry(
					"",
					"",
					"",
					"",
					"",
					""),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.entry.Id, tt.args.entry.Title, tt.args.entry.Url, tt.args.entry.UserName, tt.args.entry.Password, tt.args.entry.Note).IsEmpty()
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
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestEntry_Equals_True",
			args: args{
				entry1: Entry{
					Id:       "test_id",
					Title:    "test_title",
					Url:      "test_url",
					UserName: "test_userName",
					Password: "test_password",
					Note:     "test_note",
				},
				entry2: Entry{
					Id:       "test_id",
					Title:    "test_title",
					Url:      "test_url",
					UserName: "test_userName",
					Password: "test_password",
					Note:     "test_note",
				},
			},
			want: true,
		},
		{
			name: "TestEntry_Equals_False",
			args: args{
				entry1: Entry{
					Id:       "test_id",
					Title:    "test_title",
					Url:      "test_url",
					UserName: "test_userName",
					Password: "test_password",
					Note:     "test_note",
				},
				entry2: Entry{
					Id:       "test_id2",
					Title:    "test_title",
					Url:      "test_url",
					UserName: "test_userName",
					Password: "test_password",
					Note:     "test_note",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewEntry(tt.args.entry1.Id, tt.args.entry1.Title, tt.args.entry1.Url, tt.args.entry1.UserName, tt.args.entry1.Password, tt.args.entry1.Note)
			got2 := NewEntry(tt.args.entry2.Id, tt.args.entry2.Title, tt.args.entry2.Url, tt.args.entry2.UserName, tt.args.entry2.Password, tt.args.entry2.Note)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}
