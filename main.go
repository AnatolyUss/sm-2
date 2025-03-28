package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sm-2/main/cmd"
	"sm-2/main/internal/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env vars: %+v", err)
	}

	mode := os.Args[1]

	if mode == "--init-db" {
		fmt.Printf("Mode: %s\n", mode)
		cmd.InitDb()
	} else if mode == "--run-api" {
		fmt.Printf("Mode: %s\n", mode)
		port := utils.GetEnvVar("HTTP_PORT", "8080")
		fmt.Printf("HTTP_PORT: %v\n", port)
		err := cmd.RunApi(port)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create server: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Invalid mode: %s\n", mode)
		os.Exit(1)
	}

	//
}
