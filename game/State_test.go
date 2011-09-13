package game

import "testing"

func TestCalcCost(t *testing.T) {
	g := CreateGame("123456780",3,3)
	if calcCost(1,8,g) != 4 {
		t.Errorf("calcCost differs: expected: 4 actual: %v: %v", calcCost(1,8,g), g)
	}
}

func TestCalcallCosts(t *testing.T) {
	g := CreateGame("4231=6708",3,3)
	s := CreateState(g)
	s.board = BoardArray("0123=4678")
	result := 	s.calcAllCosts()
	expected := []int{0,1,1,3,0,4,3,1,1}

	for i,x := range result {
		if x != expected[i] {
			t.Errorf("calcCost differs at index %v: expected: %v actual: %v", i,expected[i],x)
		}
	}
}

func TestCreateStateNotGoal(t *testing.T) {
	g := CreateGame("0123=4687",3,3)
	result := CreateState(g)

	expected_cost := []int{0,1,1,3,0,4,3,0,2}
	for i,x := range result.calcAllCosts() {
		if x != expected_cost[i] {
			t.Errorf("calcCost differs at index %v: expected: %v actual: %v", i,expected_cost[i],x)
		}
	}
	if (result.IsGoal()) {
		t.Errorf("state should not be goal!")
	}
}


func TestCreateStateGoal(t *testing.T) {
	g := CreateGame("123456=80",3,3)
	result := CreateState(g)

	if (!result.IsGoal()) {
		t.Errorf("state should be goal!")
	}
}

func TestNextStates1(t *testing.T) {
	g := CreateGame("=23056=84",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 1 {
		t.Errorf("nexts length should be 1 but was %d", len(nexts))
	}
	if BoardString(nexts[0].board) != "=23506=84" {
		t.Errorf("nexts board should be %v but was %v", "=23506=84", BoardString(nexts[0].board))
	}
	if len(nexts[0].moves) != 1 {
		t.Errorf("moves should be 1 but was %v", len(nexts[0].moves))
	}
	expected_cost := []int{0,0,0,1,0,0,0,0,3}
	for i,x := range nexts[0].calcAllCosts() {
		if x != expected_cost[i] {
			t.Errorf("calcCost differs at index %v: expected: %v actual: %v nextb: %v, goalb: %v", i,expected_cost[i],x,nexts[0].board, g.goalBoard) 
		}
	}
}

func TestNextStates15(t *testing.T) {
	g := CreateGame("=23056=84",3,3)
	state := CreateState(g)
	nexts := state.NextStates()[0].NextStates()
	if len(nexts) != 3 {
		t.Errorf("nexts length should be 3 but was %d", len(nexts))
	}
}

func TestNextStates2(t *testing.T) {
	g := CreateGame("123405678",3,3)
	state := CreateState(g)
	nexts := state.NextStates()

	if len(nexts) != 4 {
		t.Errorf("nexts length should be 4 but was %d", len(nexts))
	}
}

func TestNextStates3(t *testing.T) {
	g := CreateGame("023415678",3,3)
	state := CreateState(g)
	nexts := state.NextStates()

	if len(nexts) != 2 {
		t.Errorf("nexts length should be 2 but was %d", len(nexts))
	}
}

func TestNextStates31(t *testing.T) {
	g := CreateGame("203415678",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 3 {
		t.Errorf("nexts length should be 3 but was %d", len(nexts))
	}
}

func TestNextStates4(t *testing.T) {
	g := CreateGame("230415678",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 2 {
		t.Errorf("nexts length should be 2 but was %d", len(nexts))
	}
}

func TestNextStates5(t *testing.T) {
	g := CreateGame("234015678",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 3 {
		t.Errorf("nexts length should be 3 but was %d", len(nexts))
	}
}

func TestNextStates6(t *testing.T) {
	g := CreateGame("234510678",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 3 {
		t.Errorf("nexts length should be 3 but was %d", len(nexts))
	}
}

func TestNextStates7(t *testing.T) {
	g := CreateGame("234516078",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 2 {
		t.Errorf("nexts length should be 2 but was %d", len(nexts))
	}
}

func TestNextStates8(t *testing.T) {
	g := CreateGame("234516708",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 3 {
		t.Errorf("nexts length should be 3 but was %d", len(nexts))
	}
}

func TestNextStates9(t *testing.T) {
	g := CreateGame("234516780",3,3)
	state := CreateState(g)
	nexts := state.NextStates()
	if len(nexts) != 2 {
		t.Errorf("nexts length should be 2 but was %d", len(nexts))
	}
}

func TestNextStatesChained(t *testing.T) {
	g := CreateGame("12=E4D9HIF8=GN576LOABMTPKQSR0J",5,6)
	pathstring := "ULUUULLDDRDRRUUULLLDDDRDRULURRDRDLLUUUURRDDDLUUUURDDLULLLDDRRRRDD"
	path := make([]int,len(pathstring))
	for i,x := range pathstring {
		switch x {
		case 'U':
			path[i] = -5
		case 'D':
			path[i] = 5
		case 'R':
			path[i] = 1
		case 'L':
			path[i] = -1
		}
	}
	ppstate := func(board string) {
		for i := 0; i < 6; i ++ {
			t.Errorf("%s\n",board[i*5:(i+1)*5])
		}
	}
	state := CreateState(g)
	for i := 0; i < len(path) ; i++ {
		nexts := state.NextStates()
		var nextState *State = nil
		for _,nextCandidate := range nexts {
			if path[i] == nextCandidate.moves.Last() {
				nextState = nextCandidate
				break
			}
		}
		
		if (nextState == nil) {
			t.Errorf("no next state at %v:%s, moves:%v" , i,string(pathstring[i]), state.moves)
			indexOfSpace := -1
			for i,x := range state.board {
				if (x == 36) {
					indexOfSpace = i
					break
				}
			}
			t.Errorf("next moves from game: %v\n", state.game.nexts(indexOfSpace))
			t.Errorf("%v", nexts)
			for i,x := range nexts {
				t.Errorf("next Candidates :%v, %v",i , x.moves)
				ppstate(BoardString(x.board))
			}
			ppstate(BoardString(state.board))
			return
		}
		if (nextState.Score() > 65) {
			t.Errorf("score exceeds 65: %v at %d:%v", nextState.Score(),i,string(pathstring[i]))
			ppstate(BoardString(state.board))
		}
		state = nextState
	}
	
}


