package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"timewise/internal/vars"

	"github.com/jackc/pgx/v4/pgxpool"
)

// @ ConnectAndInitializeDB подключается к базе данных и создает таблицы, если они отсутствуют.
func ConnectAndInitializeDB() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(vars.DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга строки подключения: %v", err)
	}
	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

// @ Auto Create tables
func createTables(db *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

	_, err := db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы users: %v", err)
	}

	log.Println("Таблицы проверены и созданы при необходимости.")
	return nil
}
