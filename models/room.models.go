package models

import (
	"github.com/google/uuid"
)

type Rooms struct {
	RoomName string
	Members  map[string]*Member
	Count    int
}

func (r *Rooms) AddMembers(username string) string {
	member := CreateMembers(username)
	id := uuid.New()

	r.Members[id.String()] = member
	return id.String()
}

func (r *Rooms) FindMembers(userName string) bool {
	return r.Members[userName] != nil
}

func (r *Rooms) UpdatePoints(userName string, points int) {
	r.Members[userName].Points = points
}
func CreateRooms(name string) *Rooms {
	return &Rooms{
		RoomName: name,
		Members:  make(map[string]*Member),
	}
}
