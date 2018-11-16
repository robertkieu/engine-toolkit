package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/readyz", handleReady)
	http.HandleFunc("/process", handleProcess)
	if err := http.ListenAndServe("0.0.0.0:8888", nil); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleProcess(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.String())
	if !strings.HasPrefix(r.FormValue("chunkMimeType"), "image/") {
		http.Error(w, "unsupported mime type", http.StatusBadRequest)
		return
	}
	f, _, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	image, typ, err := image.Decode(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	size := image.Bounds().Size()
	log.Printf("image is %q\n", typ)
	log.Printf("image is %dx%d\n", size.X, size.Y)
	output := engineOutput{
		Series: []seriesObject{
			{
				Object: object{
					Label: fmt.Sprintf("%dx%d", size.X, size.Y),
				},
			},
		},
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type engineOutput struct {
	Series []seriesObject `json:"series"`
}

type seriesObject struct {
	Start  int    `json:"startTimeMs"`
	End    int    `json:"stopTimeMs"`
	Object object `json:"object"`
}

type object struct {
	Label      string `json:"label"`
	ObjectType string `json:"type"`
}
