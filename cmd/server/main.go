package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MCPRequest struct {
	ID     int64           `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int64       `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		var req MCPRequest
		if err := json.Unmarshal(scanner.Bytes(), &req); err != nil {
			continue
		}

		fmt.Fprintf(os.Stderr, "Received method: %s\n", req.Method)

		switch req.Method {
		case "tools":
			resp := MCPResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result: []map[string]interface{}{
					{
						"name":        "list_models",
						"description": "List all deployed models",
						"parameters":  map[string]interface{}{},
					},
					{
						"name":        "deploy_model",
						"description": "Deploy a new model",
						"parameters": map[string]interface{}{
							"name":      "string",
							"modelURI":  "string",
							"framework": "string",
						},
					},
				},
				Error: nil,
			}
			json.NewEncoder(os.Stdout).Encode(resp)
		case "list_models":
			resp := MCPResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result:  []string{"model1", "model2"},
				Error:   nil,
			}
			json.NewEncoder(os.Stdout).Encode(resp)
		case "deploy_model":
			resp := MCPResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result:  "Model deployment started",
				Error:   nil,
			}
			json.NewEncoder(os.Stdout).Encode(resp)
		case "initialize":
			resp := MCPResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result:  map[string]interface{}{"capabilities": map[string]interface{}{}},
				Error:   nil,
			}
			json.NewEncoder(os.Stdout).Encode(resp)
		case "shutdown", "exit":
			resp := MCPResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result:  nil,
				Error:   nil,
			}
			json.NewEncoder(os.Stdout).Encode(resp)
		default:
			resp := MCPResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Error: map[string]interface{}{
					"code":    -32601,
					"message": "Unknown method",
				},
			}
			json.NewEncoder(os.Stdout).Encode(resp)
		}
	}
}
