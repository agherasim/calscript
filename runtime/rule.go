package runtime

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Rule interface used to define rule handlers.
type Rule interface {
	HandleEnter(ctx antlr.ParserRuleContext) error
	HandleExit(ctx antlr.ParserRuleContext) error
}

// RuleType to store as map key in RuleRegistry
type RuleType string

// String representation of RuleType
func (rt RuleType) String() string {
	return string(rt)
}

// RuleRegistry provides a rule register.
type RuleRegistry struct {
	rules map[RuleType]Rule
}

// NewRuleRegistry returns a rule registry
func NewRuleRegistry() (*RuleRegistry, error) {
	return &RuleRegistry{
		rules: make(map[RuleType]Rule),
	}, nil
}

// Add registers a new rule listener.
func (rr *RuleRegistry) Add(rt RuleType, r Rule) error {
	rr.rules[rt] = r
	return nil
}

// Get Rule from RuleRegistry
func (rr *RuleRegistry) Get(rt RuleType) (Rule, error) {
	if rr.rules[rt] != nil {
		return rr.rules[rt], nil
	}
	return nil, fmt.Errorf("cannot find rule '%s' in registry", rt)
}

// Len lenght of rules in Registry.
func (rr *RuleRegistry) Len() (int, error) {
	return len(rr.rules), nil
}
