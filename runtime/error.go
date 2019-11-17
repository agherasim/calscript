package runtime

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ErrorListener interface
type ErrorListener interface {
	antlr.ErrorListener
	GetErrors() (int, []error)
}

// ErrorStackSize default value
const ErrorStackSize = 10

// CalscriptErrorListener struct
type CalscriptErrorListener struct {
	errors []error
	errCnt int
}

// GetErrors reported by parser
func (el *CalscriptErrorListener) GetErrors() (int, []error) {
	if el.errCnt > 0 {
		return el.errCnt, el.errors[0:el.errCnt]
	}
	return el.errCnt, nil
}

// SyntaxError handler
func (el *CalscriptErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if el.errCnt < ErrorStackSize {
		el.errors[el.errCnt] = fmt.Errorf("line %d:%d %s", line, column, msg)
		el.errCnt++
	}
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

// NewCalscriptErrorListener returns a Calscript specific error listener
func NewCalscriptErrorListener() *CalscriptErrorListener {
	return &CalscriptErrorListener{
		errors: make([]error, ErrorStackSize),
		errCnt: 0,
	}
}
