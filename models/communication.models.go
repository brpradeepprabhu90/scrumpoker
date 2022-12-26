package models

type Communuication struct {
	Message WebSocketMessage
}

type WebSocketMessage struct {
	Type     string `json:"type"`
	RoomName string `json:"roomName"`
	Username string `json:"username"`
	Points   int    `json:"points"`
}
