package dsv

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/unicode/norm"
	"os"
	"testing"
)

type TestRow struct {
	Name   string
	Age    int
	Gender string
	Active bool
}

type TestTaggedRow struct {
	Name   string `dsv:"name"`
	Age    int    `dsv:"age"`
	Gender string `dsv:"gender"`
	Active bool   `dsv:"active"`
}

func TestTsvParserWithoutHeader(t *testing.T) {
	file, err := os.Open("fixtures/example_simple.tsv")
	assert.NoError(t, err)
	defer file.Close()
	data := TestRow{}
	parser, err := NewTsvParserWithoutHeader(file, false, &data)
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		eof, err := parser.Next()
		if eof {
			return
		}
		if i == 0 {
			assert.NoError(t, err)
			assert.Equal(t, "alex", data.Name)
			assert.Equal(t, 10, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, true, data.Active)
		}
		if i == 1 {
			assert.NoError(t, err)
			assert.Equal(t, "john", data.Name)
			assert.Equal(t, 24, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, false, data.Active)
		}
		if i == 2 {
			assert.NoError(t, err)
			assert.Equal(t, "sara", data.Name)
			assert.Equal(t, 30, data.Age)
			assert.Equal(t, "female", data.Gender)
			assert.Equal(t, true, data.Active)
		}
		if i == 3 {
			assert.Error(t, err)
		}
		if i == 4 {
			assert.Error(t, err)
		}
		if i == 5 {
			assert.Error(t, err)
		}
	}
}

func TestTsvParserTaggedStructure(t *testing.T) {
	file, err := os.Open("fixtures/example.tsv")
	assert.NoError(t, err)
	defer file.Close()
	data := TestTaggedRow{}
	parser, err := NewTsvParser(file, false, &data)
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		eof, err := parser.Next()
		if eof {
			return
		}
		if i == 0 {
			assert.NoError(t, err)
			assert.Equal(t, "alex", data.Name)
			assert.Equal(t, 10, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, true, data.Active)
		}
		if i == 1 {
			assert.NoError(t, err)
			assert.Equal(t, "john", data.Name)
			assert.Equal(t, 24, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, false, data.Active)
		}
		if i == 2 {
			assert.NoError(t, err)
			assert.Equal(t, "sara", data.Name)
			assert.Equal(t, 30, data.Age)
			assert.Equal(t, "female", data.Gender)
			assert.Equal(t, true, data.Active)
		}
	}
}

func TestTsvParserNormalize(t *testing.T) {
	file, err := os.Open("fixtures/example_norm.tsv")
	assert.NoError(t, err)
	defer file.Close()
	data := TestRow{}
	parser, err := NewTsvParser(file, false, &data)
	assert.NoError(t, err)
	// Use NFC as normalization
	parser.normalize = norm.NFKC

	for i := 0; i < 3; i++ {
		eof, err := parser.Next()
		if eof {
			return
		}
		if err != nil {
			t.Error(err)
		}
		if i == 0 {
			assert.Equal(t, "アレックス", data.Name)
		}
		if i == 1 {
			assert.Equal(t, "デボラ", data.Name)
		}
		if i == 2 {
			assert.Equal(t, "デボラ", data.Name)
		}
		if i == 3 {
			assert.Equal(t, "(テスト)", data.Name)
		}
		if i == 4 {
			assert.Equal(t, "/", data.Name)
		}
	}
}

func TestCsvParserWithoutHeader(t *testing.T) {
	file, err := os.Open("fixtures/example_simple.csv")
	assert.NoError(t, err)
	defer file.Close()
	data := TestRow{}
	parser, err := NewCsvParserWithoutHeader(file, false, &data)
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		eof, err := parser.Next()
		if eof {
			return
		}
		if i == 0 {
			assert.NoError(t, err)
			assert.Equal(t, "alex", data.Name)
			assert.Equal(t, 10, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, true, data.Active)
		}
		if i == 1 {
			assert.NoError(t, err)
			assert.Equal(t, "john", data.Name)
			assert.Equal(t, 24, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, false, data.Active)
		}
		if i == 2 {
			assert.NoError(t, err)
			assert.Equal(t, "sara", data.Name)
			assert.Equal(t, 30, data.Age)
			assert.Equal(t, "female", data.Gender)
			assert.Equal(t, true, data.Active)
		}
		if i == 3 {
			assert.Error(t, err)
		}
		if i == 4 {
			assert.Error(t, err)
		}
		if i == 5 {
			assert.Error(t, err)
		}
	}
}

func TestCsvParserTaggedStructure(t *testing.T) {
	file, err := os.Open("fixtures/example.csv")
	assert.NoError(t, err)
	defer file.Close()
	data := TestTaggedRow{}
	parser, err := NewCsvParser(file, false, &data)
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		eof, err := parser.Next()
		if eof {
			return
		}
		if i == 0 {
			assert.NoError(t, err)
			assert.Equal(t, "alex", data.Name)
			assert.Equal(t, 10, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, true, data.Active)
		}
		if i == 1 {
			assert.NoError(t, err)
			assert.Equal(t, "john", data.Name)
			assert.Equal(t, 24, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, false, data.Active)
		}
		if i == 2 {
			assert.NoError(t, err)
			assert.Equal(t, "sara", data.Name)
			assert.Equal(t, 30, data.Age)
			assert.Equal(t, "female", data.Gender)
			assert.Equal(t, true, data.Active)
		}
	}
}

func TestCsvParserNormalize(t *testing.T) {
	file, err := os.Open("fixtures/example_norm.csv")
	assert.NoError(t, err)
	defer file.Close()
	data := TestRow{}
	parser, err := NewCsvParser(file, false, &data)
	assert.NoError(t, err)
	// Use NFC as normalization
	parser.normalize = norm.NFKC

	for i := 0; i < 3; i++ {
		eof, err := parser.Next()
		if eof {
			return
		}
		if err != nil {
			t.Error(err)
		}
		if i == 0 {
			assert.Equal(t, "アレックス", data.Name)
		}
		if i == 1 {
			assert.Equal(t, "デボラ", data.Name)
		}
		if i == 2 {
			assert.Equal(t, "デボラ", data.Name)
		}
		if i == 3 {
			assert.Equal(t, "(テスト)", data.Name)
		}
		if i == 4 {
			assert.Equal(t, "/", data.Name)
		}
	}
}
