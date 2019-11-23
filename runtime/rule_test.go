package runtime

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestAddRule(t *testing.T) {
	rr, _ := NewRuleRegistry()
	rt, r := getTestRule("test")
	err := rr.Add(rt, r)
	assert.Nil(t, err)
}

func TestLenRegistry(t *testing.T) {
	rr, _ := NewRuleRegistry()
	rt, r := getTestRule("test")
	err := rr.Add(rt, r)
	assert.Nil(t, err)

	l, err := rr.Len()
	assert.Nil(t, err)

	assert.Equal(t, 1, l)
}

func TestGetRuleFromRegistry(t *testing.T) {
	rr, _ := NewRuleRegistry()
	rt, r := getTestRule("test")
	err := rr.Add(rt, r)
	assert.Nil(t, err)

	rt2 := RuleType("test")
	r, err = rr.Get(rt2)
	assert.Nil(t, err)
}

func TestNewRule(t *testing.T) {
	rr, _ := NewRuleRegistry()
	rt, r := getTestRule("*calscript_lang.CalscriptContext")
	err := rr.Add(rt, r)

	assert.Nil(t, err)
}

// Helpers
// ------------------------------------------------------------------
type TestRule struct{}

func (tr *TestRule) HandleEnter(ctx antlr.ParserRuleContext) error {
	return nil
}

func (tr *TestRule) HandleExit(ctx antlr.ParserRuleContext) error {
	return nil
}

func getTestRule(t string) (RuleType, Rule) {
	rt := RuleType(t)
	r := &TestRule{}
	return rt, r
}
