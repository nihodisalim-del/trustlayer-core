package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "github.com/nihodisalim-del/trustlayer-core/internal/approval"
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

    // Инициализация движка подтверждений
    approvalEngine := approval.NewEngine(db)

    // HTTP роутер
    mux := http.NewServeMux()

    // Проверка здоровья
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"healthy","service":"trustlayer-core"}`))
    })

    // Эндпоинт для создания запроса на одобрение операции
    mux.HandleFunc("/api/v1/approvals", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var req struct {
            TenantID    string          `json:"tenant_id"`
            RequesterID string          `json:"requester_id"`
            Action      string          `json:"action"`
            Payload     json.RawMessage `json:"payload"`
        }

        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        approvalReq, err := approvalEngine.CreateRequest(req.TenantID, req.RequesterID, req.Action, req.Payload)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w.Encode(approvalReq))
    })

    log.Printf("TRUSTLAYER Core 1.0 server is running on port %s...", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
