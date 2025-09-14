package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// CodeExecutionService handles secure Go code execution
type CodeExecutionService struct {
	dockerImage string
}

// NewCodeExecutionService creates a new code execution service
func NewCodeExecutionService() *CodeExecutionService {
	return &CodeExecutionService{
		dockerImage: "go-executor:latest",
	}
}

// ExecuteCode runs Go code in a secure Docker container
func (s *CodeExecutionService) ExecuteCode(code string) (*CodeExecutionResponse, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Run the Docker container
	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "-i", s.dockerImage)
	
	// Set up stdin to send the code
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %v", err)
	}

	// Set up stdout and stderr capture
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %v", err)
	}

	// Send the code to stdin
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, code)
	}()

	// Wait for completion
	err = cmd.Wait()

	// Prepare response
	response := &CodeExecutionResponse{
		Output: stdout.String(),
	}

	// Handle errors
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			response.Error = "Execution timeout exceeded"
		} else {
			response.Error = fmt.Sprintf("Execution error: %v", err)
		}
		response.Output += stderr.String()
	}

	return response, nil
}

// ExecuteCodeFallback provides a fallback execution method for development
func (s *CodeExecutionService) ExecuteCodeFallback(code string) (*CodeExecutionResponse, error) {
	// For development, we'll use a simple approach
	// In production, this should always use Docker
	
	// Create a temporary file
	tmpFile := "/tmp/go_code_" + fmt.Sprintf("%d", time.Now().UnixNano()) + ".go"
	
	// Write code to file
	if err := os.WriteFile(tmpFile, []byte(code), 0644); err != nil {
		return nil, fmt.Errorf("failed to write temp file: %v", err)
	}
	defer os.Remove(tmpFile)

	// Execute with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", tmpFile)
	output, err := cmd.CombinedOutput()

	response := &CodeExecutionResponse{
		Output: string(output),
	}

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			response.Error = "Execution timeout exceeded"
		} else {
			response.Error = fmt.Sprintf("Execution error: %v", err)
		}
	}

	return response, nil
}
