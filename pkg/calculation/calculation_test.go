package calculation_test

import (
	"testing"

	"github.com/kirshir/Calculator_server/pkg/calculation"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "+",
			expression:     "2+2",
			expectedResult: 4,
		},
		{
			name:           "-",
			expression:     "5-2",
			expectedResult: 3,
		},
		{
			name:           "*",
			expression:     "4*2",
			expectedResult: 8,
		},
		{
			name:           "/",
			expression:     "9/2",
			expectedResult: 4.5,
		},
		{
			name:           "priority",
			expression:     "2+2*3",
			expectedResult: 8,
		},
		{
			name:           "brackets",
			expression:     "3*(4+1)",
			expectedResult: 15,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "simple",
			expression: "1+2*",
		},
		{
			name:       "priority",
			expression: "2+3**4",
		},
		{
			name:       "unclosed brackets",
			expression: "(2 + 2",
		},
		{
			name:       "division by 0",
			expression: "15/0",
		},
		{
			name:       "wrong character",
			expression: "11-k",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result  %f was obtained", testCase.expression, val)
			}
		})
	}
}
