package dsv

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPsvParser(t *testing.T) {
	psv := `name|age|gender|active
alex|10|male|true
john|24|male|false
sara|30|female|true
taro|40|male|true|overflow|record
hanako||female|no
mike|55|male
`
	psvReader := bytes.NewReader([]byte(psv))
	data := TestTaggedRow{}
	Delimiters["psv"] = '|'
	parser, err := NewParser(psvReader, false, &data, "psv")
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		eof, err := parser.Next()
		assert.NoError(t, err)
		assert.False(t, eof)

		if i == 0 {
			assert.Equal(t, "alex", data.Name)
			assert.Equal(t, 10, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, true, data.Active)
		}
		if i == 1 {
			assert.Equal(t, "john", data.Name)
			assert.Equal(t, 24, data.Age)
			assert.Equal(t, "male", data.Gender)
			assert.Equal(t, false, data.Active)
		}
		if i == 2 {
			assert.Equal(t, "sara", data.Name)
			assert.Equal(t, 30, data.Age)
			assert.Equal(t, "female", data.Gender)
			assert.Equal(t, true, data.Active)
		}
	}
}
