package handler

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
   )

func Handler(w http.ResponseWriter, r *http.Request) {
	// --- CORS Headers ---
	// Set the Access-Control-Allow-Origin header to '*' to allow requests from any domain.
	// This makes the API public.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specify the allowed methods.
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	// Specify the allowed headers.
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle pre-flight OPTIONS requests.
	// Browsers send an OPTIONS request to check CORS permissions before making the actual request.
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// --- End CORS Headers ---


	// Seed the random number generator to ensure different reasons are returned on each request.
	rand.Seed(time.Now().UnixNano())

	// Read the reasons from the text file.
	reasons, err := readReasons("reasons.txt")
	if err != nil {
		// If there's an error reading the file, return a 500 Internal Server Error.
		http.Error(w, "Could not read reasons file", http.StatusInternalServerError)
		return
	}

	// If there are no reasons in the file, return an error.
	if len(reasons) == 0 {
		http.Error(w, "No reasons found", http.StatusInternalServerError)
		return
	}

	// Select a random reason from the list.
	randomReason := reasons[rand.Intn(len(reasons))]

	// Set the content type header to text/plain.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Write the random reason to the response.
	fmt.Fprintln(w, randomReason)
}

// readReasons reads a file line by line and returns a slice of strings.
func readReasons(filePath string) ([]string, error) {
	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// Ensure the file is closed when the function returns.
	defer file.Close()

	var lines []string
	// Create a new scanner to read the file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line to the slice.
		lines = append(lines, scanner.Text())
	}

	// Check for any errors that occurred during scanning.
	return lines, scanner.Err()
}
