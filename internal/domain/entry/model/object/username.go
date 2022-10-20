package domain_entry_model_object

type UserName string

func (u UserName) String() string {
	return string(u)
}

func NewUserName(userName string) UserName {
	return UserName(userName)
}

func (u UserName) IsEmpty() bool {
	return u.String() == ""
}

func (u UserName) Equals(userName UserName) bool {
	return u.String() == userName.String()
}
