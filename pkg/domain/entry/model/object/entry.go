package domain_entry_model_object

type Entry struct {
	Title    Title
	Url      Url
	UserName UserName
	Password Password
	Note     Note
}

func (e Entry) String() string {
	return e.Title.String() + " " + e.Url.String() + " " + e.UserName.String() + " " + e.Password.String() + " " + e.Note.String()
}

func NewEntry(title Title, url Url, userName UserName, password Password, note Note) Entry {
	return Entry{
		Title:    title,
		Url:      url,
		UserName: userName,
		Password: password,
		Note:     note,
	}
}

func (e Entry) IsEmpty() bool {
	isEmpty := false
	if e.Title.IsEmpty() && e.Url.IsEmpty() && e.UserName.IsEmpty() && e.Password.IsEmpty() && e.Note.IsEmpty() {
		isEmpty = true
	}
	return isEmpty
}

func (e Entry) Equals(entry Entry) bool {
	isEqual := false
	if e.Title.Equals(entry.Title) && e.Url.Equals(entry.Url) && e.UserName.Equals(entry.UserName) && e.Password.Equals(entry.Password) && e.Note.Equals(entry.Note) {
		isEqual = true
	}
	return isEqual
}

//TODO: ExpireDate
