package pipeline_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/rwx-yxu/lab/go/tpg-tools/pipeline"
	"golang.org/x/exp/slices"
)

func TestFromFile(t *testing.T) {
	t.Parallel()
	want := []byte("Hello, world\n")
	p := pipeline.FromFile("testdata/hello.txt")
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got, err := io.ReadAll(p.Reader)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFromFileColumn(t *testing.T) {
	t.Parallel()
	want := "203.0.113.17\n203.0.113.1\n203.0.113.1\n203.0.113.10\n203.0.113.10\n203.0.113.19\n203.0.113.250\n203.0.113.250\n203.0.113.14\n203.0.113.16\n203.0.113.19\n203.0.113.2\n203.0.113.250\n203.0.113.19\n203.0.113.250\n203.0.113.250\n203.0.113.10\n203.0.113.2\n203.0.113.2\n203.0.113.2\n203.0.113.253\n203.0.113.86\n203.0.113.20\n203.0.113.86\n203.0.113.9\n"
	p := pipeline.FromFile("testdata/log")
	got, err := p.Column(1).String()
	if err != nil {
		t.Fatal(p.Error)
	}

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}

func TestFromFileInvalid(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("error.txt")
	if p.Error == nil {
		t.Fatal("want error opening non-existent file, but got nil")
	}
}

func TestStdout(t *testing.T) {
	t.Parallel()
	want := "Hellos, world\n"
	p := pipeline.FromString(want)
	buf := &bytes.Buffer{}
	p.Output = buf
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStdoutError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world\n")
	p.Error = errors.New("og no")
	buf := &bytes.Buffer{}
	p.Output = buf
	p.Stdout()
	got := buf.String()
	if got != "" {
		t.Errorf("want no output from Stdout after error, but for %q", got)
	}
}

func TestString(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.FromString(want)
	got, err := p.String()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello world\n")
	p.Error = errors.New("Oh no\n")
	_, err := p.String()
	if err == nil {
		t.Error("want error from String when pipeline has error, but got nil")
	}

}

func TestColumn(t *testing.T) {
	t.Parallel()
	input := "1 2 3\n1 2 3\n1 2 3\n"
	p := pipeline.FromString(input)
	want := "2\n2\n2\n"
	got, err := p.Column(2).String()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q\n", want, got)
	}

}

func TestColumnError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("1 2 3\n")
	p.Error = errors.New("Oh no\n")
	//do not call String() because the error has no knowledge of the previous action.
	//io.ReadAll is basically what String() does anyway
	data, err := io.ReadAll(p.Column(1).Reader)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) > 0 {
		t.Errorf("want no output from column after error, but got %q\n", data)
	}
}

func TestColumnInvalid(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("")
	p.Column(-1)
	if p.Error == nil {
		t.Error("want error on non-positive Column but got nil")
	}
}
