package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"reflect"
	"testing"
)

// テスト用のセットアップとクリーンアップを行うヘルパー関数
func setupTest(t *testing.T) func() {
	// 元のロガーを保存
	originalHandler := slog.Default().Handler()

	// テスト中はログを無効化
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))

	// クリーンアップ関数を返す
	return func() {
		// 元のロガーに戻す
		slog.SetDefault(slog.New(originalHandler))
	}
}

func TestInitializeHandler_Handle(t *testing.T) {
	// テスト用のセットアップとクリーンアップ
	cleanup := setupTest(t)
	defer cleanup()

	handler := &InitializeHandler{}
	params := json.RawMessage(`{
		"protocolVersion": "2024-11-05",
		"capabilities": {
			"roots": {
				"listChanged": true
			}
		},
		"clientInfo": {
			"name": "test-client",
			"version": "1.0.0"
		}
	}`)

	result, err := handler.Handle(params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	initResponse, ok := result.(InitResponse)
	if !ok {
		t.Fatalf("Expected result to be InitResponse, got %T", result)
	}

	if initResponse.ProtocolVersion != "2024-11-05" {
		t.Errorf("Expected protocolVersion to be 2024-11-05, got %s", initResponse.ProtocolVersion)
	}

	if initResponse.ServerInfo.Name != "mcp-inspector" {
		t.Errorf("Expected serverInfo.name to be mcp-inspector, got %s", initResponse.ServerInfo.Name)
	}

	if initResponse.ServerInfo.Version != "0.0.1" {
		t.Errorf("Expected serverInfo.version to be 0.0.1, got %s", initResponse.ServerInfo.Version)
	}
}

func TestInitializedNotificationHandler_Handle(t *testing.T) {
	// テスト用のセットアップとクリーンアップ
	cleanup := setupTest(t)
	defer cleanup()

	handler := &InitializedNotificationHandler{}
	params := json.RawMessage(`{}`)

	result, err := handler.Handle(params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected result to be nil, got %v", result)
	}
}

func TestToolListHandler_Handle(t *testing.T) {
	// テスト用のセットアップとクリーンアップ
	cleanup := setupTest(t)
	defer cleanup()

	handler := &ToolListHandler{}
	params := json.RawMessage(`{}`)

	result, err := handler.Handle(params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	toolListResponse, ok := result.(ToolListResponse)
	if !ok {
		t.Fatalf("Expected result to be ToolListResponse, got %T", result)
	}

	if len(toolListResponse.Tools) != 1 {
		t.Fatalf("Expected 1 tool, got %d", len(toolListResponse.Tools))
	}

	tool := toolListResponse.Tools[0]
	if tool.Name != "hello" {
		t.Errorf("Expected tool name to be hello, got %s", tool.Name)
	}

	if tool.Description != "hello" {
		t.Errorf("Expected tool description to be hello, got %s", tool.Description)
	}

	if tool.InputSchema.Type != "object" {
		t.Errorf("Expected inputSchema.type to be object, got %s", tool.InputSchema.Type)
	}

	if len(tool.InputSchema.Properties) != 1 {
		t.Fatalf("Expected 1 property, got %d", len(tool.InputSchema.Properties))
	}

	nameProp, exists := tool.InputSchema.Properties["name"]
	if !exists {
		t.Fatalf("Expected name property to exist")
	}

	if nameProp.Type != "string" {
		t.Errorf("Expected name property type to be string, got %s", nameProp.Type)
	}

	if nameProp.Description != "name" {
		t.Errorf("Expected name property description to be name, got %s", nameProp.Description)
	}
}

func TestToolCallHandler_Handle(t *testing.T) {
	// テスト用のセットアップとクリーンアップ
	cleanup := setupTest(t)
	defer cleanup()

	handler := &ToolCallHandler{}

	// テストケース1: 正常系 - helloツール
	t.Run("hello tool with valid params", func(t *testing.T) {
		params := json.RawMessage(`{
			"name": "hello",
			"arguments": {
				"name": "test-user"
			}
		}`)

		result, err := handler.Handle(params)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		toolCallResponse, ok := result.(ToolCallResponse)
		if !ok {
			t.Fatalf("Expected result to be ToolCallResponse, got %T", result)
		}

		if len(toolCallResponse.Content) != 1 {
			t.Fatalf("Expected 1 content, got %d", len(toolCallResponse.Content))
		}

		content := toolCallResponse.Content[0]
		if content.Typ != "text" {
			t.Errorf("Expected content type to be text, got %s", content.Typ)
		}

		expectedText := "Hello, test-user!"
		if content.Text != expectedText {
			t.Errorf("Expected content text to be %q, got %q", expectedText, content.Text)
		}
	})

	// テストケース2: 異常系 - 不明なツール
	t.Run("unknown tool", func(t *testing.T) {
		params := json.RawMessage(`{
			"name": "unknown-tool",
			"arguments": {}
		}`)

		_, err := handler.Handle(params)
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}

		expectedErrMsg := "unknown tool: unknown-tool"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message %q, got %q", expectedErrMsg, err.Error())
		}
	})

	// テストケース3: 異常系 - 不正なJSON
	t.Run("invalid json", func(t *testing.T) {
		params := json.RawMessage(`{invalid json`)

		_, err := handler.Handle(params)
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})
}

func TestHandleRequest(t *testing.T) {
	// テスト用のセットアップとクリーンアップ
	cleanup := setupTest(t)
	defer cleanup()

	// テストケース1: 正常系 - initialize
	t.Run("initialize method", func(t *testing.T) {
		req := JSONRPCRequest{
			JSONRPC: "2.0",
			ID:      1,
			Method:  "initialize",
			Params:  json.RawMessage(`{}`),
		}

		resp, err := handleRequest(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.JSONRPC != "2.0" {
			t.Errorf("Expected jsonrpc to be 2.0, got %s", resp.JSONRPC)
		}

		if resp.ID != 1 {
			t.Errorf("Expected id to be 1, got %d", resp.ID)
		}

		if resp.Error != nil {
			t.Errorf("Expected error to be nil, got %v", resp.Error)
		}

		if resp.Result == nil {
			t.Fatalf("Expected result not to be nil")
		}
	})

	// テストケース2: 正常系 - notifications/initialized (レスポンスなし)
	t.Run("initialized notification", func(t *testing.T) {
		req := JSONRPCRequest{
			JSONRPC: "2.0",
			ID:      2,
			Method:  "notifications/initialized",
			Params:  json.RawMessage(`{}`),
		}

		resp, err := handleRequest(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// 通知の場合は空のレスポンスが返る
		emptyResp := JSONRPCResponse{}
		if !reflect.DeepEqual(resp, emptyResp) {
			t.Errorf("Expected empty response, got %v", resp)
		}
	})

	// テストケース3: 異常系 - 不明なメソッド
	t.Run("unknown method", func(t *testing.T) {
		req := JSONRPCRequest{
			JSONRPC: "2.0",
			ID:      3,
			Method:  "unknown-method",
			Params:  json.RawMessage(`{}`),
		}

		resp, err := handleRequest(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.JSONRPC != "2.0" {
			t.Errorf("Expected jsonrpc to be 2.0, got %s", resp.JSONRPC)
		}

		if resp.ID != 3 {
			t.Errorf("Expected id to be 3, got %d", resp.ID)
		}

		if resp.Error == nil {
			t.Fatalf("Expected error not to be nil")
		}

		if resp.Error.Code != -32601 {
			t.Errorf("Expected error code to be -32601, got %d", resp.Error.Code)
		}

		if resp.Error.Message != "Method not found" {
			t.Errorf("Expected error message to be 'Method not found', got %q", resp.Error.Message)
		}
	})
}

func TestCreateErrorResponse(t *testing.T) {
	// テスト用のセットアップとクリーンアップ
	cleanup := setupTest(t)
	defer cleanup()

	resp := createErrorResponse(1, 100, "Test error", "error data")

	if resp.JSONRPC != "2.0" {
		t.Errorf("Expected jsonrpc to be 2.0, got %s", resp.JSONRPC)
	}

	if resp.ID != 1 {
		t.Errorf("Expected id to be 1, got %d", resp.ID)
	}

	if resp.Error == nil {
		t.Fatalf("Expected error not to be nil")
	}

	if resp.Error.Code != 100 {
		t.Errorf("Expected error code to be 100, got %d", resp.Error.Code)
	}

	if resp.Error.Message != "Test error" {
		t.Errorf("Expected error message to be 'Test error', got %q", resp.Error.Message)
	}

	if resp.Error.Data != "error data" {
		t.Errorf("Expected error data to be 'error data', got %v", resp.Error.Data)
	}
}
