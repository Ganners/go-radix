package radix

import "testing"

type bitwiseTestCase []struct {
	Input  int
	Output int
}

func getBitwiseTestCases() bitwiseTestCase {

	testCase := bitwiseTestCase{
		{
			// 0 0 0 0 0 1 1 1
			Input:  7,
			Output: 3,
		},
		{
			// 0 0 0 0 0 0 1 0
			Input:  2,
			Output: 1,
		},
		{
			// 0 0 1 0 1 0 1 1
			Input:  51,
			Output: 4,
		},
		{
			// 0 0 1 1 1 1 1 1
			Input:  63,
			Output: 6,
		},
	}

	return testCase
}

// Tests the 32 bit version of the function (converts explicitly the int
// to int32)
func TestNumBitsSet32(t *testing.T) {

	testCases := getBitwiseTestCases()

	for _, test := range testCases {
		input := uint32(test.Input)
		output := test.Output

		res := numBitsSet32(input)
		if res != output {
			t.Errorf("Expected number of bits set to be %d, got %d",
				output, res)
		}
	}
}

// Tests the 64 bit version of the function (converts explicitly the int
// to int64)
func TestNumBitsSet64(t *testing.T) {

	testCases := getBitwiseTestCases()

	for _, test := range testCases {
		input := uint64(test.Input)
		output := test.Output

		res := numBitsSet64(input)
		if res != output {
			t.Errorf("Expected number of bits set to be %d, got %d",
				output, res)
		}
	}
}
