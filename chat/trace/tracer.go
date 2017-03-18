package trace

import (
	"io"
	"fmt"
)

type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
}

func Off() Tracer {
	return &silentTracer{}
}

type silentTracer struct {
}

func (t *silentTracer) Trace(a ...interface{}) {
}
