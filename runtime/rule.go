package runtime

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule interface used to define rule handlers.
type Rule interface {
	HandleRule(ctx antlr.ParserRuleContext, ast AST) error
}

// RuleType to store as map key for rules
type RuleType string

// RuleRegistry provides a rule register.
type RuleRegistry struct {
	rules map[*RuleType]*Rule
}

// Add registers a new rule listener.
func (rr *RuleRegistry) Add(rt *RuleType, r *Rule) error {
	rr.rules[r] = rl

	return nil
}
