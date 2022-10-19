package domain_entry_model_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestUrl_String",
			args: args{
				url: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUrl(tt.args.url).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUrl_NewUrl(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want Url
	}{
		{
			name: "TestUrl_NewUrl",
			args: args{
				url: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUrl(tt.args.url)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUrl_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestUrl_IsEmpty",
			args: args{
				url: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUrl(tt.args.url).IsEmpty()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUrl_Equals(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		url1 string
		url2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestUrl_Equals_True",
			args: args{
				url1: "test",
				url2: "test",
			},
			want: true,
		},
		{
			name: "TestUrl_Equals_False",
			args: args{
				url1: "test1",
				url2: "test2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewUrl(tt.args.url1)
			got2 := NewUrl(tt.args.url2)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)

		})
	}
}
