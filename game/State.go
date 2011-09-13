package game

import "container/vector"

type State struct {
	board [] int
	moves vector.IntVector 
	game *Game
	addedCost int
}

func calcCost(value int, index int, game *Game) int {
	for i,x := range game.goalBoard {
		if x == value {
			return game.distances[i][index]
		}
	}
	return -1 
}

func (state *State)LastMove() int {
	if len(state.moves) == 0 {
		return 0
	}
	return state.moves.Last()
}

func (state *State)movesString() string {
	ret := make([]int,0,len(state.moves))
	for _,x := range state.moves {
		switch x {
		case -1:
			ret = append(ret, 'L')
		case 1:
			ret = append(ret, 'R')
		case state.game.w:
			ret = append(ret, 'D')
		case -state.game.w:
			ret = append(ret, 'U')
		}
	}
	return string(ret)
}

func (state *State)Score() int {
	return len(state.moves) + state.addedCost
}

func (state *State)calcAllCosts() []int {
	costs := make([]int,len(state.board))
	state.addedCost = 0
	for index, value := range state.board {
		if value > 0 && value < 36 {
			costs[index] = calcCost(value, index, state.game)
			state.addedCost += costs[index]
		}
	}
	return costs
}

func CreateState(game *Game) *State {
	s := State{game.rootBoard, []int{}, game, 0}
	s.calcAllCosts()
	return &s
}

func (s *State)Copy() *State {
	return &State{s.board, s.moves.Copy(), s.game, s.addedCost}
}

func (state *State)IsGoal() bool {
	return state.addedCost == 0
}

func (state *State)StepBack() *State {
	indexOfSpace := -1
	for i,x := range state.board {
		if (x == 36) {
			indexOfSpace = i
			break
		}
	}
	nextIndex := indexOfSpace - state.LastMove()
	newBoard := make([]int, len(state.board))
	copy(newBoard, state.board)
	newBoard[indexOfSpace], newBoard[nextIndex] = newBoard[nextIndex],newBoard[indexOfSpace]
	newMoves := make([]int,len(state.moves)-1)
	for i := 0; i < len(state.moves) -1; i ++ {
		newMoves[i] = state.moves[i]
	}
	nextState := State{newBoard, newMoves, state.game, 0}
	nextState.calcAllCosts()
	return &nextState
}

func (state *State)NextStates() [] *State {
	indexOfSpace := -1
	for i,x := range state.board {
		if (x == 36) {
			indexOfSpace = i
			break
		}
	}
	ret := make([]*State,0)
	for _, next := range state.game.nexts(indexOfSpace) {
		if next != - state.LastMove() {
			nextIndex := indexOfSpace + next
			newBoard := make([]int, len(state.board))
			copy(newBoard, state.board)
			newBoard[indexOfSpace], newBoard[nextIndex] = newBoard[nextIndex],newBoard[indexOfSpace]
			newMoves := make([]int,len(state.moves)+1)
			for i,x := range state.moves {
				newMoves[i] = x
			}
			newMoves[len(state.moves)] = next
			nextState := State{newBoard, newMoves, state.game, 0}
			nextState.calcAllCosts()
			ret = append(ret,&nextState)
		} 
	}
	return ret
}