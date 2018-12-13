package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		arg  int
		want []int
	}{
		{1, []int{}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{5, []int{5}},
		{6, []int{2, 3}},
		{7, []int{7}},
		{8, []int{2, 2, 2}},
		{9, []int{3, 3}},
		{2 * 2 * 3 * 3 * 5 * 7 * 11 * 11 * 13, []int{2, 2, 3, 3, 5, 7, 11, 11, 13}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.arg), func(t *testing.T) {
			if got := primeFactors(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("primeFactors(%d) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func primeFactors(n int) []int {
	factors := []int{}
	for divisor := 2; n > 1; divisor++ {
		for ; n%divisor == 0; n /= divisor {
			factors = append(factors, divisor)
		}
	}
	return factors
}
