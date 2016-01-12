Fuzzy Radix Tree implemented in Go
==================================

This is an implementation of a Radix tree, which is a compact prefix tree.
There is no delete, it's not a requirement for my implementation.

There is, as well as the standard prefix search, and implementation of a fuzzy
search. The word fuzzy should be used very loosely, it should be considered as
a method for performing non-prefix searches on the trie, with some
simplifications which mean duplicate letters may be acceptable (for example).

Implementation details exist throughout the code.
