package radix

import (
	"math"
	"testing"
)

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
		Input        []byte
		ExpectedBits []uint32
	}{
		{
			Input:        []byte{'a'},
			ExpectedBits: []uint32{1},
		},
		{
			Input:        []byte{'A'},
			ExpectedBits: []uint32{1},
		},
		{
			Input:        []byte{'z'},
			ExpectedBits: []uint32{26},
		},
		{
			Input:        []byte{'Z'},
			ExpectedBits: []uint32{26},
		},
		{
			Input:        []byte{'a', 'b', 'c', 'A', 'B', 'C'},
			ExpectedBits: []uint32{26, 27, 28},
		},
		{
			Input:        []byte{'0'},
			ExpectedBits: []uint32{27},
		},
		{
			Input:        []byte{'1'},
			ExpectedBits: []uint32{27},
		},
		{
			Input:        []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
			ExpectedBits: []uint32{27, 28, 29, 30, 31},
		},
		{
			Input:        []byte{'.', ',', '-'},
			ExpectedBits: []uint32{32},
		},
		{
			Input:        []byte{uint8(math.Min(255.0, float64('∂')))},
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

func TestBitMaskContains(t *testing.T) {

	testCases := []struct {
		Haystack uint32
		Needle   uint32
		Expected bool
	}{
		{
			Haystack: 7,
			Needle:   3,
			Expected: true,
		},
		{
			Haystack: 7,
			Needle:   1,
			Expected: true,
		},
		{
			Haystack: 7,
			Needle:   8,
			Expected: false,
		},
		{
			Haystack: 8,
			Needle:   7,
			Expected: false,
		},
	}

	for _, test := range testCases {

		res := bitMaskContains(test.Haystack, test.Needle)
		if res != test.Expected {
			t.Errorf("Error checking bit mask for %d in %d, got %t",
				test.Needle, test.Haystack, res)
		}
	}
}
