package main

import (
	"os"
	"testing"
)

func TestAnalyzeLogFile(t *testing.T) {
	file, err := os.CreateTemp("", "test.log")
	if err != nil {
		t.Fatalf("Make log file error: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString("INFO: Test log 1\nERROR: Test log 2\nINFO: Test log 3\n")
	if err != nil {
		t.Fatalf("Writing to file error: %v", err)
	}
	file.Close()

	stats, err := analyzeLogFile(file.Name(), "ERROR")
	if err != nil {
		t.Fatalf("Analyze file error: %v", err)
	}

	if stats["ERROR"] != 1 {
		t.Errorf("Waiting for 1 log ERROR, got %d", stats["ERROR"])
	}
}

func TestOutputStats(t *testing.T) {
	outputFile, err := os.CreateTemp("", "output.txt")
	if err != nil {
		t.Fatalf("Make txt file error: %v", err)
	}
	defer os.Remove(outputFile.Name())

	stats := map[string]int{"INFO": 2, "ERROR": 1}

	err = outputStats(stats, outputFile.Name())
	if err != nil {
		t.Fatalf("Output stats error: %v", err)
	}

	content, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Writing to file error: %v", err)
	}

	expected := "Log level: INFO, Number of entries: 2\nLog level: ERROR, Number of entries: 1\n"
	if string(content) != expected {
		t.Errorf("Want:\n%s\nGot:\n%s", expected, string(content))
	}
}
