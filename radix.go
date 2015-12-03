package radix

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
		node.NewChild(input, struct{}{})
	}

	// Loop through children
	for _, child := range node.Children() {
		for i := 0; i < len(child.Key()); i++ {
		}
	}

	return node
}

// String generates an ASCII tree to allow the data structure to be
// visualised
func (rt *RadixTree) String() string {
	return ""
}
