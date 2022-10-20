package domain_entry_model_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitle_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestTitle_String",
			args: args{
				title: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTitle(tt.args.title).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestTitle_NewTitle(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want Title
	}{
		{
			name: "TestTitle_NewTitle",
			args: args{
				title: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTitle(tt.args.title)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestTitle_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestTitle_IsEmpty",
			args: args{
				title: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTitle(tt.args.title).IsEmpty()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestTitle_Equals(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		title1 string
		title2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestTitle_Equals_True",
			args: args{
				title1: "test",
				title2: "test",
			},
			want: true,
		},
		{
			name: "TestTitle_Equals_False",
			args: args{
				title1: "test1",
				title2: "test2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewTitle(tt.args.title1)
			got2 := NewTitle(tt.args.title2)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}
