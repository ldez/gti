package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

type animation struct {
	rl     bool
	height int
	length int
	frames [][]string
}

type drawer struct {
	termWidth int
	frameTime time.Duration
}

func newDrawer(speed int64) drawer {
	termWidth, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	drw := drawer{termWidth: termWidth}

	drw.frameTime = time.Duration(int64(10000000000) / (int64(termWidth) + speed + 1))

	return drw
}

func (d drawer) draw() {
	anim := d.selectAnimation(os.Args)

	fmt.Print(strings.Repeat("\n", anim.height))

	if anim.rl {
		// R -> L
		var index int
		for i := d.termWidth; i > 0; i-- {
			if len(anim.frames) <= index {
				index = 0
			}
			d.move(anim.height, anim.length, i, anim.frames, index)
			index++
		}
	} else {
		// L -> R
		var index int
		for i := -20; i < d.termWidth; i++ {
			if len(anim.frames) <= index {
				index = 0
			}
			d.move(anim.height, anim.length, i, anim.frames, index)
			if i%10 == 0 {
				index++
			}
		}
	}

	moveToTop(anim.height)
}

func (d drawer) move(height, length, x int, frames [][]string, index int) {
	moveToTop(height)

	for _, v := range frames[index] {
		d.lineAt(x, v)
	}

	time.Sleep(d.frameTime)
	d.clearCar(x, height, length)
}

func (d drawer) selectAnimation(args []string) animation {
	golf := getGolf()

	for _, value := range args {
		if anim, ok := golf[value]; ok {
			return anim
		}
	}

	return golf[""]
}

func (d drawer) clearCar(x, height, length int) {
	moveToTop(height)

	for i := 0; i < height; i++ {
		d.lineAt(x, strings.Repeat(" ", length))
	}
}

func (d drawer) lineAt(startX int, in string) {
	if startX > 1 {
		moveTo(startX)
	}

	content := in
	if startX+len(in) > d.termWidth {
		content = string([]rune(in)[:d.termWidth-startX])
	}

	fmt.Println(content)
}

func moveToTop(height int) {
	fmt.Printf("\x1b[%dA", height)
}

func moveTo(x int) {
	fmt.Printf("\x1b[%dC", x)
}
