package engine

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func BuildRuleFromString(ruleString string) error {

	antlr.NewInputStream(ruleString)
	// TODO

	return nil
}