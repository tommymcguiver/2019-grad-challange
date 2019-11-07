package scorer

import (
	"fmt"
	"sort"

	"github.com/gonum/stat/combin"
)

type Customer string
type Game []int
type WinningNumbers []int
type Division int

const (
	NoDivision Division = iota
	DivisionOne
	DivisionTwo
	DivisionThree
	DivisionFour
)

type Ticket struct {
	customer Customer
	games    []Game
}

type SystemTicket struct {
	customer Customer
	game     Game
}

const numberSetSize = 6
const bufferSize = 100

func (t *Ticket) addGame(g Game) {
	t.games = append(t.games, g)
}

func NewTicket(ticket SystemTicket) Ticket {
	regularTicket := Ticket{ticket.customer, nil}
	regularTicket.games = make([]Game, 0, bufferSize)
	cs := combin.Combinations(len(ticket.game), numberSetSize)
	for _, c := range cs {
		regularTicket.addGame(
			Game{
				ticket.game[c[0]],
				ticket.game[c[1]],
				ticket.game[c[2]],
				ticket.game[c[3]],
				ticket.game[c[4]],
				ticket.game[c[5]],
			},
		)
	}

	return regularTicket
}

func Display(customer Customer, division Division, gameNumber int, matches []int, game Game) string {
	return fmt.Sprintf(
		"%s wins a division %d on game #%d with matches %v in game %v",
		customer,
		division,
		gameNumber,
		matches,
		game,
	)
}

func DivisionNumber(matches int) Division {

	switch {
	case matches == 6:
		return DivisionOne
	case matches == 5:
		return DivisionTwo
	case matches == 4:
		return DivisionThree
	case matches == 3:
		return DivisionFour
	}

	return NoDivision
}

func (g Game) Matches(winningNumbers WinningNumbers) []int {
	matches := make([]int, 0, len(winningNumbers))
	sort.Ints(g)
	for _, number := range g {
		for _, winner := range winningNumbers {
			if number < winner {
				break
			}
			if number == winner {
				matches = append(matches, winner)
			}
		}
	}
	return matches
}

func (t Ticket) Score(winningNumbers WinningNumbers) []string {

	result := make([]string, 0, bufferSize)

	for gameIndex, game := range t.games {
		matches := game.Matches(winningNumbers)
		matchCount := len(matches)

		if matchCount != 0 {
			divisionNumber := DivisionNumber(matchCount)

			if divisionNumber == 0 {
				continue
			}

			result = append(result, Display(t.customer, divisionNumber, gameIndex+1, matches, game))
		}
	}

	return result
}

func (st SystemTicket) Score(winningNumbers WinningNumbers) []string {
	return NewTicket(st).Score(winningNumbers)
}
