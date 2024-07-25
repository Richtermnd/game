package convertor

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/Richtermnd/game/internal/field"
)

var (
	dictionary = map[byte]byte{
		' ': field.VOID,
		'_': field.EMPTY,
		'#': field.OBSTACLE,
	}
)

func OutputFilename(inputFilename string) string {
	return strings.TrimSuffix(inputFilename, ".txt") + ".gf" // gf - game field
}

func Read(r io.Reader, w, h int) ([]byte, error) {
	result := make([]byte, 0, w*h)
	scanner := bufio.NewScanner(r)
	var rowsCounter int
	for scanner.Scan() {
		row := scanner.Bytes()
		if len(row) != w {
			return nil, fmt.Errorf("wrong width. Expected: %d Got: %d", w, len(row))
		}
		result = append(result, row...)
		rowsCounter++
	}
	if rowsCounter != h {
		return nil, fmt.Errorf("wrong height. Expected: %d Got: %d", h, rowsCounter)
	}
	return result, nil
}

func Convert(content []byte) ([]byte, error) {
	result := make([]byte, 0, len(content))
	for _, c := range content {
		t, ok := dictionary[c]
		if !ok {
			return nil, fmt.Errorf("unknown symbol: %s", string([]byte{c}))
		}
		result = append(result, t)
	}
	return result, nil
}
