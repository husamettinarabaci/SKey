package domain_model_object

type Note string

func (n Note) String() string {
	return string(n)
}

func NewNote(note string) Note {
	return Note(note)
}

func (n Note) IsEmpty() bool {
	return n.String() == ""
}

func (n Note) Equals(note Note) bool {
	return n.String() == note.String()
}
