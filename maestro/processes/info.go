package processes

import (
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Get Age
//||------------------------------------------------------------------------------------------------||

func StepInfo(step int, status, info string) {
	fmt.Printf("[STEP %d] - %s %s\n", step, status, info)
}

func StepError(step int, status, info string) {
	fmt.Println("ERROR:", status, info)
}
