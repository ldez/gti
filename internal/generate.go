package main

import (
	"bufio"
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

const (
	// filePrefix the files related to animation must have this extension.
	filePrefix = ".txt"

	// rlPrefix if a file is prefixed by this prefix, the animation will be run from right to left.
	rlPrefix = "rl_"

	// multiplierPattern f a file is suffixed by this suffix, the sprite will be repeated as many times as the number defines after the letter 'x'.
	multiplierPattern = `_x(\d+).txt$`
)

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
			"Title": cases.Title(language.Und).String,
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

	err = os.WriteFile("./animations.go", source, 0o644)
	if err != nil {
		panic(err)
	}
}

func readCarTypes(rootDir string) (map[string]map[string]descriptor, error) {
	infos, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	result := map[string]map[string]descriptor{}

	for _, item := range infos {
		if !item.IsDir() {
			continue
		}

		carType := item.Name()

		typeDir := filepath.Join(rootDir, carType)

		animations, err := readCommands(typeDir)
		if err != nil {
			return nil, err
		}

		result[carType] = animations
	}

	return result, nil
}

func readCommands(dirPath string) (map[string]descriptor, error) {
	items, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var maxHeight int
	var maxLength int

	result := map[string]descriptor{}

	for _, item := range items {
		if item.IsDir() && !strings.HasPrefix(item.Name(), ".") {
			cmdName := item.Name()

			cmdDir := filepath.Join(dirPath, cmdName)

			anims, err := readAnimations(cmdDir)
			if err != nil {
				return nil, err
			}

			result[cmdName] = anims
		} else if strings.HasSuffix(item.Name(), filePrefix) {
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
				Frames: appendFrames(item.Name(), result[""].Frames, data),
			}
		}
	}

	return result, nil
}

func readAnimations(dirPath string) (descriptor, error) {
	items, err := os.ReadDir(dirPath)
	if err != nil {
		return descriptor{}, err
	}

	var rl bool
	var maxHeight int
	var maxLength int

	var frames [][]string
	for _, item := range items {
		if item.IsDir() || !strings.HasSuffix(item.Name(), filePrefix) {
			continue
		}

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

		frames = appendFrames(item.Name(), frames, data)
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

func appendFrames(name string, src [][]string, data []string) [][]string {
	frames := append(src, data)

	rawMulti := regexp.MustCompile(multiplierPattern).FindStringSubmatch(name)
	if len(rawMulti) == 2 {
		multi, _ := strconv.Atoi(rawMulti[1])
		for i := 0; i < multi; i++ {
			frames = append(frames, data)
		}
	}

	return frames
}
