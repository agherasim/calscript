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
	listener      Listener
	errorListener ErrorListener
}

// Parse input calscript string.
func (p *CalscriptParser) Parse(s string) (Listener, error) {
	stream := antlr.NewInputStream(s)
	lexer := calscript_lang.NewCalscriptLexer(stream)
	lexer.AddErrorListener(p.errorListener)

	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parser := calscript_lang.NewCalscriptParser(tokens)
	parser.BuildParseTrees = true

	parser.AddErrorListener(p.errorListener)
	antlr.ParseTreeWalkerDefault.Walk(p.listener, parser.Calscript())

	lenErr, err := p.errorListener.GetErrors()
	if lenErr > 0 {
		return p.listener, err[0]
	}

	return p.listener, p.listener.GetError()
}

// SetListener for parser
func (p *CalscriptParser) SetListener(l Listener) {
	p.listener = l
}

// SetErrorListener for parser
func (p *CalscriptParser) SetErrorListener(el ErrorListener) {
	p.errorListener = el
}

// NewCalscriptParser returns a *CalscriptParser instance.
func NewCalscriptParser(l Listener, el ErrorListener) (*CalscriptParser, error) {
	p := &CalscriptParser{}
	p.SetListener(l)
	p.SetErrorListener(el)
	return p, nil
}
