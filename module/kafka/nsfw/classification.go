package nsfw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type DetectionResult struct {
	Class string  `json:"class"`
	Score float64 `json:"score"`
	Box   []int   `json:"box"`
}

// ConfigureModel returns a function that takes a base64-encoded string
// and returns the classification result.
func ConfigureModel() (func(string) (*[]DetectionResult, error), error) {
	classifyFunc := func(imageBase64 string) (*[]DetectionResult, error) {
		// Use the absolute path to the renamed Python script
		scriptPath := "./kafka/nsfw/classify_nsfw.py"

		// Run the Python script using the global Python installation
		cmd := exec.Command("python3", scriptPath, imageBase64)

		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		err := cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("Error running Python script: %v, stderr: %s", err, errb.String())
		}

		// Parse the JSON result
		var classificationResults []DetectionResult
		err = json.Unmarshal(outb.Bytes(), &classificationResults)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshaling JSON: %v", err)
		}

		return &classificationResults, nil
	}

	return classifyFunc, nil
}
