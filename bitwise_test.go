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

// Tests that generated bitmasks appear how we expect
func TestGenBitMask(t *testing.T) {

	testCases := []struct {
		Input        []rune
		ExpectedBits []uint32
	}{
		{
			Input:        []rune{'a'},
			ExpectedBits: []uint32{1},
		},
		{
			Input:        []rune{'A'},
			ExpectedBits: []uint32{1},
		},
		{
			Input:        []rune{'z'},
			ExpectedBits: []uint32{26},
		},
		{
			Input:        []rune{'Z'},
			ExpectedBits: []uint32{26},
		},
		{
			Input:        []rune{'a', 'b', 'c', 'A', 'B', 'C'},
			ExpectedBits: []uint32{26, 27, 28},
		},
		{
			Input:        []rune{'0'},
			ExpectedBits: []uint32{27},
		},
		{
			Input:        []rune{'1'},
			ExpectedBits: []uint32{27},
		},
		{
			Input:        []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
			ExpectedBits: []uint32{27, 28, 29, 30, 31},
		},
		{
			Input:        []rune{'.', ',', '-'},
			ExpectedBits: []uint32{32},
		},
		{
			Input:        []rune{'∂'},
			ExpectedBits: []uint32{32},
		},
	}

	for _, test := range testCases {

		// Call the function
		res := genBitMask(test.Input)

		for _, bit := range test.ExpectedBits {

			// Check if the bit is set
			if (res &^ (1 << bit)) == 0 {
				t.Errorf("Expected bit %d to be set for input %s, got uint32: %d",
					bit,
					string(test.Input),
					res)
			}
		}
	}
}
