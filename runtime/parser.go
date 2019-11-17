package runtime

import (
	"github.com/agherasim/calscript_lang"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Parser defines a parse listener.
type Parser struct {
	*calscript_lang.BaseCalscriptListener
	register *RuleListenerRegister
}

// New returns a new *Parser type instance.
func New(register *RuleListenerRegister) *Parser {
	l := make(Parser)
	l.register = register
	return &l
}

// EnterEveryRule is called when any rule is entered.
func (p *Parser) EnterEveryRule(ctx antlr.ParserRuleContext) {

}

// ExitEveryRule is called when any rule is exited.
func (p *Parser) ExitEveryRule(ctx antlr.ParserRuleContext) {

}
