package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Parser interface {
	Parse(reader io.Reader) (*ParseResult, error)
	ParseFile(filepath string) (*ParseResult, error)
}

type StandardParser struct{}

func NewParser() Parser {
	return &StandardParser{}
}

func (p *StandardParser) Parse(reader io.Reader) (*ParseResult, error) {
	var raw interface{}
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(&raw); err != nil {
		return &ParseResult{
			Value: nil,
			Error: fmt.Errorf("failed to decode JSON: %w", err),
		}, err
	}

	value, err := convertToJSONValue(raw)
	if err != nil {
		return &ParseResult{
			Value: nil,
			Error: err,
		}, err
	}

	return &ParseResult{
		Value: value,
		Error: nil,
	}, nil
}

func (p *StandardParser) ParseFile(filepath string) (*ParseResult, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return &ParseResult{
			Value: nil,
			Error: fmt.Errorf("failed to open file: %w", err),
		}, err
	}
	defer file.Close()

	return p.Parse(file)
}

func convertToJSONValue(raw interface{}) (JSONValue, error) {
	if raw == nil {
		return &JSONNull{}, nil
	}

	switch typedValue := raw.(type) {
	case map[string]interface{}:
		jsonObject := &JSONObject{Data: make(map[string]JSONValue)}
		for key, value := range typedValue {
			converted, err := convertToJSONValue(value)
			if err != nil {
				return nil, err
			}
			jsonObject.Data[key] = converted
		}
		return jsonObject, nil

	case []interface{}:
		jsonArray := &JSONArray{Elements: make([]JSONValue, 0, len(typedValue))}
		for _, value := range typedValue {
			converted, err := convertToJSONValue(value)
			if err != nil {
				return nil, err
			}
			jsonArray.Elements = append(jsonArray.Elements, converted)
		}
		return jsonArray, nil

	case string:
		return &JSONString{Value: typedValue}, nil

	case float64:
		return &JSONNumber{Value: typedValue}, nil

	case bool:
		return &JSONBoolean{Value: typedValue}, nil

	default:
		return nil, fmt.Errorf("unsupported JSON value type: %T", typedValue)
	}
}
