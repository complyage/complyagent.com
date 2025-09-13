package consumers

import (
	"os"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

func init() {
	//||------------------------------------------------------------------------------------------------||
	//|| Level 1
	//||------------------------------------------------------------------------------------------------||
	if os.Getenv("AGENT_SUPPORTS_LEVEL_ONE") == "true" {
		errL1 := app.QueueRabbit["agent"].ConsumeQueue("AgentLevelOne", func(msg []byte) {
			if err := LevelOneRouter(msg); err != nil {
				app.Log.Error("LevelOneRouter error: %v\n", err)
			}
		})
		if errL1 != nil {
			app.Log.Error("Failed to start Level One Consumer: %v", errL1)
			os.Exit(1)
		}
		app.Log.Info("Listening for Level One messages on 'AgentLevelOne' queue\n")
	} else {
		app.Log.Info("AGENT_SUPPORTS_LEVEL_ONE is not set to true, skipping Level One Consumer\n")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Level 2
	//||------------------------------------------------------------------------------------------------||
	if os.Getenv("AGENT_SUPPORTS_LEVEL_TWO") == "true" {
		errL2 := app.QueueRabbit["agent"].ConsumeQueue("AgentLevelTwo", func(msg []byte) {
			if err := LevelTwoRouter(msg); err != nil {
				app.Log.Error("LevelTwoRouter Error: %v\n", err)
			}
		})
		if errL2 != nil {
			app.Log.Error("Failed to start Level Two consumer: %v", errL2)
			os.Exit(1)
		}
		app.Log.Info("Listening for Level Two messages on 'AgentLevelTwo' queue\n")
	} else {
		app.Log.Info("AGENT_SUPPORTS_LEVEL_TWO is not set to true, skipping Level Two Consumer\n")
	}
}
