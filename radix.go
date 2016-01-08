package radix

import (
	"fmt"
	"strings"
)

// RadixTree ...
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

// Add inserts a string into the trie
func (tree *RadixTree) Add(str string, content interface{}) {

	// Bail out if null
	if str == "" {
		return
	}

	// Convert input to rune slice
	input := []rune(str)
	leaf := tree.add(tree.root, input, 0)
	tree.stringCount++

	// Set the content only on the leaf node
	leaf.SetContent(content)
}

//
func (tree *RadixTree) add(node *radixNode, input []rune, depth int) *radixNode {

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
				continue
			}

			childRune := child.Key()[i : i+1][0]
			inputRune := input[i : i+1][0]

			// If the letter is a match
			if childRune == inputRune {

				// Are we on the last character of keys?
				if i+1 == len(child.Key()) {
					// Are there more children?
					if len(child.Children()) > 0 {
						tree.add(child, input[i+1:], depth+1)
					}
				}
			} else {

				// Else it's not a match, check if we've exhausted the
				// chance of it existing in another sibling, then:
				if childIndex == len(node.Children())-1 {

					// If it's the first letter, just insert to node
					// (not child)
					if i == 0 {
						return node.NewChild(input[i:])
					}

					// If we don't have a match on the first letter,
					// insert

					// If theres a single letter difference, we actually
					// need to break one before so we have a prefix
					if input[i : i+1][0] == 0 {
						i -= 1
					}

					// Break the child at i into 2 nodes
					tree.nodeCount++
					child.Break(i)

					if len(input[i:]) > 0 {
						tree.nodeCount++
						return child.NewChild(input[i:])
					}
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
			key []rune,
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
