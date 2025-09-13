package calls

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Request / Response
//||------------------------------------------------------------------------------------------------||

type FaceDetectRequest struct {
	Image string `json:"image,omitempty"`
}

type FaceDetectDataType struct {
	Age        int     `json:"age"`
	AgeMin     int     `json:"age_min"`
	AgeMax     int     `json:"age_max"`
	Gender     string  `json:"gender"`
	Confidence float64 `json:"confidence"`
}

type FaceDetectResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    FaceDetectDataType `json:"data"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentOCR: Gets the Text from an Image
//||------------------------------------------------------------------------------------------------||

func AgentFaceDetect(image string) (FaceDetectResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(FaceDetectRequest{
		Image: image,
	})

	if err != nil {
		return FaceDetectResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Cal Vision
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := AgentCall(AGENT_FACE_DETECT, string(payload))
	if err != nil {
		return FaceDetectResponse{}, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var r FaceDetectResponse
	if err := json.Unmarshal(respBytes, &r); err != nil {
		return FaceDetectResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	return r, nil
}
