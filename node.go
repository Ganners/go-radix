package radix

// The all-important building block
type radixNode struct {

	// The key is the set of runes contained in the node
	key []rune

	// The child runes is a bit map where 0-26 is A-Z and 27 - 32 is
	// squashed into 0-9
	childRunes int32

	// Contains a link up to the parent
	parent *radixNode

	// Child nodes
	children []*radixNode

	// Content associated with this node
	content interface{}
}

// Returns the key run slice
func (rn *radixNode) Key() []rune {
	return rn.key
}

// Returns the pointer to the parent node
func (rn *radixNode) Parent() *radixNode {
	return rn.parent
}

// Returns the children contained within the parent node
func (rn *radixNode) Children() []*radixNode {
	return rn.children
}

// Sets the content of a radix node
func (rn *radixNode) SetContent(content interface{}) {
	rn.content = content
}

// Returns the content, this will need type inference when it comes out
// to be of any use
func (rn *radixNode) Content() interface{} {
	return rn.content
}

// -----------------------------------------------------------------------------

// Inserts a child node
func (rn *radixNode) NewChild(key []rune, data interface{}) {

}

type (
	terminate  bool
	walkerFunc func([]rune, int) terminate
)

func (rn *radixNode) WalkDepthFirst(wf walkerFunc) {

}

func (rn *radixNode) WalkBreadthFirst(wf walkerFunc) {

}
