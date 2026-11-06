# TRUSTLAYER Core 1.0

> Centralized control, verification, and secure execution engine for AI agents (MCP) and administrators.

## Overview

**TRUSTLAYER Core 1.0** is an enterprise-grade, zero-trust security platform designed to govern, inspect, and authorize privileged operations executed by AI agents (Model Context Protocol) and human administrators. 

### Key Features
* **Zero Trust & Least Privilege:** Enforced via declarative OPA (Open Policy Agent) rules.
* **Secure Execution Gateway:** Envoy-based proxy with mTLS and real-time payload inspection.
* **Approval Engine:** Multi-factor "Four-Eyes Principle" verification and cryptographic approvals.
* **Multi-Tenant Isolation:** PostgreSQL Row-Level Security (RLS) for strict MSP data segregation.
* **WORM Audit Trail:** Cryptographically hashed, append-only immutable audit logging.
* **Emergency Kill Switch:** Instant system-wide or agent-specific lockdown capability.

## Project Structure

```text
trustlayer-core/
├── api/                  # gRPC protos and OpenAPI specifications
├── cmd/                  # Main server and CLI utilities
├── internal/             # Core business logic (auth, policy, approval, mcp, audit, db)
└── deployments/          # Docker and Kubernetes Helm charts
