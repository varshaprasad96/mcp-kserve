# MCP-KServe

MCP-KServe is a Model Control Plane tool for deploying and managing machine learning models using KServe. It provides a REST API for model deployment, monitoring, and management.

## Features

- Deploy ML models using KServe
- Monitor model deployment status
- List deployed models
- Delete deployed models
- Health check endpoint

## Prerequisites

- Kubernetes cluster with KServe installed
- Go 1.21 or later
- kubectl configured with cluster access

## Installation

1. Clone the repository:
```bash
git clone https://github.com/varshaprasad96/mcp-kserve.git
cd mcp-kserve
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build -o mcp-kserve ./cmd/server
```

## Configuration

The application can be configured using the `config/config.yaml` file. The following settings are available:

```yaml
server:
  port: 8080
  host: "0.0.0.0"

k8s:
  namespace: "default"
  inCluster: true

logging:
  level: "info"
  format: "json"
```

## API Endpoints

### Deploy Model
```http
POST /api/v1/models
Content-Type: application/json

{
    "name": "model-name",
    "modelURI": "s3://bucket/model",
    "framework": "tensorflow"
}
```

### Get Model Status
```http
GET /api/v1/models/:name
```

### List Models
```http
GET /api/v1/models
```

### Delete Model
```http
DELETE /api/v1/models/:name
```

### Health Check
```http
GET /health
```

## Running the Server

```bash
./mcp-kserve
```

The server will start on the configured port (default: 8080).

## Development

To run the server in development mode:

```bash
go run cmd/server/main.go
```
