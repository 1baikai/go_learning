package bytefmt

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestByteSize(t *testing.T) {
	type args struct {
		bytes uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ceshi",
			args: args{8192},

			want: "8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteSize(tt.args.bytes); got != tt.want {
				t.Errorf("ByteSize() = %v, want %v", got, tt.want)
			}
			reg1 := regexp.MustCompile(`\d+`)
			fmt.Println(reg1.FindString("123kv"))

			fmt.Println(strings.TrimSpace("0\n"))
		})
	}
}
