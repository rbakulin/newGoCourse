package contact

type Contact struct {
	Name string
	Phone string
	Email string
}

type PhoneBook []*Contact

func (b PhoneBook) Len() int {
	return len(b)
}

func (b PhoneBook) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b PhoneBook) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}
