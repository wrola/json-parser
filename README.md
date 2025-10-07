# JSON Parser CLI

A command-line JSON parser built in Go to learn fundamental programming concepts through practical implementation.

## What You'll Learn

This project teaches key Go concepts through building a real-world tool:

- **Basic Types**: strings, numbers, booleans, slices, maps
- **Structs**: custom data structures for JSON representation
- **Interfaces**: `Parser`, `Validator`, `JSONValue` interfaces
- **Methods**: functions attached to types
- **Error Handling**: custom error types with wrapping
- **Pointers**: efficient memory usage with pointer receivers
- **Recursion**: processing nested JSON structures
- **File I/O**: reading and writing files with proper cleanup
- **Type Switches**: handling different JSON value types
- **Package Organization**: clean separation of concerns

## Features

- **Parse**: Read and display JSON files with pretty formatting
- **Validate**: Check JSON syntax and report errors
- **Format**: Pretty-print JSON with proper indentation
- **Stats**: Display statistics about JSON structure (object count, depth, etc.)

## Installation

```bash
# Clone the repository
cd json-parser

# Build the application
go build -o json-parser

# Run it
./json-parser help
```

## Usage

### Parse a JSON file
```bash
./json-parser parse data.json
```

### Validate JSON syntax
```bash
./json-parser validate config.json
```

### Format JSON with indentation
```bash
./json-parser format data.json
```

### Display JSON statistics
```bash
./json-parser stats data.json
```

## Project Structure

```
json-parser/
├── main.go                    # CLI entry point
├── go.mod                     # Go module definition
├── cmd/
│   ├── commands.go            # Command implementations
│   └── commands_test.go       # Tests for commands
├── internal/
│   ├── parser/
│   │   ├── types.go           # JSON data structures
│   │   ├── parser.go          # Parser interface & implementation
│   │   ├── validator.go       # Validation logic
│   │   └── errors.go          # Custom error types
│   └── output/
│       └── formatter.go       # JSON output formatting
└── testdata/                  # Sample JSON files
    ├── simple.json
    ├── nested.json
    ├── types.json
    └── invalid.json
```

## Example Output

### Parse Command
```bash
$ ./json-parser parse testdata/simple.json
Successfully parsed: testdata/simple.json

{
  "name": "John Doe",
  "age": 30,
  "email": "john@example.com",
  "active": true
}
```

### Stats Command
```bash
$ ./json-parser stats testdata/nested.json
JSON Statistics for: testdata/nested.json
---------------------------
Total Objects: 7
Total Arrays:  4
Total Strings: 13
Total Numbers: 3
Total Booleans: 5
Total Nulls:   0
Max Depth:     4
```

### Validate Command
```bash
$ ./json-parser validate testdata/simple.json
✓ testdata/simple.json is valid JSON

$ ./json-parser validate testdata/invalid.json
Error: validation failed: parse error: invalid JSON syntax
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test ./... -cover
```

Run tests verbosely:
```bash
go test ./... -v
```

## Key Concepts Demonstrated

### 1. Interfaces
```go
type JSONValue interface {
    Type() string
}
```
Different types (JSONObject, JSONArray, etc.) all implement this interface.

### 2. Custom Error Types
```go
type ParseError struct {
    Line    int
    Column  int
    Message string
    Err     error
}
```
Rich error information with error wrapping support.

### 3. Recursion
The parser handles nested structures by recursively calling itself:
```go
func convertToJSONValue(raw interface{}) (JSONValue, error) {
    // ... handles nested objects and arrays recursively
}
```

### 4. Type Switches
```go
switch typedValue := value.(type) {
case *JSONObject:
    // handle object
case *JSONArray:
    // handle array
}
```

### 5. Method Receivers
```go
func (p *StandardParser) Parse(reader io.Reader) (*ParseResult, error) {
    // p is the receiver
}
```

## Learning Path

1. **Start with types.go** - Understand the data structures
2. **Read parser.go** - See how JSON is converted to Go types
3. **Study formatter.go** - Learn about recursion and formatting
4. **Explore errors.go** - Understand custom error handling
5. **Check validator.go** - See validation in action
6. **Review commands.go** - Tie everything together

## Contributing

This is a learning project. Feel free to:
- Add new commands
- Implement additional formatters
- Enhance error messages
- Add more test cases

## License

MIT License - Feel free to use this for learning!
