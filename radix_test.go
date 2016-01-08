package radix

import (
	"strings"
	"testing"
)

// I like the ability to visualise, to help explain how things work to
// others and to find bugs that I wouldn't normally think of
func TestDrawVisualisation(t *testing.T) {

	// Example from Wikipedia includes the following words:
	// + romane
	// + romanus
	// + romulus
	// + rubens
	// + ruber
	// + rubicon
	// + rubicundus

	r := &RadixTree{
		root: &radixNode{
			children: []*radixNode{
				{
					key: []rune("r"),
					children: []*radixNode{
						{
							key: []rune("om"),
							children: []*radixNode{
								{
									key: []rune("ulus"),
								},
								{
									key: []rune("an"),
									children: []*radixNode{
										{
											key: []rune("e"),
										},
										{
											key: []rune("us"),
										},
									},
								},
							},
						},
						{
							key: []rune("ub"),
							children: []*radixNode{
								{
									key: []rune("e"),
									children: []*radixNode{
										{
											key: []rune("ns"),
										},
										{
											key: []rune("r"),
										},
									},
								},
								{
									key: []rune("ic"),
									children: []*radixNode{
										{
											key: []rune("on"),
										},
										{
											key: []rune("undus"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		stringCount: 7,
		nodeCount:   13,
	}

	// A depth-first visualisation of the tree structure
	expect := strings.Join([]string{
		``,
		`[r]`,
		`|`,
		`+- [om]`,
		`   |`,
		`   +- [ulus]`,
		`   +- [an]`,
		`      |`,
		`      +- [e]`,
		`      +- [us]`,
		`+- [ub]`,
		`   |`,
		`   +- [e]`,
		`      |`,
		`      +- [ns]`,
		`      +- [r]`,
		`   +- [ic]`,
		`      |`,
		`      +- [on]`,
		`      +- [undus]`,
		``,
	}, "\n")

	toString := r.String()
	if toString != expect {
		t.Errorf(
			strings.Join([]string{
				"String output did not match expected visualisation.",
				"Got '%s' Expected '%s'",
			}, ""), toString, expect)
	}
}

// Most basic test, can we insert two words with no overlapping characters
func TestSparseInsert(t *testing.T) {

	r := NewRadixTree()
	r.Add("chocolate", struct{}{})
	r.Add("pizza", struct{}{})

	// Expected
	expected := &RadixTree{
		root: &radixNode{
			children: []*radixNode{
				{key: []rune("chocolate"), content: struct{}{}},
				{key: []rune("pizza"), content: struct{}{}},
			},
		},
		stringCount: 2,
		nodeCount:   2,
	}

	// If it doesn't match..
	if r.String() != expected.String() {
		t.Errorf("Result %s does not match expected %s", r.String(), expected.String())
	}
}

// This will be a test which requires splitting a radix into at least two nodes
func TestInsertBreakingEnd(t *testing.T) {

	r := NewRadixTree()
	r.Add("magazine", struct{}{})
	r.Add("magsafe", struct{}{})

	// Expected
	expected := &RadixTree{
		root: &radixNode{
			children: []*radixNode{
				{
					key:     []rune("mag"),
					content: struct{}{},
					children: []*radixNode{
						{key: []rune("azine"), content: struct{}{}},
						{key: []rune("safe"), content: struct{}{}},
					},
				},
			},
		},
		stringCount: 2,
		nodeCount:   2,
	}

	// If it doesn't match..
	if r.String() != expected.String() {
		t.Errorf("Result %s does not match expected %s", r.String(), expected.String())
	}
}

// Test a prefix search
func TestTraversePrefix(t *testing.T) {

}

// Test an in-tree search (non-prefix) with some element of fuzz
func TestTraverseFuzzy(t *testing.T) {

}
