package Game

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/saket3199/TicTacToe/cell"
	"github.com/saket3199/TicTacToe/player"
)

// var player1, player2 player.Player

func GetUserName(players *[]player.Player) {

	var p1, p2 string
	fmt.Println("Let's Play Tic Tac Toe!")
	fmt.Println("Player 1, What's Your Name?")
	fmt.Scan(&p1)
label:
	fmt.Println("Player 2, What's Your Name?")
	fmt.Scan(&p2)
	player1 := player.Player{
		PlayerName: p1,
		PlayerTurn: true,
	}
	if p1 != p2 {
		player2 := player.Player{
			PlayerName: p2,
			PlayerTurn: false,
		}
		// player.New(p2, false)
		*players = append(*players, player1)
		*players = append(*players, player2)
		AskPlayerMark(*players)
	} else {
		log.Print("UserName Already Taken.", "Please Choose another name")
		goto label
	}

	// Game.Players = append(game.Players, player1)
	// Game.Players = append(game.Players, player2)
	// game.Players = append(game.Players, player1)
	// game.Players = append(game.Players, player2)
}
func GetBoardSize() uint8 {
	var size uint8
	fmt.Println("Enter Board Size between 3 to 7")
	if _, err := fmt.Scan(&size); err != nil {
		log.Print("  Scan for row failed, due to ", err)
	}
	return size
}

func AskPlayerMark(players []player.Player) {
	fmt.Println("Choose Your Mark :")
	var mark1, mark2 string
player1Mark:
	fmt.Println(players[0].GetPlayerName(), "Enter a Single Alphabet")
	if _, err := fmt.Scan(&mark1); err != nil {
		log.Print("  Scan for row failed, due to ", err)
		goto player1Mark

	}
player2Mark:
	fmt.Println(players[1].GetPlayerName(), "Enter a Single Alphabet")
	if _, err := fmt.Scan(&mark2); err != nil {
		log.Print("  Scan for row failed, due to ", err)
		goto player2Mark

	}
	players[0].SetPlayerMark(string(mark1[0]))
	players[1].SetPlayerMark(string(mark2[0]))

}
func Play() {
	size := GetBoardSize()
	var game *Game
	if size >= 3 && size <= 7 {
		game = New(size)
	} else {
		log.Print("Enter a Board Size between 3 to 7")
		Play()
	}

	var players []player.Player
	turnCount := 0
	// AskPlayerMark(game.Players)
	DrawBoard(game.GetBoard().GetAllCells())
	GetUserName(&players)
	game.Players = append(game.Players, players...)
	fmt.Println(game.Players)
	for {
		DrawBoard(game.GetAllCells())
		WhoseTurn(game.TakeInput(), game.Players)
		for {
			i := game.PutMark(UserPosition(size))
			if i == 1 || i == 2 {
				BoardValidator(i)
			} else {
				break
			}
		}
		game.SetMark()
		turnCount++
		if turnCount > 4 {
			i := game.ResultAnalysis()
			if i == 1 || i == 2 || i == 3 {
				PrintResult(i, game.Players)
				PlayAgain()
				break
			}
		} else {
			game.ChangeTurn()
		}

	}
	DrawBoard(game.GetAllCells())

}

func DrawBoard(board [][]cell.Cell) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Print(board[i][j].GetCellMark())
		}
		fmt.Println()

	}
}

func WhoseTurn(i uint8, p []player.Player) {
	switch i {
	case 1:
		fmt.Println(p[0].GetPlayerName()+"'s Turn ", p[0].GetPlayerMark())
	case 2:
		fmt.Println(p[1].GetPlayerName()+"'s Turn ", p[1].GetPlayerMark())
	}
}

func BoardValidator(i uint8) {
	switch i {
	case 1:
		fmt.Println("This Position is Out of Bounds. Try Again !")
	case 2:
		fmt.Println("Someone has Already Made a Move. Try Again ! ")
	default:
		fmt.Println("Error")
	}
}
func UserPosition(size uint8) []uint8 {
	var row, col uint8
	var pos int
takeUserPosition:
	fmt.Println("Enter a Position to Play: ")
	if _, err := fmt.Scan(&pos); err != nil {
		log.Print("Scan for row failed, due to ", err)
		goto takeUserPosition
	}
	row = uint8(pos / int(size))
	col = uint8(pos) % size

	// for i := 0; i < int(size); i++ {
	// 	for j := 0; j < int(size); j++ {
	// 		if pos == (i*int(size) + j + 1) {
	// 			row = uint8(i)
	// 			col = uint8(j)
	// 		}
	// 	}
	// }
	fmt.Println(row, col)
	// fmt.Println("Enter a row number (0, 1, or 2): ")
	// if _, err := fmt.Scan(&row); err != nil {
	// 	log.Print("  Scan for row failed, due to ", err)
	// }

	// fmt.Println("Enter a column number (0, 1, or 2): ")
	// if _, err := fmt.Scan(&col); err != nil {
	// 	log.Print("  Scan for col failed, due to ", err)
	// }

	positions := []uint8{}
	positions = append(positions, row)
	positions = append(positions, col)

	return positions
}

func PlayAgain() {

	var answer string
	fmt.Println("Do you want to Play Again?")
	fmt.Println("Enter Y to Play Again  or\n Press any other Key to exit")
	fmt.Scan(&answer)
	if answer == "y" || answer == "Y" {
		Play()
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
	}

}
func PrintResult(i uint8, p []player.Player) {
	switch i {
	case 1:
		fmt.Println(p[0].GetPlayerName(), " has Won")
	case 2:
		fmt.Println(p[1].GetPlayerName(), "has Won")
	case 3:
		fmt.Println("Its a Tie ")
	default:
		fmt.Println("Error")
	}
}
