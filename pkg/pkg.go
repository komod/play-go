package pkg

type MyString struct {
	content string
}

func (s *MyString) String() string {
	return s.content
}

func (s *MyString) Content() string {
	return s.content
}

func (s *MyString) SetContent(str string) {
	s.content = str
}
