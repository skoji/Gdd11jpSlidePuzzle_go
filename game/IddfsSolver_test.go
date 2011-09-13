package game
import "testing"

func TestRun(t *testing.T) {

	solver := CreateIddfsSolver("168452=30",3,3,200,1024*1024)
	result := make(chan string)
	go solver.Run(result)
	actual := <-result
	if (actual != "ULDRUULDDRUULDDR") {
		t.Errorf("actual is %s", actual)
	}
	
}

 