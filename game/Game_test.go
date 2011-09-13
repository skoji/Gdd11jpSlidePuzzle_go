package game

import "testing"

func TestBoardArray(t *testing.T) { 
	expected := []int{1,9,10,-1,34,35,36,-1}
	a := BoardArray("19A=YZ0=")
	if len(a) != 8 { t.Errorf("length of a is not 8.") }
	for i,_ := range a {
		if a[i] != expected[i] {
			t.Errorf("differs at index %d: result: %d, expected: %d.",i,a[i], expected[i])
		}
	}
}

func TestBoardString(t *testing.T) {
	expected := "19A=YZ0="
	result := BoardString([]int{1,9,10,-1,34,35,36,-1})
	if result != expected { t.Errorf("result : %s expected: %s.", result, expected) }
}


func TestGoalArray(t *testing.T) {
	board := BoardArray("ZY0=12=34")
	goal := GoalArray(board)
	expected := [...]int{1,2,3,-1,4,34,-1,35,36}
	if len(goal) != len(board) {
		t.Errorf("length differs: actual:%d, expected: %d.", len(goal), len(board))
	}
	for i,_ := range goal {
			if goal[i] != expected[i] {
				t.Errorf("differs at index %d: result: %d, expected: %d.",i,goal[i], expected[i])
			}
	}
}

func TestDistances(t *testing.T) {
	expected := [][]int{
		[]int { 0,1,2,1,0,3,2,3,4 },
		[]int { 1,0,1,2,0,2,3,4,3 } ,
		[]int { 2,1,0,3,0,1,4,3,2 } ,
		[]int { 1,2,3,0,0,4,1,2,3 } ,
		[]int { 0,0,0,0,0,0,0,0,0 } ,
		[]int { 3,2,1,4,0,0,3,2,1 } ,
		[]int { 2,3,4,1,0,3,0,1,2 } ,
		[]int { 3,4,3,2,0,2,1,0,1 } ,
		[]int { 4,3,2,3,0,1,2,1,0 } }

	g := CreateGame("1234=6780",3,3)
	g.calcDistances()
	actual := g.distances
	if len(actual) != len(expected) {
		t.Errorf("length differs: actual %d, expected %d.", len(actual), len(expected))
	}

	for i,x := range actual {
		for j,y := range x {
			if y != expected[i][j] {
				t.Errorf("data differs at [%v,%v]: actual %d, expected %d.",i,j,y, expected[i][j])
			}
		}
	}
}

