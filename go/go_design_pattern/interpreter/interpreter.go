package interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// IExpression Expression interface
type IExpression interface {
	Interpret(stats map[string]float64) bool
}

// AndExpression struct
type AndExpression struct {
	expressions []IExpression
}

// Interpret AndExpression Interpret function
func (a AndExpression) Interpret(stats map[string]float64) bool {
	for _, expression := range a.expressions {
		if !expression.Interpret(stats) {
			return false
		}
	}

	return true
}

// NewAndExpression create an AndExpression function
func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	expressions := make([]IExpression, len(exps))
	for i, e := range exps {
		var expression IExpression
		var err error

		switch {
		case strings.Contains(e, ">"):
			expression, err = NewGreaterExpression(e)
		case strings.Contains(e, "<"):
			expression, err = NewLessExpression(e)
		default:
			err = fmt.Errorf("exp is invalid: %s", exp)
		}

		if err != nil {
			return nil, err
		}

		expressions[i] = expression
	}

	return &AndExpression{expressions: expressions}, nil
}

// AlertRule alert rule struct
type AlertRule struct {
	expression IExpression
}

// NewAlertRule create an AlertRule
func NewAlertRule(rule string) (*AlertRule, error) {
	exp, err := NewAndExpression(rule)
	return &AlertRule{expression: exp}, err
}

// Interpret AlertRule Interpret function
func (a AlertRule) Interpret(stats map[string]float64) bool {
	return a.expression.Interpret(stats)
}

// GreaterExpression struct
type GreaterExpression struct {
	key   string
	value float64
}

// Interpret GreaterExpression  Interpret function
func (g GreaterExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}

	return v > g.value
}

// NewGreaterExpression create an GreaterExpression function
func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is ivalid: %s", exp)
	}

	return &GreaterExpression{
		key:   data[0],
		value: val,
	}, nil
}

// LessExpression struct
type LessExpression struct {
	key   string
	value float64
}

// Interpret LessExpression Interpret function
func (l LessExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[l.key]
	if !ok {
		return false
	}
	return v < l.value
}

// NewLessExpression create an LessExpression function
func NewLessExpression(exp string) (*LessExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != "<" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &LessExpression{
		key:   data[0],
		value: val,
	}, nil
}
