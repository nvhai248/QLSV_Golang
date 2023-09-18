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

# devide directory
mkdir common
# - pagin.go
# - response.go
# - sql_model.go

mkdir component
# - app_context.go

mkdir modules
# microservices
# with each service
# - directory biz (business)
# - model 
# - storage
# - transport

# get AWS to upload photos and files
go get github.com/aws/aws-sdk-go

# handle image upload
go get github.com/disintegration/imaging

#secure id of user
go get -u github.com/btcsuite/btcutil/base58

#jwt install
go get github.com/dgrijalva/jwt-go

# for debugging 
go install -v github.com/go-delve/delve/cmd/dlv@latest
