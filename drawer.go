package main

import (
	"fmt"
	"os"
	"strconv"
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

	speed, err := strconv.Atoi(os.Getenv("GTI_SPEED"))
	if err != nil {
		speed = 1000
	}

	drw.frameTime = time.Duration(10000000000 / (termWidth + speed + 1))

	return drw
}

func (d drawer) draw() {
	drawFunc := d.selectCommand(os.Args)

	fmt.Print("\n\n\n\n\n\n\n")

	for i := -20; i < d.termWidth; i++ {
		drawFunc(i)
		d.clearCar(i)
	}
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
	moveToTop()

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
	moveToTop()

	d.lineAt(x, "   __      ,---------------.")
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
	moveToTop()

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

func (d drawer) clearCar(x int) {
	moveToTop()

	d.lineAt(x, "  ")
	d.lineAt(x, "  ")
	d.lineAt(x, "  ")
	d.lineAt(x, "  ")
	d.lineAt(x, "  ")
	d.lineAt(x, "  ")
	d.lineAt(x, "  ")
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

func moveToTop() {
	fmt.Printf("\x1b[7A")
}

func moveTo(x int) {
	fmt.Printf("\x1b[%dC", x)
}
