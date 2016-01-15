Fuzzy Radix Tree implemented in Go
==================================

[![Build Status](https://travis-ci.org/Ganners/go-radix.svg?branch=master)](https://travis-ci.org/Ganners/go-radix)

This is an implementation of a Radix tree, which is a compact prefix tree.
There is no delete, it's not a requirement for my implementation.

There is, as well as the standard prefix search, and implementation of a fuzzy
search. The word fuzzy should be used very loosely, it should be considered as
a method for performing non-prefix searches on the trie, with some
simplifications which mean duplicate letters may be acceptable (for example).

Implementation details exist throughout the code.

Example Usage
-------------

### Create the tree and insert some keys with content

	r := NewRadixTree()
	r.Add("romane", struct{}{})
	r.Add("romanus", struct{}{})
	r.Add("romulus", struct{}{})
	r.Add("ruber", struct{}{})
	r.Add("rubens", struct{}{})
	r.Add("rubicon", struct{}{})
	r.Add("rubicundus", struct{}{})

### Want to see how it looks?

    fmt.Printf("%s", r.String())

### Example output:

    [r]
    |
    +- [om]
       |
       +- [an]
          |
          +- [e]
          +- [us]
       +- [ulus]
    +- [ub]
       |
       +- [e]
          |
          +- [r]
          +- [ns]
       +- [ic]
          |
          +- [on]
          +- [undus]

### Prefix search

    searchStrings := r.PrefixSearch("rom")

### Fuzzy search

    fuzzyStrings := r.FuzzySearch("us")
