package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	// Define the number of customers for the test
	numCustomers := 2

	// Run the main.go file with the specified number of customers
	cmd := exec.Command("go", "run", "main.go", "-numCustomers", fmt.Sprint(numCustomers))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Start the process
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run main.go: %v\nOutput:\n%s", err, out.String())
	}

	// Check the output for expected behavior
	output := out.String()
	if !contains(output, "Coffee shop closed.") {
		t.Errorf("Expected output to contain 'Coffee shop closed.', but got:\n%s", output)
	}
}

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
