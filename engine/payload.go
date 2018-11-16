package main

import (
	"encoding/json"
	"errors"
	"os"
)

// Payload is the task payload that describes a task to perform.
type Payload struct {
	// ApplicationID is the ID to which the specified Recording belongs.
	ApplicationID string `json:"applicationId"`
	// RecordingID is the ID of the Recording on which the engine should operate.
	RecordingID string `json:"recordingId"`
	// JobID is the ID of the Job to which this task belongs.
	JobID string `json:"jobId"`
	// TaskID is the ID of this engine's Task.
	TaskID string `json:"taskId"`
	// Mode is the mode for the task, such as "library-train" and "library-run".
	Mode string `json:"mode"`
	// LibraryID is the ID of the library to use.
	LibraryID string `json:"libraryId"`
	// LibraryEngineModelID is the library engine model ID to use.
	LibraryEngineModelID string `json:"libraryEngineModelId"`
	// Token is the single-use API token provided to the engine.
	// All engine requests to the Veritone API must use this token.
	Token string `json:"token"`
	// TaskPayload holds additional configuration parameters provided
	// by the media owner in the Create Job request.
	TaskPayload map[string]interface{} `json:"taskPayload"`
	// VeritoneAPIBaseURL is the base URL that should be used to access
	// the APIs.
	VeritoneAPIBaseURL string `json:"veritoneApiBaseUrl"`
}

// EnvPayload gets the Payload from the standard PAYLOAD_FILE or PAYLOAD_JSON
// environment variables.
func EnvPayload() (*Payload, error) {
	payloadFilename := os.Getenv("PAYLOAD_FILE")
	if payloadFilename != "" {
		return envPayloadFile(payloadFilename)
	}
	payloadJSON := os.Getenv("PAYLOAD_JSON")
	if payloadJSON != "" {
		return envPayloadJSON(payloadJSON)
	}
	return nil, ErrNoPayload
}

func envPayloadFile(payloadFile string) (*Payload, error) {
	f, err := os.Open(payloadFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var p Payload
	if err := json.NewDecoder(f).Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func envPayloadJSON(payloadJSON string) (*Payload, error) {
	var p Payload
	if err := json.Unmarshal([]byte(payloadJSON), &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// ErrNoPayload is returned by EnvPayload when there is no payload
// file.
var ErrNoPayload = errors.New("missing environment variables PAYLOAD_FILE or PAYLOAD_JSON")
