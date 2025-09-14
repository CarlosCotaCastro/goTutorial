package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Read Go code from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var code strings.Builder
	
	for scanner.Scan() {
		code.WriteString(scanner.Text() + "\n")
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
	
	// Create a temporary file for the Go code
	tmpDir := "/app/code"
	err := os.MkdirAll(tmpDir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temp directory: %v\n", err)
		os.Exit(1)
	}
	
	tmpFile := filepath.Join(tmpDir, "main.go")
	err = ioutil.WriteFile(tmpFile, []byte(code.String()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing temp file: %v\n", err)
		os.Exit(1)
	}
	
	// Execute the Go code with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	cmd := exec.CommandContext(ctx, "go", "run", tmpFile)
	cmd.Dir = tmpDir
	
	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Fprintf(os.Stderr, "Execution timeout exceeded\n")
		} else {
			fmt.Fprintf(os.Stderr, "Execution error: %v\n", err)
		}
		fmt.Fprintf(os.Stderr, "Output: %s\n", output)
		os.Exit(1)
	}
	
	// Print the output
	fmt.Print(string(output))
	
	// Clean up
	os.Remove(tmpFile)
}
