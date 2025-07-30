package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// ✅ Initialize database (creates accounts table if not exists)
	if err := store.Init(); err != nil {
		log.Fatal("failed to initialize database:", err)
	}

	server := NewApiServer(":3000", store)
	server.Run()
}
