# automatabot
![](https://user-images.githubusercontent.com/212941/59635448-1fbfb700-9106-11e9-8675-21ca2f382258.png)

# 👋 Meet Automatabot

Automatabot was born when an intern accidentally put a ream of graph paper into the punchcard hopper on the Noops AI programmer. Automatabot loves grids, and especially loves to play with [cellular automata](https://en.wikipedia.org/wiki/Cellular_automaton).

Automatabot invites you to try your hand at its many two-dimensional challenges.

There are lots of different types of cellular automata but Automatabot likes to keep it simple and stick with the "Life-like" variations, so-called because they use the same rules framework as [Conway's Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life), the most famous cellular automaton ever created.

The state of each cell in the next generation is computed based on whether it is currently alive and how many of its neighbors are alive.

For example, Conway's Life rules are:

- Any alive cell with fewer than two live neighbors dies.
- Any alive cell with two or three live neighbors lives on to the next generation. ("Survival")
- Any alive cell with more than three live neighbors dies.
- Any dead cell with exactly three live neighbors becomes a live cell. ("Birth")

Automatabot has a library of automaton rules and is excited to share them with you.

## ✳️ How to play

Get a challenge from Automatabot, and once you've figured out the answer, send it back to Automatabot for verification.


# 🤖 API

Automatabot

Automatabot has the following endpoints:

`GET /automatabot/rules`

Get a list of all the rules that automatabot knows about.

`GET /automatabot/challenges/new`

Create a new challenge to solve.

`POST /automatabot/challenges/{id}`

Send your solution to automatabot for checking.

`GET /automatabot/challenges/{id}`

Get the same challenge again, using the path provided by `automatabot/challenged/new`

`GET /automatabot/rules/{rulename}/random`

Generate a random setup with the provided ruleset.

Each challenge has a ruleset, an initial grid of cells, and the number of generations of evolution to apply to get the answer.
Automatabot represents a grid of cells as an array of arrays, where `1` represents an alive cell, and `0` is a dead one:

An example challenge with the Conway's Life ruleset looks like this:
```
{
  "rules": {
     "name": "conway"
     "birth": [3],
     "survival": [2, 3],
   },
  "cells": [
     [ 1, 1, 1, 0, 0 ],
     [ 1, 0, 1, 0, 0 ],
     [ 0, 1, 0, 1, 1 ],
     [ 0, 0, 1, 1, 0 ],
     [ 0, 0, 0, 1, 0 ]
  ],
  generations: 12
}
```

Each of Automatabot's challenges has a fixed width and height. Cells that are outside the boundaries are considered dead and the size of the grid does not grow. The size of the grid changes with each challenge.


Once you have the solution, POST it back to the provided `challengePath` to check your work.

# 🖼️ Starter

A [go example client](./starters/automatabot.go) is included to help you get started. This client fetches a challenge and prints out the grid of cells.

Run it with `go run automatabot.go`

It's up to you to add the logic to compute the result of the challenge and send it back to the API.

# ✨ A few ideas

- **Keep going:** Each challenge has a fixed number of generations, but many of these rulesets will generate patterns that repeat forever. Try to keep going until a static pattern emerges. Does a pattern always emerge? Why or why not?

- **Animate it:** Create a visualization with these patterns. Incorporate color. Try changing the colors of cells based on their position in the grid or how long they have been alive.

- **Make it boundless:** Although these grids have a fixed size, some automata will grow larger and larger with each generation. Try to support either an unbounded grid of cells, or one that wraps around from left to right and top to bottom.

More about Automatabot on [noopschallenge.com](https://noopschallenge.com/challenges/automatabot)

See the Golfbot solutions to Conway's Game of Life (written in less than 256 characters) in [this repository](https://github.com/noops-challenge/golfbot/tree/master/challenges/conways-life).