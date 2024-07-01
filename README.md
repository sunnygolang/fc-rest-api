# FC - Rest API

## Repository Contents
- Rest API
- Go CHI
- JWT
- Testify
- GORM
- Swag

## Folder Structure
```
fc-rest-api
├── README.md
├── api
│   ├── product.http
│   └── user.http
├── cmd
│   └── server
│       ├── main.go
│       └── test.db
├── configs
│   └── config.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── dto
│   │   └── dto.go
│   ├── entity
│   │   ├── product.go
│   │   └── user.go
│   └── infra
│       ├── database
│       │   ├── interfaces.go
│       │   ├── product_db.go
│       │   └── user_db.go
│       └── webserver
│           └── handlers
│               ├── product_handlers.go
│               └── user_handlers.go
├── pkg
│   └── entity
│       └── id.go
└── test
    ├── product_db_test.go
    ├── product_test.go
    ├── user_db_test.go
    └── user_test.go

15 directories, 24 files
```