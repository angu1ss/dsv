package dsv

import (
	"encoding/csv"
	"errors"
	"golang.org/x/text/unicode/norm"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// Parser has information for parser
type Parser struct {
	Headers    []string
	Reader     *csv.Reader
	Data       interface{}
	ref        reflect.Value
	indices    []int // indices are field index list of header array
	structMode bool
	normalize  norm.Form
}

// NewParser creates new DSV parser with given io.Reader and delimiter type
func NewParser(reader io.Reader, lazyQuotes bool, data interface{}, t string) (p *Parser, err error) {
	var delimiter rune
	delimiter, ok := Delimiters[t]
	if !ok {
		err = errors.New("unsupported dsv type: " + t)
		return nil, err
	}

	r := csv.NewReader(reader)
	r.Comma = delimiter
	r.LazyQuotes = lazyQuotes

	// first line should be fields
	headers, err := r.Read()

	if err != nil {
		return nil, err
	}

	for i, header := range headers {
		headers[i] = header
	}

	p = &Parser{
		Reader:     r,
		Headers:    headers,
		Data:       data,
		ref:        reflect.ValueOf(data).Elem(),
		indices:    make([]int, len(headers)),
		structMode: false,
		normalize:  -1,
	}

	// get type information
	tp := p.ref.Type()

	for i := 0; i < tp.NumField(); i++ {
		// get DSV tag
		tag := tp.Field(i).Tag.Get("dsv")
		if tag != "" {
			// find DSV position by header
			for j := 0; j < len(headers); j++ {
				if headers[j] == tag {
					// indices are 1 start
					p.indices[j] = i + 1
					p.structMode = true
				}
			}
		}
	}

	if !p.structMode {
		for i := 0; i < len(headers); i++ {
			p.indices[i] = i + 1
		}
	}

	return p, nil
}

// NewParserWithoutHeader creates new DSV parser with given io.Reader and delimiter type
func NewParserWithoutHeader(reader io.Reader, lazyQuotes bool, data interface{}, t string) (p *Parser, err error) {
	var delimiter rune
	delimiter, ok := Delimiters[t]
	if !ok {
		err = errors.New("unsupported dsv type: " + t)
		return nil, err
	}

	r := csv.NewReader(reader)
	r.Comma = delimiter
	r.LazyQuotes = lazyQuotes

	p = &Parser{
		Reader:    r,
		Data:      data,
		ref:       reflect.ValueOf(data).Elem(),
		normalize: -1,
	}

	return p, nil
}

// Next puts reader forward by a line
func (p *Parser) Next() (eof bool, err error) {

	// Get next record
	var records []string

	for {
		// read until valid record
		records, err = p.Reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				return true, nil
			}
			return false, err
		}
		if len(records) > 0 {
			break
		}
	}

	if len(p.indices) == 0 {
		p.indices = make([]int, len(records))
		// mapping simple index
		for i := 0; i < len(records); i++ {
			p.indices[i] = i + 1
		}
	}

	// record should be a pointer
	for i, record := range records {
		idx := p.indices[i]
		if idx == 0 {
			// skip empty index
			continue
		}
		// get target field
		field := p.ref.Field(idx - 1)
		switch field.Kind() {
		case reflect.String:
			// Normalize text
			if p.normalize >= 0 {
				record = p.normalize.String(record)
			}
			field.SetString(record)
		case reflect.Bool:
			if record == "" {
				field.SetBool(false)
			} else {
				col, err := strconv.ParseBool(record)
				if err != nil {
					return false, err
				}
				field.SetBool(col)
			}
		case reflect.Int:
			if record == "" {
				field.SetInt(0)
			} else {
				col, err := strconv.ParseInt(record, 10, 0)
				if err != nil {
					return false, err
				}
				field.SetInt(col)
			}
		case reflect.Int8:
			if record == "" {
				field.SetInt(0)
			} else {
				col, err := strconv.ParseInt(record, 10, 8)
				if err != nil {
					return false, err
				}
				field.SetInt(col)
			}
		case reflect.Int16:
			if record == "" {
				field.SetInt(0)
			} else {
				col, err := strconv.ParseInt(record, 10, 16)
				if err != nil {
					return false, err
				}
				field.SetInt(col)
			}
		case reflect.Int32:
			if record == "" {
				field.SetInt(0)
			} else {
				col, err := strconv.ParseInt(record, 10, 32)
				if err != nil {
					return false, err
				}
				field.SetInt(col)
			}
		case reflect.Int64:
			if record == "" {
				field.SetInt(0)
			} else {
				col, err := strconv.ParseInt(record, 10, 64)
				if err != nil {
					return false, err
				}
				field.SetInt(col)
			}
		case reflect.Uint:
			if record == "" {
				field.SetUint(0)
			} else {
				col, err := strconv.ParseUint(record, 10, 0)
				if err != nil {
					return false, err
				}
				field.SetUint(col)
			}
		case reflect.Uint8:
			if record == "" {
				field.SetUint(0)
			} else {
				col, err := strconv.ParseUint(record, 10, 8)
				if err != nil {
					return false, err
				}
				field.SetUint(col)
			}
		case reflect.Uint16:
			if record == "" {
				field.SetUint(0)
			} else {
				col, err := strconv.ParseUint(record, 10, 16)
				if err != nil {
					return false, err
				}
				field.SetUint(col)
			}
		case reflect.Uint32:
			if record == "" {
				field.SetUint(0)
			} else {
				col, err := strconv.ParseUint(record, 10, 32)
				if err != nil {
					return false, err
				}
				field.SetUint(col)
			}
		case reflect.Uint64:
			if record == "" {
				field.SetUint(0)
			} else {
				col, err := strconv.ParseUint(record, 10, 64)
				if err != nil {
					return false, err
				}
				field.SetUint(col)
			}
		case reflect.Float32:
			if record == "" {
				field.SetFloat(0)
			} else {
				// Prevent strconv error for records like "1,989.00"
				record = strings.ReplaceAll(record, ",", "")
				col, err := strconv.ParseFloat(record, 32)
				if err != nil {
					return false, err
				}
				field.SetFloat(col)
			}
		case reflect.Float64:
			if record == "" {
				field.SetFloat(0)
			} else {
				// Prevent strconv error for records like "1,991.00"
				record = strings.ReplaceAll(record, ",", "")
				col, err := strconv.ParseFloat(record, 64)
				if err != nil {
					return false, err
				}
				field.SetFloat(col)
			}
		case reflect.Complex64:
			if record == "" {
				field.SetComplex(0)
			} else {
				col, err := strconv.ParseComplex(record, 64)
				if err != nil {
					return false, err
				}
				field.SetComplex(col)
			}
		case reflect.Complex128:
			if record == "" {
				field.SetComplex(0)
			} else {
				col, err := strconv.ParseComplex(record, 128)
				if err != nil {
					return false, err
				}
				field.SetComplex(col)
			}
		default:
			return false, errors.New("unsupported field type: " + field.Type().String())
		}
	}

	return false, nil
}
