package radix

import "errors"

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

	// Is this something which was inserted?
	doCollect bool

	// The bit mask for all child letters (excluding itself)
	bitMask uint32
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

// OrBitMask will take a bit mask (uint32) and OR it (logical inclusive)
// to the current bit mask that is set.
func (rn *radixNode) OrBitMask(bitMask uint32) {
	rn.bitMask |= bitMask
}

// IsBitMaskSet performs a check to see if the bit mask is set
func (rn *radixNode) IsBitMaskSet(bitMask uint32) bool {
	return bitMaskContains(rn.bitMask, bitMask)
}

// BitMask returns the bit mask which is set, should only have practical uses
// in testing
func (rn *radixNode) BitMask() uint32 {
	return rn.bitMask
}

// Sets the node to be collected (this means it's a string that was
// inserted)
func (rn *radixNode) SetToCollect() {
	rn.doCollect = true
}

// Returns if this is a node which should be collected or not
func (rn *radixNode) Collect() bool {
	return rn.doCollect
}

// -----------------------------------------------------------------------------

// Inserts a child node
func (rn *radixNode) NewChild(key []rune) *radixNode {

	newNode := &radixNode{
		key:        key,
		childRunes: 0,
		parent:     rn,
		bitMask:    genBitMask(key),
	}
	rn.children = append(rn.children, newNode)

	return newNode
}

// Break will split a node into two nodes at a given index
func (rn *radixNode) Break(index int) (*radixNode, error) {

	if index > len(rn.Key()) {
		return nil, errors.New("Index exceeds key length")
	}

	// Split the string
	preKey := rn.Key()[:index]
	sufKey := rn.Key()[index:]
	content := rn.Content()
	children := rn.Children()

	// Set the vars, move children and add the child
	rn.key = preKey
	rn.content = nil
	rn.children = make([]*radixNode, 0)

	child := rn.NewChild(sufKey)
	child.children = children
	child.SetContent(content)

	// Rebuild the child bit mask (contain itself and it's children)
	child.OrBitMask(genBitMask(child.Key()))
	for _, childsChild := range child.Children() {
		child.OrBitMask(childsChild.BitMask())
	}

	// Generate a bitmask on the parent, should have it's child's runes
	// set too
	rn.OrBitMask(genBitMask(sufKey))

	return rn, nil
}

type (
	terminate  bool
	walkerFunc func([]rune, int, bool, bool, int) terminate
)

// WalkDepthFirst will execute a function for
// each node visited depth first in the tree
func (rn *radixNode) WalkDepthFirst(wf walkerFunc, depth int) {

	isFirst := true
	numCurrentChildren := len(rn.Children())

	for i, childNode := range rn.Children() {

		// Is this the last node at this depth?
		isLast := false
		numChildren := len(childNode.Children())
		if i == numCurrentChildren-1 {
			isLast = true
		}

		stop := wf(
			childNode.Key(),
			depth,
			isFirst,
			isLast,
			numChildren,
		)

		isFirst = false
		if stop == terminate(true) {
			return
		}
		childNode.WalkDepthFirst(wf, depth+1)
	}
}
