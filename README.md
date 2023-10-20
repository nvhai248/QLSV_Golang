# Back-end for Student management Web application
* Author: Nguyễn Văn Hải
* nvhai248

## Overview
This is the simple back-end for student management web application that provides a simple interface for teacher to manage their students. It's provided some basic APIs to manage students, register to the classes and register account.

I used microservices architecture for this project to simple scale up.

## Technologies Used

* `Go (Golang)` : Go is a statically typed, compiled language known for its performance and simplicity. It serves as the foundation for this project.
* `gin` : gin is a library that help me write apis for this project.
* `sqlx` : sqlx is a library that help me connect to the database `mySQL` and handle the data.
* `jwt` : jwt is a library that helps authenticate and authorize.

## Features

* Provide APIs for student management web applications based on RESTful API
* Authenticate and authorized by `jwt`
* Message broker by `pubsub` to improve the performance when a lot of students register to the class and vice versa

## Usage

1. Clone the repository:

```
git clone https://github.com/nvhai248/QLSV_Golang
```

2. Install dependencies (Optional if vendor folder included):

```
go get -u all
```

3. Start the server:

```
go run ./app.go
```

## Contact

If you have any questions or suggestions, please feel free to contact us at nvhai2408@gmail.com
