# init file app.go
echo. > app.go

#add go.mod
go mod init studyGoApp

# get libraries for secure (environment)
go get github.com/joho/godotenv

# libraries to connect to MySQL DB and working 
go get -u github.com/go-sql-driver/mysql
go get github.com/jmoiron/sqlx

# REST API
go get -u github.com/gin-gonic/gin