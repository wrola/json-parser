package output

import (
	"fmt"
	"io"
	"strings"

	"github.com/wojciech/json-parser/internal/parser"
)

type Formatter struct {
	Indent string
}

func NewFormatter() *Formatter {
	return &Formatter{Indent: "  "}
}

func (f *Formatter) Format(value parser.JSONValue, writer io.Writer) error {
	return f.formatWithIndent(value, writer, 0)
}

func (f *Formatter) formatWithIndent(value parser.JSONValue, writer io.Writer, depth int) error {
	indent := strings.Repeat(f.Indent, depth)
	nextIndent := strings.Repeat(f.Indent, depth+1)

	switch typedValue := value.(type) {
	case *parser.JSONObject:
		fmt.Fprint(writer, "{\n")
		index := 0
		for key, childValue := range typedValue.Data {
			fmt.Fprintf(writer, "%s\"%s\": ", nextIndent, key)
			f.formatWithIndent(childValue, writer, depth+1)
			if index < len(typedValue.Data)-1 {
				fmt.Fprint(writer, ",")
			}
			fmt.Fprint(writer, "\n")
			index++
		}
		fmt.Fprintf(writer, "%s}", indent)

	case *parser.JSONArray:
		fmt.Fprint(writer, "[\n")
		for index, element := range typedValue.Elements {
			fmt.Fprint(writer, nextIndent)
			f.formatWithIndent(element, writer, depth+1)
			if index < len(typedValue.Elements)-1 {
				fmt.Fprint(writer, ",")
			}
			fmt.Fprint(writer, "\n")
		}
		fmt.Fprintf(writer, "%s]", indent)

	case *parser.JSONString:
		fmt.Fprintf(writer, "\"%s\"", typedValue.Value)

	case *parser.JSONNumber:
		fmt.Fprintf(writer, "%v", typedValue.Value)

	case *parser.JSONBoolean:
		fmt.Fprintf(writer, "%t", typedValue.Value)

	case *parser.JSONNull:
		fmt.Fprint(writer, "null")

	default:
		return fmt.Errorf("unknown JSON value type")
	}

	return nil
}
