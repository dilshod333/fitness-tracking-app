package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fitness-tracking/storage"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
)

func main() {
	connString := "postgres://postgres:dilshod@localhost:5432/fitness_tracking?sslmode=disable"
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Error("failed to connect", slog.String("error", err.Error()))
		os.Exit(1)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		logger.Error("failed to ping", slog.String("error", err.Error()))
		os.Exit(1)
	}

	queries := storage.New(db)
	ctx := context.Background()
	m := map[string]any{
		"age": 10,
		"bio": "string",
	}

	b, err := json.Marshal(m)
	if err != nil {
		logger.Error("failed to marshal JSON", slog.String("error", err.Error()))
		os.Exit(1)
	}

	err = queries.CreateUser(ctx, storage.CreateUserParams{
		Username:     sql.NullString{String: "test", Valid: true},
		Email:        sql.NullString{String: "test@gmail.com", Valid: true},
		PasswordHash: sql.NullString{String: "hashed", Valid: true},
		Profile:      pqtype.NullRawMessage{RawMessage: b, Valid: true},
	})

	if err != nil {
		logger.Error("failed to create user", slog.String("error", err.Error()))
		os.Exit(1)
	}

	users, err := queries.ListUsers(ctx)

	
	if err != nil {
		logger.Error("failed to get users", slog.String("error", err.Error()))
		os.Exit(1)
	}

	for _, v := range users{
		s := v.Profile.RawMessage
		fmt.Printf("user: %+v\n", string(s))
	}

	fmt.Println("users", users)

	err = queries.DeleteUser(ctx, 1)

	if err != nil {
		logger.Error("failed to delete user")
		os.Exit(1)
	}
}
