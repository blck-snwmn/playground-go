package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type RequestLog struct {
	Index      int
	RequestAt  time.Time
	ResponseAt time.Time
	Duration   time.Duration
	Body       string
	Interval   time.Duration // 前回のリクエストからの間隔
}

func runClient(serverURL string, requestCount int, limiter *rate.Limiter) ([]RequestLog, error) {
	var logs []RequestLog
	var lastRequestTime time.Time

	ctx := context.Background()

	for i := 0; i < requestCount; i++ {
		// rate limiterで待機
		if err := limiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("limiter wait failed: %w", err)
		}

		requestTime := time.Now()

		// 前回のリクエストからの間隔を計算
		var interval time.Duration
		if !lastRequestTime.IsZero() {
			interval = requestTime.Sub(lastRequestTime)
		}
		lastRequestTime = requestTime

		// HTTPリクエストを送信
		resp, err := http.Get(serverURL)
		if err != nil {
			log.Printf("Request %d failed: %v\n", i, err)
			continue
		}

		responseTime := time.Now()
		duration := responseTime.Sub(requestTime)

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		reqLog := RequestLog{
			Index:      i,
			RequestAt:  requestTime,
			ResponseAt: responseTime,
			Duration:   duration,
			Body:       string(body),
			Interval:   interval,
		}

		logs = append(logs, reqLog)

		log.Printf("Request %d: interval=%v, response_time=%v\n", i, interval, duration)
	}

	return logs, nil
}

// リクエスト間隔が設定されたレート制限を守っているか検証
func validateRateLimit(logs []RequestLog, expectedMinInterval time.Duration, burstSize int) bool {
	allValid := true
	fmt.Println("\n=== Rate Limit Validation ===")
	fmt.Printf("Expected minimum interval: %v\n", expectedMinInterval)
	fmt.Printf("Burst size: %d (first %d requests can be sent immediately)\n", burstSize, burstSize)
	fmt.Println()

	for i, log := range logs {
		if i == 0 {
			// 最初のリクエストはスキップ
			continue
		}

		// バーストサイズまでのリクエストは間隔チェックをスキップ
		if i < burstSize {
			fmt.Printf("Request %d: interval=%v (burst, not checked)\n", log.Index, log.Interval)
			continue
		}

		// 許容誤差を考慮（1msの誤差を許容）
		tolerance := 1 * time.Millisecond
		isValid := log.Interval >= (expectedMinInterval - tolerance)

		status := "✓ OK"
		if !isValid {
			status = "✗ VIOLATION"
			allValid = false
		}

		fmt.Printf("Request %d: interval=%v %s\n", log.Index, log.Interval, status)
	}

	fmt.Println()
	if allValid {
		fmt.Println("✓ All requests after burst respected the rate limit!")
	} else {
		fmt.Println("✗ Some requests violated the rate limit!")
	}
	fmt.Println()

	return allValid
}
