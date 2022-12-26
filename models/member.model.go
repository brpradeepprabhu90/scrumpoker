package models

type Member struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func CreateMembers(userName string) *Member {
	return &Member{
		Name:   userName,
		Points: 0,
	}
}
