package radix

import "testing"

func BenchmarkBuildTrie(b *testing.B) {

	for i := 0; i < b.N; i++ {
		prebuiltIntegrationTree = nil
		buildIntegrationTree()
	}
}

// Benchmarks a prefix search for 'Som'
func BenchmarkPrefixSom(b *testing.B) {

	trie := buildIntegrationTree()

	for i := 0; i < b.N; i++ {
		trie.PrefixSearch("som")
	}
}

// Benchmarks a fuzzy search for 'Som'
func BenchmarkFuzzySom(b *testing.B) {

	trie := buildIntegrationTree()

	for i := 0; i < b.N; i++ {
		trie.FuzzySearch("som")
	}
}

// Benchmarks a prefix search for 'Somer'
func BenchmarkPrefixSomer(b *testing.B) {

	trie := buildIntegrationTree()

	for i := 0; i < b.N; i++ {
		trie.PrefixSearch("somer")
	}
}

// Benchmarks a fuzzy search for 'Somer'
func BenchmarkFuzzySomer(b *testing.B) {

	trie := buildIntegrationTree()

	for i := 0; i < b.N; i++ {
		trie.FuzzySearch("somer")
	}
}
