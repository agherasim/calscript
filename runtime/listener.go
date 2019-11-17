package runtime

import (
	"github.com/agherasim/calscript_lang"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Listener defines a parse listener.
type Listener struct {
	*calscript_lang.BaseCalscriptListener
	registry *RuleRegistry
}

// NewListener returns a new *Parser type instance.
func NewListener(rr *RuleRegistry) (*Listener, error) {
	l := new(Listener)
	l.registry = rr
	return l, nil
}

// GetRule from embedded RuleRegistry
func (p *Listener) GetRule(n string) (Rule, error) {
	rt := RuleType(n)
	r, err := p.registry.Get(rt)
	return r, err
}

// EnterEveryRule is called when any rule is entered.
func (p *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	r, err := p.GetRule("test")
	if err != nil {
		r.HandleEnter(ctx)
	}
}

// ExitEveryRule is called when any rule is exited.
func (p *Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	r, err := p.GetRule("test")
	if err != nil {
		r.HandleExit(ctx)
	}
}
