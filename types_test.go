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
		Float32 float32 `dsv:"float32"`
		Float64 float64 `dsv:"float64"`
		Comma   float64 `dsv:"comma"`
	}
	raw := []byte("name,float32,float64,comma\nsergey,363.4,-3.27,\"1,989.12\"\nanton,1.82,715.0,\n")

	r := bytes.NewReader(raw)
	data := testRow{}
	parser, err := NewCsvParser(r, true, &data)
	assert.NoError(t, err)

	eof, err := parser.Next()
	assert.NoError(t, err)
	assert.False(t, eof)

	assert.Equal(t, "sergey", data.Name)
	assert.Equal(t, float32(363.4), data.Float32)
	assert.Equal(t, -3.27, data.Float64)
	assert.Equal(t, 1989.12, data.Comma)

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
		Complex64  complex64  `tsv:"complex64"`
		Complex128 complex128 `tsv:"complex128"`
	}
	raw := []byte("name,complex64,complex128\nsergey,3634+327i,-1.5+17.0i\n")

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
		Name   string `tsv:"name"`
		Uint   uint   `tsv:"unit"`
		Uint8  uint8  `tsv:"unit8"`
		Uint16 uint16 `tsv:"unit16"`
		Uint32 uint32 `tsv:"unit32"`
		Uint64 uint64 `tsv:"unit64"`
	}
	raw := []byte("name,unit,unit8,unit16,unit32,unit64\nsergey,327,222,327,327,327\nanton,-3634,-222,-3634,-3634,-3634\n")

	r := bytes.NewReader(raw)
	data := testRow{}
	parser, err := NewCsvParser(r, false, &data)
	assert.NoError(t, err)

	eof, err := parser.Next()
	assert.NoError(t, err)
	assert.False(t, eof)

	assert.Equal(t, "sergey", data.Name)
	assert.Equal(t, uint(327), data.Uint)
	assert.Equal(t, uint8(222), data.Uint8)
	assert.Equal(t, uint16(327), data.Uint16)
	assert.Equal(t, uint32(327), data.Uint32)
	assert.Equal(t, uint64(327), data.Uint64)

	eof, err = parser.Next()
	assert.Error(t, err)
}

func TestIntParser(t *testing.T) {
	type testRow struct {
		Name  string `tsv:"name"`
		Int   int    `tsv:"int"`
		Int8  int8   `tsv:"int8"`
		Int16 int16  `tsv:"int16"`
		Int32 int32  `tsv:"int32"`
		Int64 int64  `tsv:"int64"`
	}
	raw := []byte("name,int,int8,int16,int32,int64\nsergey,327,-111,327,327,-327\nanton,-3634,-222,-3634,-3634,-3634\n")

	r := bytes.NewReader(raw)
	data := testRow{}
	parser, err := NewCsvParser(r, false, &data)
	assert.NoError(t, err)

	eof, err := parser.Next()
	assert.NoError(t, err)
	assert.False(t, eof)

	assert.Equal(t, "sergey", data.Name)
	assert.Equal(t, 327, data.Int)
	assert.Equal(t, int8(-111), data.Int8)
	assert.Equal(t, int16(327), data.Int16)
	assert.Equal(t, int32(327), data.Int32)
	assert.Equal(t, int64(-327), data.Int64)

	eof, err = parser.Next()
	assert.Error(t, err)
}
