package parser

import "fmt"

type ParseError struct {
	Line    int
	Column  int
	Message string
	Err     error
}

func (e *ParseError) Error() string {
	if e.Line > 0 && e.Column > 0 {
		return fmt.Sprintf("parse error at line %d, column %d: %s", e.Line, e.Column, e.Message)
	}
	return fmt.Sprintf("parse error: %s", e.Message)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

type ValidationError struct {
	Path    string
	Message string
}

func (e *ValidationError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("validation error at %s: %s", e.Path, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

type FileError struct {
	Filepath string
	Op       string
	Err      error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("file error (%s) on %s: %v", e.Op, e.Filepath, e.Err)
}

func (e *FileError) Unwrap() error {
	return e.Err
}

func NewParseError(message string, err error) *ParseError {
	return &ParseError{
		Message: message,
		Err:     err,
	}
}

func NewValidationError(path, message string) *ValidationError {
	return &ValidationError{
		Path:    path,
		Message: message,
	}
}

func NewFileError(filepath, op string, err error) *FileError {
	return &FileError{
		Filepath: filepath,
		Op:       op,
		Err:      err,
	}
}
