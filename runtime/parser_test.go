package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseShouldError(t *testing.T) {
	pp := getSimpleParser("test")
	errs := pp.Parse("test")

	assert.EqualError(t, errs[0], "line 1:0 missing 'calscript' at 'test'")
}

func TestParseShouldSucceed(t *testing.T) {
	script := `
	calscript test_script
		on 2019/12/06 "I have an ëvent scheduled" // and this is a comment
	`
	pp := getSimpleParser("test")
	errs := pp.Parse(script)
	assert.Empty(t, errs)
}

func TestParseSimpleEvent(t *testing.T) {
	script := `
	calscript test_script
		on 2019/12/06 "I have an ëvent scheduled" // and this is a comment
	`
	rr, err := NewRuleRegistry()
	assert.NoError(t, err)

	l, err := NewListener(rr)
	assert.NoError(t, err)

	el := NewErrorListener(DefaultErrorStackSize)

	pp, err := NewParser(l, el)
	assert.NoError(t, err)

	errs := pp.Parse(script)
	assert.Empty(t, errs)
}

func getSimpleParser(ruleName string) *CalscriptParser {
	rr, _ := NewRuleRegistry()
	rt, r := getTestRule(ruleName)
	rr.Add(rt, r)
	l, _ := NewListener(rr)
	el := NewErrorListener(DefaultErrorStackSize)
	pp, _ := NewParser(l, el)
	return pp
}
