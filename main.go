package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/component/uploadprovider"
	"studyGoApp/middleware"
	"studyGoApp/modules/class/classtransport/ginclass"
	classregistrationgrpc "studyGoApp/modules/classregister/storage/grpc"
	"studyGoApp/modules/classregister/transport/ginclassregister"
	"studyGoApp/modules/student/studenttransport/ginstudent"
	"studyGoApp/modules/upload/uploadtransport/ginupload"
	"studyGoApp/proto"
	"studyGoApp/pubsub/pblocal"
	"studyGoApp/skio"
	"studyGoApp/subscriber"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func ConnectToDB(dns string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// Check connection to DB by Ping
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	fmt.Println("Connected to the database!")

	return db
}

func runServices(db *sqlx.DB, secretKey string, upProvider uploadprovider.UploadProvider) {

	appCtx := component.NewAppContext(db, secretKey, upProvider, pblocal.NewPubSub())

	router := gin.Default()

	// call pubsub
	//subscriber.Setup(appCtx)

	rtEngine := skio.NewEngine()

	if err := rtEngine.Run(appCtx, router); err != nil {
		log.Fatalln("Failed to start server: ", err)
	}

	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("...Hello Client...")

	router.Use(middleware.Recover(appCtx))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping 5555",
		})
	})

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	//router.StaticFile("/demo/", "./demo.html")
	v1 := router.Group("/v1")

	// upload photo
	v1.POST("/upload", ginupload.Upload(appCtx))

	// CRUD student
	// authentication and authorization with JWT
	v1.POST("/students/register", ginstudent.CreateStudent(appCtx))
	v1.POST("/students/login", ginstudent.Login(appCtx))

	MainStudent(appCtx, v1)

	/* students := v1.Group("/students", middleware.RequireAuth(appCtx))
	{
		students.GET("/profile", ginstudent.GetProfile(appCtx))

		students.GET("", ginstudent.ListStudent(appCtx))
		students.GET("/:id", ginstudent.DetailStudent(appCtx))
		students.PATCH("/:id", ginstudent.UpdateStudent(appCtx))
		students.DELETE("/:id", ginstudent.SoftDeleteStudent(appCtx))
	} */

	// check role admin
	admin := v1.Group("/admin", middleware.RequireAuth(appCtx), middleware.RequireRoles(appCtx, "admin"))
	{
		admin.GET("/profile", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("He is a admin!"))
		})
	}

	classes := v1.Group("/classes", middleware.RequireAuth(appCtx))
	{
		classes.GET("", ginclass.ListClass(appCtx, cc))
		classes.GET("/:id", ginclass.FindClass(appCtx))
		classes.POST("", ginclass.CreateClass(appCtx))
		classes.DELETE("/:id/cancel_registration", ginclassregister.StudentCancelRegisterClass(appCtx))
		classes.POST("/:id/register", ginclassregister.StudentRegisterClass(appCtx))
		classes.GET("/:id/registered_student", ginclass.GetListRegisteredStudents(appCtx))
	}

	// Create a listener on TCP port
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	proto.RegisterClassRegistrationServiceServer(s, classregistrationgrpc.NewgRPCServer(db))

	go func() {
		// Serve gRPC Server
		log.Println("Serving gRPC on ", address)
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	router.Run(":8080")
}

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	dns := os.Getenv("DBConnectionStr")

	// Connect to the database
	db := ConnectToDB(dns)

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3ApiKey := os.Getenv("S3ApiKey")
	s3Secret := os.Getenv("S3Secret")
	s3Domain := os.Getenv("S3Domain")
	s3upProvider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, s3Secret, s3Domain)
	secretKey := os.Getenv("SYSTEM_SECRET")

	runServices(db, secretKey, s3upProvider)
}

/* func startSocketIoServer(engine *gin.Engine, appCtx component.AppContext) {

	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500") // Set the allowed origin(s)
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

	// Create a new Socket.IO server
	server := socketio.NewServer(nil)

	// Handle connections
	server.OnConnect("/", func(s socketio.Conn) error {
		// Set up CORS
		s.SetContext("")
		s.RemoteHeader().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")

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

	// Handle authentication
	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		// Implement your authentication logic here
		db := appCtx.GetMainDBConnection()
		store := studentstorage.NewSQLStore(db)

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

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

		s.Emit("your profile", user)
	})

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
}
*/
