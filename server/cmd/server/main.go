// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"timewise/internal/database"
	"timewise/internal/user"
	"timewise/internal/vars"
	"timewise/timewise/pb"

	"google.golang.org/grpc"
)

func main() {
	db, err := database.ConnectAndInitializeDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, user.NewUserServiceServer(db))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", vars.PORT))
	if err != nil {
		log.Fatalf("Не удалось начать прослушивание на порту %s: %v", vars.PORT, err)
	}

	//@ Run server
	runServer(server, lis)
}

func runServer(server *grpc.Server, lis net.Listener) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("GRPC сервер запущен на порту %s...\n", vars.PORT)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	<-stop
	log.Println("Получен сигнал завершения, выключение сервера...")
	shutdownServer(server)
}

func shutdownServer(server *grpc.Server) {
	server.GracefulStop()
	log.Println("Сервер выключен")
}
