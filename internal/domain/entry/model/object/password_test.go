package domain_entry_model_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestPassword_String",
			args: args{
				password: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPassword(tt.args.password).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestPassword_NewPassword(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want Password
	}{
		{
			name: "TestPassword_NewPassword",
			args: args{
				password: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPassword(tt.args.password)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestPassword_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestPassword_IsEmpty",
			args: args{
				password: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPassword(tt.args.password).IsEmpty()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestPassword_Equals(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		password1 string
		password2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestPassword_Equals_True",
			args: args{
				password1: "test",
				password2: "test",
			},
			want: true,
		},
		{
			name: "TestPassword_Equals_False",
			args: args{
				password1: "test1",
				password2: "test2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewPassword(tt.args.password1)
			got2 := NewPassword(tt.args.password2)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}
