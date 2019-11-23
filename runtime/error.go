package runtime

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ErrorListener interface
type ErrorListener interface {
	antlr.ErrorListener
	GetErrors() []error
}

// DefaultErrorStackSize specify how many errors to keep in stack
// More errors reported than ErrorStackSize will be discarded
const DefaultErrorStackSize = 10

// CalscriptErrorListener struct
type CalscriptErrorListener struct {
	errors []error
}

// GetErrors reported by parser
func (el *CalscriptErrorListener) GetErrors() []error {
	return el.errors
}

// AppendError to error stack
func (el *CalscriptErrorListener) AppendError(line, column int, msg string) {
	if len(el.errors) < cap(el.errors) {
		e := fmt.Errorf("line %d:%d %s", line, column, msg)
		el.errors = append(el.errors, e)
	}
}

// SyntaxError handler
func (el *CalscriptErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.AppendError(line, column, msg)
}

// ReportAmbiguity handler
func (el *CalscriptErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

// ReportAttemptingFullContext handler
func (el *CalscriptErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

// ReportContextSensitivity handler
func (el *CalscriptErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

// NewErrorListener returns a Calscript specific error listener
func NewErrorListener(capacity int) *CalscriptErrorListener {
	return &CalscriptErrorListener{
		errors: make([]error, 0, capacity),
	}
}
