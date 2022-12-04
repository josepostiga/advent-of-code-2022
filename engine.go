package main

const (
	winPoints  = 6
	drawPoints = 3
	losePoints = 0

	rockPoints     = 1
	paperPoints    = 2
	scissorsPoints = 3

	rockOpponent     = "A"
	rockMe           = "X"
	paperOpponent    = "B"
	paperMe          = "Y"
	scissorsOpponent = "C"
	scissorsMe       = "Z"
)

type PlayResult struct {
	round      StrategyPlay
	movePoints int
	points     int
	score      int
}

func Play(round StrategyPlay) PlayResult {
	var result PlayResult
	result.round = round
	result.movePoints = calculateMovePoints(round.me)

	if isDraw(round) {
		result.points = drawPoints
	} else if isWin(round) {
		result.points = winPoints
	} else {
		result.points = losePoints
	}

	result.score = result.movePoints + result.points

	return result
}

func isWin(round StrategyPlay) bool {
	if round.opponent == rockOpponent && round.me == paperMe {
		return true
	} else if round.opponent == scissorsOpponent && round.me == rockMe {
		return true
	} else if round.opponent == paperOpponent && round.me == scissorsMe {
		return true
	}
	return false
}

func isDraw(round StrategyPlay) bool {
	if round.opponent == rockOpponent && round.me == rockMe {
		return true
	} else if round.opponent == scissorsOpponent && round.me == scissorsMe {
		return true
	} else if round.opponent == paperOpponent && round.me == paperMe {
		return true
	}
	return false
}

func calculateMovePoints(move string) int {
	switch move {
	case paperMe:
		return paperPoints
	case rockMe:
		return rockPoints
	case scissorsMe:
		return scissorsPoints
	default:
		return 0
	}
}
