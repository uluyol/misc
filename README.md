
About
===
This repository contains toy implementations of various things for personal interest. I would not recommend using anything here for anything important though everything is under the MIT license.

Incomplete table with descriptions

| File           | Description |
|:-------------- |:---------------------------------------------------------- |
| `self-repl.py`   | A self-replicating python script |
| `self-repl.c`    | A C program which will generate a self-replicating C program. The generated C program is identical to the original except for whitespace differences and the way the string is constructed.
| `self-repl-mem.c` | Identical to `self-repl.c` except this contains a memory indicating which "generation" the current copy of the program is from. The initial value is -1 so that the first "truly replicating" program starts from 0 (ie. `self-repl-mem.c` is the parent of the 0th generation of self-replicating programs).
| `switch-card-game.go` | Simulate (badly) the game where under one card is a prize, and the player is allowed to make a choice, then an incorrect card is removed, and the player is allowed to switch. |
| `quicksort.c` | Quicksort implementation |
