package main

//go:generate go run gen/gen_funcs.go

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/bmatcuk/doublestar/v4"

	"github.com/jonas747/template"
)

// See https://github.community/t5/GitHub-Actions/set-output-Truncates-Multiline-Strings/td-p/37870.
var replacer = strings.NewReplacer("%", "%25", "\n", "%0A", "\r", "%0D")

func main() {
	registerProblemMatcher()

	failures := checkFiles(os.Getenv("INPUT_INCLUDE"))
	var buf strings.Builder
	for i, f := range failures {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(f.String())
	}

	out := buf.String()
	fmt.Println(out)
	fmt.Println("::set-output name=output::" + replacer.Replace(out))
	if len(failures) > 0 {
		os.Exit(1)
	}
}

func registerProblemMatcher() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("error getting user home dir: ", err)
	}

	in, err := os.Open("/check_yag_tmpl_syntax.json")
	if err != nil {
		log.Fatalln("error reading syntax matcher file: ", err)
	}
	defer in.Close()

	dst := path.Join(homedir, "check_yag_tmpl_syntax.json")
	out, err := os.Create(dst)
	if err != nil {
		log.Fatalln("error creating file under user home dir: ", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		log.Fatalln("error copying problem matcher to user home dir: ", err)
	}
	fmt.Println("::add-matcher::" + dst)
}

func checkFiles(pattern string) []CheckFailure {
	matches, err := doublestar.FilepathGlob(pattern)
	if err != nil {
		log.Fatalln("glob matching failed: ", err)
	}

	var failures []CheckFailure
	for _, path := range matches {
		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatalln("error reading file: ", err)
		}

		_, err = template.New("").Funcs(funcs).Parse(string(content))
		if err != nil {
			failures = append(failures, CheckFailure{path, err})
		}
	}
	return failures
}

type CheckFailure struct {
	Path string
	Err  error
}

func (c CheckFailure) String() string {
	return c.Path + ": " + c.Err.Error()
}
