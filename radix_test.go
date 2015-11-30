package radix

import (
	"strings"
	"testing"
)

func TestInsert(t *testing.T) {

}

func TestTraversePrefix(t *testing.T) {

}

func TestTraverseFuzzy(t *testing.T) {

}

// I like the ability to visualise, to help explain how things work to
// others and to find bugs that I wouldn't normally think of
func TestDrawVisualisation(t *testing.T) {

	// Example from Wikipedia
	r := NewRadixTree()
	r.Add("romane")
	r.Add("romanus")
	r.Add("romulus")
	r.Add("rubens")
	r.Add("ruber")
	r.Add("rubicon")
	r.Add("rubicundus")

	// Visualise. Prior knowledge to build is the largest sequence length and
	// the maximum depth, then it is just a case of walking the tree breadth
	// first and appending each value (padded) to the string.
	expect := strings.Join([]string{
		`                                    ....r                                  `,
		`                                      ^                                    `,
		`                                     / \                                   `,
		`               ...om                                   ...ub               `,
		`                 ^                                       ^                 `,
		`                / \                                     / \                `,
		`     .ulus                ...an               ....e               ...ic    `,
		`       ^                    ^                   ^                   ^      `,
		`      / \                  / \                 / \                 / \     `,
		`.....     .....     ....e     ...us     ...ns     ....r     ...on     undus`,
	}, "\n")

	toString := r.String()
	if toString != expect {
		t.Errorf(
			"String output did not match expected visualisation. Got\n%s\n",
			toString)
	}
}
