package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

// メソッドハンドラーのインターフェース
type MethodHandler interface {
	Handle(params json.RawMessage) (interface{}, error)
}

// ハンドラーマップ
var methodHandlers = map[string]MethodHandler{
	"initialize":                &InitializeHandler{},
	"notifications/initialized": &InitializedNotificationHandler{},
	"tools/list":                &ToolListHandler{},
	"tools/call":                &ToolCallHandler{},
}

func main() {
	// Setup slog with file handler
	setupLogger()
	slog.Info("Starting server")

	// メインループ
	processRequests()
}

func setupLogger() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func processRequests() {
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		slog.Info("Waiting for request")

		// リクエストの読み取り
		var req JSONRPCRequest
		if err := decoder.Decode(&req); err != nil {
			slog.Error("Failed to decode request", "error", err)
			return
		}
		slog.Info("Received request", "request", req)

		// リクエスト処理
		response, err := handleRequest(req)
		if err != nil {
			slog.Error("Error handling request", "error", err)
			response = createErrorResponse(req.ID, -32603, "Internal error", err.Error())
		}

		// レスポンス送信（レスポンスがある場合）
		if response != (JSONRPCResponse{}) {
			slog.Info("Sending response", "response", response)
			if err := encoder.Encode(response); err != nil {
				slog.Error("Failed to encode response", "error", err)
			}
		}
	}
}

func handleRequest(req JSONRPCRequest) (JSONRPCResponse, error) {
	method := req.Method
	id := req.ID

	slog.Info("Processing request", "method", method, "id", id)

	// ハンドラーの取得
	handler, exists := methodHandlers[method]
	if !exists {
		slog.Error("Unknown method", "method", method)
		return createErrorResponse(id, -32601, "Method not found", nil), nil
	}

	// ハンドラーの実行
	result, err := handler.Handle(req.Params)
	if err != nil {
		return JSONRPCResponse{}, err
	}

	// 通知の場合はレスポンスを返さない
	if method == "notifications/initialized" {
		return JSONRPCResponse{}, nil
	}

	// 正常レスポンスの作成
	return JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}, nil
}

func createErrorResponse(id int, code int, message string, data interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &Error{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
}

// InitializeHandler は initialize メソッドを処理する
type InitializeHandler struct{}

func (h *InitializeHandler) Handle(params json.RawMessage) (interface{}, error) {
	return InitResponse{
		ProtocolVersion: "2024-11-05",
		ServerInfo: struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{
			Name:    "mcp-inspector",
			Version: "0.0.1",
		},
		Capabilities: ServerCapabilities{
			Tools:     ServerCapabilitiesTools{},
			Prompts:   ServerCapabilitiesPrompts{},
			Resources: ServerCapabilitiesResources{},
		},
	}, nil
}

// InitializedNotificationHandler は notifications/initialized メソッドを処理する
type InitializedNotificationHandler struct{}

func (h *InitializedNotificationHandler) Handle(params json.RawMessage) (interface{}, error) {
	slog.Info("Received initialized notification")
	return nil, nil
}

// ToolListHandler は tools/list メソッドを処理する
type ToolListHandler struct{}

func (h *ToolListHandler) Handle(params json.RawMessage) (interface{}, error) {
	return ToolListResponse{
		Tools: []Tool{
			{
				Name:        "hello",
				Description: "hello",
				InputSchema: ToolInputSchema{
					Type: "object",
					Properties: map[string]ToolInputSchemaProperty{
						"name": {
							Type:        "string",
							Description: "name",
						},
					},
				},
			},
		},
	}, nil
}

// ToolCallHandler は tools/call メソッドを処理する
type ToolCallHandler struct{}

func (h *ToolCallHandler) Handle(params json.RawMessage) (interface{}, error) {
	var toolCallParams ToolCallParams
	if err := json.Unmarshal(params, &toolCallParams); err != nil {
		slog.Error("Failed to unmarshal tool call params", "error", err)
		return nil, err
	}
	slog.Info("Tool call params", "params", toolCallParams)

	name, args := toolCallParams.Name, toolCallParams.Arguments
	switch name {
	case "hello":
		result := fmt.Sprintf("Hello, %s!", args["name"])
		return ToolCallResponse{
			Content: []TextContent{
				{
					Typ:  "text",
					Text: result,
				},
			},
		}, nil
	default:
		slog.Error("Unknown tool", "tool", name)
		return nil, fmt.Errorf("unknown tool: %s", name)
	}
}

// JSONRPCResponse はJSON-RPCレスポンスを表す
type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

// Error はJSON-RPCエラーを表す
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// InitResponse は initialize メソッドのレスポンスを表す
type InitResponse struct {
	ProtocolVersion string `json:"protocolVersion"`
	ServerInfo      struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"serverInfo"`
	Capabilities ServerCapabilities `json:"capabilities"`
}

// ServerCapabilities はサーバーの機能を表す
type ServerCapabilities struct {
	Prompts   ServerCapabilitiesPrompts   `json:"prompts"`
	Resources ServerCapabilitiesResources `json:"resources"`
	Tools     ServerCapabilitiesTools     `json:"tools"`
}

type ServerCapabilitiesPrompts struct{}

type ServerCapabilitiesResources struct{}

type ServerCapabilitiesTools struct{}

// ToolListResponse は tools/list メソッドのレスポンスを表す
type ToolListResponse struct {
	Tools []Tool `json:"tools"`
}

// Tool はツールを表す
type Tool struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	InputSchema ToolInputSchema `json:"inputSchema"`
}

// ToolCallResponse は tools/call メソッドのレスポンスを表す
type ToolCallResponse struct {
	Content []TextContent `json:"content"`
}

// TextContent はテキストコンテンツを表す
type TextContent struct {
	Typ  string `json:"type"`
	Text string `json:"text"`
}

// ToolInputSchema はツール入力スキーマを表す
type ToolInputSchema struct {
	Type       string                             `json:"type"`
	Properties map[string]ToolInputSchemaProperty `json:"properties"`
	Required   []string                           `json:"required,omitempty"`
}

// ToolInputSchemaProperty はツール入力スキーマのプロパティを表す
type ToolInputSchemaProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// JSONRPCRequest はJSON-RPCリクエストを表す
type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

// ToolCallParams は tools/call メソッドのパラメータを表す
type ToolCallParams struct {
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

// InitParams は initialize メソッドのパラメータを表す
type InitParams struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ClientCapabilities `json:"capabilities"`
	ClientInfo      ClientInfo         `json:"clientInfo"`
}

// ClientCapabilities はクライアントの機能を表す
type ClientCapabilities struct {
	Sampling map[string]interface{} `json:"sampling,omitempty"`
	Roots    Roots                  `json:"roots"`
}

// Roots はルートを表す
type Roots struct {
	ListChanged bool `json:"listChanged"`
}

// ClientInfo はクライアント情報を表す
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
