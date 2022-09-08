# gin-gorm

A [golang](https://go.dev/) web server built with [gin](https://gin-gonic.com/) web framework and [gorm](https://gorm.io/index.html) ORM. Uses [sqlite](https://sqlite.org/index.html) as the database.

## Usage

```
# Clone this project 
https://github.com/chilldenaya/gin-gorm.git

# Run locally
make run-dev
```

## Main Structure
```
├── routes
│   └── student.go // student routes and controller callers
│   └── main.go // routes initiator
├── controllers
│   └── student.go // student controllers
├── models
│   ├── student.go // student model
│   ├── db.go // db initiator
├── dto
│   ├── student.go // student data transfer object (request and response)
└── main.go
```

