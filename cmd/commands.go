package cmd

import (
	"fmt"
	"os"

	"github.com/wojciech/json-parser/internal/output"
	"github.com/wojciech/json-parser/internal/parser"
)

func ParseCommand(filepath string) error {
	jsonParser := parser.NewParser()

	result, err := jsonParser.ParseFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	fmt.Printf("Successfully parsed: %s\n\n", filepath)

	formatter := output.NewFormatter()
	if err := formatter.Format(result.Value, os.Stdout); err != nil {
		return fmt.Errorf("failed to format output: %w", err)
	}
	fmt.Println()

	return nil
}

func ValidateCommand(filepath string) error {
	validator := parser.NewValidator()

	if err := validator.ValidateFile(filepath); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	fmt.Printf("✓ %s is valid JSON\n", filepath)
	return nil
}

func FormatCommand(filepath string) error {
	jsonParser := parser.NewParser()
	result, err := jsonParser.ParseFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	formatter := output.NewFormatter()
	if err := formatter.Format(result.Value, os.Stdout); err != nil {
		return fmt.Errorf("failed to format output: %w", err)
	}
	fmt.Println()

	return nil
}

func StatsCommand(filepath string) error {
	jsonParser := parser.NewParser()
	result, err := jsonParser.ParseFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	stats := calculateStats(result.Value)

	fmt.Printf("JSON Statistics for: %s\n", filepath)
	fmt.Println("---------------------------")
	fmt.Printf("Total Objects: %d\n", stats.Objects)
	fmt.Printf("Total Arrays:  %d\n", stats.Arrays)
	fmt.Printf("Total Strings: %d\n", stats.Strings)
	fmt.Printf("Total Numbers: %d\n", stats.Numbers)
	fmt.Printf("Total Booleans: %d\n", stats.Booleans)
	fmt.Printf("Total Nulls:   %d\n", stats.Nulls)
	fmt.Printf("Max Depth:     %d\n", stats.MaxDepth)

	return nil
}

type JSONStats struct {
	Objects  int
	Arrays   int
	Strings  int
	Numbers  int
	Booleans int
	Nulls    int
	MaxDepth int
}

func calculateStats(value parser.JSONValue) *JSONStats {
	stats := &JSONStats{}
	calculateStatsHelper(value, stats, 0)
	return stats
}

func calculateStatsHelper(value parser.JSONValue, stats *JSONStats, depth int) {
	if depth > stats.MaxDepth {
		stats.MaxDepth = depth
	}

	switch typedValue := value.(type) {
	case *parser.JSONObject:
		stats.Objects++
		for _, childValue := range typedValue.Data {
			calculateStatsHelper(childValue, stats, depth+1)
		}
	case *parser.JSONArray:
		stats.Arrays++
		for _, element := range typedValue.Elements {
			calculateStatsHelper(element, stats, depth+1)
		}
	case *parser.JSONString:
		stats.Strings++
	case *parser.JSONNumber:
		stats.Numbers++
	case *parser.JSONBoolean:
		stats.Booleans++
	case *parser.JSONNull:
		stats.Nulls++
	}
}
