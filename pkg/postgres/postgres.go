package postgres

import (
	"fmt"
	"log"
	"shoshilinch/internal/config"

	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/jackc/pgx/v4/stdlib"
	// _ "github.com/lib/pq"

)

// DB return database connection
func DB(cfg *config.Postgres) (*sqlx.DB, error) {
	var err error
	// log.Println(cfg)
	// psqlString := fmt.Sprintf(
	// 	`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
	// 	cfg.Host,
	// 	cfg.Port,
	// 	cfg.Username,
	// 	cfg.Password,
	// 	cfg.Database,
	// )	
	// // connConfig := pgx.ConnConfig{
	// // 	Host:     "localhost",
	// // 	Port:     5432,
	// // 	Database: "your_database_name",
	// // 	User:     "your_username",
	// // 	Password: "your_password",
	// // }
	// // connConfig := pgx.ConnConfig{psqlString,
	// // 	// Host:     cfg.Host,
	// // 	// Port:     cfg.Port,
	// // 	// Database: cfg.Database,
	// // 	// User:     cfg.User,
	// // 	// Password: cfg.Password,
	// // }
	
	// instance, err := sqlx.Connect("postgres",psqlString)
	instance, err := sqlx.Open("sqlite3", "shoshilinch.db")
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "pgx.Connect")
	}
	err = instance.PingContext(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
		return nil, err
	}

	fmt.Println("Connected to database")

	return instance, nil
}