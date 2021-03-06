package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return from New should not fail", tracer)
	} else {
		tracer.Trace("Hello trace package")
		if buf.String() != "Hello trace package" {
			t.Errorf("Trace should not write '%s", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()

	if silentTracer == nil {
		t.Error("Off should return a Tracer, not nil")
	}
}
