package game

import (
	"fmt"
	"math"
)

type IddfsSolver struct {
	game *Game
	maxDepth int
	maxVisit int
	minScoreState *State
}

func CreateIddfsSolver(boardString string, w int, h int, maxDepth int, maxVisit int) *IddfsSolver {
	solver := IddfsSolver{CreateGame(boardString, w, h), maxDepth, maxVisit,  nil}
	return &solver
}

func (this *IddfsSolver)dfs(state *State, depth int, currentMaxDepth int, visited *int, minScore *int, maxScore *int) string {

	if (state.IsGoal()) {
		fmt.Printf("goal found at depth %d\n", depth)
		return state.movesString()
	}

	if (state.Score() > currentMaxDepth) {
		return ""
	}

	if (depth > currentMaxDepth) {
		return ""
	}

	if (*visited > this.maxVisit) {
		return ""
	}


	*visited ++

	if (*maxScore < state.addedCost) {
		*maxScore = state.addedCost
	}

	if (state.addedCost == *minScore) {
		if len(this.minScoreState.moves) > len(state.moves) {
			this.minScoreState = state
		}
	}
	if (state.addedCost < *minScore) {
		*minScore = state.addedCost
		this.minScoreState = state
	} 

	var s string = ""
	nexts := state.NextStates()

	for _, next := range nexts {
		s = this.dfs(next, depth +1, currentMaxDepth, visited, minScore, maxScore)
		if s != "" {
			return s
		}
	}
	return ""
}

func (this *IddfsSolver)Run(result chan string) {
	rootState := CreateState(this.game)
	currentMaxDepth := rootState.Score() + 1
	visited := 0
	startState := rootState
	for ;currentMaxDepth < this.maxDepth; {
		minScore := math.MaxInt32
		maxScore := 0
		visited = 0
		s := this.dfs(startState, 0, currentMaxDepth, &visited, &minScore, &maxScore)
		if s != "" { 
			result <- s
			return
		}
		if (visited >= this.maxVisit) {
			if BoardString(startState.board) == BoardString(this.minScoreState.board) {
				fmt.Println("can't purge anymore.")
				result <- ""
				return
			} else {
				fmt.Printf("purge tree. start from minimum score state\n")
				startState = this.minScoreState
				currentMaxDepth = startState.Score() + 1
			}
		} else {
			currentMaxDepth += 2
		}
	}

	result <- ""
}
