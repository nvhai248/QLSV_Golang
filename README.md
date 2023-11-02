# Back-end for Student Management Web Application
* Author: Nguyễn Văn Hải
* GitHub: [nvhai248](https://github.com/nvhai248)

## Overview
This is a simple back-end for a student management web application that provides an interface for teachers to manage their students. It offers basic APIs for managing students, registering for classes, and creating accounts. The project is designed using a microservices architecture for easy scalability.

## Technologies Used

* **Go (Golang)**: Go is a statically typed, compiled language known for its performance and simplicity, serving as the foundation for this project.
* **gin**: Gin is a library that facilitates API development for this project.
* **sqlx**: Sqlx is a library that enables database connectivity to MySQL and data handling.
* **jwt**: JWT is a library used for authentication and authorization.

## Features

* Provides RESTful APIs for student management web applications.
* Implements authentication and authorization using JWT.
* Utilizes a message broker with `pubsub` to enhance performance, especially when handling a large number of student registrations.
* Dockerized for easy deployment. [Docker Repository](https://hub.docker.com/repository/docker/nvhai248/student-management/general)
* Utilizes gRPC for improved API performance.

## Usage

1. Clone the repository:

```
git clone https://github.com/nvhai248/QLSV_Golang
```


2. Install dependencies (optional if the vendor folder is included):

```
go get -u all
```


3. Start the server:

```
go run ./main.go
```


## Contact

If you have any questions or suggestions, please feel free to contact us at nvhai2408@gmail.com.
