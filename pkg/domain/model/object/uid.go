package domain_model_object

type UId string

func (u UId) String() string {
	return string(u)
}

func NewUId(uid string) UId {
	return UId(uid)
}

func (u UId) IsEmpty() bool {
	return u.String() == ""
}

func (u UId) Equals(uid UId) bool {
	return u.String() == uid.String()
}

//TODO: Is Valid UUID
