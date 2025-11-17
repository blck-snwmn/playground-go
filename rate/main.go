package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// サーバー設定
	serverAddr := "localhost:8080"
	serverURL := "http://" + serverAddr

	// サーバーを起動（別ゴルーチンで）
	go func() {
		if err := startServer(serverAddr); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// サーバーが起動するまで少し待つ
	time.Sleep(500 * time.Millisecond)

	// Rate Limiter設定
	// 1秒あたり2リクエスト、バーストサイズは1（すべてのリクエストが等間隔）
	rps := 2.0
	burst := 1
	limiter := rate.NewLimiter(rate.Limit(rps), burst)

	// 期待される最小間隔（1秒あたり2リクエスト = 500ms間隔）
	expectedMinInterval := time.Second / time.Duration(rps)

	fmt.Printf("=== Rate Limiting Test ===\n")
	fmt.Printf("Rate: %.1f req/sec\n", rps)
	fmt.Printf("Burst: %d\n", burst)
	fmt.Printf("Expected min interval: %v\n\n", expectedMinInterval)

	// リクエスト数
	requestCount := 10

	// クライアントを実行
	logs, err := runClient(serverURL, requestCount, limiter)
	if err != nil {
		log.Fatalf("Client failed: %v", err)
	}

	// レート制限の検証
	validateRateLimit(logs, expectedMinInterval, burst)
}
