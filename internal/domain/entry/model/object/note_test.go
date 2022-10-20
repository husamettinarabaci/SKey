package domain_entry_model_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNote_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		note string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestNote_String",
			args: args{
				note: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNote(tt.args.note).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestNote_NewNote(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		note string
	}
	tests := []struct {
		name string
		args args
		want Note
	}{
		{
			name: "TestNote_NewNote",
			args: args{
				note: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNote(tt.args.note)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestNote_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		note string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestNote_IsEmpty",
			args: args{
				note: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNote(tt.args.note).IsEmpty()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestNote_Equals(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		note1 string
		note2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestNote_Equals_True",
			args: args{
				note1: "test",
				note2: "test",
			},
			want: true,
		},
		{
			name: "TestNote_Equals_False",
			args: args{
				note1: "test1",
				note2: "test2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewNote(tt.args.note1)
			got2 := NewNote(tt.args.note2)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}
