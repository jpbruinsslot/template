package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestSubstituteEnv(t *testing.T) {
	if err := os.Setenv("ENVVARTEST", "this is a test"); err != nil {
		t.Error(err)
	}

	var w bytes.Buffer
	r := strings.NewReader("Test environment variable: {{ .ENVVARTEST }}")

	if err := Substitute(r, &w, Env()); err != nil {
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

func TestSubstituteDataJSON(t *testing.T) {
	data := strings.NewReader("{\"test\": \"hello, world\"}")

	df, err := DataFile(data)
	if err != nil {
		t.Error(err)
	}

	var w bytes.Buffer
	template := strings.NewReader("Test data file variable {{ .test }}")

	if err := Substitute(template, &w, df); err != nil {
		t.Error(err)
	}

	if strings.Contains(w.String(), "hello, world") != true {
		t.Error("data file variable not present in destination file")
	}

	// TODO: when it isn't valid json
}

func TestSubstituteDataENV(t *testing.T) {
	data := strings.NewReader("test=hello, world")

	df, err := DataFile(data)
	if err != nil {
		t.Error(err)
	}

	var w bytes.Buffer
	template := strings.NewReader("Test data file variable {{ .test }}")

	if err := Substitute(template, &w, df); err != nil {
		t.Error(err)
	}

	if strings.Contains(w.String(), "hello, world") != true {
		t.Error("data file variable not present in destination file")
	}
}

func TestDataFileErrorENV(t *testing.T) {
	data := strings.NewReader("hello, world")
	_, err := DataFile(data)
	if err == nil {
		t.Error("expected error")
	}
}

func TestDataFileJSON(t *testing.T) {
	data := strings.NewReader("{\"test\"}")
	_, err := DataFile(data)
	if err == nil {
		t.Error("expected error")
	}
}
