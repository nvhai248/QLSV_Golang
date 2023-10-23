package skstudent

import (
	"log"
	"studyGoApp/component"

	socketio "github.com/googollee/go-socket.io"
)

type locationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnStudentUpdateLocation(appCtx component.AppContext) func(s socketio.Conn, location locationData) {
	return func(s socketio.Conn, location locationData) {
		log.Println("location: ", location)
	}
}
