package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	VERSION = "0.1.0"
	USAGE   = `NAME: 
    template - use environment variables in Go templates

USAGE:
    template -i [input-file] -o [output-file]

EXAMPLES:

    $ template -i input.tpml -o output.txt

    $ echo "{{ .PWD }}" | template -o output.txt

    $ echo "{{ .PWD }}" | template

    $ template -i input.tmpl > output.txt

VERSION:
    %s

WEBSITE:
    https://github.com/erroneousboat/template		

GLOBAL OPTIONS:
    -i, -input [input-file]     input file
    -o, -output [output-file]   output file
    -h, -help
`
)

var (
	flgInput  string
	flgOutput string
)

func init() {
	flag.StringVar(
		&flgInput,
		"i",
		"",
		"input file",
	)

	flag.StringVar(
		&flgInput,
		"input",
		"",
		"input file",
	)

	flag.StringVar(
		&flgOutput,
		"o",
		"",
		"output file",
	)

	flag.StringVar(
		&flgOutput,
		"output",
		"",
		"output file",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION)
	}

}

func main() {
	flag.Parse()

	var err error

	var r io.Reader
	if flgInput != "" {
		r, err = os.Open(flgInput)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r = os.Stdin
	}

	var fp *os.File
	if flgOutput != "" {
		fp, err = os.Create(flgOutput)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
	} else {
		fp = os.Stdout
	}

	Substitute(r, fp)
}

func Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func Substitute(r io.Reader, w io.Writer) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	tmpl, err := template.New("template").Parse(string(b))
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, Env())
	if err != nil {
		return err
	}

	return nil
}
