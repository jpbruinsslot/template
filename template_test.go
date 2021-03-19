package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestSubstitute(t *testing.T) {
	if err := os.Setenv("ENVVARTEST", "this is a test"); err != nil {
		t.Error(err)
	}

	r := strings.NewReader("Test environment variable: {{ .ENVVARTEST }}")
	var w bytes.Buffer

	if err := Substitute(r, &w); err != nil {
		t.Error(err)
	}

	envVar := os.Getenv("ENVVARTEST")
	if strings.Contains(w.String(), envVar) != true {
		t.Error("environment variable not present in destination file")
	}

	if err := os.Unsetenv("ENVVARTEST"); err != nil {
		t.Error(err)
	}
}
