package main

import (
	"flag"
	"math"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var os string

func main() {
	var offset int
	flag.IntVar(&offset, "offset", 0, "Offset in tenth of a second, default is 0")
	flag.Parse()

	offset = int(math.Abs(float64(offset)))

	colorTargetTenthSec := 10-offset
	colorTargetSeconds := []int{57,58,59}
	if offset > 0 {
		for _, v := range colorTargetSeconds {
			v--
		}
	}

	os = runtime.GOOS

	cw := color.New(color.BgBlack, color.FgWhite)
	cg := color.New(color.BgBlack, color.FgGreen)
	cy := color.New(color.BgBlack, color.FgYellow)
	cb := color.New(color.BgBlack, color.FgRed)
	c := cw

	c.Println("Close with Ctrl-C.")
	for {
		curTime := time.Now()
		tenth := curTime.Nanosecond() / 100000000
		curSec := curTime.Second()
		
		if (curSec == colorTargetSeconds[0] && tenth == colorTargetTenthSec) {
			c = cg
		} else if (curSec == colorTargetSeconds[1] && tenth == colorTargetTenthSec) {
			c = cy
		} else if (curSec == colorTargetSeconds[2] && tenth == colorTargetTenthSec) {
			c = cb
		} else if !contains(colorTargetSeconds, curSec) {
			c = color.New(color.BgBlack, color.FgWhite)
		}
		c.Printf("\r%s.%d %s", curTime.Format("15:04:05"), tenth, printGraph(tenth, offset))
		time.Sleep(time.Millisecond * 50)
	}

}

func contains(a []int, n int) bool {
	for _, v := range a {
		if v == n {
			return true
		}
	}
	return false
}

func printGraph(count int, offset int) string {
	j := (count + offset) % 10
	s := ""

	printSpace := func() {
		for i := 0; i < 9-j; i++ {
			s += " "
		}
	}
	printDot := func() {
		for i := 0; i < j; i++ {
			s += "·"
		}
	}

	printDot()
	printSpace()
	if j == 0 {
		switch os {
		case "windows":
			s += "X"
		default:
			s += "💥"
		}
	} else {
		switch os {
		case "windows":
			s += "o"
		default:
			s += "💣"
		}
	}
	printSpace()
	printDot()

	return s
}
