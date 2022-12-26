package models

type Member struct {
	Name   string
	Points int
}

func CreateMembers(userName string) *Member {
	return &Member{
		Name:   userName,
		Points: 0,
	}
}
