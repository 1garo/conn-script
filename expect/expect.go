package expect

import (
	"bytes"
	"errors"
	shell "github.com/kballard/go-shellquote"
	"github.com/kr/pty"
	"io"
	"os"
	"os/exec"
	//	"regexp"
	//	"time"
	//	"unicode/utf8"
)

type ExpectSubprocess struct {
	Cmd          *exec.Cmd
	buf          *buffer
	outputBuffer []byte
}

type buffer struct {
	f       *os.File
	b       bytes.Buffer
	collect bool

	collection bytes.Buffer
}

func (expect *ExpectSubprocess) Interact() {
	defer expect.Cmd.Wait()
	io.Copy(os.Stdout, &expect.buf.b)
	go io.Copy(os.Stdout, expect.buf.f)
	go io.Copy(expect.buf.f, os.Stdin)
}

func Spawn(command string) (*ExpectSubprocess, error) {
	expect, err := _spawn(command)
	if err != nil {
		return nil, err
	}
	return _start(expect)
}

func _start(expect *ExpectSubprocess) (*ExpectSubprocess, error) {
	f, err := pty.Start(expect.Cmd)
	if err != nil {
		return nil, err
	}
	expect.buf.f = f

	return expect, nil
}

func _spawn(command string) (*ExpectSubprocess, error) {
	wrapper := new(ExpectSubprocess)

	wrapper.outputBuffer = nil

	splitArgs, err := shell.Split(command)
	if err != nil {
		return nil, err
	}
	numArguments := len(splitArgs) - 1
	if numArguments < 0 {
		return nil, errors.New("gexpect: No command given to spawn")
	}
	path, err := exec.LookPath(splitArgs[0])
	if err != nil {
		return nil, err
	}

	if numArguments >= 1 {
		wrapper.Cmd = exec.Command(path, splitArgs[1:]...)
	} else {
		wrapper.Cmd = exec.Command(path)
	}
	wrapper.buf = new(buffer)

	return wrapper, nil
}
