package vm

import (
	"github.com/ZeroBl21/go-interpreter/code"
	"github.com/ZeroBl21/go-interpreter/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}