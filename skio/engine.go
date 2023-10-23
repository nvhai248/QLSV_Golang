package skio

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"studyGoApp/component"
	"studyGoApp/component/tokenprovider/jwt"
	"studyGoApp/modules/student/studentstorage"
	"studyGoApp/modules/student/studenttransport/skstudent"
	"sync"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(ctx component.AppContext, engine *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *rtEngine) saveAppSocket(userId int, appSkt AppSocket) {
	engine.locker.Lock()

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSkt)
	} else {
		engine.storage[userId] = []AppSocket{appSkt}
	}

	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket

	if scks, ok := engine.storage[userId]; ok {
		return scks
	}

	return sockets
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)

	for _, s := range sockets {
		s.conn.Emit(key, data)
	}

	return nil
}

func (en *rtEngine) authenticationForClientBySk(ctx component.AppContext, engine *gin.Engine) func(s socketio.Conn, token string) {
	return func(s socketio.Conn, token string) {
		// Implement your authentication logic here
		db := ctx.GetMainDBConnection()
		store := studentstorage.NewSQLStore(db)

		tokenProvider := jwt.NewTokenJWTProvider(ctx.SecretKey())

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			s.Emit("authentication failed", err.Error())
			s.Close()
			return
		}

		user, err := store.DetailStudent(context.Background(), payload.UserId)

		if err != nil {
			s.Emit("authentication failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication failed", errors.New("you had been banned/deleted"))
			s.Close()
			return
		}

		appSck := NewAppSocket(s, user)
		en.saveAppSocket(user.Id, *appSck)

		s.Emit("your profile", user)
	}
}

func (en *rtEngine) Run(ctx component.AppContext, engine *gin.Engine) error {
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENTS")) // Set the allowed origin(s)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	})

	server := socketio.NewServer(nil)

	en.server = server

	// Handle connections
	server.OnConnect("/", func(s socketio.Conn) error {
		// Set up CORS
		s.SetContext("")
		s.RemoteHeader().Set("Access-Control-Allow-Origin", os.Getenv("CLIENTS"))

		// Log the connection
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())

		return nil
	})

	// Handle errors
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error: ", e)
	})

	// Handle disconnections
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("Closed: ", reason)
		// Remove socket from the socket engine (from app context) if necessary
	})

	// setup
	server.OnEvent("/", "", skstudent.OnStudentUpdateLocation(ctx))

	// Handle authentication
	server.OnEvent("/", "location", en.authenticationForClientBySk(ctx, engine))

	// Handle test event
	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
		log.Println(msg)
	})

	// Define a Person struct
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Handle notice event
	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
		fmt.Println("server received notice:", p.Name, p.Age)
		p.Age = 33
		s.Emit("notice", p)
	})

	// Start the Socket.IO server
	go server.Serve()

	// Handle HTTP requests for the Socket.IO server
	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))

	return nil
}
