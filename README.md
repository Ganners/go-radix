Fuzzy Radix Tree (PATRICIA) implemented in Go
=============================================

Primary source for knowledge:
[Wikipedia](https://en.wikipedia.org/wiki/Radix_tree)

This is an implementation of a Radix tree, which is a compact prefix tree.
There is no delete, it's not a requirement for my implementation.

The goal is to be able to generate a space and time efficient autocomplete for
large sparse address strings.

Fuzzy matching will be produce with a compressed bitmap (squashing into 32
bits) with lookaheads.

It will be able to output a visualisation, although it won't be recommended for
a trie of any great length, it will help with basic tests.
