package gamerules

import "fmt"

func GetWaterResources(numberOfPlayers int) (int, error) {

	if numberOfPlayers < 3 {
		return 0, fmt.Errorf("not enough players to start the game")
	}

	switch numberOfPlayers {
	case 3:
		return 6, nil
	case 4:
		return 8, nil
	case 5:
		return 10, nil
	case 6:
		return 12, nil
	case 7:
		return 14, nil
	case 8:
		return 16, nil
	case 9:
		return 18, nil
	case 10:
		return 20, nil
	case 11:
		return 22, nil
	case 12:
		return 24, nil
	default:
		return 0, fmt.Errorf("something is wrong with the number of players")
	}
}
