package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	arg1 int
	arg2 int
	want int
}

func TestAdd(t *testing.T) {

	result := Add(1, 2)
	want := 3
	if result != want {
		t.Errorf("Add(1, 3) FAILED. Expected %d, got %d\n", 8, result)
	}

}
func TestSubtract(t *testing.T) {

	result := Subtract(8, 2)
	want := 6
	if result != want {
		t.Errorf("Subtract(8, 2) FAILED. Expected %d, got %d\n", 8, result)
	}

}

// func TestMultiply(t *testing.T) {
// 	got := Multiply(2, 3)
// 	want := 7

// 	if want != got {
// 		t.Errorf("Expected '%d', but got '%d'", want, got)
// 	}
// }

//Table driven tests in Go

func TestMultiply(t *testing.T) {
	cases := []testCase{
		{2, 3, 6},
		{10, 5, 50},
		{-8, -3, 24},
		{0, 9, 0},
		{-7, 6, -42},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d*%d=%d", tc.arg1, tc.arg2, tc.want), func(t *testing.T) {
			got := Multiply(tc.arg1, tc.arg2)
			if tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {

	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("Divide(5, 0) FAILED. Expected %f, got %f\n", 0.0, result)
	} else {
		t.Logf("Divide(5, 0) PASSED. Expected %f, got %f\n", 0.0, result)
	}
}
