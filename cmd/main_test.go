package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/simonnagl/eolctl/test"
)

func testMain(t *testing.T, arg string, expected string) {
	cleanup := test.ClearCommandLine()
	defer cleanup()

	os.Args = append(os.Args, arg)

	o, err := captureOutput(main)
	if err != nil {
		t.Fatal("Could not capture output of main().", err)
	}

	if expected != o {
		t.Errorf("Expected output does not match: %s", diff.Diff(expected, o))
	}
}

func captureOutput(f func()) (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	stderr := os.Stderr
	stdout := os.Stdout
	os.Stderr = w
	os.Stdout = w
	defer func() {
		os.Stderr = stderr
		os.Stdout = stdout
	}()

	f()

	_ = w.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func TestMain_Usage(t *testing.T) {
	e := `Usage: eolctl [-hv]

Options:
  -h	Print this usage note
  -v	Print version info
`
	testMain(t, "-h", e)
}

func TestMain_Version(t *testing.T) {
	testMain(t, "-v", "eolctl 0.0.1\n")
}
