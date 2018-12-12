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
	return []int{}
}
