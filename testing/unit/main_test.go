package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("ERROR: got %d expected %d", total, 10)
	}

	matrix := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range matrix {
		total := Sum(item.a, item.b)

		if total != item.n {
			t.Errorf("ERROR Sum: got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	matrix := []struct {
		a int
		b int
		n int
	}{
		{2, 5, 5},
		{4, 2, 4},
		{2, 2, 2},
	}

	for _, item := range matrix {
		max := GetMax(item.a, item.b)

		if max != item.n {
			t.Errorf("ERROR GetMax: received a=%d, b=%d expecting the return %d",
				item.a, item.b, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	matrix := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{40, 102334155},
	}

	for _, item := range matrix {
		fib := Fibonacci(item.a)

		if fib != item.n {
			t.Errorf("ERROR Fibonacci: got %d expected %d", fib, item.n)
		}
	}
}
