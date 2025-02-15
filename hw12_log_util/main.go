package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Config struct {
	fileFlag   string
	levelFlag  string
	outputFlag string
}

func main() {
	var config Config
	if err := os.Setenv("LOG_ANALYZER_FILE", "asd.log"); err != nil {
		logrus.Fatalf("set env LOG_ANALYZER_FILE error: %s", err)
	}
	if err := os.Setenv("LOG_ANALYZER_LEVEL", "error"); err != nil {
		logrus.Fatalf("set env LOG_ANALYZER_LEVEL error: %s", err)
	}
	if err := os.Setenv("LOG_ANALYZER_OUTPUT", "b.txt"); err != nil {
		logrus.Fatalf("set env LOG_ANALYZER_OUTPUT error: %s", err)
	}

	pflag.StringVarP(&config.fileFlag, "file", "f", "", "log file path")
	pflag.StringVarP(&config.levelFlag, "level", "l", "info", "chooses level to analyze")
	pflag.StringVarP(&config.outputFlag, "output", "o", "", "path to log file write")

	pflag.Lookup("file").NoOptDefVal = os.Getenv("LOG_ANALYZER_FILE")
	pflag.Lookup("level").NoOptDefVal = os.Getenv("LOG_ANALYZER_LEVEL")
	pflag.Lookup("output").NoOptDefVal = os.Getenv("LOG_ANALYZER_OUTPUT")

	pflag.Parse()

	stats, err := analyzeLogFile(config.fileFlag, config.levelFlag)
	if err != nil {
		logrus.Fatalf("cant analyze file: %s", err)
	}

	err = outputStats(stats, config.outputFlag)
	if err != nil {
		logrus.Fatalf("output error: %s", err)
	}
}

func analyzeLogFile(filePath, logLevel string) (map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	stats := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, logLevel) {
			stats[logLevel]++
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}

	return stats, nil
}

func outputStats(stats map[string]int, outputFile string) error {
	var output *os.File
	var err error

	if outputFile != "" {
		output, err = os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("create file error: %w", err)
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	for level, count := range stats {
		fmt.Fprintf(output, "Log level: %s, Number of entries: %d\n", level, count)
	}

	return nil
}
