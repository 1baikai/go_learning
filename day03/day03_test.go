package day03

import (
	"reflect"
	"testing"
)

func Test_reversal(t *testing.T) {
	type args struct {
		list []rune
		k    int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "cehsim",
			args: args{
				list: []rune{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []rune{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversal(tt.args.list, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		num []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "原地删除",
			args: args{
				num: []int{5,1, 2, 3, 45, 5, 5},
				val: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.num, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}