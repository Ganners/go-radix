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
	leaf := tree.add(tree.root, input)
	tree.stringCount++

	// Set the content only on the leaf node
	leaf.SetContent(content)
}

//
// apple
// appleby
// apboy
// orange
// orangina
//
// nodes empty? yes - insert apple
// search through children
//   -> search through runes (left to right)
//       -> break when found a match
//       -> is the match in the middle of a word?
//            -> break the word into two nodes
//            -> add remaining runes and return leaf?
//       -> else is our match a partial? return the recursion

func (tree *RadixTree) add(node *radixNode, input []rune) *radixNode {

	// Recursion down to 0 means we're all out and we should return
	if len(input) == 0 {
		return node
	}

	// If there are no children, FIRST
	if len(node.Children()) == 0 {
		tree.nodeCount++
		return node.NewChild(input)
	}

	// Loop through children
	for _, child := range node.Children() {
		for i := 0; i < len(child.Key()); i++ {

			// If this not the rune you are looking for...
			if string(child.Key()[i:i+1]) != string(input[i:i+1]) {
				if i > 0 {
					// Have we already got some depth? If so we need to
					// do something
				} else {
					// Else this is new and a sibling of the children
					tree.nodeCount++
					return node.NewChild(input)
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
