# Tetris Optimizer

Tetris Optimizer is a high-performance Go-based solver designed to arrange a list of tetrominoes into the smallest possible square. The project utilizes bitboard-accelerated backtracking with smart pruning to handle complex configurations efficiently.

## Features
- **Optimal Solving**: Finds the absolute smallest square area for any given set of tetrominoes.
- **Bitboard Efficiency**: Uses row-based bitboards (`[16]uint16`) for lightning-fast placement checks.
- **Robust Validation**: Strictly validates input files, ensuring connectivity and correct 4x4 formatting for each piece.
- **Audit Ready**: Meets all requirements for the Tetris Optimizer project, including correct handling of required spaces and uppercase Latin letter labeling.

## Installation

Ensure you have [Go](https://go.dev/dl/) installed on your system.

```bash
# Clone the repository
git clone https://01.tomorrow-school.ai/git/azhakysh/tetris-optimizer.git
cd tetris-optimizer

# Build the executable
go build -o tetris-optimizer .
```

## Usage

The program expects a single argument: the path to a text file containing the tetrominoes.

```bash
./tetris-optimizer sample.txt
```

### Input File Format
Each tetromino must be defined in a 4x4 grid of `.` (empty) and `#` (block), separated by a newline.

Example `sample.txt`:
```
#...
#...
#...
#...

....
....
..##
..##
```

### Output
The program prints the smallest square found, with each tetromino represented by a unique uppercase letter based on its order in the input file.

```
ABB.
ABB.
A...
A...
```

## Testing

The project includes unit tests for the core logic.

```bash
go test -v ./...
```

## Project Objectives
This project was developed to explore:
- Efficient backtracking algorithms.
- Bitwise operations and bitboards.
- File I/O and strict data validation in Go.
- Standard Go project structure and best practices.

## License
MIT
