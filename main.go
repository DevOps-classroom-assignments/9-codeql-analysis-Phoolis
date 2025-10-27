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

	// Resolve the absolute path for allowedDir and the requested file
	allowedBaseDir, err := filepath.Abs(allowedDir)
	if err != nil {
		http.Error(w, "Server misconfiguration", 500)
		return
	}
	path := filepath.Join(allowedBaseDir, filename)
	absPath, err := filepath.Abs(path)
	if err != nil {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}
	// Ensure the file is inside the allowed directory
	// Use filepath.Clean to avoid path traversal
	// Ensure that absPath starts with allowedBaseDir + separator
	if !strings.HasPrefix(absPath, allowedBaseDir+string(filepath.Separator)) {
		http.Error(w, "File not allowed", http.StatusBadRequest)
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
