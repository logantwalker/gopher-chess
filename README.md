# Gopher Chess - Advanced Golang Chess Engine (1953 ELO)

![gopher chess auto gameplay](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExY2gydDVhaGVrYjd4ZXNtNjF6dXQybjJvdnI2d2UzYmU1bGlhbDRybSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/QXDQ8N8DkuFUlMO9sD/giphy.gif)


Gopher Chess is an advanced Chess Engine developed in Golang, achieving a competitive ELO rating of 1950 on the Lichess platform. The engine uses a 0x88 board representation and incorporates a number of advanced algorithms and techniques such as iterative deepening, Principal Variation Search, Quiescence Search, and Check Extensions for optimal gameplay. Evaluation techniques include Material Balance, Piece-Square Tables, and Doubled Pawn, improving the engine's decision-making process.

You can check out some of the games it has played [here](https://lichess.org/@/gopher307).

[You can play against the engine without downloading the project by clicking here.](https://gopher-frontend-nq3ygk4l7a-uc.a.run.app)

## Installation

```bash
git clone <repository_url>
cd <repository_directory>
go mod download
go mod tidy
```

## Usage

You can interact with Gopher Chess by running the gopher.go file in your terminal. The engine accepts both standard UCI commands and simplified terminal-friendly commands. Here are some of the commands you can use:

### UCI Commands

- **uci** : Prints out the engine's details.
- **isready** : Checks if the engine is ready to receive commands.
- **setoption name Move Overhead value <ms>** : Sets the time overhead for each move.
- **position fen <fen_string>** : Sets the board to the given position represented by the FEN string.
- **position startpos moves e2e4 e7e5...** : Starts a new game and makes the given moves from the starting position.
- **go** : The engine makes the best move it can find.
- **ucinewgame** : Starts a new game.

### Terminal Commands

- **movestring (e.g. e2e4)** : makes the provided move on the board so long as it is a legal move.
- **go** : The engine makes the best move it can find.
- **m or moves** : Prints out all the legal moves in the current position.
- **turn** : Prints out which side's turn it is.
- **perft** : Runs perft testing on preset positions.
- **f or fen** : Prints the current position in FEN format.
- **u or undo** : Undoes the last move.
- **p or print** : Prints out the current board.
- **s or search** : Initiates a search for the best move in the current position.
- **e or eval** : Evaluates the current position.
- **a or auto** : The engine plays against itself automatically.
- **fen <fen_string>** : Sets the board to the given position represented by the FEN string.
- **n** : Starts a new game.
- **quit or q** : Quits the game.

### Example Game

```bash
go run gopher.go
e2e4
making move
go
best move: d7d5
print
a  b  c  d  e  f  g  h
8  ♖  ♘  ♗  ♕  ♔  ♗  ♘  ♖  
7  ♙  ♙  ♙  .  ♙  ♙  ♙  ♙  
6  .  ,  .  ,  .  ,  .  ,  
5  ,  .  ,  ♙* ,  .  ,  .       (2) white's move
4  .  ,  .  ,  ♟  ,  .  ,       Casteling: KQkq
3  ,  .  ,  .  ,  .  ,  .  
2  ♟  ♟  ♟  ♟  .  ♟  ♟  ♟  
1  ♜  ♞  ♝  ♛  ♚  ♝  ♞  ♜  
a  b  c  d  e  f  g  h       d7d5
```
note: the chesspieces may seem to have inverted colors if you play with a white background terminal. This is because the chesspieces are represented by unicode characters that are white in color.
### Online Play

Gopher Chess can be integrated with the Lichess API, enabling it to compete against other players and bots online. 
