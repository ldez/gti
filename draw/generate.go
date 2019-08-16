package main

import (
	"bufio"
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

const srcTemplate = `package main
import (
	"time"
)

{{range .Draws }}
func (d drawer) draw{{ .Name }}(x int) {
	if x%2 != 0 {
		{{range .Draw01 -}}
		d.lineAt(x, {{.}})
		{{end -}}
	} else {
		{{range .Draw02 -}}
		d.lineAt(x, {{.}})
		{{end -}}
	}

	time.Sleep(d.frameTime)
}
{{end}}
`

type model struct {
	Name   string
	Draw01 []string
	Draw02 []string
}

func main() {
	var draws []model
	for _, name := range []string{"std", "pull", "push"} {
		m, err := getDraw(name)
		if err != nil {
			panic(err)
		}

		draws = append(draws, m)
	}

	tmpl, err := template.New("draws.go").Parse(srcTemplate)
	if err != nil {
		panic(err)
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, map[string][]model{"Draws": draws})
	if err != nil {
		panic(err)
	}

	source, err := format.Source(b.Bytes())
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./draws.go", source, 0644)
	if err != nil {
		panic(err)
	}
}

func getDraw(name string) (model, error) {
	m := model{
		Name: strings.Title(name),
	}

	lines, err := readDraw("./draw/" + name + "_01.txt")
	if err != nil {
		return model{}, err
	}
	m.Draw01 = lines

	lines, err = readDraw("./draw/" + name + "_02.txt")
	if err != nil {
		return model{}, err
	}
	m.Draw02 = lines

	return m, nil
}

func readDraw(filename string) ([]string, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = inFile.Close() }()

	var lines []string

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			lines = append(lines, strconv.Quote(text[:len(text)-1]))
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
