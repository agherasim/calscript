package runtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListenerGetRule(t *testing.T) {
	rt, r := getTestRule("test")
	rr, _ := NewRuleRegistry()
	rr.Add(rt, r)
	length, err := rr.Len()
	assert.Equal(t, 1, length)
	assert.Nil(t, err)

	l, err := NewCalscriptListener(rr)
	assert.Nil(t, err)
	r1, err := l.GetRule("test")
	assert.Nil(t, err)
	assert.ObjectsAreEqual(r, r1)
}

func TestParseShouldError(t *testing.T) {
	pp := getCalscriptParser("test")
	_, err := pp.Parse("test")

	assert.EqualError(t, err, "line 1:0 missing 'calscript' at 'test'")
}

func TestParseShouldSucceed(t *testing.T) {
	script := `
	calscript test_script
		on 2019/12/06 "I have an ëvent scheduled" // and this is a comment
	`
	pp := getCalscriptParser("*calscript_lang.CalscriptContext")
	_, err := pp.Parse(script)
	assert.Nil(t, err)
}

func getCalscriptParser(ruleName string) *CalscriptParser {
	rr, _ := NewRuleRegistry()
	rt, r := getTestRule(ruleName)
	rr.Add(rt, r)
	l, _ := NewCalscriptListener(rr)
	el := NewCalscriptErrorListener()
	pp, _ := NewCalscriptParser(l, el)
	return pp
}
