package eval

import (
	"errors"
	"eval/stack"
	"log"
	"strconv"
	"strings"
)

type Operator struct {
	v        string
	priority int
}

func (o Operator) String() string {
	return o.v
}

func (o Operator) HasHigherOrEqualPriorityThan(other Operator) bool {
	return o.priority >= other.priority
}

func ParseOperator(s string) (Operator, error) {
	switch s {
	case Plus.String():
		return Plus, nil
	case Minus.String():
		return Minus, nil
	case Multiply.String():
		return Multiply, nil
	case Divide.String():
		return Divide, nil
	}

	return Operator{}, errors.New("invalid operator")
}

var (
	Plus     = Operator{v: "+", priority: 1}
	Minus    = Operator{v: "-", priority: 1}
	Multiply = Operator{v: "*", priority: 2}
	Divide   = Operator{v: "/", priority: 2}
)

type Expression interface {
	Evaluate() (int, error)
}

type expression struct {
	FirstOperand  Expression
	Operator      Operator
	SecondOperand Expression
}

func (e *expression) Evaluate() (int, error) {
	a, err := e.FirstOperand.Evaluate()
	if err != nil {
		return 0, err
	}
	b, err := e.SecondOperand.Evaluate()
	if err != nil {
		return 0, err
	}

	switch e.Operator {
	case Plus:
		return a + b, nil
	case Minus:
		return a - b, nil
	case Multiply:
		return a * b, nil
	case Divide:
		return a / b, nil
	}
	log.Println(a, b, e.Operator.String())

	return 0, errors.New("invalid operator")
}

type operand int

func ParseOperand(exp string) (*operand, error) {
	num, err := strconv.Atoi(exp)
	if err != nil {
		return nil, err
	}

	o := operand(num)
	return &o, nil
}

func (o operand) Evaluate() (int, error) {
	return int(o), nil
}

func ParseInfixExpression(infixExp string) (Expression, error) {
	postfixExp, err := infixToPostfix(infixExp)
	if err != nil {
		return nil, err
	}
	log.Println(postfixExp)

	var (
		tokens      = strings.Split(postfixExp, " ")
		expressions stack.Stack[Expression]
	)

	for _, token := range tokens {
		if o, err := strconv.Atoi(token); err == nil {
			expressions.Push(operand(o))
			continue
		}

		operator, err := ParseOperator(token)
		if err != nil {
			return nil, errors.New("invalid syntax")
		}

		rightExp, ok := expressions.Pop()
		if !ok {
			return nil, errors.New("invalid syntax")
		}
		leftExp, ok := expressions.Pop()
		if !ok {
			return nil, errors.New("invalid syntax")
		}

		expressions.Push(&expression{
			FirstOperand:  leftExp,
			Operator:      operator,
			SecondOperand: rightExp,
		})

	}

	exp, ok := expressions.Pop()
	if !ok {
		return nil, errors.New("invalid syntax")
	}

	return exp, nil
}

func infixToPostfix(exp string) (string, error) {
	tokens := strings.Split(exp, " ")

	var (
		postfixExp string
		operators  stack.Stack[Operator]
	)
	for _, token := range tokens {
		if _, err := strconv.Atoi(token); err == nil {
			postfixExp += token + " "
			continue
		}

		if operator, err := ParseOperator(token); err == nil {
			if operators.IsEmpty() {
				operators.Push(operator)
				continue
			}

			for topOperator, _ := operators.Top(); topOperator.HasHigherOrEqualPriorityThan(operator); topOperator, _ = operators.Top() {
				postfixExp += topOperator.String() + " "
				operators.Pop()
			}

			operators.Push(operator)
		}

		if token == "(" {
			operators.Push(Operator{v: token})
		}
		if token == ")" {
			for topOperator, _ := operators.Top(); topOperator != (Operator{v: "("}); topOperator, _ = operators.Top() {
				postfixExp += topOperator.String() + " "
				operators.Pop()
			}
			operators.Pop()
		}
	}

	for {
		operator, ok := operators.Pop()
		if !ok {
			break
		}

		postfixExp += operator.String() + " "
	}

	return strings.TrimSpace(postfixExp), nil
}
