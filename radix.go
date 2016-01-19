package radix

import (
	"fmt"
	"strings"
)

const (
	fuzzyIterationLimit = 2
)

// RadixTree wraps the root, provides all functionality to add, search
// and so on.
type RadixTree struct {
	root        *radixNode
	stringCount int
	nodeCount   int
}

// NewRadixTree sets up and returns a RadixTree struct
func NewRadixTree() *RadixTree {

	// Build the zero-value radix tree
	return &RadixTree{
		root: &radixNode{},
	}
}

// FuzzySearch launches the fuzzySearch() method if the search is valid. Here,
// unlike with prefix search, we do not accept an empty string.
func (tree *RadixTree) FuzzySearch(
	str string,
) ([]string, []interface{}) {

	if len(tree.root.Children()) == 0 {
		return []string{}, []interface{}{}
	}

	if str == "" {
		return []string{}, []interface{}{}
	}

	return tree.fuzzySearch(
		tree.stringToBytes(str),
		tree.root,
		0,
		0,
		0,
		[]byte{})
}

// fuzzySearch performs a non-prefix search with some element of 'fuzz',
// although this is somewhat restricted in that it won't cater for large sparse
// string segments.
//
// The fuzziness is achieved through bitwise operations that check if under a
// given node, the letters we are searching for exist. If they do then descend
func (tree *RadixTree) fuzzySearch(
	str []byte,
	node *radixNode,
	index int,
	iteration int,
	lastIncrement int,
	found []byte,
) ([]string, []interface{}) {

	searchBitMask := genBitMask(str[index:])
	collectedKeys := []string{}
	collectedContent := []interface{}{}

	if len(node.Children()) == 0 {
		return []string{}, []interface{}{}
	}

	startIndex := index

	for _, child := range node.Children() {

		// Reset index for each iteration of the child
		index = startIndex

		// If this is the case, then somewhere inside the depth of this
		// node there MIGHT exist what we're looking for, or it could
		// be shallow
		if child.IsBitMaskSet(searchBitMask) {

			// Iterate letters
			for _, letter := range child.Key() {
				if index < len(str) {
					if letter == str[index] {
						lastIncrement = iteration
						index++
					}

					// Small optimization, break early
					if index >= len(str) {
						break
					}
				}
				iteration++
			}

			if index >= len(str) {

				colKeys, colContent := tree.collect(
					child,
					append(found, child.Key()...),
				)
				collectedKeys = append(collectedKeys, colKeys...)
				collectedContent = append(collectedContent, colContent...)
			} else {

				colKeys, colContent := tree.fuzzySearch(
					str,
					child,
					index,
					iteration,
					lastIncrement,
					append(found, child.Key()...),
				)
				collectedKeys = append(collectedKeys, colKeys...)
				collectedContent = append(collectedContent, colContent...)
			}
		} else {
			// Not set, can't do anything here really
			continue
		}
	}

	return collectedKeys, collectedContent
}

// PrefixSearch executes the fastest form of search, whereby it iterates
// through the letters starting from index 0
//
// Search with an empty string should return everything
func (tree *RadixTree) PrefixSearch(
	str string,
) ([]string, []interface{}) {

	if len(tree.root.Children()) == 0 {
		return []string{}, []interface{}{}
	}

	return tree.prefixSearch(
		tree.stringToBytes(str),
		tree.root,
		0,
		[]byte{})
}

// Recursively prefix-searches
func (tree *RadixTree) prefixSearch(
	str []byte,
	node *radixNode,
	index int,
	found []byte,
) ([]string, []interface{}) {

	if index+1 > len(str) {

		// Collect below node with found prefix
		return tree.collect(node, found)
	}

	searchLetter := str[index]

	for _, child := range node.Children() {
		for _, letter := range child.Key() {

			// A matching letter has been found
			if searchLetter == letter {

				// Recurse, iterate by the number of keys in this node.
				// There's a guarantee with prefix trees so we don't
				// actually need to look at all letters
				newIndex := index + len(child.Key())
				toAppend := child.Key()

				return tree.prefixSearch(
					str,
					child,
					newIndex,
					append(found, toAppend...))
			}
		}
	}

	return []string{}, []interface{}{}
}

// The collection will, starting from a given node, recurse and generate
// strings from every leaf
func (tree *RadixTree) collect(
	node *radixNode,
	prefix []byte,
) ([]string, []interface{}) {

	if len(node.Children()) == 0 {
		return []string{string(prefix)},
			[]interface{}{node.Content()}
	}

	collectedStrings := []string{}
	collectedContent := []interface{}{}

	if node.Collect() {
		collectedStrings = append(collectedStrings, string(prefix))
		collectedContent = append(collectedContent, node.Content())
	}

	// Recursively append
	for _, child := range node.Children() {
		bytes := append(prefix, child.Key()...)
		colKeys, colContent := tree.collect(child, bytes)
		collectedStrings = append(collectedStrings, colKeys...)
		collectedContent = append(collectedContent, colContent...)
	}

	return collectedStrings, collectedContent
}

// Add inserts a string into the trie, it returns the node that it
// inserts (for testing purposes)
func (tree *RadixTree) Add(
	str string, content interface{}) *radixNode {

	// Bail out if null
	if str == "" {
		return &radixNode{}
	}

	// Convert input to byte slice
	input := tree.stringToBytes(str)

	bitMask := genBitMask(input)
	leaf := tree.add(tree.root, input, bitMask, 0)
	tree.stringCount++

	// Set the content only on the leaf node
	leaf.SetToCollect()

	leaf.SetContent(content)
	return leaf
}

// The brains behind the adding, handles all cases for adding new keys
func (tree *RadixTree) add(
	node *radixNode,
	input []byte,
	bitMask uint32,
	depth int,
) *radixNode {

	// Recursion down to 0 means we're all out and we should return
	if len(input) == 0 {
		return node
	}

	// If there are no children, I.e. this is the top element
	if len(node.Children()) == 0 {
		tree.nodeCount++
		return node.NewChild(input)
	}

	// Loop through children
	for childIndex, child := range node.Children() {

		// Loop through the letters that make up the key
		for i := 0; i < len(child.Key()); i++ {

			// Add some safety to the input
			if i > len(input) {
				break
			}

			var inputbyte byte
			if i+1 <= len(input) {
				inputbyte = input[i : i+1][0]
			}

			childbyte := child.Key()[i : i+1][0]

			// If the letter is a match
			if childbyte == inputbyte {

				child.OrBitMask(genBitMask(input[i:]))

				// Are we on the last character of keys?
				if i+1 == len(child.Key()) {

					// Are there more children?
					if len(child.Children()) > 0 {
						return tree.add(child, input[i+1:], bitMask, depth+1)
					}
				}
			} else {

				// Else it's not a match, check if we've exhausted the
				// chance of it existing in another sibling, then:
				if i > 0 {

					// Break the child at i into 2 nodes
					tree.nodeCount++
					child.Break(i)

					if len(input[i:]) > 0 {
						tree.nodeCount++
						newNode := child.NewChild(input[i:])
						return newNode
					} else {

						// If the break is less than the input then
						// return the child (which is the parent of any
						// new child)
						return child
					}
				} else {

					// If there are more nodes to be seen, continue
					if childIndex+1 < len(node.Children()) {
						break
					}

					// If it's the first letter, just insert to node
					// (not child)
					tree.nodeCount++
					newNode := node.NewChild(input[i:])
					return newNode
				}
			}
		}
	}

	return node
}

// String generates an ASCII tree to allow the data structure to be
// visualised
func (rt *RadixTree) String() string {

	output := "\n"
	first := true

	rt.root.WalkDepthFirst(
		func(
			key []byte,
			depth int,
			firstAtDepth bool,
			lastAtDepth bool,
			numChildren int,
		) terminate {

			if !first && firstAtDepth {
				output += strings.Repeat(" ", (depth*3)-3)
				output += "|\n"
			}

			if depth > 0 {
				output += strings.Repeat(" ", (depth*3)-3)
				output += "+- "
			}

			output += fmt.Sprintf("[%s]\n", string(key))

			first = false
			return terminate(false)
		}, 0)

	return output
}

// Because we're converting from utf8 down, we'll max out at 255 on the
// letter's value as to not overflow a byte
func (rt *RadixTree) stringToBytes(str string) []byte {

	bytes := make([]byte, len(str))

	for i, letter := range str {
		if letter > rune(255) {
			// This is ignored
			// bytes[i] = 0
		} else {
			bytes[i] = uint8(letter)
		}
	}

	return bytes
}
