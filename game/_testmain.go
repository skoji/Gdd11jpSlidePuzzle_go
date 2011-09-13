package main

import "game"
import "testing"
import __os__ "os"
import __regexp__ "regexp"

var tests = []testing.InternalTest{
	{"game.TestBoardArray", game.TestBoardArray},
	{"game.TestBoardString", game.TestBoardString},
	{"game.TestGoalArray", game.TestGoalArray},
	{"game.TestDistances", game.TestDistances},
	{"game.TestRun", game.TestRun},
	{"game.TestCalcCost", game.TestCalcCost},
	{"game.TestCalcallCosts", game.TestCalcallCosts},
	{"game.TestCreateStateNotGoal", game.TestCreateStateNotGoal},
	{"game.TestCreateStateGoal", game.TestCreateStateGoal},
	{"game.TestNextStates1", game.TestNextStates1},
	{"game.TestNextStates15", game.TestNextStates15},
	{"game.TestNextStates2", game.TestNextStates2},
	{"game.TestNextStates3", game.TestNextStates3},
	{"game.TestNextStates31", game.TestNextStates31},
	{"game.TestNextStates4", game.TestNextStates4},
	{"game.TestNextStates5", game.TestNextStates5},
	{"game.TestNextStates6", game.TestNextStates6},
	{"game.TestNextStates7", game.TestNextStates7},
	{"game.TestNextStates8", game.TestNextStates8},
	{"game.TestNextStates9", game.TestNextStates9},
	{"game.TestNextStatesChained", game.TestNextStatesChained},
}

var benchmarks = []testing.InternalBenchmark{}

var matchPat string
var matchRe *__regexp__.Regexp

func matchString(pat, str string) (result bool, err __os__.Error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = __regexp__.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

func main() {
	testing.Main(matchString, tests, benchmarks)
}
