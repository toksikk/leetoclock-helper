package main

import (
	"flag"
	"math"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var goos string

func main() {
	var offset int
	flag.IntVar(&offset, "offset", 0, "Offset in tenth of a second (1-9), default is 0")
	flag.Parse()

	offset = int(math.Abs(float64(offset)))

	goos = runtime.GOOS

	cw := color.New(color.BgBlack, color.FgWhite)
	cg := color.New(color.BgBlack, color.FgGreen)
	cy := color.New(color.BgBlack, color.FgYellow)
	cr := color.New(color.BgBlack, color.FgRed)
	c := cw

	c.Println("Close with Ctrl-C.")
	for {
		curTime := time.Now()
		curTenthSec := curTime.Nanosecond() / 100000000
		curSec := curTime.Second()

		t := curSec*10 + curTenthSec

		if t >= 570-offset && t < 580-offset {
			c = cg
		} else if t >= 580-offset && t < 590-offset {
			c = cy
		} else if t >= 590-offset && t < 600-offset {
			c = cr
		} else if (t >= 0-offset && t < 570-offset) || t >= 600-offset {
			c = cw
		}

		c.Printf("\r%s.%d | %s |", curTime.Format("15:04:05"), curTenthSec, printGraph(curTenthSec, offset))
		time.Sleep(time.Millisecond * 50)
	}

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
		switch goos {
		case "windows":
			s += "X"
		default:
			s += "💥"
		}
	} else {
		switch goos {
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
