package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	fmt.Println(workingDir)
	// Define the marker and new content
	fixtureFile := "../../testdata/fixtures/marker_example.md"
	marker := "commands"
	newContent := "new text whatever"

	// Simulate command-line arguments
	runArgs := []string{"testprogram", "-file", fixtureFile, "-marker", marker, newContent}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	err = run(runArgs)
	if err != nil {
		t.Fatalf("Error running main function: %v", err)
	}

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the captured output
	out, _ := ioutil.ReadAll(r)
	capturedOutput := string(out)

	expectedContent := `# Marker Example

This file contains just an example for testing with a "commands" marker.

<!--commands-->
new text whatever
<!--/commands-->`

	if strings.TrimSpace(string(capturedOutput)) != strings.TrimSpace(expectedContent) {
		t.Fatalf("Content mismatch. Expected:\n%s\nGot:\n%s\n", expectedContent, string(capturedOutput))
	}
}
