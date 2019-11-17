package runtime

import (
	"github.com/agherasim/calscript_lang"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Listener interface type
type Listener interface {
	GetRule(n string) (Rule, error)
	GetError() error
	antlr.ParseTreeListener
}

// CalscriptListener defines a parse listener.
type CalscriptListener struct {
	*calscript_lang.BaseCalscriptListener
	rules *RuleRegistry
	err   error
}

// NewCalscriptListener returns a new *Parser type instance.
func NewCalscriptListener(rr *RuleRegistry) (*CalscriptListener, error) {
	l := new(CalscriptListener)
	l.rules = rr
	return l, nil
}

// GetRule from embedded RuleRegistry
func (l *CalscriptListener) GetRule(n string) (Rule, error) {
	rt := RuleType(n)
	r, err := l.rules.Get(rt)
	return r, err
}

// EnterEveryRule is called when any rule is entered.
func (l *CalscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	r, err := l.GetRule("test")
	if err != nil {
		l.err = err
	} else {
		r.HandleEnter(ctx)
	}

}

// ExitEveryRule is called when any rule is exited.
func (l *CalscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	r, err := l.GetRule("test")
	if err != nil {
		l.err = err
	} else {
		r.HandleExit(ctx)
	}
}

// GetError from tree listener
func (l *CalscriptListener) GetError() error {
	return l.err
}
