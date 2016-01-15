package radix

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// Used for the data assigned to the struct for something to compare
// against
type identifier struct {
	Id string
}

// Makes sure that the keys returned match the content
func compareKeysAndContent(
	keys []string, content []interface{}, t *testing.T) {

	if len(keys) != len(content) {
		t.Fatalf("Keys (%d) and Content (%d) do not match in length",
			len(keys), len(content))
	}

	for i, _ := range keys {
		ident, ok := content[i].(identifier)
		if !ok {
			t.Errorf("Key %s did not return valid content", keys[i])
		}
		if keys[i] != ident.Id {
			t.Errorf("Keys and content do not match. (Keys: %+v Content: %+v)",
				keys, content)
		}
	}
}

// Returns a pre-built version of the Wikipedia example tree
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
											key:       []rune("e"),
											content:   identifier{"romane"},
											doCollect: true,
										},
										{
											key:       []rune("us"),
											content:   identifier{"romanus"},
											doCollect: true,
										},
									},
								},
								{
									key:       []rune("ulus"),
									content:   identifier{"romulus"},
									doCollect: true,
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
											key:       []rune("r"),
											content:   identifier{"ruber"},
											doCollect: true,
										},
										{
											key:       []rune("ns"),
											content:   identifier{"rubens"},
											doCollect: true,
										},
									},
								},
								{
									key: []rune("ic"),
									children: []*radixNode{
										{
											key:       []rune("on"),
											content:   identifier{"rubicon"},
											doCollect: true,
										},
										{
											key:       []rune("undus"),
											content:   identifier{"rubicundus"},
											doCollect: true,
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
	r.Add("chocolate", identifier{"chocolate"})
	r.Add("pizza", identifier{"pizza"})

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
	r.Add("magazine", identifier{"magazine"})
	r.Add("magsafe", identifier{"magsafe"})

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
	r.Add("rabbit", identifier{"rabbit"})
	r.Add("rabbi", identifier{"rabbi"})

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
	r.Add("rabbit", identifier{"rabbit"})
	r.Add("rab", identifier{"rab"})

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
	r.Add("romane", identifier{"romane"})
	r.Add("romanus", identifier{"romanus"})
	r.Add("romulus", identifier{"romulus"})
	r.Add("ruber", identifier{"ruber"})
	r.Add("rubens", identifier{"rubens"})
	r.Add("rubicon", identifier{"rubicon"})
	r.Add("rubicundus", identifier{"rubicundus"})

	// Expected
	expected := getWikipediaExampleTree()

	// If it doesn't match..
	if r.String() != expected.String() {
		t.Errorf("Result %s does not match expected %s", r.String(), expected.String())
	}
}

// Test a prefix search
func TestPrefixSearch(t *testing.T) {

	// Grab pre-created tree
	r := NewRadixTree()
	r.Add("romane", identifier{"romane"})
	r.Add("romanus", identifier{"romanus"})
	r.Add("romulus", identifier{"romulus"})
	r.Add("ruber", identifier{"ruber"})
	r.Add("rubens", identifier{"rubens"})
	r.Add("rubicon", identifier{"rubicon"})
	r.Add("rubicundus", identifier{"rubicundus"})

	// Search for 'rom' which is 2 nodes deep
	{
		expected := []string{
			"romane",
			"romanus",
			"romulus",
		}

		keys, content := r.PrefixSearch("rom")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
	}

	// Search for the two 'rubi' which are 3 nodes deep
	{
		expected := []string{
			"rubicon",
			"rubicundus",
		}

		keys, content := r.PrefixSearch("rubi")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
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

		keys, content := r.PrefixSearch("r")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
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

		keys, content := r.PrefixSearch("")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
	}
}

// In the case we input 'rabbit' and then 'rabbi', we should be able to
// search for both
func TestPrefixSearchDiffByOne(t *testing.T) {

	r := NewRadixTree()
	r.Add("rabbit", identifier{"rabbit"})
	r.Add("rabbi", identifier{"rabbi"})

	expected := []string{
		"rabbi",
		"rabbit",
	}

	keys, content := r.PrefixSearch("rab")

	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("Prefix result %+v does not matched expected %+v",
			keys, expected)
	}

	compareKeysAndContent(keys, content, t)
}

// Test the bit masks propogate down the nodes as we would expect
func TestAddBitMaskSet(t *testing.T) {

	r := NewRadixTree()
	root := r.Add("november", identifier{"november"})
	r.Add("nova", identifier{"nova"})
	r.Add("niagra falls", identifier{"falls"})
	r.Add("noel", identifier{"noel"})

	// The first node should the letters below n
	{
		expected := genBitMask([]rune{
			'n', 'o', 'v', 'e', 'm', 'b', 'r',
			'a',
			'i', 'g', ' ', 'f', 'l', 's'})

		if root.BitMask() != expected {
			expectedStr := strconv.FormatInt(int64(expected), 2)
			gotStr := strconv.FormatInt(int64(root.BitMask()), 2)
			t.Errorf("Bitmask %s did not match %s",
				gotStr, expectedStr)
		}
	}

	// The letters below 'o' should contain everything but those in
	// niagra falls
	{
		node := root.Children()[0]

		expected := genBitMask([]rune{
			'o', 'v', 'e', 'm', 'b', 'r',
			'a', 'l'})

		if node.BitMask() != expected {
			expectedStr := strconv.FormatInt(int64(expected), 2)
			gotStr := strconv.FormatInt(int64(node.BitMask()), 2)
			t.Errorf("Bitmask %s did not match %s",
				gotStr, expectedStr)
		}
	}

	// The next node down should be a v, which drops the l and o
	{
		node := root.Children()[0].Children()[0]

		expected := genBitMask([]rune{
			'v', 'e', 'm', 'b', 'r', 'a'})

		if node.BitMask() != expected {
			expectedStr := strconv.FormatInt(int64(expected), 2)
			gotStr := strconv.FormatInt(int64(node.BitMask()), 2)
			t.Errorf("Bitmask %s did not match %s",
				gotStr, expectedStr)
		}
	}

	// The next node down should be a ember, which has no children
	{
		node := root.Children()[0].Children()[0].Children()[0]

		expected := genBitMask([]rune{
			'e', 'm', 'b', 'r'})

		if node.BitMask() != expected {
			expectedStr := strconv.FormatInt(int64(expected), 2)
			gotStr := strconv.FormatInt(int64(node.BitMask()), 2)
			t.Errorf("Bitmask %s did not match %s",
				gotStr, expectedStr)
		}
	}

}

// Test an in-tree search (non-prefix) with some element of fuzz
func TestFuzzySearch(t *testing.T) {

	// Grab pre-created tree
	r := NewRadixTree()
	r.Add("romane", identifier{"romane"})
	r.Add("romanus", identifier{"romanus"})
	r.Add("romulus", identifier{"romulus"})
	r.Add("ruber", identifier{"ruber"})
	r.Add("rubens", identifier{"rubens"})
	r.Add("rubicon", identifier{"rubicon"})
	r.Add("rubicundus", identifier{"rubicundus"})

	// Search for 'us', which is contained at the end of 3
	{
		expected := []string{
			"romanus",
			"romulus",
			"rubens", // <-- Fuzzy!
			"rubicundus",
		}

		keys, content := r.FuzzySearch("us")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
	}

	// Search for 'an' which is contained in 2
	{
		expected := []string{
			"romane",
			"romanus",
		}

		keys, content := r.FuzzySearch("an")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
	}

	// Search for 'rubicundus' which is an exact match
	{
		expected := []string{
			"rubicundus",
		}

		keys, content := r.FuzzySearch("rubicundus")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
	}

	// Empty search should return zilch
	{
		expected := []string{}

		keys, content := r.FuzzySearch("")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
	}

	// r search should return all
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

		keys, content := r.FuzzySearch("r")

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Prefix result %+v does not matched expected %+v",
				keys, expected)
		}

		compareKeysAndContent(keys, content, t)
	}
}

// Checks if a given string is contained within a result
func resultsShouldContain(res []string, contain string) bool {

	for _, str := range res {
		if str == contain {
			return true
		}
	}

	return false
}

// Run some integration fuzzy searches against some expectations from
// our test_tree
func TestFuzzyIntegration(t *testing.T) {

	testCases := []struct {
		Search string
		Expect string
	}{
		{
			Search: "som",
			Expect: "somerset road, royal borough of kingston upon thames",
		},
		{
			Search: "kingston",
			Expect: "livesey close, royal borough of kingston upon thames",
		},
		{
			Search: "pablo iglesia",
			Expect: "avenida de pablo iglesias, alcobendas",
		},
		{
			Search: "madrid",
			Expect: "calle de berástegui, pueblo nuevo, madrid",
		},
	}

	r := buildIntegrationTree()

	for _, test := range testCases {
		res, _ := r.FuzzySearch(test.Search)
		if !resultsShouldContain(res, test.Expect) {
			t.Errorf("Search '%s' did not contain '%s'",
				test.Search,
				test.Expect)
		}
	}
}
