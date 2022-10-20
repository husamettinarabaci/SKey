package domain_entry_model_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntry_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		title    Title
		url      Url
		userName UserName
		password Password
		note     Note
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestEntry_String",
			args: args{
				title:    NewTitle("test_title"),
				url:      NewUrl("test_url"),
				userName: NewUserName("test_userName"),
				password: NewPassword("test_password"),
				note:     NewNote("test_note"),
			},
			want: "test_title test_url test_userName test_password test_note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.title, tt.args.url, tt.args.userName, tt.args.password, tt.args.note).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestEntry_NewEntry(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		title    Title
		url      Url
		userName UserName
		password Password
		note     Note
	}
	tests := []struct {
		name string
		args args
		want Entry
	}{
		{
			name: "TestEntry_NewEntry",
			args: args{
				title:    NewTitle("test_title"),
				url:      NewUrl("test_url"),
				userName: NewUserName("test_userName"),
				password: NewPassword("test_password"),
				note:     NewNote("test_note"),
			},
			want: Entry{
				Title:    NewTitle("test_title"),
				Url:      NewUrl("test_url"),
				UserName: NewUserName("test_userName"),
				Password: NewPassword("test_password"),
				Note:     NewNote("test_note"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.title, tt.args.url, tt.args.userName, tt.args.password, tt.args.note)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestEntry_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		title    Title
		url      Url
		userName UserName
		password Password
		note     Note
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestEntry_IsEmpty_Full",
			args: args{
				title:    NewTitle("test_title"),
				url:      NewUrl("test_url"),
				userName: NewUserName("test_userName"),
				password: NewPassword("test_password"),
				note:     NewNote("test_note"),
			},
			want: false,
		},
		{
			name: "TestEntry_IsEmpty_One",
			args: args{
				title:    NewTitle("test_title"),
				url:      NewUrl(""),
				userName: NewUserName(""),
				password: NewPassword(""),
				note:     NewNote(""),
			},
			want: false,
		},
		{
			name: "TestEntry_IsEmpty_None",
			args: args{
				title:    NewTitle(""),
				url:      NewUrl(""),
				userName: NewUserName(""),
				password: NewPassword(""),
				note:     NewNote(""),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEntry(tt.args.title, tt.args.url, tt.args.userName, tt.args.password, tt.args.note).IsEmpty()
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
					Title:    NewTitle("test_title"),
					Url:      NewUrl("test_url"),
					UserName: NewUserName("test_userName"),
					Password: NewPassword("test_password"),
					Note:     NewNote("test_note"),
				},
				entry2: Entry{
					Title:    NewTitle("test_title"),
					Url:      NewUrl("test_url"),
					UserName: NewUserName("test_userName"),
					Password: NewPassword("test_password"),
					Note:     NewNote("test_note"),
				},
			},
			want: true,
		},
		{
			name: "TestEntry_Equals_False",
			args: args{
				entry1: Entry{
					Title:    NewTitle("test_title"),
					Url:      NewUrl("test_url"),
					UserName: NewUserName("test_userName"),
					Password: NewPassword("test_password"),
					Note:     NewNote("test_note"),
				},
				entry2: Entry{
					Title:    NewTitle("test_title_1"),
					Url:      NewUrl("test_url"),
					UserName: NewUserName("test_userName"),
					Password: NewPassword("test_password"),
					Note:     NewNote("test_note"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewEntry(tt.args.entry1.Title, tt.args.entry1.Url, tt.args.entry1.UserName, tt.args.entry1.Password, tt.args.entry1.Note)
			got2 := NewEntry(tt.args.entry2.Title, tt.args.entry2.Url, tt.args.entry2.UserName, tt.args.entry2.Password, tt.args.entry2.Note)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}
