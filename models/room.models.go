package models

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
)

type Rooms struct {
	RoomName  string             `json:"roomName"`
	Members   map[string]*Member `json:"members"`
	Count     int                `json:"count"`
	IsVisible bool               `json:"isVisible"`
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

func (r *Rooms) UpdateIsVisible() {
	r.IsVisible = true
}

func (r *Rooms) ResetPoints() {
	r.IsVisible = false
	count := 0
	keys := reflect.ValueOf(r.Members).MapKeys()

	for count < len(keys) {
		key := fmt.Sprint(reflect.ValueOf(keys[count]).Interface())
		r.Members[key].Points = 0
		count += 1
	}
}
func CreateRooms(name string) *Rooms {
	return &Rooms{
		RoomName: name,
		Members:  make(map[string]*Member),
	}
}
