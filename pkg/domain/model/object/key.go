package domain_model_object

type Key struct {
	Title    Title
	Url      Url
	UserName UserName
	Password Password
	Note     Note
}

func (k Key) String() string {
	return k.Title.String() + k.Url.String() + k.UserName.String() + k.Password.String() + k.Note.String()
}

func NewKey(title Title, url Url, userName UserName, password Password, note Note) Key {
	return Key{
		Title:    title,
		Url:      url,
		UserName: userName,
		Password: password,
		Note:     note,
	}
}

func (k Key) IsEmpty() bool {
	isEmpty := false
	if k.Title.IsEmpty() && k.Url.IsEmpty() && k.UserName.IsEmpty() && k.Password.IsEmpty() && k.Note.IsEmpty() {
		isEmpty = true
	}
	return isEmpty
}

func (k Key) Equals(key Key) bool {
	isEqual := false
	if k.Title.Equals(key.Title) && k.Url.Equals(key.Url) && k.UserName.Equals(key.UserName) && k.Password.Equals(key.Password) && k.Note.Equals(key.Note) {
		isEqual = true
	}
	return isEqual
}
