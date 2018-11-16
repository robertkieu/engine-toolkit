package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

func newRequestFromMediaChunk(client *http.Client, processURL string, msg mediaChunkMessage) (*http.Request, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	w.WriteField("chunkMimeType", msg.MIMEType)
	w.WriteField("chunkIndex", strconv.Itoa(msg.ChunkIndex))
	//w.WriteField("chunkUUID", msg.ChunkUUID)
	//w.WriteField("jobID", msg.JobID)
	//w.WriteField("taskID", msg.TaskID)
	//w.WriteField("tdoID", msg.TDOID)
	w.WriteField("startOffsetMS", strconv.Itoa(msg.StartOffsetMS))
	w.WriteField("endOffsetMS", strconv.Itoa(msg.EndOffsetMS))
	w.WriteField("width", strconv.Itoa(msg.Width))
	w.WriteField("height", strconv.Itoa(msg.Height))
	w.WriteField("libraryId", msg.TaskPayload.LibraryID)
	w.WriteField("libraryEngineModelId", msg.TaskPayload.LibraryEngineModelID)
	w.WriteField("cacheURI", msg.CacheURI)
	w.WriteField("veritoneApiBaseUrl", msg.TaskPayload.VeritoneAPIBaseURL)
	w.WriteField("token", msg.TaskPayload.Token)
	//w.WriteField("content", msg.Content)
	if msg.CacheURI != "" {
		f, err := w.CreateFormFile("chunk", "chunk.data")
		if err != nil {
			return nil, err
		}
		resp, err := client.Get(msg.CacheURI)
		if err != nil {
			return nil, errors.Wrap(err, "download chunk file from source")
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Wrapf(err, "download chunk file from source: status code %v", resp.Status)
		}
		if _, err := io.Copy(f, resp.Body); err != nil {
			return nil, errors.Wrap(err, "read chunk file from source")
		}
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, processURL, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "veritone-engine-toolkit/1.0")
	req.Header.Set("Content-Type", w.FormDataContentType())
	return req, nil
}
