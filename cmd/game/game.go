package game

import "fmt"

type Game struct {
	Player1 *Player
	Player2 *Player
	Winner  *Player
	// Board   [10][10]int
}

type Player struct {
	Name  string
	Board [10][10]int
	Ships []Ship
}

type Ship struct {
	X    int
	Y    int
	Size int
	Name string
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}

func NewGame(player1 *Player, player2 *Player) *Game {
	return &Game{
		Player1: player1,
		Player2: player2,
	}
}

func (g *Game) CreateAndPlaceShip(player *Player, nameOfShip string, x, y, size int, horizontal bool) (string, error) {
	if size > 5 {
		return "", fmt.Errorf("How big do you want your ship please?")
	} else if size < 1 {
		size = 1
	}

	if len(player.Ships) == 5 {
		return "", fmt.Errorf("Maximum number amount of ships reached!")
	}

	ship := Ship{
		X:    x,
		Y:    y,
		Size: size,
		Name: nameOfShip,
	}

	if y > 5 && size > 4 {
		return "", fmt.Errorf("Your ship is too big for the space")
	} else if x > 5 && size > 4 {
		return "", fmt.Errorf("Your ship is too big for the space")
	}

	for i := 0; i < size; i++ {
		if horizontal {
			player.Board[x+i][y] = 1
		} else {
			player.Board[x][y+i] = 1
		}
	}

	player.Ships = append(player.Ships, ship)

	return fmt.Sprintf("%s is placed", nameOfShip), nil
}
