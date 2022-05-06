package core

import "math"

func (board *Board) NewRoot(depth int) int {
	return board.alphaBetaMax(math.MinInt32, math.MaxInt32, depth)
}

func (board *Board) alphaBetaMax(alpha, beta, depthLeft int) int {
	if depthLeft == 0 {
		return board.evaluate()
	}
	for { // (all moves)
		//board.makeMove()
		score := board.alphaBetaMin(alpha, beta, depthLeft-1)
		//board.unmakeMove()
		if score >= beta {
			return beta
		}
		if score > alpha {
			alpha = score
		}
	}
	return alpha
}

func (board *Board) alphaBetaMin(alpha, beta, depthLeft int) int {
	if depthLeft == 0 {
		return board.evaluate()
	}
	for { // (all moves)
		//board.makeMove()
		score := board.alphaBetaMax(alpha, beta, depthLeft-1)
		//board.unmakeMove()
		if score <= alpha {
			return alpha
		}
		if score < beta {
			beta = score
		}
	}
	return beta
}
