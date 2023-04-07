package convert_test

import (
	"pet/pkg/convert"
	"reflect"
	"testing"
)

func TestByteToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []int
	}{

		{"test 1", []byte{0, 1, 127, 128, 255}, []int{0, 1, 127, 128, 255}},
		{"test 2", []byte{255, 254, 128, 1, 0}, []int{255, 254, 128, 1, 0}},
		//{"test 3", []byte{}, []int{}},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := convert.ByteToInt(tt.input)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("unexpected output for %v. expected: %v, got: %v", tt.input, tt.expected, output)
			}
		})
	}
}

func TestIntToStr(t *testing.T) {
	testDT := []struct {
		name     string
		input    []int
		expected string
	}{
		{"test 1", []int{0, 1, 127, 128, 255}, "0,1,127,128,255"},
		{"test 2", []int{255, 254, 128, 1, 0}, "255,254,128,1,0"},
	}
	for _, tt := range testDT {
		t.Run(tt.name, func(t *testing.T) {
			output := convert.IntToStr(tt.input)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("unexpected output for %v. expected: %v, got: %v", tt.input, tt.expected, output)
			}
		})
	}
}

func TestStrToByte(t *testing.T) {
	testDT := []struct {
		name     string
		input    string
		expected []byte
	}{
		{"test 1", "0,1,127,128,255", []byte{0, 1, 127, 128, 255}},
		{"test 2", "255,254,128,1,0", []byte{255, 254, 128, 1, 0}},
	}
	for _, tt := range testDT {
		t.Run(tt.name, func(t *testing.T) {
			output := convert.StrToByte(tt.input)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("unexpected output for %v. expected: %v, got: %v", tt.input, tt.expected, output)
			}
		})
	}
}
