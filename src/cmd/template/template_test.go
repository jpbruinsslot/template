package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	args := []string{"/tmp/env.tmpl:/tmp/env.conf"}
	src, dst := parseArgs(args)
	if src == "" || dst == "" {
		t.Error()
	}
}

func TestGenerateFile(t *testing.T) {
	// set source and destination
	tmpDir := os.TempDir()
	src := fmt.Sprintf("%s/%s", tmpDir, "env.tmpl")
	dst := fmt.Sprintf("%s/%s", tmpDir, "env.conf")

	// set env variable
	if err := os.Setenv("ENVVARTEST", "this is a test"); err != nil {
		t.Error(err)
	}

	// create template file
	testTemplate := []byte(`Test environment variable: {{ .Env.ENVVARTEST }}`)
	err := ioutil.WriteFile(src, testTemplate, 0644)
	if err != nil {
		t.Error(err)
	}

	// generate the config file
	generateFile(src, dst)

	// check whether the environment variable is present in the file
	envVar := os.Getenv("ENVVARTEST")
	file, err := ioutil.ReadFile(dst)
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(file), envVar) != true {
		t.Error("environment variable not present in destination file")
	}

	// remove everything
	if err := os.Unsetenv("ENVVARTEST"); err != nil {
		t.Error(err)
	}

	os.Remove(src)
	os.Remove(dst)
}
