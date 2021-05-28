package main

import (
	"testing"
)

func Test_f(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "defer",
			args: args{x: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.x); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}