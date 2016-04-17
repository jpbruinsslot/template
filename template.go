package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	usage = `Usage: template [TEMPLATE FILE]:[CONFIG FILE]`
)

type Context struct {
}

// Env returns a map with environment variables.
// These can be referenced from a template file by using the
// following notation: {{ .Env.SHELL }}
func (c *Context) Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" {
		fmt.Println(usage)
	} else {
		src, dst := parseArgs(args)
		generateFile(src, dst)
	}
}

// parseArgs will parse the arguments given from the cli
// format will be source_file.tmpl > dest_file.conf
func parseArgs(args []string) (string, string) {
	strs := strings.Split(args[0], ":")
	if len(strs) != 2 {
		log.Fatalf("malformed arguments: %s. expected src.tmpl > dst.conf", args)
	}
	return strs[0], strs[1]
}

// generateFile will generate a file (dstPath) given a template (srcPath)
func generateFile(srcPath, dstPath string) bool {
	tmpl, err := template.New(filepath.Base(srcPath)).ParseFiles(srcPath)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	dstFile := os.Stdout
	if dstPath != "" {
		if dstFile, err = os.Create(dstPath); err != nil {
			log.Fatalf("creating file: %s", err)
		}
		defer dstFile.Close()
	}

	err = tmpl.ExecuteTemplate(dstFile, filepath.Base(srcPath), &Context{})
	if err != nil {
		log.Fatalf("template error: %s", err)
	}

	return true
}
