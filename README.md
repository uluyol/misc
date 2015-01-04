
About
===
This repository contains toy implementations of various things for personal interest. I would not recommend using anything here for anything important though everything is under the MIT license.

Incomplete table with descriptions:

| File           | Description |
|:-------------- |:----------- |
| bit-manipulation.c | Fun with bits in C. Implemented unsigned long printing with base 2 and insertion of N into M's [i,j]th bits. This assumes that N fits within those bits. |
| bloom.cc | Basic bloom filter implementation in C++ |
| bloom.go | Basic bloom filter implementation in Go |
| bst.go | Basic binary search tree implementation in Go |
| cache.go | LRU Cache implementation in Go (hashtable and doubly linked lists) |
| data-structs.c  | Implementation of a stack and hashtable in C |
| data-structs.cc | Implementation of a stack and hashtable in C++ |
| huffman.go | Basic Huffman Coding |
| inverted_index.cc | An interactive program to query for files that contain certain text (exact string matches) |
| malloc.go | A malloc implementation (maybe?) in Go. Please **don't** use this. |
| self-repl.py   | A self-replicating python script |
| self-repl.c    | A C program which will generate a self-replicating C program. The generated C program is identical to the original except for whitespace differences and the way the string is constructed.
| self-repl-mem.c | Identical to self-repl.c except this contains a memory indicating which "generation" the current copy of the program is from. The initial value is -1 so that the first "truly replicating" program starts from 0 (ie. self-repl-mem.c is the parent of the 0th generation of self-replicating programs).
| switch-card-game.go | Simulate (badly) the game where under one card is a prize, and the player is allowed to make a choice, then an incorrect card is removed, and the player is allowed to switch. |
| quicksort.c | Quicksort implementation |
