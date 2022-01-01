package core

import (
	"io"

	"github.com/kiteco/kiteco/kite-go/event"
	"github.com/kiteco/kiteco/kite-golib/kitectx"
)

// The Driver interface is a simple API for handling events and collecting
// output generated by handling of events.
type Driver interface {
	HandleEvent(kitectx.Context, *event.Event) string
	// CollectOutput should return quickly
	CollectOutput() []interface{}
}

// HandlerOutput is an interface used to emit output from a Handler.
type HandlerOutput interface {
	Emit(interface{})
}

// FileDriver is an extension of Driver that allows retrieval of
// file buffer contents and cursor position.
type FileDriver interface {
	Driver
	// Bytes returns a COPY of the underlying bytes in the file
	Bytes() []byte
	Cursor() int64
	SetContents([]byte)
	// ResendText returns true if the text for the file needs to be resent to the backend.
	ResendText() bool
}

// The Diagnoser interface is implemented by drivers that expose a diagnostics stream
type Diagnoser interface {
	// StartDiagnostics causes the driver to start writing diagnostics to the given stream
	// on the next call to HandleEvent.
	StartDiagnostics(io.Writer)
}