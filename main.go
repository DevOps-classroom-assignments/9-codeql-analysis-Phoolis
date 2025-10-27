package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)

const allowedDir = "./safe-files"

func main() {
	http.HandleFunc("/readfile", readFileHandler)
	http.HandleFunc("/exec", execHandler)

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func readFileHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
			http.Error(w, "missing file parameter", http.StatusBadRequest)
			return
	}

	allowedBaseDir, err := filepath.Abs(allowedDir)
	if err != nil {
		http.Error(w, "Server misconfiguration", 500)
		return
	}

	requestedPath := filepath.Join(allowedBaseDir, filename)
	absPath, err := filepath.Abs(requestedPath)
	if err != nil {
		http.Error(w, "Invalid file path", 400)
		return
	}

	// Use relative filepath
	rel, err := filepath.Rel(allowedBaseDir, absPath)
	if err != nil || strings.HasPrefix(rel, "..") || rel == "." {
		http.Error(w, "File not allowed", 403)
		return
	}

	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	w.Write(data)
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")

	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		http.Error(w, "Command failed", 500)
		return
	}
	w.Write(out)
}
