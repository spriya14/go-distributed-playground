package main

import (
	"bytes"
	"encoding/json"
	"goURL-shortie/rpc-toy/common"
	"net/http"
	"testing"
)

func BenchmarkHTTPJSON_PayloadSizes(b *testing.B) {
	b.Run("16KB", func(b *testing.B) {
		benchmarkHTTPJSON_with_provided_payloadSize(b, 16*1024)
	})
	b.Run("32KB", func(b *testing.B) {
		benchmarkHTTPJSON_with_provided_payloadSize(b, 32*1024)
	})
	b.Run("64KB", func(b *testing.B) {
		benchmarkHTTPJSON_with_provided_payloadSize(b, 64*1024)
	})
	b.Run("256KB", func(b *testing.B) {
		benchmarkHTTPJSON_with_provided_payloadSize(b, 256*1024)
	})
	b.Run("1024KB", func(b *testing.B) {
		benchmarkHTTPJSON_with_provided_payloadSize(b, 1024*1024)
	})
}

func benchmarkHTTPJSON_with_provided_payloadSize(b *testing.B, payloadSize int) {
	calculatorURL := "http://localhost:9000/add"
	buf := make([]byte, payloadSize)
	args := common.Args{A: 12, B: 89, Payload: buf}
	jsonData, err := json.Marshal(args)
	if err != nil {
		b.Fatal("Failed to marshal JSON: ", err)
	}
	b.ReportAllocs()
	b.ResetTimer() // reset timer to exclude setup time

	for i := 0; i < b.N; i++ {
		resp, err := http.Post(calculatorURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			b.Fatal("HTTP POST request failed: ", err)
		}
		if resp.StatusCode != http.StatusOK {
			b.Fatalf("Unexpected HTTP status: %s", resp.Status)
			resp.Body.Close()
		}
		var jsonResponse common.JsonResponse
		err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
		if err != nil {
			b.Fatal("Failed to decode JSON response: ", err)
			resp.Body.Close()
		}

	}

}
