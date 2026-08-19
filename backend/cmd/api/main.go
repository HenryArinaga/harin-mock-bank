package main

import (
	"context"
	"fmt"
	"harin-mock-bank/backend/internal/db"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../.env")

	pool, err := db.NewPool(ctx)
	if err != nil {
		fmt.Printf("Could not connect to pool: %v\n", err)
		return
	}

	pool.Close()
	fmt.Printf("Success\n")

	return
}
