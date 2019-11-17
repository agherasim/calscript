package runtime

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestAddRule(t *testing.T) {
	rr, _ := NewRuleRegistry()
	rt, r := GetTestRule()
	err := rr.Add(rt, r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLenRegistry(t *testing.T) {
	rr, _ := NewRuleRegistry()
	rt, r := GetTestRule()
	err := rr.Add(rt, r)
	if err != nil {
		t.Fatal(err)
	}
	l, err := rr.Len()

	if err != nil {
		t.Fatal(err)
	}

	if l != 1 {
		t.Fatalf("Invalid list len, expected 1 got %d", l)
	}
}

func TestGetRuleFromRegistry(t *testing.T) {
	rr, _ := NewRuleRegistry()
	rt, r := GetTestRule()
	err := rr.Add(rt, r)
	if err != nil {
		t.Fatal(err)
	}
	rt2 := RuleType("test")
	r, err = rr.Get(rt2)
	if err != nil {
		t.Fatalf("Cannot find key %s\n", rt)
	}
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

func GetTestRule() (RuleType, Rule) {
	rt := RuleType("test")
	r := &TestRule{}
	return rt, r
}
