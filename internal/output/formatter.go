package output

import (
	"fmt"
	"io"
	"strings"

	"github.com/wojciech/json-parser/internal/parser"
)

const (
	ColorReset   = "\033[0m"
	ColorKey     = "\033[36m" // Cyan for keys
	ColorString  = "\033[32m" // Green for strings
	ColorNumber  = "\033[33m" // Yellow for numbers
	ColorBoolean = "\033[35m" // Magenta for booleans
	ColorNull    = "\033[90m" // Gray for null
	ColorBrace   = "\033[37m" // White for braces/brackets
)

type Formatter struct {
	Indent string
	Color  bool
}

func NewFormatter() *Formatter {
	return &Formatter{Indent: "  ", Color: true}
}

func (f *Formatter) Format(value parser.JSONValue, writer io.Writer) error {
	return f.formatWithIndent(value, writer, 0)
}

func (f *Formatter) formatWithIndent(value parser.JSONValue, writer io.Writer, depth int) error {
	indent := strings.Repeat(f.Indent, depth)
	nextIndent := strings.Repeat(f.Indent, depth+1)

	switch typedValue := value.(type) {
	case *parser.JSONObject:
		f.printColored(writer, "{", ColorBrace)
		fmt.Fprint(writer, "\n")
		index := 0
		for key, childValue := range typedValue.Data {
			fmt.Fprint(writer, nextIndent)
			f.printColored(writer, fmt.Sprintf("\"%s\"", key), ColorKey)
			fmt.Fprint(writer, ": ")
			f.formatWithIndent(childValue, writer, depth+1)
			if index < len(typedValue.Data)-1 {
				fmt.Fprint(writer, ",")
			}
			fmt.Fprint(writer, "\n")
			index++
		}
		fmt.Fprintf(writer, "%s", indent)
		f.printColored(writer, "}", ColorBrace)

	case *parser.JSONArray:
		f.printColored(writer, "[", ColorBrace)
		fmt.Fprint(writer, "\n")
		for index, element := range typedValue.Elements {
			fmt.Fprint(writer, nextIndent)
			f.formatWithIndent(element, writer, depth+1)
			if index < len(typedValue.Elements)-1 {
				fmt.Fprint(writer, ",")
			}
			fmt.Fprint(writer, "\n")
		}
		fmt.Fprintf(writer, "%s", indent)
		f.printColored(writer, "]", ColorBrace)

	case *parser.JSONString:
		f.printColored(writer, fmt.Sprintf("\"%s\"", typedValue.Value), ColorString)

	case *parser.JSONNumber:
		f.printColored(writer, fmt.Sprintf("%v", typedValue.Value), ColorNumber)

	case *parser.JSONBoolean:
		f.printColored(writer, fmt.Sprintf("%t", typedValue.Value), ColorBoolean)

	case *parser.JSONNull:
		f.printColored(writer, "null", ColorNull)

	default:
		return fmt.Errorf("unknown JSON value type")
	}

	return nil
}

func (f *Formatter) printColored(writer io.Writer, text, color string) {
	if f.Color {
		fmt.Fprintf(writer, "%s%s%s", color, text, ColorReset)
	} else {
		fmt.Fprint(writer, text)
	}
}
