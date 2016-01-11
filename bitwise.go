// Keeping all of the bitwise operations together, these include
// counting the number of bits, creating bit masks and comparing bit
// masks.
package radix

// Counts the number of bits that are set using voodoo black magic from
// the gates of hell. Simplified as much as possible
//
// This is specific to 32 bits (8 pairs of hex goodness)
func numBitsSet32(n uint32) int {

	var a uint32 = 0x55555555
	var b uint32 = 0x33333333
	var c uint32 = 0x0F0F0F0F
	var d uint32 = 0x01010101

	n = n - ((n >> 1) & a)
	n = (n & b) + ((n >> 2) & b)
	n = (((n + (n >> 4)) & c) * d) >> 24

	return int(n)
}

// Similar to the 32, except with variables doubled, in case we want to
// increase precision a little later on.
func numBitsSet64(n uint64) int {

	var a uint64 = 0x5555555555555555
	var b uint64 = 0x3333333333333333
	var c uint64 = 0x0F0F0F0F0F0F0F0F
	var d uint64 = 0x0101010101010101

	n = n - ((n >> 1) & a)
	n = (n & b) + ((n >> 2) & b)
	n = (((n + (n >> 4)) & c) * d) >> 56

	return int(n)
}
