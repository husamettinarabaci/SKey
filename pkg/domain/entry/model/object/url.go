package domain_entry_model_object

type Url string

func (u Url) String() string {
	return string(u)
}

func NewUrl(url string) Url {
	return Url(url)
}

func (u Url) IsEmpty() bool {
	return u.String() == ""
}

func (u Url) Equals(url Url) bool {
	return u.String() == url.String()
}
