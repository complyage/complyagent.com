package consumers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"errors"
	"fmt"

	"complyagent.com/processes"
	"complyagent.com/publish"
)

//||------------------------------------------------------------------------------------------------||
//|| Level 1
//||------------------------------------------------------------------------------------------------||

func LevelOneRouter(msg []byte) error {

	fmt.Println("Level1Router: Received message for Level 1 processing")

	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal
	//||------------------------------------------------------------------------------------------------||

	var av publish.AgentVerification
	if err := json.Unmarshal(msg, &av); err != nil {
		return fmt.Errorf("failed to unmarshal AgentVerification: %w", err)
	}

	fmt.Printf("Parsed AgentVerification: %+v\n", av)

	//||------------------------------------------------------------------------------------------------||
	//|| Route by Process
	//||------------------------------------------------------------------------------------------------||

	switch av.Process {
	case publish.ProcessVerifyID:
		return processes.HandleVerifyID(av)
	case publish.ProcessFacialAge:
		return processes.HandleFacialAge(av)
	default:
		return errors.New("unknown process type: " + string(av.Process))
	}
}
