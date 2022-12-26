package main

import (
	"encoding/json"
	"fmt"
	"github.com/brpradeepprabhu90/scrumpoker/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)

var rooms = make(map[string]*models.Rooms)

func createModel(roomName string) bool {
	if !checkModel(roomName) {
		rooms[roomName] = models.CreateRooms(roomName)
		return true
	}
	return false
}
func isUserPresent(roomName string, userName string) bool {
	d := rooms[roomName]
	return d.FindMembers(userName)

}

func checkUserInRoom(roomName string, userName string) string {
	if checkModel(roomName) {
		d := rooms[roomName]
		valid := isUserPresent(roomName, userName)
		if !valid {
			uuid := d.AddMembers(userName)
			return uuid
		} else {
			return "false"
		}
	}
	return "false"
}
func checkModel(roomName string) bool {
	_, found := rooms[roomName]
	return found
}

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var clients = make(map[*websocket.Conn]*client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func implementMessage(message string) {
	s := models.Communuication{}
	if err := json.Unmarshal([]byte(message), &s); err != nil {
		log.Println(err)
	}
	d := rooms[s.Message.RoomName]
	switch s.Message.Type {
	case "updateUser":
		d.UpdatePoints(s.Message.Username, s.Message.Points)
	case "revealCards":
		d.UpdateIsVisible()
	case "resetCards":
		d.ResetPoints()
	}

}

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}
			log.Println("connection registered")

		case message := <-broadcast:
			log.Println("message received:", message)
			// Send the message to all clients
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) { // send to each client in parallel so we don't block on a slow client
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}

					implementMessage(message)

					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Println("write error:", err)

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("connection unregistered")
		}
	}
}
func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/createRoom", func(c *fiber.Ctx) error {
		payload := struct {
			RoomName string `json:"roomName"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		isRoomCreated := createModel(payload.RoomName)
		if isRoomCreated {
			return c.JSON(models.Message{
				Message: "Room " + payload.RoomName + " is created",
			})
		}
		return c.Status(500).JSON(models.Message{
			Message: "Room " + payload.RoomName + " is already created",
		})
	})

	app.Post("/createUser", func(c *fiber.Ctx) error {
		payload := struct {
			RoomName string `json:"roomName"`
			UserName string `json:"userName"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		if !checkModel(payload.RoomName) {
			return c.Status(500).JSON(models.Message{
				Message: "Room " + payload.RoomName + " is not created",
			})
		}
		uuid := checkUserInRoom(payload.RoomName, payload.UserName)
		if uuid != "false" {
			return c.Status(200).JSON(models.Message{
				Message: uuid,
			})
		}

		return c.Status(500).JSON(models.Message{
			Message: "User " + payload.UserName + " is already present",
		})
	})
	app.Get("isUserPresent/:roomName/:userName", func(c *fiber.Ctx) error {
		roomName := c.Params("roomName")
		userName := c.Params("userName")
		if !checkModel(roomName) {
			return c.Status(500).JSON(models.Message{
				Message: "Room " + roomName + " is not created",
			})
		}
		if isUserPresent(roomName, userName) {
			return c.JSON(models.Message{
				Message: "User " + userName + " is added to room",
			})
		}
		return c.Status(500).JSON(models.Message{
			Message: "User " + userName + " is not present",
		})
	})
	app.Get("getUsers/:roomName", func(c *fiber.Ctx) error {
		roomName := c.Params("roomName")
		if !checkModel(roomName) {
			return c.Status(500).JSON(models.Message{
				Message: "Room " + roomName + " is not created",
			})
		}
		return c.JSON(rooms[roomName])

	})
	/*** Webs sockets ***************/
	go runHub()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:roomName", websocket.New(func(c *websocket.Conn) {
		fmt.Println("connected") // "Localhost:3000"
		//userName := c.Params("userName")
		defer func() {
			unregister <- c
			c.Close()
		}()
		register <- c
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			s := models.Communuication{}
			if err := json.Unmarshal(msg, &s); err != nil {
				log.Println(err)
			}

			broadcast <- string(msg)

			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}
