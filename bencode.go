package bencode

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var ErrInvalidFormat = errors.New("invalid format")

func isNumber(r rune) bool {
	return int(r) >= ('0') || int(r) <= int('9')
}

func isInt(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isNumber(rune(s[i])) {
			return false
		}
	}
	return true
}

func parseDictionary(reader *bufio.Reader) (map[string]interface{}, error) {
	parsedDict := make(map[string]interface{})

	r, _, err := reader.ReadRune()
	if err != nil {
		return nil, err
	}

	if r != 'd' {
		return nil, errors.New("'d' expected")
	}

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return nil, err
		}

		if r == 'e' {
			break
		}

		reader.UnreadRune()

		key, err := parseString(reader)
		if err != nil {
			return nil, err
		}

		val, err := parse(reader)
		if err != nil {
			return nil, err
		}

		parsedDict[key] = val
	}

	return parsedDict, nil
}

func parseString(reader *bufio.Reader) (string, error) {
	var len uint64

	str := ""
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}

		if r == ':' {
			break
		}

		str += string(r)
	}

	if !isInt(str) {
		return "", errors.New("integer expected")
	}

	len, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return "", err
	}

	str = ""
	for i := uint64(0); i < len; i++ {
		r, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}

		str += string(r)
	}

	return str, nil
}

func parseList(reader *bufio.Reader) ([]interface{}, error) {
	parsedList := make([]interface{}, 0)

	r, _, err := reader.ReadRune()
	if err != nil {
		return nil, err
	}

	if r != 'l' {
		return nil, errors.New("'l' expected")
	}

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return nil, err
		}

		if r == 'e' {
			break
		}
		reader.UnreadRune()

		val, err := parse(reader)
		if err != nil {
			return nil, err
		}

		parsedList = append(parsedList, val)
	}

	return parsedList, nil
}

func parseInt(reader *bufio.Reader) (int64, error) {
	neg := false

	r, _, err := reader.ReadRune()
	if err != nil {
		return 0, err
	}

	if r != 'i' {
		return 0, errors.New("'i' expected")
	}

	r, _, err = reader.ReadRune()

	if r != '-' {
		reader.UnreadRune()
	} else {
		neg = true
	}

	str := ""
	for {
		r, _, err = reader.ReadRune()
		if err != nil {
			return 0, err
		}

		if r == 'e' {
			break
		}

		str += string(r)
	}

	if len(str) == 0 {
		return 0, ErrInvalidFormat
	}

	// assert int
	if !isInt(str) {
		return 0, errors.New("integer expected")
	}

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	// zero cannot have a neg sign
	if num == 0 && neg {
		return 0, ErrInvalidFormat
	}

	// non zero number should not start with zero
	if num != 0 && rune(str[0]) == '0' {
		return 0, ErrInvalidFormat
	}

	if neg {
		num *= -1
	}

	return num, nil
}

func parse(reader *bufio.Reader) (interface{}, error) {
	r, _, err := reader.ReadRune()
	reader.UnreadRune()

	if err != nil {
		return nil, err
	}

	switch r {
	case 'd':
		return parseDictionary(reader)
	case 'i':
		return parseInt(reader)
	case 'l':
		return parseList(reader)
	default:
		if isNumber(r) {
			return parseString(reader)
		}

		return nil, ErrInvalidFormat
	}
}

func Parse(reader io.Reader) (interface{}, error) {
	return parse(bufio.NewReader(reader))
}

func ParseString(str string) (interface{}, error) {
	return Parse(strings.NewReader(str))
}
