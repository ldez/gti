package main

import (
	"time"
)

func (d drawer) drawStd(x int) {
	if x%2 != 0 {
		d.lineAt(x, "   ,---------------.            ")
		d.lineAt(x, "  /  /``````|``````\\\\           ")
		d.lineAt(x, " /  /_______|_______\\\\________  ")
		d.lineAt(x, "|]      GTI |'       |        |]")
		d.lineAt(x, "=  .-:-.    |________|  .-:-.  =")
		d.lineAt(x, " `  -+-  --------------  -+-  ' ")
		d.lineAt(x, "   '-:-'                '-:-'   ")
	} else {
		d.lineAt(x, "   ,---------------.            ")
		d.lineAt(x, "  /  /``````|``````\\\\           ")
		d.lineAt(x, " /  /_______|_______\\\\________  ")
		d.lineAt(x, "|]      GTI |'       |        |]")
		d.lineAt(x, "=  .:-:.    |________|  .:-:.  =")
		d.lineAt(x, " `   X   --------------   X   ' ")
		d.lineAt(x, "   ':-:'                ':-:'   ")
	}

	time.Sleep(d.frameTime)
}

func (d drawer) drawPull(x int) {
	if x%2 != 0 {
		d.lineAt(x, "   ,---------------.               __  ")
		d.lineAt(x, "  /  /``````|``````\\\\             /--\\ ")
		d.lineAt(x, " /  /_______|_______\\\\________    \\__/ ")
		d.lineAt(x, "|]      GTI |'       |        |] >-||  ")
		d.lineAt(x, "=  .-:-.    |________|  .-:-.  = >-||  ")
		d.lineAt(x, " `  -+-  --------------  -+-  '    ||  ")
		d.lineAt(x, "   '-:-'                '-:-'      ||  ")
	} else {
		d.lineAt(x, "   ,---------------.               __  ")
		d.lineAt(x, "  /  /``````|``````\\\\             /--\\ ")
		d.lineAt(x, " /  /_______|_______\\\\________    \\__/ ")
		d.lineAt(x, "|]      GTI |'       |        |] >-||  ")
		d.lineAt(x, "=  .:-:.    |________|  .:-:.  = >-||  ")
		d.lineAt(x, " `   X   --------------   X   '   /  \\ ")
		d.lineAt(x, "   ':-:'                ':-:'    /    \\")
	}

	time.Sleep(d.frameTime)
}

func (d drawer) drawPush(x int) {
	if x%2 != 0 {
		d.lineAt(x, "   __      ,---------------.            ")
		d.lineAt(x, "  /--\\    /  /``````|``````\\\\           ")
		d.lineAt(x, "  \\__/   /  /_______|_______\\\\________  ")
		d.lineAt(x, "   ||-< |]      GTI |'       |        |]")
		d.lineAt(x, "   ||-< =  .-:-.    |________|  .-:-.  =")
		d.lineAt(x, "   ||    `  -+-  --------------  -+-  ' ")
		d.lineAt(x, "   ||      '-:-'                '-:-'   ")
	} else {
		d.lineAt(x, "   __      ,---------------.            ")
		d.lineAt(x, "  /--\\    /  /``````|``````\\\\           ")
		d.lineAt(x, "  \\__/   /  /_______|_______\\\\________  ")
		d.lineAt(x, "   ||-< |]      GTI |'       |        |]")
		d.lineAt(x, "   ||-< =  .:-:.    |________|  .:-:.  =")
		d.lineAt(x, "   /\\    `   X   --------------   X   ' ")
		d.lineAt(x, "  /  \\     ':-:'                ':-:'   ")
	}

	time.Sleep(d.frameTime)
}
