package bencode_test

import (
	"io"
	"strings"
	"testing"

	"github.com/ArunMurugan78/bencode"
	"github.com/stretchr/testify/assert"
)

func Test_Bencode(t *testing.T) {
	t.Run("string parsing", func(t *testing.T) {
		val, err := bencode.ParseString("4:arun")
		assert.Nil(t, err)
		assert.EqualValues(t, val, "arun")
		val, err = bencode.ParseString("0:")
		assert.Nil(t, err)
		assert.EqualValues(t, val, "")
		val, err = bencode.Parse(strings.NewReader("6:rosebphaha"))
		assert.Nil(t, err)
		assert.EqualValues(t, val, "rosebp")
		_, err = bencode.Parse(strings.NewReader("-0:sjldflsd"))
		assert.NotNil(t, err)
	})

	t.Run("list parsing", func(t *testing.T) {
		val, err := bencode.ParseString("l4:arun3:abii356ed4:name4:arunee")
		assert.Nil(t, err)
		list, ok := val.([]interface{})
		assert.EqualValues(t, ok, true)
		assert.EqualValues(t, list[0], "arun")
		assert.EqualValues(t, list[1], "abi")
		assert.EqualValues(t, list[2], 356)
		mp, ok := list[3].(map[string]interface{})
		assert.EqualValues(t, ok, true)
		assert.EqualValues(t, mp["name"], "arun")
		_, err = bencode.Parse(strings.NewReader("l4:rose3:onei6969e"))
		assert.NotNil(t, err)
	})

	t.Run("integer parsing", func(t *testing.T) {
		val, err := bencode.Parse(strings.NewReader("i69e"))
		assert.Nil(t, err)
		assert.EqualValues(t, val, 69)
		val, err = bencode.Parse(strings.NewReader("i0e"))
		assert.Nil(t, err)
		assert.EqualValues(t, val, 0)
		_, err = bencode.ParseString("i-0e")
		assert.NotNil(t, err)
		_, err = bencode.ParseString("i083673e")
		assert.NotNil(t, err)
		val, err = bencode.ParseString("i-6969e")
		assert.Nil(t, err)
		assert.EqualValues(t, val, -6969)
		_, err = bencode.ParseString("i6347387")
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, io.EOF)
		_, err = bencode.ParseString("i78.33e")
		assert.NotNil(t, err)
		_, err = bencode.ParseString("ie")
		assert.NotNil(t, err)
	})

	t.Run("test dictionary", func(t *testing.T) {
		val, err := bencode.ParseString("d4:name4:rosee")
		mp, ok := val.(map[string]interface{})
		assert.EqualValues(t, ok, true)
		assert.Nil(t, err)
		assert.EqualValues(t, mp["name"], "rose")
		val, err = bencode.ParseString("d4:listl4:rose3:onei78eee")
		mp, ok = val.(map[string]interface{})
		assert.EqualValues(t, ok, true)
		assert.Nil(t, err)
		list, ok := mp["list"].([]interface{})
		assert.EqualValues(t, ok, true)
		assert.EqualValues(t, list[0], "rose")
		assert.EqualValues(t, list[1], "one")
		assert.EqualValues(t, list[2], 78)
		_, err = bencode.Parse(strings.NewReader("d4:namee"))
		assert.NotNil(t, err)
		_, err = bencode.Parse(strings.NewReader("d4:name4:rose"))
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, io.EOF)
		_, err = bencode.Parse(strings.NewReader("di36e4:rosee"))
		assert.NotNil(t, err)
	})
}
