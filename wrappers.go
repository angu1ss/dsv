package dsv

import "io"

// Delimiters contains a map like "extension": "delimiter"
var Delimiters = map[string]rune{
	"tsv": '\t',
	"csv": ',',
}

// NewTsvParser creates new TSV parser with given io.Reader
func NewTsvParser(reader io.Reader, lazyQuotes bool, data interface{}) (*Parser, error) {
	return NewParser(reader, lazyQuotes, data, "tsv")
}

// NewCsvParser creates new CSV parser with given io.Reader
func NewCsvParser(reader io.Reader, lazyQuotes bool, data interface{}) (*Parser, error) {
	return NewParser(reader, lazyQuotes, data, "csv")
}

// NewTsvParserWithoutHeader creates new TSV parser without header with given io.Reader
func NewTsvParserWithoutHeader(reader io.Reader, lazyQuotes bool, data interface{}) (*Parser, error) {
	return NewParserWithoutHeader(reader, lazyQuotes, data, "tsv")
}

// NewCsvParserWithoutHeader creates new CSV parser without header with given io.Reader
func NewCsvParserWithoutHeader(reader io.Reader, lazyQuotes bool, data interface{}) (*Parser, error) {
	return NewParserWithoutHeader(reader, lazyQuotes, data, "csv")
}
