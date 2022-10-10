package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Post(url string, data string, contentType string) (string, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("internal/util error: Post fail %w", err)
	}
	defer resp.Body.Close()

	result, _ := io.ReadAll(resp.Body)
	return string(result), nil
}
