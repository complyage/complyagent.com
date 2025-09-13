package calls

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Request / Response
//||------------------------------------------------------------------------------------------------||

type OllamaRequest struct {
	Prompt string `json:"prompt,omitempty"`
	Stream bool   `json:"stream,omitempty"`
}

type OllamaRawResponse struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	DoneReason         string `json:"done_reason"`
	Context            []int  `json:"context"`
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int64  `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int64  `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
}

type OllamaDataType struct {
	Text string `json:"text"`
}

type OllamaResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    OllamaDataType `json:"data"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentOCR: Gets the Text from an Image
//||------------------------------------------------------------------------------------------------||

func AgentOllama(prompt string) (OllamaResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(OllamaRequest{
		Prompt: prompt,
		Stream: false,
	})

	if err != nil {
		return OllamaResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Cal Vision
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := AgentCall(AGENT_FACE_DETECT, string(payload))
	if err != nil {
		return OllamaResponse{}, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var r OllamaRawResponse
	if err := json.Unmarshal(respBytes, &r); err != nil {
		return OllamaResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Altered Response
	//||------------------------------------------------------------------------------------------------||

	altered := OllamaResponse{
		Success: false,
		Message: "OK",
		Data: OllamaDataType{
			Text: r.Response,
		},
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	return altered, nil
}
