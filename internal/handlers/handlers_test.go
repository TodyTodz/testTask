package handlers

import (
	"github.com/Maldris/mathparse"
	"testing"
)

//Just some tests for imported package
func TestMath(t *testing.T) {
	expression := "2+2-3-5+1"

	p := mathparse.NewParser(expression)

	p.Resolve()
	if p.FoundResult() {
		var result float64
		result = p.GetValueResult()
		t.Log(result)
	} else {
		var expression string
		expression = p.GetExpressionResult()
		t.Log(expression)
	}
}
