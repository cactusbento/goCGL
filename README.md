# Conway's Game of Life

This is an implementation of Conway's Game of Life (CGL) in GO in the terminal.

## Rules

1. Any live cell with fewer than two live neighbors dies
2. Any live cell with two or three live neighbors lives 
3. Any live cell with more than three live neighbors dies
4. Any dead cell with exactly three live neighbors becomes alive

Simplified:
1. Any live cell with two or three live neighbours survives.
2. Any dead cell with three live neighbours becomes a live cell.
3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
