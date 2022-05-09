# 0. Overview
간단한 Ecommerce Project입니다...

Golang Version : 1.17

Docker version : 20.10.14

mongo:5.0.3

image: mongo-express:0.54

Gin FrameWork : v1.7.7

# 1. Folders
```bash
├── Ecommerce-sw
│   ├── README.md
│   ├── controllers
│   │   ├── address.go
│   │   ├── cart.go
│   │   └── controllers.go
│   ├── database
│   │   ├── cart.go
│   │   └── databasetup.go
│   ├── docker-compose.yaml
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── middleware
│   │   └── middleware.go
│   ├── models
│   │   └── models.go
│   ├── routes
│   │   └── routes.go
│   └── tokens
│       └── tokengen.go
└── docker-compose.yaml
```

# 2. Installation

~~~
docker-compose up -d
go run main.go
~~~

# 3. Use Postman
	
#### 1) 회원가입 : `POST` "localhost:8080/users/signup"
#### 2) 로그인 : `POST` "localhost:8080/users/login"
#### 3) 물건 등록: `POST` "localhost:8080/admin/addproduct"
#### 4) 등록된 물건 확인 : `GET` "localhost:8080/users/productview"
#### 5) 검색 : `GET` "localhost:8080/users/search"
