package convertor_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Richtermnd/game/cmd/mapconvertor/convertor"
	"github.com/Richtermnd/game/internal/field"
)

func TestOutputFilename(t *testing.T) {
	input := "test.txt"
	if convertor.OutputFilename(input) != "test.gf" {
		t.Fatal()
	}
}

func TestRead(t *testing.T) {
	testCases := []struct {
		desc    string
		w, h    int
		input   []byte
		output  []byte
		wantErr bool
	}{
		{
			desc:   "Valid input",
			w:      3,
			h:      3,
			input:  []byte("***\n***\n***"),
			output: []byte("*********"),
		},
		{
			desc:    "wrong height",
			w:       3,
			h:       2,
			input:   []byte("***\n***\n***"),
			wantErr: true,
		},
		{
			desc:    "wrong width",
			w:       2,
			h:       3,
			input:   []byte("***\n***\n***"),
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			output, err := convertor.Read(bytes.NewReader(tc.input), tc.w, tc.h)
			if err != nil && !tc.wantErr {
				t.FailNow()
			}
			if !reflect.DeepEqual(tc.output, output) {
				t.FailNow()
			}
		})
	}
}

func TestConvert(t *testing.T) {
	testCases := []struct {
		desc    string
		input   []byte
		output  []byte
		wantErr bool
	}{
		{
			desc:   "valid input",
			input:  []byte("# # #"),
			output: []byte{field.OBSTACLE, field.VOID, field.OBSTACLE, field.VOID, field.OBSTACLE},
		},
		{
			desc:    "unknown symbol",
			input:   []byte("xdd"),
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ouput, err := convertor.Convert(tc.input)
			if err != nil && !tc.wantErr {
				t.FailNow()
			}
			if !reflect.DeepEqual(tc.output, ouput) {
				t.FailNow()
			}
		})
	}
}
