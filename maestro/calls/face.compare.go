package calls

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Request / Response
//||------------------------------------------------------------------------------------------------||

type FaceCompareRequest struct {
	Image string `json:"image,omitempty"`
}

type FaceCompareDataType struct {
	Match      bool   `json:"match"`
	Confidence string `json:"confidence"`
}

type FaceCompareResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    FaceCompareDataType `json:"data"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentOCR: Gets the Text from an Image
//||------------------------------------------------------------------------------------------------||

func AgentFaceCompare(image string) (FaceCompareResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(FaceCompareRequest{
		Image: image,
	})

	if err != nil {
		return FaceCompareResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Cal Vision
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := AgentCall(AGENT_FACE_COMPARE, string(payload))
	if err != nil {
		return FaceCompareResponse{}, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var r FaceCompareResponse
	if err := json.Unmarshal(respBytes, &r); err != nil {
		return FaceCompareResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	return r, nil
}
