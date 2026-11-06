package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        dbURL = "postgres://postgres:postgres@localhost:5432/trustlayer?sslmode=disable"
    }

    // Подключение к PostgreSQL
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Printf("Warning: Could not connect to PostgreSQL yet: %v", err)
    } else {
        log.Println("Successfully connected to PostgreSQL database!")
    }

    // HTTP роутер
    mux := http.NewServeMux()
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"healthy","service":"trustlayer-core"}`))
    })

    log.Printf("TRUSTLAYER Core 1.0 server is running on port %s...", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
