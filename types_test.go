package dsv

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFloatParser(t *testing.T) {
	type testRow struct {
		Name    string  `dsv:"name"`
		Float32 float32 `dsv:"f32"`
		Float64 float64 `dsv:"f64"`
	}
	raw := []byte("name,f32,f64\nsergey,363.4,-3.27\nanton,1.82,715.0\n")

	r := bytes.NewReader(raw)
	data := testRow{}
	parser, err := NewCsvParser(r, false, &data)
	assert.NoError(t, err)

	eof, err := parser.Next()
	assert.NoError(t, err)
	assert.False(t, eof)

	assert.Equal(t, "sergey", data.Name)
	assert.Equal(t, float32(363.4), data.Float32)
	assert.Equal(t, -3.27, data.Float64)

	eof, err = parser.Next()
	assert.NoError(t, err)
	assert.False(t, eof)
	assert.Equal(t, "anton", data.Name)
	assert.Equal(t, float32(1.82), data.Float32)
	assert.Equal(t, float64(715), data.Float64)

	log.Printf("%v", complex(1, 5))
}

func TestComplexParser(t *testing.T) {
	type testRow struct {
		Name       string     `tsv:"name"`
		Complex64  complex64  `tsv:"c64"`
		Complex128 complex128 `tsv:"c128"`
	}
	raw := []byte("name,c64,c128\nsergey,3634+327i,-1.5+17.0i\n")

	r := bytes.NewReader(raw)
	data := testRow{}
	parser, err := NewCsvParser(r, false, &data)
	assert.NoError(t, err)

	eof, err := parser.Next()
	assert.NoError(t, err)
	assert.False(t, eof)

	assert.Equal(t, "sergey", data.Name)
	assert.Equal(t, complex64(complex(3634, 327)), data.Complex64)
	assert.Equal(t, -1.5+17.0i, data.Complex128)
}

func TestUintParser(t *testing.T) {
	type testRow struct {
		Name string `tsv:"name"`
		Unit uint   `tsv:"ut"`
	}
	raw := []byte("name,ut\nsergey,327\nanton,-3634\n")

	r := bytes.NewReader(raw)
	data := testRow{}
	parser, err := NewCsvParser(r, false, &data)
	assert.NoError(t, err)

	eof, err := parser.Next()
	assert.NoError(t, err)
	assert.False(t, eof)

	assert.Equal(t, "sergey", data.Name)
	assert.Equal(t, uint(327), data.Unit)

	eof, err = parser.Next()
	assert.Error(t, err)
}
