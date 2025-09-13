package calls

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Request / Response
//||------------------------------------------------------------------------------------------------||

type OCRRequest struct {
	Image string `json:"image,omitempty"`
}

type OCRDataType struct {
	Text string `json:"text"`
}

type OCRResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    OCRDataType `json:"data"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentOCR: Gets the Text from an Image
//||------------------------------------------------------------------------------------------------||

func AgentOCR(image string) (OCRResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(OCRRequest{
		Image: image,
	})
	if err != nil {
		return OCRResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Cal Vision
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := AgentCall(AGENT_OCR, string(payload))
	if err != nil {
		return OCRResponse{}, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var r OCRResponse
	if err := json.Unmarshal(respBytes, &r); err != nil {
		return OCRResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	return r, nil
}
