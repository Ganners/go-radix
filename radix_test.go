package radix

import (
	"reflect"
	"strings"
	"testing"
)

func getWikipediaExampleTree() *RadixTree {

	return &RadixTree{
		root: &radixNode{
			children: []*radixNode{
				{
					key: []rune("r"),
					children: []*radixNode{
						{
							key: []rune("om"),
							children: []*radixNode{
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
								{
									key: []rune("ulus"),
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
											key: []rune("r"),
										},
										{
											key: []rune("ns"),
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
	}
}

// I like the ability to visualise, to help explain how things work to
// others and to find bugs that I wouldn't normally think of
func TestDrawVisualisation(t *testing.T) {

	r := getWikipediaExampleTree()

	// A depth-first visualisation of the tree structure
	expect := strings.Join([]string{
		``,
		`[r]`,
		`|`,
		`+- [om]`,
		`   |`,
		`   +- [an]`,
		`      |`,
		`      +- [e]`,
		`      +- [us]`,
		`   +- [ulus]`,
		`+- [ub]`,
		`   |`,
		`   +- [e]`,
		`      |`,
		`      +- [r]`,
		`      +- [ns]`,
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
	}

	// If it doesn't match..
	if r.String() != expected.String() {
		t.Errorf("Result %s does not match expected %s", r.String(), expected.String())
	}
}

// Test to split with a shorter word (must still prefix with letter)
func TestInsertShorter(t *testing.T) {

	r := NewRadixTree()
	r.Add("rabbit", struct{}{})
	r.Add("rabbi", struct{}{})

	// Expected
	expected := &RadixTree{
		root: &radixNode{
			children: []*radixNode{
				{
					key:     []rune("rabbi"),
					content: struct{}{},
					children: []*radixNode{
						{key: []rune("t"), content: struct{}{}},
					},
				},
			},
		},
	}

	// If it doesn't match..
	if r.String() != expected.String() {
		t.Errorf("Result %s does not match expected %s", r.String(), expected.String())
	}
}

// More extreme version of the above test
func TestInsertEvenShorter(t *testing.T) {

	r := NewRadixTree()
	r.Add("rabbit", struct{}{})
	r.Add("rab", struct{}{})

	// Expected
	expected := &RadixTree{
		root: &radixNode{
			children: []*radixNode{
				{
					key:     []rune("rab"),
					content: struct{}{},
					children: []*radixNode{
						{key: []rune("bit"), content: struct{}{}},
					},
				},
			},
		},
	}

	// If it doesn't match..
	if r.String() != expected.String() {
		t.Errorf("Result %s does not match expected %s", r.String(), expected.String())
	}
}

// Test generating the wikipedia example tree
func TestWikipediaExample(t *testing.T) {

	r := NewRadixTree()
	r.Add("romane", struct{}{})
	r.Add("romanus", struct{}{})
	r.Add("romulus", struct{}{})
	r.Add("ruber", struct{}{})
	r.Add("rubens", struct{}{})
	r.Add("rubicon", struct{}{})
	r.Add("rubicundus", struct{}{})

	// Expected
	expected := getWikipediaExampleTree()

	// If it doesn't match..
	if r.String() != expected.String() {
		t.Errorf("Result %s does not match expected %s", r.String(), expected.String())
	}
}

// Test a prefix search
func TestPrefixSearch(t *testing.T) {

	r := NewRadixTree()
	r.Add("romane", struct{}{})
	r.Add("romanus", struct{}{})
	r.Add("romulus", struct{}{})
	r.Add("ruber", struct{}{})
	r.Add("rubens", struct{}{})
	r.Add("rubicon", struct{}{})
	r.Add("rubicundus", struct{}{})

	// Search for 'rom' which is 2 nodes deep
	{
		expected := []string{
			"romane",
			"romanus",
			"romulus",
		}

		res := r.PrefixSearch("rom")

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				res, expected)
		}
	}

	// Search for the two 'rubi' which are 3 nodes deep
	{
		expected := []string{
			"rubicon",
			"rubicundus",
		}

		res := r.PrefixSearch("rubi")

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				res, expected)
		}
	}

	// Searching for 'r' should return everything in this trie
	{
		expected := []string{
			"romane",
			"romanus",
			"romulus",
			"ruber",
			"rubens",
			"rubicon",
			"rubicundus",
		}

		res := r.PrefixSearch("r")

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				res, expected)
		}
	}

	// Empty searches should look at all nodes and return
	{
		expected := []string{
			"romane",
			"romanus",
			"romulus",
			"ruber",
			"rubens",
			"rubicon",
			"rubicundus",
		}

		res := r.PrefixSearch("")

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				res, expected)
		}
	}
}

// Test an in-tree search (non-prefix) with some element of fuzz
func TestTraverseFuzzy(t *testing.T) {

}
