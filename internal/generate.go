package main

import (
	"bufio"
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

const srcTemplate = `package main

// Code generated. DO NOT EDIT.

{{range $carType, $commands := . -}}
func get{{ Title $carType }}() map[string]animation{
	return map[string]animation{
		{{range $cmdName, $descriptor := $commands -}}
		"{{ $cmdName }}": {
			rl: {{ $descriptor.RL }},
			height: {{ $descriptor.Height }},
			length: {{ $descriptor.Length }},
			frames: [][]string{
				{{range $frame := $descriptor.Frames -}}
				{
					{{range $frame -}}
					{{ . }},
					{{end -}}
				},
				{{end -}}
			},
		},
		{{end -}}
	}
}

{{end -}}
`

const rlPrefix = "rl_"

type descriptor struct {
	RL     bool
	Height int
	Length int
	Frames [][]string
}

func main() {
	cars, err := readCarTypes("./internal/")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("animations.go").
		Funcs(template.FuncMap{
			"Title": strings.Title,
		}).
		Parse(srcTemplate)
	if err != nil {
		panic(err)
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, cars)
	if err != nil {
		panic(err)
	}

	source, err := format.Source(b.Bytes())
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./animations.go", source, 0644)
	if err != nil {
		panic(err)
	}
}

func readCarTypes(rootDir string) (map[string]map[string]descriptor, error) {
	infos, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	result := map[string]map[string]descriptor{}

	for _, item := range infos {
		if item.IsDir() {
			carType := item.Name()

			typeDir := filepath.Join(rootDir, carType)

			animations, err := readCommands(typeDir)
			if err != nil {
				return nil, err
			}

			result[carType] = animations
		}
	}

	return result, nil
}

func readCommands(dirPath string) (map[string]descriptor, error) {
	items, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var maxHeight int
	var maxLength int

	result := map[string]descriptor{}

	for _, item := range items {
		if item.IsDir() {
			cmdName := item.Name()

			cmdDir := filepath.Join(dirPath, cmdName)

			anims, err := readAnimations(cmdDir)
			if err != nil {
				return nil, err
			}
			result[cmdName] = anims
		} else {
			data, length, err := readDraw(filepath.Join(dirPath, item.Name()))
			if err != nil {
				return nil, err
			}
			if maxHeight < len(data) {
				maxHeight = len(data)
			}

			if maxLength < length {
				maxLength = length
			}

			result[""] = descriptor{
				RL:     result[""].RL || strings.HasPrefix(item.Name(), rlPrefix),
				Height: maxHeight,
				Length: maxLength,
				Frames: append(result[""].Frames, data),
			}
		}
	}

	return result, nil
}

func readAnimations(dirPath string) (descriptor, error) {
	items, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return descriptor{}, err
	}

	var rl bool
	var maxHeight int
	var maxLength int

	var frames [][]string
	for _, item := range items {
		if !item.IsDir() {
			data, length, err := readDraw(filepath.Join(dirPath, item.Name()))
			if err != nil {
				return descriptor{}, err
			}

			rl = rl || strings.HasPrefix(item.Name(), rlPrefix)

			if maxHeight < len(data) {
				maxHeight = len(data)
			}

			if maxLength < length {
				maxLength = length
			}

			frames = append(frames, data)
		}
	}

	return descriptor{
		RL:     rl,
		Height: maxHeight,
		Length: maxLength,
		Frames: frames,
	}, nil
}

func readDraw(filename string) ([]string, int, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer func() { _ = inFile.Close() }()

	var maxLength int
	var lines []string

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		text := scanner.Text()
		if maxLength < len(text) {
			maxLength = len(text)
		}

		if len(text) > 0 {
			lines = append(lines, strconv.Quote(text[:len(text)-1]))
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, 0, err
	}

	return lines, maxLength, nil
}
