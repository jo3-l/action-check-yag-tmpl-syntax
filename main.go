package main

//go:generate go run gen/gen_funcs.go

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jonas747/template"
)

func checkFile(path string) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error reading file: ", err)
	}

	_, err = template.New("").Funcs(funcs).Parse(string(contents))
	return err
}

func main() {
	matches, err := filepath.Glob(os.Getenv("INPUT_INCLUDE"))
	if err != nil {
		log.Fatal("invalid glob pattern: ", err)
	}

	var sb strings.Builder
	for _, path := range matches {
		err = checkFile(path)
		if err != nil {
			fmtErr := fmt.Sprintf("%s: %s", path, err)
			if sb.Len() > 0 {
				sb.WriteByte('\n')
			}
			sb.WriteString(fmtErr)
		}
	}

	out := sb.String()
	fmt.Println(out)

	// See https://github.community/t5/GitHub-Actions/set-output-Truncates-Multiline-Strings/td-p/37870
	out = strings.ReplaceAll(out, "%", "%25")
	out = strings.ReplaceAll(out, "\n", "%0A")
	out = strings.ReplaceAll(out, "\r", "%0D")
	fmt.Println("::set-output name=output::" + out)

	if out == "" {
		os.Exit(0)
	}
	os.Exit(1)
}
