package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

type drawer struct {
	termWidth int
	frameTime time.Duration
}

func newDrawer() drawer {
	termWidth, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	drw := drawer{termWidth: termWidth}

	speed, err := strconv.ParseInt(os.Getenv("GTI_SPEED"), 10, 64)
	if err != nil {
		speed = 1000
	}

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

func (d drawer) drawStd(x int) {
	d.lineAt(x, "   ,---------------.")
	d.lineAt(x, "  /  /``````|``````\\\\")
	d.lineAt(x, " /  /_______|_______\\\\________")
	d.lineAt(x, "|]      GTI |'       |        |]")
	if x%2 != 0 {
		d.lineAt(x, "=  .-:-.    |________|  .-:-.  =")
		d.lineAt(x, " `  -+-  --------------  -+-  '")
		d.lineAt(x, "   '-:-'                '-:-'  ")
	} else {
		d.lineAt(x, "=  .:-:.    |________|  .:-:.  =")
		d.lineAt(x, " `   X   --------------   X   '")
		d.lineAt(x, "   ':-:'                ':-:'  ")
	}

	time.Sleep(d.frameTime)
}

func (d drawer) drawPush(x int) {
	d.lineAt(x, "   __     ,---------------.")
	d.lineAt(x, "  /--\\   /  /``````|``````\\\\")
	d.lineAt(x, "  \\__/  /  /_______|_______\\\\________")
	d.lineAt(x, "   ||-< |]      GTI |'       |        |]")
	if x%2 != 0 {
		d.lineAt(x, "   ||-< =  .-:-.    |________|  .-:-.  =")
		d.lineAt(x, "   ||    `  -+-  --------------  -+-  '")
		d.lineAt(x, "   ||      '-:-'                '-:-'  ")
	} else {
		d.lineAt(x, "   ||-< =  .:-:.    |________|  .:-:.  =")
		d.lineAt(x, "   /\\    `   X   --------------   X   '")
		d.lineAt(x, "  /  \\     ':-:'                ':-:'  ")
	}

	time.Sleep(d.frameTime * 10)
}

func (d drawer) drawPull(x int) {
	d.lineAt(x, "   ,---------------.               __")
	d.lineAt(x, "  /  /``````|``````\\\\             /--\\")
	d.lineAt(x, " /  /_______|_______\\\\________    \\__/")
	d.lineAt(x, "|]      GTI |'       |        |] >-||")
	if x%2 != 0 {
		d.lineAt(x, "=  .-:-.    |________|  .-:-.  = >-||")
		d.lineAt(x, " `  -+-  --------------  -+-  '    || ")
		d.lineAt(x, "   '-:-'                '-:-'      ||  ")
	} else {
		d.lineAt(x, "=  .:-:.    |________|  .:-:.  = >-||")
		d.lineAt(x, " `   X   --------------   X   '   /  \\")
		d.lineAt(x, "   ':-:'                ':-:'    /    \\")
	}

	time.Sleep(d.frameTime * 8)
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
