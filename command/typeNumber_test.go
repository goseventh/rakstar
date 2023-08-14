package command

import "testing"
import "reflect"

func TestStrFuncsTypeNumber(t *testing.T) {
	type result struct {
		value bool
	}
	testCases := []struct {
		name      string
		typeTest  uint8
		arg       string
		condValue interface{}
		expected  result
	}{
		{"valueStrEqual", MustEqual, "30", 30, result{value: true}},
		{"valueStrEqualFail", MustEqual, "30", 10, result{value: false}},
		{"valueStrBeBetween", MustBeBetween, "7", []int{10, 5}, result{value: true}},
		{"valueStrBeBetweenFail", MustBeBetween, "5", []int{3, 10}, result{value: false}},
		{"valueStrBeLess", MustBeLessThan, "3", 7, result{value: true}},
		{"valueStrBeLessFail", MustBeLessThan, "7", 7, result{value: false}},
		{"valueStrBeGreeter", MustBeGreaterThan, "120", 30, result{value: true}},
		{"valueStrBeGreeterFail", MustBeGreaterThan, "30", 50, result{value: false}},
		{"valueStrSquareRootOf", MustBeSquareRootOf, "2", 4, result{value: true}},
		{"valueStrSquareRootOfFail", MustBeSquareRootOf, "2", 16, result{value: false}},
		{"valueStrDivisibleBy", MustBeDivisibleBy, "10", 2, result{value: true}},
		{"valueStrDivisibleByFail", MustBeDivisibleBy, "50", 3, result{value: false}},
		{"valueStrMultipleOfBy", MustBeMultipleOf, "10", 2, result{value: true}},
		{"valueStrMultipleOfByFail", MustBeMultipleOf, "2", 10, result{value: false}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.typeTest {
			case MustEqual:
				got := valueStrEqual(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeBetween:
				got := valueStrBeBetween(tc.arg, tc.condValue.([]int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeGreaterThan:
				got := valueStrBeGreeter(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeLessThan:
				got := valueStrBeLess(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeMultipleOf:
				got := valueStrDivisibleBy(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeSquareRootOf:
				got := valueStrSquareRootOf(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			}
		})
	}
}
