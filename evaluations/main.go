package main

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
)

func basicUsage() {
	// basic use
	expr, err := govaluate.NewEvaluableExpression("true == true")
	if err != nil {
		fmt.Printf("govaluate.NewEvaluableExpression failed: %v\n", err)
		return
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		fmt.Printf("unable to evaluate expression: %v", err)
		return
	}

	fmt.Printf("result: %v\n", result)
}

type Result struct {
	Labels map[string]string
	Value  int
}

func getByLabel(args ...interface{}) (interface{}, error) {
	result := args[0].(Result)
	key := args[1].(string)

	return result.Labels[key], nil
}

func (r Result) Func1(key string) string {
	return r.Labels[key]
}

func (r Result) Func2() string {
	return "test"
}

func functionUsage() {
	functions := map[string]govaluate.ExpressionFunction{
		"getByLabel": getByLabel,
	}

	expr, err := govaluate.NewEvaluableExpressionWithFunctions("getByLabel(result, 'name') == 'test'", functions)
	if err != nil {
		fmt.Printf("govaluate.NewEvaluableExpression failed: %v\n", err)
		return
	}

	input := Result{
		Labels: map[string]string{
			"name": "test",
		},
		Value: 1,
	}

	result, err := expr.Evaluate(map[string]interface{}{
		"result": input,
	})
	if err != nil {
		fmt.Printf("unable to evaluate expression: %v", err)
		return
	}

	fmt.Printf("result: %v\n", result)
}

type dummyParameter struct {
	String    string
	Int       int
	BoolFalse bool
	Nil       interface{}
	// Nested    dummyNestedParameter
}

func (this dummyParameter) Func() string {
	return "funk"
}

var dummyParameterInstance = dummyParameter{
	String:    "string!",
	Int:       101,
	BoolFalse: false,
	Nil:       nil,
}

func methodUsage() {
	expr, err := govaluate.NewEvaluableExpression("foo.Func() + 'hi'")
	if err != nil {
		fmt.Printf("govaluate.NewEvaluableExpression failed: %v\n", err)
		return
	}

	result, err := expr.Evaluate(map[string]interface{}{
		"foo": dummyParameterInstance,
	})
	if err != nil {
		fmt.Printf("unable to evaluate expression: %v", err)
		return
	}

	fmt.Printf("result: %v\n", result)
}

func stringInSlice(args ...interface{}) (interface{}, error) {
	s := args[0].(string)
	dstr := args[1].(string)

	ds := strings.Split(dstr, ",")

	for _, d := range ds {
		if s == d {
			return true, nil
		}
	}

	return false, nil
}

func stringNotInSlice(args ...interface{}) (interface{}, error) {
	s := args[0].(string)
	dstr := args[1].(string)

	ds := strings.Split(dstr, ",")

	for _, d := range ds {
		if s == d {
			return false, nil
		}
	}

	return true, nil
}

func regUsage() {
	functions := map[string]govaluate.ExpressionFunction{
		"stringInSlice":    stringInSlice,
		"stringNotInSlice": stringNotInSlice,
	}

	expr, err := govaluate.NewEvaluableExpressionWithFunctions("(1<2) && ! stringInSlice('asi_zjk_core_b', 'asi_sh_flink_h01,asi_sh_flink_a01')", functions)
	if err != nil {
		fmt.Printf("govaluate.NewEvaluableExpression failed: %v\n", err)
		return
	}

	result, err := expr.Evaluate(map[string]interface{}{
		"foo": dummyParameterInstance,
	})
	if err != nil {
		fmt.Printf("unable to evaluate expression: %v", err)
		return
	}

	fmt.Printf("result: %v\n", result)
}

func main() {
	regUsage()
}
