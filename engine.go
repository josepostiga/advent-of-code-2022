package main

const (
	winPoints  = 6
	drawPoints = 3
	losePoints = 0

	rockPoints     = 1
	paperPoints    = 2
	scissorsPoints = 3

	rock     = "A"
	paper    = "B"
	scissors = "C"
)

type PlayResult struct {
	round StrategyPlay
	score int
}

func Play(round StrategyPlay) PlayResult {
	var result PlayResult
	result.round = round
	result.score = calculateScore(round)

	return result
}

func calculateScore(round StrategyPlay) int {
	score := 0

	if isDraw(round.action) {
		score += drawPoints + drawMovePoints(round)
	} else if isWin(round.action) {
		score += winPoints + winMovePoints(round)
	} else {
		score += losePoints + loseMovePoints(round)
	}

	return score
}

func loseMovePoints(round StrategyPlay) int {
	if round.opponent == rock {
		return scissorsPoints
	} else if round.opponent == paper {
		return rockPoints
	} else if round.opponent == scissors {
		return paperPoints
	}
	return 0
}

func winMovePoints(round StrategyPlay) int {
	if round.opponent == rock {
		return paperPoints
	} else if round.opponent == paper {
		return scissorsPoints
	} else if round.opponent == scissors {
		return rockPoints
	}
	return 0
}

func drawMovePoints(round StrategyPlay) int {
	if round.opponent == rock {
		return rockPoints
	} else if round.opponent == paper {
		return paperPoints
	} else if round.opponent == scissors {
		return scissorsPoints
	}
	return 0
}

func isWin(action string) bool {
	if action == "Z" {
		return true
	}

	return false
}

func isDraw(action string) bool {
	if action == "Y" {
		return true
	}
	return false
}
