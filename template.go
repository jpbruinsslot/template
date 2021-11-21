package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode/utf8"
)

const (
	VERSION = "0.2.0"
	USAGE   = `NAME: 
    template - use environment variables in Go templates

USAGE:
    template -t [template-file] -o [output-file] -d [data-file]

EXAMPLES:

	# Use system wide environment variables
    $ template -t input.tpml -o output.txt
    $ template -t input.tmpl > output.txt

	# Use data files (support for env and json files)
    $ template -t input.tpml -o output.txt -d data.env
    $ template -t input.tpml -o output.txt -d data.json

	# Use stdin for template file
    $ cat input.tmpl | template -o output.txt    # output txt file
	$ cat input.tmpl | template -o -             # output to stdout
    $ cat input.tmpl | template                  # output to stdout

	# Use stdin for data file
	$ cat data.env | template -t input.tmpl -o output.txt -d -  # output txt file
	$ cat data.env | template -t input.tmpl -d -                # output to stdout

VERSION:
    %s

WEBSITE:
    https://github.com/erroneousboat/template		

GLOBAL OPTIONS:
    -t, -template [template-file]    template file
    -o, -output [output-file]        output file
	-d, -data [data-file]            data file
    -h, -help
`
)

var (
	flgInput  string
	flgOutput string
	flgData   string
)

func init() {
	flag.StringVar(
		&flgInput,
		"t",
		"",
		"template file",
	)

	flag.StringVar(
		&flgInput,
		"template",
		"",
		"template file",
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

	flag.StringVar(
		&flgData,
		"d",
		"",
		"data file",
	)

	flag.StringVar(
		&flgData,
		"data",
		"",
		"data file",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION)
	}
}

func main() {
	flag.Parse()

	var err error

	var tmpl io.Reader
	if flgInput == "-" {
		tmpl = os.Stdin
	} else if flgInput != "" {
		tmpl, err = os.Open(flgInput)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tmpl = os.Stdin
	}

	var output *os.File
	if flgOutput == "-" {
		output = os.Stdout
	} else if flgOutput != "" {
		output, err = os.Create(flgOutput)
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	var df map[string]interface{}
	if flgData == "-" {
		df, err = DataFile(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else if flgData != "" {
		fp, err := os.Open(flgData)
		if err != nil {
			log.Fatal(err)
		}

		df, err = DataFile(fp)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		df = Env()
	}

	Substitute(tmpl, output, df)
}

func DataFile(r io.Reader) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	r, isJSON, err := isJSON(r)
	if err != nil {
		return nil, err
	}

	if isJSON {
		err = json.NewDecoder(r).Decode(&data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	r, isENV, err := isEnv(r)
	if err != nil {
		return nil, err
	}

	if isENV {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			sep := strings.Index(line, "=")

			if sep < 0 {
				return nil, fmt.Errorf("invalid line: %s", line)
			}

			data[line[0:sep]] = line[sep+1:]
		}
		return data, nil
	}

	return nil, fmt.Errorf("unknown data format")
}

func isJSON(r io.Reader) (io.Reader, bool, error) {
	buf := make([]byte, 1)

	n, err := io.ReadAtLeast(r, buf[:], len(buf))
	if err != nil {
		return nil, false, err
	}

	isJSON := bytes.HasPrefix(buf, []byte("{"))
	return io.MultiReader(bytes.NewReader(buf[:n]), r), isJSON, nil
}

func isEnv(r io.Reader) (io.Reader, bool, error) {
	buf := make([]byte, 3)

	n, err := io.ReadAtLeast(r, buf[:], len(buf))
	if err != nil {
		return nil, false, err
	}

	isENV := utf8.Valid(buf)
	return io.MultiReader(bytes.NewReader(buf[:n]), r), isENV, nil
}

func Env() map[string]interface{} {
	env := make(map[string]interface{})
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func Substitute(r io.Reader, w io.Writer, df map[string]interface{}) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	tmpl, err := template.New("template").Parse(string(b))
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, df)
	if err != nil {
		return err
	}

	return nil
}
