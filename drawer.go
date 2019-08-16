package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

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
	drawFunc := d.selectCommand(os.Args)

	height := 7

	fmt.Print(strings.Repeat("\n", height))

	for i := -20; i < d.termWidth; i++ {
		moveToTop(height)
		drawFunc(i)
		d.clearCar(i, height, 2)
	}

	moveToTop(height)
}

func (d drawer) selectCommand(args []string) func(int) {
	for _, value := range args {
		switch value {
		case "push":
			return d.drawPush
		case "pull":
			return d.drawPull
		}
	}
	return d.drawStd
}

func (d drawer) clearCar(x int, height int, length int) {
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
