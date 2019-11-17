package runtime

import "github.com/antlr/antlr4/runtime/Go/antlr"

import "github.com/agherasim/calscript_lang"

// Parser interface type.
type Parser interface {
	Parse(s string) error
	SetRuleRegistry(rr *RuleRegistry)
}

// CalscriptParser executes antlr for parsing a calscript string.
type CalscriptParser struct {
	listener Listener
}

// Parse input calscript string.
func (p *CalscriptParser) Parse(s string) (Listener, error) {
	errorListener := NewCalscriptErrorListener()

	stream := antlr.NewInputStream(s)
	lexer := calscript_lang.NewCalscriptLexer(stream)
	lexer.AddErrorListener(errorListener)

	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parser := calscript_lang.NewCalscriptParser(tokens)
	parser.BuildParseTrees = true

	parser.AddErrorListener(errorListener)
	antlr.ParseTreeWalkerDefault.Walk(p.listener, parser.Calscript())

	lenErr, err := errorListener.GetErrors()
	if lenErr > 0 {
		return p.listener, err[0]
	}

	return p.listener, p.listener.GetError()
}

// SetListener for parser
func (p *CalscriptParser) SetListener(l Listener) {
	p.listener = l
}

// NewCalscriptParser returns a *CalscriptParser instance.
func NewCalscriptParser(l Listener) (*CalscriptParser, error) {
	p := &CalscriptParser{}
	p.SetListener(l)
	return p, nil
}
