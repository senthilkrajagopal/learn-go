package senthil

type user struct {
	name string
}

func (u *user) setName(name string) {
	u.name = name
}

func (u user) getName() string {
	return u.name
}
