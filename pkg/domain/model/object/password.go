package domain_model_object

type Password string

func (p Password) String() string {
	return string(p)
}

func NewPassword(password string) Password {
	return Password(password)
}

func (p Password) IsEmpty() bool {
	return p.String() == ""
}

func (p Password) Equals(password Password) bool {
	return p.String() == password.String()
}
