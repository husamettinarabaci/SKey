package domain_entry_model_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserName_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestUserName_String",
			args: args{
				userName: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserName(tt.args.userName).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUserName_NewUserName(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		want UserName
	}{
		{
			name: "TestUserName_NewUserName",
			args: args{
				userName: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserName(tt.args.userName)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUserName_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestUserName_IsEmpty",
			args: args{
				userName: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserName(tt.args.userName).IsEmpty()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUserName_Equals(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		userName1 string
		userName2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestUserName_Equals_True",
			args: args{
				userName1: "test",
				userName2: "test",
			},
			want: true,
		},
		{
			name: "TestUserName_Equals_False",
			args: args{
				userName1: "test1",
				userName2: "test2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewUserName(tt.args.userName1)
			got2 := NewUserName(tt.args.userName2)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}
