package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./spl -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("spl version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./spl foo foo:bar:baz a:b:c:d:e:f:g:h:i:j:k", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := `1  foo

1  foo
2  bar
3  baz

 1  a
 2  b
 3  c
 4  d
 5  e
 6  f
 7  g
 8  h
 9  i
10  j
11  k

`
	if !reflect.DeepEqual(outStream.Bytes(), []byte(expected)) {
		t.Errorf("\ngot:\n%s\nexpected:\n%s", string(outStream.Bytes()), expected)
	}
}

func TestRun_nonumFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./spl -nonum foo foo:bar:baz", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := `foo

foo
bar
baz

`
	if !reflect.DeepEqual(outStream.Bytes(), []byte(expected)) {
		t.Errorf("\ngot:\n%s\nexpected:\n%s", string(outStream.Bytes()), expected)
	}
}

func TestRun_delimFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./spl -delim , foo foo,bar,baz", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := `1  foo

1  foo
2  bar
3  baz

`
	if !reflect.DeepEqual(outStream.Bytes(), []byte(expected)) {
		t.Errorf("\ngot:\n%s\nexpected:\n%s", string(outStream.Bytes()), expected)
	}
}
