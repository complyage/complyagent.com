package calls

import (
	"encoding/json"
	"fmt"

	"github.com/complyage/base/verify"
)

//||------------------------------------------------------------------------------------------------||
//|| Request / Response
//||------------------------------------------------------------------------------------------------||

type DOBRequest struct {
	Text string `json:"text,omitempty"`
}

type DOBDataType struct {
	RawDOB string     `json:"raw_dob"`
	DOB    verify.DOB `json:"text"`
}

type DOBResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    DOBDataType `json:"data"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentOCR: Gets the Text from an Image
//||------------------------------------------------------------------------------------------------||

func AgentDOB(text string) (DOBResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(DOBRequest{
		Text: text,
	})
	if err != nil {
		return DOBResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Cal Vision
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := AgentCall(AGENT_DOB, string(payload))
	if err != nil {
		return DOBResponse{}, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var r DOBResponse
	if err := json.Unmarshal(respBytes, &r); err != nil {
		return DOBResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Marshal
	//||------------------------------------------------------------------------------------------------||

	return r, nil
}
