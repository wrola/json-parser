package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Validator interface {
	Validate(reader io.Reader) error
	ValidateFile(filepath string) error
}

type SyntaxValidator struct{}

func NewValidator() Validator {
	return &SyntaxValidator{}
}

func (v *SyntaxValidator) Validate(reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	var raw interface{}
	if err := decoder.Decode(&raw); err != nil {
		return NewParseError("invalid JSON syntax", err)
	}

	if decoder.More() {
		return NewParseError("unexpected data after JSON value", nil)
	}

	return nil
}

func (v *SyntaxValidator) ValidateFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return NewFileError(filepath, "open", err)
	}
	defer file.Close()

	return v.Validate(file)
}

type StructureValidator struct {
	MaxDepth int
	MaxSize  int64
}

func NewStructureValidator(maxDepth int, maxSize int64) Validator {
	return &StructureValidator{
		MaxDepth: maxDepth,
		MaxSize:  maxSize,
	}
}

func (v *StructureValidator) Validate(reader io.Reader) error {
	syntaxValidator := NewValidator()
	if err := syntaxValidator.Validate(reader); err != nil {
		return err
	}

	return nil
}

func (v *StructureValidator) ValidateFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return NewFileError(filepath, "open", err)
	}
	defer file.Close()

	if v.MaxSize > 0 {
		info, err := file.Stat()
		if err != nil {
			return NewFileError(filepath, "stat", err)
		}
		if info.Size() > v.MaxSize {
			return NewValidationError("", fmt.Sprintf("file size %d exceeds maximum %d", info.Size(), v.MaxSize))
		}
	}

	return v.Validate(file)
}

func ValidateDepth(value JSONValue, maxDepth int) error {
	return validateDepthHelper(value, 0, maxDepth)
}

func validateDepthHelper(value JSONValue, currentDepth, maxDepth int) error {
	if currentDepth > maxDepth {
		return NewValidationError("", fmt.Sprintf("nesting depth exceeds maximum of %d", maxDepth))
	}

	switch typedValue := value.(type) {
	case *JSONObject:
		for key, childValue := range typedValue.Data {
			if err := validateDepthHelper(childValue, currentDepth+1, maxDepth); err != nil {
				return NewValidationError(key, err.Error())
			}
		}
	case *JSONArray:
		for index, childValue := range typedValue.Elements {
			if err := validateDepthHelper(childValue, currentDepth+1, maxDepth); err != nil {
				return NewValidationError(fmt.Sprintf("[%d]", index), err.Error())
			}
		}
	}

	return nil
}
