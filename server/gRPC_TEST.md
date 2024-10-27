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