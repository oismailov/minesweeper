# Minesweeper

It is a game with command line interface. There is no UI, however, you can still enjoy playing it using your favorite terminal.
## Usage

Open the terminal (command line interface) and run this command:
```bash
go run cmd/main.go
```
- There will be a prompt to enter grid height. Only integers are expected.
```bash
Enter grid height:
```
- Then prompt to enter grid width. Only integers are expected.
```bash
Enter grid width:
```
- Then prompt to enter amount of black holes in grid. Only integers are expected.
```bash
Enter amount of black holes in grid:
```

Once all parameters are set - game will start and there will be an output with initial grid.

CLI input is still active and user can provide coordinates of cell to be revealed.   
Expected format is `x,y`   
The message is:
```bash
 Enter coordinates separated by comma:
```
After each input - new updated grid will be printed until the user win or lose the game.

#### Please enjoy :)

## Tech Stack
- Golang 1.19

## Possible improvements
- There is very little validation of input parameters. Please make sure you provided valid data format
- Input can be scanned by using flags
- Unit tests can be added
- Create UI, deploy to remote server, share with friends and have a fun ;)

## Game Screenshots

Game screenshots can be found in `/screenshots` folder
