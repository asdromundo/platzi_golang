package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d\n", total, 10)
	}
}

func TestTableSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Positivos",
			args: args{1, 1000},
			want: 1000,
		},
		{
			name: "Negativos",
			args: args{0, -1000},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMax(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Max Fib in int64",
			args: args{50},
			want: 12586269025,
		},
		{
			name: "Fib(50)",
			args: args{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseTailFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Base case 0",
			args: args{0},
			want: 0,
		},
		{
			name: "Base case 1",
			args: args{1},
			want: 1,
		},
		{
			name: "Max case 50",
			args: args{50},
			want: 12586269025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseTailFibonacci(tt.args.n); got != tt.want {
				t.Errorf("UseTailFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
