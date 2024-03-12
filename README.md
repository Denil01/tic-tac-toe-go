# Tic Tac Toe Game in Go

This is a simple implementation of the classic Tic Tac Toe game written in Go. It allows two players to play the game from the terminal.

## Features

- Simple terminal-based interface.
- Players take turns making moves.
- The game checks for wins horizontally, vertically, and diagonally.
- Detects a draw when the board is full.

## How It Works

The game is implemented in Go, with different functions handling various aspects of the game:

### `initializeBoard()`

This function initializes the game board with empty spaces and sets the current player to "X".

- **Purpose**: Prepare the game board for a new game session.
- **Functionality**:
  - Iterates over each cell of the 3x3 board.
  - Sets each cell to contain an empty space `" "`.
  - Sets the `currentPlayer` variable to `"X"`.

### `displayBoard()`

This function clears the terminal screen and prints the current state of the game board.

- **Purpose**: Provide a visual representation of the game board to the players.
- **Functionality**:
  - Calls `clearScreen()` to clear the terminal screen.
  - Iterates over each row of the board.
  - Prints the contents of each cell, separated by a pipe character `|`.
  - Prints a horizontal line to separate rows, except for the last row.

### `clearScreen()`

This function clears the terminal screen.

- **Purpose**: Ensure a clean display before printing the updated game board.
- **Functionality**:
  - Uses system-specific commands to clear the terminal screen.

### `makeMove()`

This function prompts the current player to enter their move and updates the game board with the move if it's valid.

- **Purpose**: Allow players to make their moves and update the game state accordingly.
- **Functionality**:
  - Prompts the current player to enter their move, specifying the row and column.
  - Validates the move:
    - Checks if the specified row and column are within the board boundaries.
    - Checks if the specified cell is empty.
  - Updates the board with the player's move if it's valid.

### `checkWin()`

This function checks if there's a winner by examining rows, columns, and diagonals of the game board.

- **Purpose**: Determine if a player has won the game.
- **Functionality**:
  - Checks rows, columns, and diagonals for matching symbols (`"X"` or `"O"`).
  - Returns `true` if any row, column, or diagonal contains three matching symbols.

### `checkDraw()`

This function checks if the game ends in a draw by verifying if the board is full.

- **Purpose**: Determine if the game has ended in a draw.
- **Functionality**:
  - Iterates over each cell of the board.
  - Returns `true` if every cell contains a symbol (i.e., the board is full).

### `switchPlayer()`

This function switches the current player for the next turn.

- **Purpose**: Alternate between players for successive moves.
- **Functionality**:
  - Checks the `currentPlayer` variable.
  - Switches the player from "X" to "O" or from "O" to "X" for the next turn.

## Prerequisites

- Go (version 1.16 or later)

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/your-username/Tic-Tac-Toe.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Tic-Tac-Toe
    ```

3. Run the game:

    ```bash
    go run main.go
    ```

## How to Play

1. The game starts with an empty 3x3 board.
2. Players take turns entering their moves by specifying the row and column (e.g., `row[1-3] col[1-3]`). First input the row and press `ENTER` and then input the column.
3. The game continues until one player wins or the board is full resulting in a draw.
4. The winner is announced at the end of the game.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- This project was inspired by the desire to practice Go programming and create a simple interactive game.
- Thanks to the Go programming language community for their valuable resources and documentation.
