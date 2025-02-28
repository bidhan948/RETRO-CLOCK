package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

type placeholder [5]string

var (
	digits = [...]placeholder{
		{ // 0
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		{ // 1
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		{ // 2
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		{ // 3
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		{ // 4
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		{ // 5
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		{ // 6
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		{ // 7
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		{ // 8
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		{ // 9
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
	}

	colon = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
)

func main() {
	screen.Clear()

	// pre-allocate a builder for each frame
	var b strings.Builder

	for {
		b.Reset()
		screen.MoveTopLeft()

		now := time.Now()
		h, m, s := now.Hour(), now.Minute(), now.Second()

		// Build a slice of the placeholders to display.
		clocks := []placeholder{
			digits[h/10], digits[h%10],
			colon,
			digits[m/10], digits[m%10],
			colon,
			digits[s/10], digits[s%10],
		}

		// Toggle the colon on even seconds.
		showColon := s%2 != 0

		// Build each line of the clock output.
		for line := 0; line < 5; line++ {
			for i, ph := range clocks {
				lineStr := ph[line]
				if (i == 2 || i == 5) && !showColon {
					lineStr = "   "
				}
				b.WriteString(lineStr)
				b.WriteString("  ") // gap between placeholders
			}
			b.WriteByte('\n')
		}

		// Print the complete frame.
		fmt.Print(b.String())

		time.Sleep(time.Second)
	}
}
