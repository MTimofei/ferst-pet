package convert_test

import (
	"reflect"
	"testing"
)

func TestByteToInt(t *testing.T) {
	testCases := []struct {
		input          []byte
		expectedOutput []int
	}{
		{
			input:          []byte{1, 2, 3},
			expectedOutput: []int{1, 2, 3},
		},
		{
			input:          []byte{255, 0, 128},
			expectedOutput: []int{255, 0, 128},
		},
		{
			input:          []byte{},
			expectedOutput: []int{},
		},
	}

	for _, testCase := range testCases {
		actualOutput := ByteToInt(testCase.input)
		if !reflect.DeepEqual(actualOutput, testCase.expectedOutput) {
			t.Errorf("Unexpected output. Input: %v. Expected: %v. Actual: %v", testCase.input, testCase.expectedOutput, actualOutput)
		}
	}

}

func ByteToInt(b []byte) {
	panic("unimplemented")
}
