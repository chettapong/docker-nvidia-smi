package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	xj "github.com/basgys/goxml2json"
)

const LISTEN_ADDRESS = ":9202"
const NVIDIA_SMI_PATH = "nvidia-smi"

var testMode string

func filterNumber(value string) string {
	r := regexp.MustCompile("[^0-9.]")
	return r.ReplaceAllString(value, "")
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	log.Print("Serving /json")

	var cmd *exec.Cmd
	if testMode == "1" {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		cmd = exec.Command("/bin/cat", dir+"/test.xml")
	} else {
		cmd = exec.Command(NVIDIA_SMI_PATH, "-q", "-x")
	}

	// Execute system command
	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}

	// Parse JSON
	xml := bytes.NewReader(stdout)
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	// Output
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, json.String())
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	log.Print("Serving /index")
	html := `<!doctype html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Nvidia SMI Exporter</title>
    </head>
    <body>
        <h1>Nvidia SMI Exporter</h1>
        <p><a href="/json">JSON</a></p>
    </body>
</html>`
	io.WriteString(w, html)
}

func main() {
	testMode = os.Getenv("TEST_MODE")
	if testMode == "1" {
		log.Print("Test mode is enabled")
	}

	log.Print("Nvidia SMI exporter listening on " + LISTEN_ADDRESS)
	http.HandleFunc("/", getIndex)
	http.HandleFunc("/json", getJSON)
	http.ListenAndServe(LISTEN_ADDRESS, nil)
}
