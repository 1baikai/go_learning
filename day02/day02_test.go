package day02

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		stringList []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "公共前缀",
			args: args{stringList: []string{"haha","ha","haydh","hasha","hahaha"}},
			want: "ha",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.stringList); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}