package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	//Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
	Message        string `json:"message"`
}

var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket",
	"You are so lucky on this game!",
}

var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day.",
	"You should think computer next move!",
}

var drawMessages = []string{
	"Greet minds think alike",
	"Uh oh. Try again.",
	"Nobody wins, but you can try again.",
	"It's kinda boring, try win or shut the game!",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	//winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	messageInt := rand.Intn(4)

	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		//winner = DRAW
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		//winner = PLAYERWINS
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins!"
		//winner = COMPUTERWINS
		message = loseMessages[messageInt]
	}

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result

}
