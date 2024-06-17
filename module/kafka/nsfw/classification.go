package nsfw

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type Result struct {
	Unsafe float64 `json:"unsafe"`
}

func ConfigureModel() (func(string) (Result, error), error) {
	classifyFunc := func(imageBase64 string) (Result, error) {
		// Run the Python script
		cmd := exec.Command("python3", "nudenet.py", imageBase64)
		output, err := cmd.Output()
		if err != nil {
			return Result{}, fmt.Errorf("Error running Python script: %v", err)
		}

		// Parse the JSON result
		var classificationResult map[string]Result
		err = json.Unmarshal(output, &classificationResult)
		if err != nil {
			return Result{}, fmt.Errorf("Error unmarshaling JSON: %v", err)
		}

		// Assuming there's only one result in the map and returning the first one
		for _, res := range classificationResult {
			return res, nil
		}

		return Result{}, fmt.Errorf("No result found")
	}
	log.Println("Finish configuring NSFW model")
	return classifyFunc, nil
}
