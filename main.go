package main

import (
	"flag"
	"math"
	"time"

	"github.com/fatih/color"
)

func main() {
	var offset int
	flag.IntVar(&offset, "offset", 0, "Offset, default is 0")
	flag.Parse()

	c := color.New(color.BgBlack, color.FgWhite)
	c.Println("Close with Ctrl-C.")
	for {
		curTime := time.Now()
		firstNanoDigit := curTime.Nanosecond() / 100000000
		if curTime.Second() == 57 {
			c = color.New(color.BgBlack, color.FgGreen)
		} else if curTime.Second() == 58 {
			c = color.New(color.BgBlack, color.FgYellow)
		} else if curTime.Second() == 59 {
			c = color.New(color.BgBlack, color.FgRed)
		} else {
			c = color.New(color.BgBlack, color.FgWhite)
		}
		c.Printf("\r%s:%d %s", curTime.Format("15:04:05"), firstNanoDigit, printGraph(firstNanoDigit, offset))
		time.Sleep(time.Millisecond * 100)
	}

}

func printGraph(count int, offset int) string {
	j := (count + int(math.Abs(float64(offset)))) % 10
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
		s += "💥"
	} else {
		s += "💣"
	}
	printSpace()
	printDot()

	return s
}
