# gin-gorm

A [golang](https://go.dev/) web server built with [gin](https://gin-gonic.com/) web framework and [gorm](https://gorm.io/index.html) ORM. Uses [sqlite](https://sqlite.org/index.html) as the database.

## Usage

```
# Clone this project 
https://github.com/chilldenaya/gin-gorm.git

# Run locally
make run-dev
```

## Structure
```
├── routes
│   └── student.go // student routes
│   └── main.go // routes initiator
├── services
│   └── student.go // student service or controller
├── models
│   ├── student.go // student model
│   ├── db.go // db initiator
├── dto
│   ├── student.go // student data transfer object
├── helpers
│   └── response.go
└── main.go
```

