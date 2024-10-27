package user

import (
	"context"
	"fmt"
	"time"

	// "timewise/pb/timewise/pb"

	"github.com/golang-jwt/jwt/v4"

	"timewise/internal/vars"
	"timewise/timewise/pb"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
	db *pgxpool.Pool
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(db *pgxpool.Pool) *UserServiceServer {
	return &UserServiceServer{db: db}
}

func (s *UserServiceServer) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", req.Username, req.Email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{Success: true, Message: "Пользователь успешно зарегистрврован"}, nil
}

func (s *UserServiceServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	var id int32
	var username, hashedPassword string
	var createdAt time.Time

	err := s.db.QueryRow(ctx, "SELECT id, username, password_hash, created_at FROM users WHERE email = $1", req.Email).Scan(&id, &username, &hashedPassword, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	token, err := generateJWT(id, username, req.Email, createdAt)
	if err != nil {
		return nil, fmt.Errorf("could not generate token: %w", err)
	}

	return &pb.LoginUserResponse{Token: token, Message: "Успешная авторизация"}, nil
}

func generateJWT(id int32, username, email string, createdAt time.Time) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    id,
		"username":   username,
		"email":      email,
		"created_at": createdAt.Format(time.RFC3339),
		"exp":        time.Now().Add(time.Hour * 720).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(vars.JWT_SECRET))
}
