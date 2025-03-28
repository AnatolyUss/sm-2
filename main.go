package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sm-2/main/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env vars: %+v", err)
	}

	mode := os.Args[1]

	if mode == "--init-db" {
		fmt.Printf("Mode: %s\n", mode)
		os.Exit(0)
		//
	} else if mode == "--run-api" {
		fmt.Printf("Mode: %s\n", mode)
		port := os.Getenv("HTTP_PORT")

		if port == "" {
			port = "8080"
		}

		fmt.Printf("HTTP_PORT: %v\n", port)
		cmd.RunApi(":" + port)
	} else {
		fmt.Printf("Invalid mode: %s\n", mode)
		os.Exit(1)
	}

	//
}
