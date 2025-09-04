package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os/exec"
)

func getVideoAspectRatio(filePath string) (string, error) {
	const tolerance = 0.01
	type Streams struct {
		Streams []struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"streams"`
	}

	if filePath == "" {
		return "", errors.New("No file path provided")
	}
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)

	buff := new(bytes.Buffer)
	cmd.Stdout = buff

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output := Streams{}
	decoder := json.NewDecoder(buff)
	if err := decoder.Decode(&output); err != nil {
		return "", fmt.Errorf("Error decoding probe result: %w", err)
	}

	width := output.Streams[0].Width
	height := output.Streams[0].Height

	ratio := float64(width) / float64(height)
	if math.Abs(ratio-16.0/9.0) < tolerance {
		return "16:9", nil
	}

	if math.Abs(ratio-9.0/16.0) < tolerance {
		return "9:16", nil
	}

	return "other", nil
}
