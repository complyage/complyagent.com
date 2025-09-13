package calls

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func AgentCall(agent AgentStruct, jsonBody string) ([]byte, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Generate API Call
	//||------------------------------------------------------------------------------------------------||

	url := agent.GenerateURL()

	//||------------------------------------------------------------------------------------------------||
	//|| Post
	//||------------------------------------------------------------------------------------------------||

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	//||------------------------------------------------------------------------------------------------||
	//|| Failed
	//||------------------------------------------------------------------------------------------------||

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	//||------------------------------------------------------------------------------------------------||
	//|| Read Body
	//||------------------------------------------------------------------------------------------------||

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Failed
	//||------------------------------------------------------------------------------------------------||

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("model call failed: status=%d, body=%s", resp.StatusCode, string(body))
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Body
	//||------------------------------------------------------------------------------------------------||

	return body, nil
}
