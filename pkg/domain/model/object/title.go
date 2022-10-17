package domain_model_object

type Title string

func (t Title) String() string {
	return string(t)
}

func NewTitle(title string) Title {
	return Title(title)
}

func (t Title) IsEmpty() bool {
	return t.String() == ""
}

func (t Title) Equals(title Title) bool {
	return t.String() == title.String()
}
