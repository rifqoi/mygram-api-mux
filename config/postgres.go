package config

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectPostgres() (*pgxpool.Pool, error) {
	dbHost := GetEnv("DB_HOST")
	dbUser := GetEnv("DB_USER")
	dbPass := GetEnv("DB_PASS")
	dbName := GetEnv("DB_NAME")
	dbPort := GetEnv("DB_PORT")

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(dbUser, dbPass),
		Host:   fmt.Sprintf("%s:%s", dbHost, dbPort),
		Path:   dbName,
	}

	fmt.Println(dsn.String())
	pool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Cannot ping the database", err)
	}

	return pool, nil
}
