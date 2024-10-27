# Simple gRPC server

## Server
# Test gRPC


## Generate code from proto

```bash
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

## Register
```bash
grpcurl -plaintext -d '{"username":"testuser", "email":"test@example.com", "password":"password123"}' localhost:50051 pb.UserService/RegisterUser
```

Request example:
```bash
$ grpcurl -plaintext -d '{"username":"testuser", "email":"test@example.com", "password":"password123"}' localhost:50051 pb.UserService/RegisterUser
{
  "success": true,
  "message": "Пользователь успешно зарегистрврован"
}
```

## Login
```bash
grpcurl -plaintext -d '{"email":"test@example.com", "password":"password123"}' localhost:50051 pb.UserService/LoginUser
```

Request example:
```bash
$ grpcurl -plaintext -d '{"email":"test@example.com", "password":"password123"}' localhost:50051 pb.UserService/LoginUser
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJleHAiOjE3MzAyNDcwNTF9.CXIE75JRVU2iNxta29Va_tpipUl2dpIRd8h9j3zTCng",
  "message": "Успешная авторизация"
}
```

## Client
# Data for test
```js
const userData = {
    username: 'testuser',
    email: 'test@example.com',
    password: 'password123',
  };
```
## Run server
```bash
$ node index.js 
Register response: { success: true, message: 'Пользователь успешно зарегистрврован' }
Login response: {
  token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyNC0xMC0yN1QwNToxNDo1OFoiLCJlbWFpbCI6InRlc3Q
zQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyNTg3MzAwLCJ1c2VyX2lkIjo5LCJ1c2VybmFtZSI6InRlc3QzdXNlciJ9.ueTy0CSiTAlJQ1wRQyywD4Vn
Y7ECNIgvDNXXX5JL1kc',
  message: 'Успешная авторизация'
}
```
# JWT encoded format
```
{
  "created_at": "2024-10-27T05:14:58Z",
  "email": "test3@example.com",
  "exp": 1732587300,
  "user_id": 9,
  "username": "test3user"
}
```

---

## files tree
```
.
|-- Readme.md
|-- client
|   |-- README.md
|   |-- index.js
|   |-- package-lock.json
|   |-- package.json
|   `-- proto
|       `-- user.proto
`-- server
    |-- Makefile
    |-- arch.md
    |-- cmd
    |   `-- server
    |       `-- main.go
    |-- gRPC_TEST.md
    |-- go.mod
    |-- go.sum
    |-- internal
    |   |-- auth
    |   |-- database
    |   |   `-- postgresql.go
    |   |-- user
    |   |   `-- user.go
    |   `-- vars
    |       `-- vars.go
    |-- proto
    |   `-- user.proto
    |-- timewise
    |   `-- pb
    |       |-- user.pb.go
    |       `-- user_grpc.pb.go
    `-- tmp
        |-- build-errors.log
        `-- main.exe

14 directories, 20 files
```