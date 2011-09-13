package main

import(
	"game"
	"fmt"
	"flag"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func readLine(r *bufio.Reader) string {
	line, _, err := r.ReadLine()
	if err != nil {
		return ""
	}
	return string(line)
}

func main() {
	var datafile string
	var solvedfile string
	var outputfile string
	var maxVisit int
	var maxDepth int
	var startFrom int
	flag.StringVar(&datafile, "datafile", "data", "input data file")
	flag.StringVar(&solvedfile, "solved", "result_solved", "already solved result file")
	flag.StringVar(&outputfile, "output", "result_go", "file to output result")
	flag.IntVar(&maxVisit, "maxvisit", 1024*1024*4, "max number of nodes to visit")
	flag.IntVar(&maxDepth, "maxdepth", 250, "max depth")
	flag.IntVar(&startFrom, "startfrom", 1, "start from")
	flag.Parse()

	var err os.Error
	var data,solved,output *os.File
	data,err = os.Open(datafile)
	if (err != nil) {
		fmt.Println("error:" + err.String())
		return
	}
	defer data.Close()
	solved,err = os.Open(solvedfile)
	if (err != nil) {
		fmt.Println("error:" + err.String())
		return
	}
	defer solved.Close()
	output, err = os.Create(outputfile)
	if (err != nil) {
		fmt.Println("error:" + err.String())
		return
	}
	defer output.Close()

	dataReader := bufio.NewReader(data)
	solvedReader := bufio.NewReader(solved)
	fmt.Printf("start to solve: input %s, solved %s, output %s\n",datafile, solvedfile, outputfile)
	
	dataReader.ReadLine()
	dataReader.ReadLine()

	ct := 0
	solvedct := 0
	newsolvedct := 0
	for {
		ct ++
		line, prefix, err := dataReader.ReadLine()
		if (err != nil) {
			if err != os.EOF {
				fmt.Println("error:%s", err.String())
			}
			return 
		}
		if prefix {
			fmt.Printf("prefixed!")
			return
		}
		solvedline := readLine(solvedReader)
		if (solvedline != "" || ct < startFrom) {
			if (solvedline != "") { solvedct ++ }
			output.WriteString(solvedline + "\n")
		} else {
			question := strings.Split(string(line), ",")
			w,h,board := question[0], question[1], question[2]
			wval, _ := strconv.Atoi(w)
			hval,_ := strconv.Atoi(h)
			solver := game.CreateIddfsSolver(board,wval,hval,maxDepth,maxVisit)
			result := make(chan string)
			fmt.Println("-------")
			go solver.Run(result)
			s := <-result
			fmt.Printf("%s\n",s)
			output.WriteString(s + "\n")
			if s != "" { solvedct ++; newsolvedct++ }
		}
		if ct % 100 == 0 {
			fmt.Printf("%v/%v (%v)\n", solvedct, ct, newsolvedct)
		}
	}
}
