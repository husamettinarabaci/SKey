package domain_entry_model_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUId_String(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		uid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestUId_String",
			args: args{
				uid: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUId(tt.args.uid).String()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUId_GenerateUId(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
	}{
		{
			name: "TestUId_GenerateUId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateUId()
			assert.NotEmpty(got, tt.name)
		})
	}
}

func TestUId_NewUId(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		uid string
	}
	tests := []struct {
		name string
		args args
		want UId
	}{
		{
			name: "TestUId_NewUId",
			args: args{
				uid: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUId(tt.args.uid)
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUId_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		uid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestUId_IsEmpty",
			args: args{
				uid: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUId(tt.args.uid).IsEmpty()
			assert.Equal(tt.want, got, tt.name)
		})
	}
}

func TestUId_Equals(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		uid1 string
		uid2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestUId_Equals_True",
			args: args{
				uid1: "test",
				uid2: "test",
			},
			want: true,
		},
		{
			name: "TestUId_Equals_False",
			args: args{
				uid1: "test1",
				uid2: "test2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := NewUId(tt.args.uid1)
			got2 := NewUId(tt.args.uid2)
			assert.Equal(tt.want, got1.Equals(got2), tt.name)
		})
	}
}

func TestUId_IsValidUUID(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
	}{
		{
			name: "TestUId_IsValidUUID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateUId()
			assert.True(got.IsValidUUID(), tt.name)
		})
	}
}
