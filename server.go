package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
)

func contains(slc []string, target string) bool {
	for _, item := range slc {
		if item == target {
			return true
		}
	}
	return false
}

func healthHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/health" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(writer, "Status OK")
}

func wavHandler(writer http.ResponseWriter, request *http.Request) {
	// Parse form
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Fprintf(writer, "ParseForm error: %v", err)
		return
	}

	// Try retrieve file
	file, handler, err := request.FormFile("file")
	if err != nil {
		errString := fmt.Sprintf("Error retrieving the file %v", err)
		http.Error(writer, errString, http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Check content type
	fileType := handler.Header.Get("Content-Type")
	validTypes := []string{
		"audio/wav",
		"audio/wave",
	}
	if !contains(validTypes, fileType) {
		errString := fmt.Sprintf("Invalid filetype %v, must be .wav", fileType)
		http.Error(writer, errString, http.StatusUnsupportedMediaType)
		return
	}

	// Save file locally
	var fileBuffer bytes.Buffer
	io.Copy(&fileBuffer, file)
	tmpFileNumber := fmt.Sprintf("%v", rand.Intn(10000))
	tmpFileName := fmt.Sprintf("%v.wav", tmpFileNumber)
	tmpFile, err := os.Create(tmpFileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := tmpFile.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := tmpFile.Write(fileBuffer.Bytes()); err != nil {
		panic(err)
	}

	// Run praat script
	// https://stackoverflow.com/questions/6182369/exec-a-shell-command-in-go/7786922#7786922
	program := "praat"

	arg0 := "--run"
	arg1 := "getPitchTier.praat"
	arg2 := tmpFileNumber
	cmd := exec.Command(program, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}
	fmt.Println(string(stdout))

	// Open praat output
	outfileName := fmt.Sprintf("%v.csv", tmpFileNumber)
	outfile, err := os.Open(outfileName)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	// Return .csv file
	writer.Header().Set("Content-Disposition", "attachment; filename=tmp.csv")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Content-Type", "text/csv")
	io.Copy(writer, outfile)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/extract_pitch/wav", wavHandler)

	fmt.Printf("Starting on port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
