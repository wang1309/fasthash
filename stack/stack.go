package stack

import (
	"fmt"
	"io"
	"runtime"

	pkgErr "github.com/pkg/errors"
)

type StackTrace = pkgErr.StackTrace

type Frame = pkgErr.Frame

type Stack []uintptr

// Callers makes the depth customizable.
func Callers(depth int) *Stack {
	const numFrames = 32
	var pcs [numFrames]uintptr
	n := runtime.Callers(2+depth, pcs[:])
	var st Stack = pcs[0:n]
	return &st
}

// Caller return only one Frame
func Caller(depth int) Frame {
	const numFrames = 1
	var pcs [numFrames]uintptr
	_ = runtime.Callers(2+depth, pcs[:])
	return Frame(pcs[0])
}

func (s *Stack) Format(st fmt.State, verb rune) {
	for _, pc := range *s {
		f := Frame(pc)
		io.WriteString(st, "\n")
		f.Format(st, verb)
	}
}

func (s *Stack) StackTrace() pkgErr.StackTrace {
	f := make([]pkgErr.Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = pkgErr.Frame((*s)[i])
	}
	return f
}
