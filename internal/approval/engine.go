package approval

import (
    "database/sql"
    "encoding/json"
    "errors"
    "time"
)

type RequestStatus string

const (
    StatusPending  RequestStatus = "PENDING"
    StatusApproved RequestStatus = "APPROVED"
    StatusRejected RequestStatus = "REJECTED"
    StatusExpired  RequestStatus = "EXPIRED"
)

type ApprovalRequest struct {
    ID          string        `json:"id"`
    TenantID    string        `json:"tenant_id"`
    RequesterID string        `json:"requester_id"`
    Action      string        `json:"action"`
    Payload     json.RawMessage `json:"payload"`
    Status      RequestStatus `json:"status"`
    CreatedAt   time.Time     `json:"created_at"`
}

type Engine struct {
    db *sql.DB
}

func NewEngine(db *sql.DB) *Engine {
    return &Engine{db: db}
}

// CreateRequest создает новый запрос на подтверждение операции
func (e *Engine) CreateRequest(tenantID, requesterID, action string, payload json.RawMessage) (*ApprovalRequest, error) {
    query := `
        INSERT INTO audit_logs (tenant_id, actor_id, action, payload, previous_hash, current_hash)
        VALUES ($1, $2, $3, $4, 'INIT', 'PENDING_HASH')
        RETURNING id, created_at
    `
    // Упрощенная логика интеграции с таблицей запросов/аудита для MVP
    req := &ApprovalRequest{
        TenantID:    tenantID,
        RequesterID: requesterID,
        Action:      action,
        Payload:     payload,
        Status:      StatusPending,
        CreatedAt:   time.Now(),
    }

    return req, nil
}
