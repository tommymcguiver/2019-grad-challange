package scorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	type args struct {
		who         Customer
		division    Division
		game        int
		matches     []int
		gameNumbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Basic",
			args{
				Customer("John"),
				3,
				1,
				[]int{7, 24, 33, 40},
				[]int{7, 9, 13, 24, 33, 40},
			},
			"John wins a division 3 on game #1 with matches [7 24 33 40] in game [7 9 13 24 33 40]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Display(tt.args.who, tt.args.division, tt.args.game, tt.args.matches, tt.args.gameNumbers)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestScore(t *testing.T) {
	type args struct {
		ticket         Ticket
		winningNumbers WinningNumbers
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Basic",
			args{
				Ticket{
					"John",
					[]Game{
						{7, 9, 13, 24, 33, 40},
						{16, 19, 22, 29, 31, 39},
						{1, 7, 18, 22, 30, 36},
					},
				},
				WinningNumbers{
					7, 22, 24, 31, 33, 40,
				},
			},
			[]string{
				"John wins a division 3 on game #1 with matches [7 24 33 40] in game [7 9 13 24 33 40]",
			},
		},
		{
			"Mary",
			args{
				Ticket{
					"Mary",
					[]Game{
						{2, 22, 13, 24, 32, 39},
						{7, 22, 24, 31, 33, 40},
						{3, 7, 18, 21, 37, 38},
					},
				},
				WinningNumbers{
					7, 22, 24, 31, 33, 40,
				},
			},
			[]string{
				"Mary wins a division 1 on game #2 with matches [7 22 24 31 33 40] in game [7 22 24 31 33 40]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ticket.Score(tt.args.winningNumbers)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMatches(t *testing.T) {
	type args struct {
		game           Game
		winningNumbers WinningNumbers
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Basic",
			args{
				Game{7, 9, 13, 24, 33, 40},
				WinningNumbers{7, 22, 24, 31, 33, 40},
			},
			[]int{7, 24, 33, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.game.Matches(tt.args.winningNumbers)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestScoreSystem(t *testing.T) {
	type args struct {
		ticket         SystemTicket
		winningNumbers WinningNumbers
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Basic",
			args{
				SystemTicket{
					"Jack",
					Game{3, 5, 7, 11, 22, 24, 31, 34, 40},
				},
				WinningNumbers{7, 22, 24, 31, 33, 40},
			},
			[]string{
				"Jack wins a division 4 on game #1 with matches [7 22 24] in game [3 5 7 11 22 24]",
				"Jack wins a division 4 on game #2 with matches [7 22 31] in game [3 5 7 11 22 31]",
				"Jack wins a division 4 on game #4 with matches [7 22 40] in game [3 5 7 11 22 40]",
				"Jack wins a division 4 on game #5 with matches [7 24 31] in game [3 5 7 11 24 31]",
				"Jack wins a division 4 on game #7 with matches [7 24 40] in game [3 5 7 11 24 40]",
				"Jack wins a division 4 on game #9 with matches [7 31 40] in game [3 5 7 11 31 40]",
				"Jack wins a division 3 on game #11 with matches [7 22 24 31] in game [3 5 7 22 24 31]",
				"Jack wins a division 4 on game #12 with matches [7 22 24] in game [3 5 7 22 24 34]",
				"Jack wins a division 3 on game #13 with matches [7 22 24 40] in game [3 5 7 22 24 40]",
				"Jack wins a division 4 on game #14 with matches [7 22 31] in game [3 5 7 22 31 34]",
				"Jack wins a division 3 on game #15 with matches [7 22 31 40] in game [3 5 7 22 31 40]",
				"Jack wins a division 4 on game #16 with matches [7 22 40] in game [3 5 7 22 34 40]",
				"Jack wins a division 4 on game #17 with matches [7 24 31] in game [3 5 7 24 31 34]",
				"Jack wins a division 3 on game #18 with matches [7 24 31 40] in game [3 5 7 24 31 40]",
				"Jack wins a division 4 on game #19 with matches [7 24 40] in game [3 5 7 24 34 40]",
				"Jack wins a division 4 on game #20 with matches [7 31 40] in game [3 5 7 31 34 40]",
				"Jack wins a division 4 on game #21 with matches [22 24 31] in game [3 5 11 22 24 31]",
				"Jack wins a division 4 on game #23 with matches [22 24 40] in game [3 5 11 22 24 40]",
				"Jack wins a division 4 on game #25 with matches [22 31 40] in game [3 5 11 22 31 40]",
				"Jack wins a division 4 on game #28 with matches [24 31 40] in game [3 5 11 24 31 40]",
				"Jack wins a division 4 on game #31 with matches [22 24 31] in game [3 5 22 24 31 34]",
				"Jack wins a division 3 on game #32 with matches [22 24 31 40] in game [3 5 22 24 31 40]",
				"Jack wins a division 4 on game #33 with matches [22 24 40] in game [3 5 22 24 34 40]",
				"Jack wins a division 4 on game #34 with matches [22 31 40] in game [3 5 22 31 34 40]",
				"Jack wins a division 4 on game #35 with matches [24 31 40] in game [3 5 24 31 34 40]",
				"Jack wins a division 3 on game #36 with matches [7 22 24 31] in game [3 7 11 22 24 31]",
				"Jack wins a division 4 on game #37 with matches [7 22 24] in game [3 7 11 22 24 34]",
				"Jack wins a division 3 on game #38 with matches [7 22 24 40] in game [3 7 11 22 24 40]",
				"Jack wins a division 4 on game #39 with matches [7 22 31] in game [3 7 11 22 31 34]",
				"Jack wins a division 3 on game #40 with matches [7 22 31 40] in game [3 7 11 22 31 40]",
				"Jack wins a division 4 on game #41 with matches [7 22 40] in game [3 7 11 22 34 40]",
				"Jack wins a division 4 on game #42 with matches [7 24 31] in game [3 7 11 24 31 34]",
				"Jack wins a division 3 on game #43 with matches [7 24 31 40] in game [3 7 11 24 31 40]",
				"Jack wins a division 4 on game #44 with matches [7 24 40] in game [3 7 11 24 34 40]",
				"Jack wins a division 4 on game #45 with matches [7 31 40] in game [3 7 11 31 34 40]",
				"Jack wins a division 3 on game #46 with matches [7 22 24 31] in game [3 7 22 24 31 34]",
				"Jack wins a division 2 on game #47 with matches [7 22 24 31 40] in game [3 7 22 24 31 40]",
				"Jack wins a division 3 on game #48 with matches [7 22 24 40] in game [3 7 22 24 34 40]",
				"Jack wins a division 3 on game #49 with matches [7 22 31 40] in game [3 7 22 31 34 40]",
				"Jack wins a division 3 on game #50 with matches [7 24 31 40] in game [3 7 24 31 34 40]",
				"Jack wins a division 4 on game #51 with matches [22 24 31] in game [3 11 22 24 31 34]",
				"Jack wins a division 3 on game #52 with matches [22 24 31 40] in game [3 11 22 24 31 40]",
				"Jack wins a division 4 on game #53 with matches [22 24 40] in game [3 11 22 24 34 40]",
				"Jack wins a division 4 on game #54 with matches [22 31 40] in game [3 11 22 31 34 40]",
				"Jack wins a division 4 on game #55 with matches [24 31 40] in game [3 11 24 31 34 40]",
				"Jack wins a division 3 on game #56 with matches [22 24 31 40] in game [3 22 24 31 34 40]",
				"Jack wins a division 3 on game #57 with matches [7 22 24 31] in game [5 7 11 22 24 31]",
				"Jack wins a division 4 on game #58 with matches [7 22 24] in game [5 7 11 22 24 34]",
				"Jack wins a division 3 on game #59 with matches [7 22 24 40] in game [5 7 11 22 24 40]",
				"Jack wins a division 4 on game #60 with matches [7 22 31] in game [5 7 11 22 31 34]",
				"Jack wins a division 3 on game #61 with matches [7 22 31 40] in game [5 7 11 22 31 40]",
				"Jack wins a division 4 on game #62 with matches [7 22 40] in game [5 7 11 22 34 40]",
				"Jack wins a division 4 on game #63 with matches [7 24 31] in game [5 7 11 24 31 34]",
				"Jack wins a division 3 on game #64 with matches [7 24 31 40] in game [5 7 11 24 31 40]",
				"Jack wins a division 4 on game #65 with matches [7 24 40] in game [5 7 11 24 34 40]",
				"Jack wins a division 4 on game #66 with matches [7 31 40] in game [5 7 11 31 34 40]",
				"Jack wins a division 3 on game #67 with matches [7 22 24 31] in game [5 7 22 24 31 34]",
				"Jack wins a division 2 on game #68 with matches [7 22 24 31 40] in game [5 7 22 24 31 40]",
				"Jack wins a division 3 on game #69 with matches [7 22 24 40] in game [5 7 22 24 34 40]",
				"Jack wins a division 3 on game #70 with matches [7 22 31 40] in game [5 7 22 31 34 40]",
				"Jack wins a division 3 on game #71 with matches [7 24 31 40] in game [5 7 24 31 34 40]",
				"Jack wins a division 4 on game #72 with matches [22 24 31] in game [5 11 22 24 31 34]",
				"Jack wins a division 3 on game #73 with matches [22 24 31 40] in game [5 11 22 24 31 40]",
				"Jack wins a division 4 on game #74 with matches [22 24 40] in game [5 11 22 24 34 40]",
				"Jack wins a division 4 on game #75 with matches [22 31 40] in game [5 11 22 31 34 40]",
				"Jack wins a division 4 on game #76 with matches [24 31 40] in game [5 11 24 31 34 40]",
				"Jack wins a division 3 on game #77 with matches [22 24 31 40] in game [5 22 24 31 34 40]",
				"Jack wins a division 3 on game #78 with matches [7 22 24 31] in game [7 11 22 24 31 34]",
				"Jack wins a division 2 on game #79 with matches [7 22 24 31 40] in game [7 11 22 24 31 40]",
				"Jack wins a division 3 on game #80 with matches [7 22 24 40] in game [7 11 22 24 34 40]",
				"Jack wins a division 3 on game #81 with matches [7 22 31 40] in game [7 11 22 31 34 40]",
				"Jack wins a division 3 on game #82 with matches [7 24 31 40] in game [7 11 24 31 34 40]",
				"Jack wins a division 2 on game #83 with matches [7 22 24 31 40] in game [7 22 24 31 34 40]",
				"Jack wins a division 3 on game #84 with matches [22 24 31 40] in game [11 22 24 31 34 40]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ticket.Score(tt.args.winningNumbers)
			assert.Equal(t, tt.want, got)
		})
	}
}
