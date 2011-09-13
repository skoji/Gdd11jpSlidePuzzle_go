package game

import (
	"sort"
	"container/vector"
)

type Game struct {
	rootBoard [] int
	goalBoard [] int
	w,h int
	distances [][] int
}

func BoardString(a [] int) string {
	temp := make([]int, len(a))
	for i,x := range a {
		if (x >= 1 && x <= 9) {
			temp[i] = x + '1' - 1
		} else if x == 36 {
			temp[i] = '0'
		} else if x >= 10 && x <= 35 {
			temp[i] = x + 55
		} else {
			temp[i] = '='
		}
	}
	return string(temp)
}


func BoardArray(s string) [] int {
	ret := make([]int, len(s))
	for i,x := range s {
		if x >= '1' && x <= '9' {
			ret[i] = x - '1' + 1
		} else if x == '0' {
			ret[i] =  36
		} else if x >= 'A' && x <= 'Z' {
			ret[i] = x - 55
		} else {
			ret[i] = -1
		}
	}
	return ret
}

func GoalArray(board [] int) [] int {
	ret := make([]int, len(board))
	temp := make([]int, 0)
	for _,x := range board {
		if x >= 1 && x <= 36 {
			temp = append(temp, x)
		}
	}

	sort.Ints(temp)
	ct := 0
	for i,x := range board {
		if x >= 1 && x <= 36 {
			ret[i] = temp[ct]
			ct ++
		} else {
			ret[i] = x
		}
	}
	return ret
}

func (g *Game) nexts(index int) [] int {
	var r vector.IntVector = make([]int,0)
	x := index % g.w
	y := index / g.w
	if y < g.h - 1 && g.rootBoard[index + g.w] > 0 { r.Push(g.w) }
	if y > 0 && g.rootBoard[index - g.w] > 0 { r.Push(-g.w) }
	if x < g.w - 1 && g.rootBoard[index + 1 ] > 0 { r.Push(1) }
	if x > 0 && g.rootBoard[index -1] > 0 { r.Push(-1) }
	return r
}

func (g *Game) calcDistances() {
	q := new(vector.Vector)
	g.distances = make([][]int, len(g.rootBoard))
	for i,x := range g.rootBoard {
		g.distances[i] = make([]int, len(g.rootBoard))
		if x >= 0 {
			visited := make(map[int] bool)
			q.Push([]int{i,0})
			for ;q.Len() > 0; {
				cell := q.At(0).([]int)
				index, step := cell[0], cell[1]
				q.Delete(0)
				visited[index] = true
				g.distances[i][index] = step
				for _, next := range g.nexts(index) {
					if visited[next + index] != true {
						q.Push([]int{next + index, step + 1})
					}
				}
			}
		}
	}
}

func CreateGame(boardstr string, w int, h int) *Game {
	board := BoardArray(boardstr)
	g := Game{board, GoalArray(board), w, h, nil}
	g.calcDistances()
	return &g
}