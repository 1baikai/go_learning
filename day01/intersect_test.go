package day01

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums  []int
		nums1 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "取交集",
			args: args{
				nums:  []int{4,9,5},
				nums1: []int{9,4,9,8,4},
			},
			want: []int{4,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums, tt.args.nums1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect_1(t *testing.T) {
	type args struct {
		nums []int
		num1 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "取交集",
			args: args{
				nums:  []int{1,2,3,4,4,13},
				num1: []int{1,2,3,9,10},
			},
			want: []int{1,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect_1(tt.args.nums, tt.args.num1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect3(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "取交集",
			args: args{
				nums1:  []int{1,2,3,4,4,13},
				nums2: []int{1,2,3,9,10},
			},
			want: []int{1,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect3(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect5(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "取交集",
			args: args{
				nums1:  []int{1,2,3,4,4,13},
				nums2: []int{1,2,3,9,10},
			},
			want: []int{1,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect5(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect5() = %v, want %v", got, tt.want)
			}
		})
	}
}