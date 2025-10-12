package internal

import (
	"testing"
)

func TestPassingEmptyRule(t *testing.T) {
	rules := []Rule{}

	result, err := RunRuleEngine(rules)
	
	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	if result != false {
		t.Errorf("Rule engine should return false for empty rules")
	}
}
