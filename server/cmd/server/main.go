// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net"
	"timewise/internal/database"
	"timewise/internal/user"
	"timewise/internal/vars"
	"timewise/timewise/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := database.ConnectAndInitializeDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, user.NewUserServiceServer(db))

	// Включаем отражение
	reflection.Register(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", vars.PORT))
	if err != nil {
		log.Fatalf("Не удалось начать прослушивание на порту %s: %v", vars.PORT, err)
	}

	log.Printf("GRPC сервер запущен на порту %s...\n", vars.PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
