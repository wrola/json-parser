package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name        string
		filepath    string
		shouldError bool
	}{
		{
			name:        "valid simple JSON",
			filepath:    "../testdata/simple.json",
			shouldError: false,
		},
		{
			name:        "valid nested JSON",
			filepath:    "../testdata/nested.json",
			shouldError: false,
		},
		{
			name:        "invalid JSON",
			filepath:    "../testdata/invalid.json",
			shouldError: true,
		},
		{
			name:        "non-existent file",
			filepath:    "../testdata/does-not-exist.json",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseCommand(tt.filepath)
			if tt.shouldError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestValidateCommand(t *testing.T) {
	tests := []struct {
		name        string
		filepath    string
		shouldError bool
	}{
		{
			name:        "valid simple JSON",
			filepath:    "../testdata/simple.json",
			shouldError: false,
		},
		{
			name:        "valid nested JSON",
			filepath:    "../testdata/nested.json",
			shouldError: false,
		},
		{
			name:        "valid types JSON",
			filepath:    "../testdata/types.json",
			shouldError: false,
		},
		{
			name:        "invalid JSON syntax",
			filepath:    "../testdata/invalid.json",
			shouldError: true,
		},
		{
			name:        "non-existent file",
			filepath:    "../testdata/missing.json",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCommand(tt.filepath)
			if tt.shouldError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestFormatCommand(t *testing.T) {
	tests := []struct {
		name        string
		filepath    string
		shouldError bool
	}{
		{
			name:        "valid JSON",
			filepath:    "../testdata/simple.json",
			shouldError: false,
		},
		{
			name:        "invalid file",
			filepath:    "../testdata/invalid.json",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := FormatCommand(tt.filepath)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			buf.ReadFrom(r)

			if tt.shouldError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.shouldError && buf.Len() == 0 {
				t.Errorf("expected output but got none")
			}
		})
	}
}

func TestStatsCommand(t *testing.T) {
	tests := []struct {
		name        string
		filepath    string
		shouldError bool
	}{
		{
			name:        "simple JSON stats",
			filepath:    "../testdata/simple.json",
			shouldError: false,
		},
		{
			name:        "nested JSON stats",
			filepath:    "../testdata/nested.json",
			shouldError: false,
		},
		{
			name:        "types JSON stats",
			filepath:    "../testdata/types.json",
			shouldError: false,
		},
		{
			name:        "invalid JSON",
			filepath:    "../testdata/invalid.json",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StatsCommand(tt.filepath)
			if tt.shouldError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestCalculateStats(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.json")

	testJSON := `{
		"string": "test",
		"number": 42,
		"boolean": true,
		"null": null,
		"array": [1, 2, 3],
		"object": {"nested": "value"}
	}`

	if err := os.WriteFile(testFile, []byte(testJSON), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	err := StatsCommand(testFile)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

}

func TestStatsWithKnownValues(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "known.json")

	testJSON := `{
		"a": "string1",
		"b": "string2",
		"c": 123,
		"d": 456,
		"e": true,
		"f": false,
		"g": null,
		"h": [1, 2, 3],
		"i": {"nested": "value"}
	}`

	if err := os.WriteFile(testFile, []byte(testJSON), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := StatsCommand(testFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expectedCounts := map[string]bool{
		"Total Objects: 2":   false, // root object + nested object
		"Total Arrays:  1":   false,
		"Total Strings: 3":   false, // "string1", "string2", "value"
		"Total Numbers: 5":   false, // 123, 456, 1, 2, 3
		"Total Booleans: 2":  false, // true, false
		"Total Nulls:   1":   false,
	}

	for expected := range expectedCounts {
		if strings.Contains(output, expected) {
			expectedCounts[expected] = true
		}
	}

	for expected, found := range expectedCounts {
		if !found {
			t.Errorf("expected output to contain: %s", expected)
		}
	}
}

func BenchmarkParseCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCommand("../testdata/nested.json")
	}
}

func BenchmarkValidateCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateCommand("../testdata/nested.json")
	}
}
