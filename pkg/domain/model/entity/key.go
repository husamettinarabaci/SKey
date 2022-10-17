package domain_model_entity

import (
	dmo "github.com/husamettinarabaci/SKey/pkg/domain/model/object"
)

type Key struct {
	Id  dmo.UId
	Key dmo.Key
}

func (k Key) String() string {
	return k.Id.String() + k.Key.String()
}

func NewKey(id dmo.UId, key dmo.Key) Key {
	return Key{
		Id:  id,
		Key: key,
	}
}

func (k Key) IsEmpty() bool {
	return k.Id.IsEmpty()
}

func (k Key) Equals(key Key) bool {
	return k.Id.Equals(key.Id)
}
