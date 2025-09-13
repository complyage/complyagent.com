package calls

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Request / Response
//||------------------------------------------------------------------------------------------------||

type NSFWDetectRequest struct {
	Image string `json:"image,omitempty"`
}

type NSFWDataType struct {
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Confidence string `json:"confidence"`
}

type NSFWDetectResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    NSFWDataType `json:"data"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentOCR: Gets the Text from an Image
//||------------------------------------------------------------------------------------------------||

func AgentNSFWDetect(image string) (NSFWDetectResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(NSFWDetectRequest{
		Image: image,
	})

	if err != nil {
		return NSFWDetectResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Cal Vision
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := AgentCall(AGENT_FACE_DETECT, string(payload))
	if err != nil {
		return NSFWDetectResponse{}, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var r NSFWDetectResponse
	if err := json.Unmarshal(respBytes, &r); err != nil {
		return NSFWDetectResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	return r, nil
}
