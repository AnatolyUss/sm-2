package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"sm-2/main/internal/utils"
)

func InitDb() {
	host := utils.GetEnvVar("POSTGRES_HOST", "localhost")
	port := utils.GetEnvVar("POSTGRES_PORT", "5432")
	user := utils.GetEnvVar("POSTGRES_USERNAME", "postgres")
	password := utils.GetEnvVar("POSTGRES_PASSWORD", "secret")
	dbName := utils.GetEnvVar("POSTGRES_DATABASE_NAME", "postgres")
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(), ""); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connected to database at %s\n", dbUrl)
}
