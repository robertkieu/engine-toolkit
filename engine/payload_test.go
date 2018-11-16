package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestPayloadJSON(t *testing.T) {
	is := is.New(t)

	is.NoErr(os.Setenv("PAYLOAD_FILE", ""))
	is.NoErr(os.Setenv("PAYLOAD_JSON", ""))

	p := Payload{
		ApplicationID:        "applicationId",
		RecordingID:          "recordingId",
		JobID:                "jobId",
		TaskID:               "taskId",
		Token:                "token",
		Mode:                 "mode",
		LibraryID:            "libraryId",
		LibraryEngineModelID: "libraryEngineModelId",
		VeritoneAPIBaseURL:   "https://api.veritone.com",
	}
	b, err := json.Marshal(p)
	is.NoErr(err)
	os.Setenv("PAYLOAD_JSON", string(b))

	payload, err := EnvPayload()
	is.NoErr(err)
	is.Equal(payload.ApplicationID, "applicationId")
	is.Equal(payload.RecordingID, "recordingId")
	is.Equal(payload.JobID, "jobId")
	is.Equal(payload.TaskID, "taskId")
	is.Equal(payload.Token, "token")
	is.Equal(payload.Mode, "mode")
	is.Equal(payload.LibraryID, "libraryId")
	is.Equal(payload.LibraryEngineModelID, "libraryEngineModelId")
	is.Equal(payload.VeritoneAPIBaseURL, "https://api.veritone.com")

}

func TestPayloadFile(t *testing.T) {
	is := is.New(t)

	is.NoErr(os.Setenv("PAYLOAD_FILE", ""))
	is.NoErr(os.Setenv("PAYLOAD_JSON", ""))

	payload, err := EnvPayload()
	is.Equal(err, ErrNoPayload)

	is.NoErr(
		os.Setenv("PAYLOAD_FILE", `testdata/payload.json`),
	)

	payload, err = EnvPayload()
	is.NoErr(err)
	is.Equal(payload.ApplicationID, "applicationId")
	is.Equal(payload.RecordingID, "recordingId")
	is.Equal(payload.JobID, "jobId")
	is.Equal(payload.TaskID, "taskId")
	is.Equal(payload.Token, "token")
	is.Equal(payload.Mode, "mode")
	is.Equal(payload.LibraryID, "libraryId")
	is.Equal(payload.LibraryEngineModelID, "libraryEngineModelId")
	is.Equal(payload.VeritoneAPIBaseURL, "https://api.veritone.com")
}

func TestNoPayload(t *testing.T) {
	is := is.New(t)

	is.NoErr(os.Setenv("PAYLOAD_FILE", ""))
	is.NoErr(os.Setenv("PAYLOAD_JSON", ""))

	_, err := EnvPayload()
	is.Equal(err, ErrNoPayload)
}
