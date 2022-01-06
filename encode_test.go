package bencode_test

import (
	"testing"

	"github.com/ArunMurugan78/bencode"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	t.Run("encode string", func(t *testing.T) {
		val, err := bencode.Encode("rose")
		assert.Nil(t, err)
		assert.EqualValues(t, val, "4:rose")
		val, err = bencode.Encode("")
		assert.Nil(t, err)
		assert.EqualValues(t, "0:", val)
	})

	t.Run("encode int", func(t *testing.T) {
		val, err := bencode.Encode(78)
		assert.Nil(t, err)
		assert.EqualValues(t, val, "i78e")
		val, err = bencode.Encode(0)
		assert.Nil(t, err)
		assert.EqualValues(t, "i0e", val)
		val, err = bencode.Encode(-78)
		assert.Nil(t, err)
		assert.EqualValues(t, "i-78e", val)
		_, err = bencode.Encode(78.9)
		assert.NotNil(t, err)
	})

	t.Run("encode dict", func(t *testing.T) {
		val, err := bencode.Encode(map[string]interface{}{"name": []string{"one"}})
		assert.Nil(t, err)
		assert.EqualValues(t, val, "d4:namel3:oneee")

		val, err = bencode.Encode(map[string]interface{}{"names": []string{"rose", "arun"}, "num": 20, "sub": map[string]interface{}{"subfield": []string{"one", "two"}}})
		assert.Nil(t, err)
		assert.EqualValues(t, val, "d5:namesl4:rose4:arune3:numi20e3:subd8:subfieldl3:one3:twoeee")
	})

	t.Run("encode lists", func(t *testing.T) {
		val, err := bencode.Encode([]string{"one", "four", "fivee"})
		assert.Nil(t, err)
		assert.EqualValues(t, val, "l3:one4:four5:fiveee")
		val, err = bencode.Encode([]interface{}{"rose", map[string]interface{}{"one": "two"}})
		assert.Nil(t, err)
		assert.EqualValues(t, val, "l4:rosed3:one3:twoee")
		val, err = bencode.Encode([]interface{}{})
		assert.Nil(t, err)
		assert.EqualValues(t, val, "le")
	})
}
