package eval

import "testing"

func TestParseExpression(t *testing.T) {
}

func TestExpression_Evaluate(t *testing.T) {
	for name, tc := range map[string]struct {
		exp            *expression
		wantEvaluation int
		wantErr        error
	}{
		"1+2*3": {
			exp: &expression{
				FirstOperand: operand(1),
				Operator:     Plus,
				SecondOperand: &expression{
					FirstOperand:  operand(2),
					Operator:      Multiply,
					SecondOperand: operand(3),
				},
			},
			wantEvaluation: 7,
			wantErr:        nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			evaluation, err := tc.exp.Evaluate()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if evaluation != tc.wantEvaluation {
				t.Errorf("want: %d, got: %d", tc.wantEvaluation, evaluation)
			}
		})
	}
}

func Test_infixToPostfix(t *testing.T) {
	for name, tc := range map[string]struct {
		exp     string
		wantExp string
		wantErr bool
	}{
		"1 + 2 / ( 3 * 4 - 5 )": {
			exp:     "1 + 2 / ( 3 * 4 - 5 )",
			wantExp: "1 2 3 4 * 5 - / +",
		},
	} {
		t.Run(name, func(t *testing.T) {
			exp, err := infixToPostfix(tc.exp)
			if (err != nil) != tc.wantErr {
				t.Errorf("want: %t, got: %t", tc.wantErr, err != nil)
			}

			if exp != tc.wantExp {
				t.Errorf("want: %s, got: %s", tc.wantExp, exp)
			}
		})

	}
}
