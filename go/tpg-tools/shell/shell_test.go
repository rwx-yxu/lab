package shell_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/rwx-yxu/lab/go/tpg-tools/shell"
	"golang.org/x/exp/slices"
)

func TestCmdFromStringErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestCmdFromString(t *testing.T) {
	t.Parallel()
	cmd, err := shell.CmdFromString("/bin/ls -l main.go")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !slices.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewSession(t *testing.T) {
	t.Parallel()
	stdin := os.Stdin
	stdout := os.Stdout
	stderr := os.Stderr

	want := shell.Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}

	got := *shell.NewSession(stdin, stdout, stderr)

	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	stdin := strings.NewReader("echo hello\n\n")
	stdout := &bytes.Buffer{}
	session := shell.NewSession(stdin, stdout, io.Discard)
	session.DryRun = true
	session.Run()
	fmt.Printf("%q\n", stdout.String())
	//Output:
	// > hello\n> >\nBe seeing you!\n
}
