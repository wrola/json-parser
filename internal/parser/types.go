package parser

type JSONValue interface {
	Type() string
}

type JSONObject struct {
	Data map[string]JSONValue
}

func (o *JSONObject) Type() string {
	return "object"
}

type JSONArray struct {
	Elements []JSONValue
}

func (a *JSONArray) Type() string {
	return "array"
}

type JSONString struct {
	Value string
}

func (s *JSONString) Type() string {
	return "string"
}

type JSONNumber struct {
	Value float64
}

func (n *JSONNumber) Type() string {
	return "number"
}

type JSONBoolean struct {
	Value bool
}

func (b *JSONBoolean) Type() string {
	return "boolean"
}

type JSONNull struct{}

func (n *JSONNull) Type() string {
	return "null"
}

type ParseResult struct {
	Value JSONValue
	Error error
}
