package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCaller(t *testing.T) {
	type args struct {
		depth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "depth_0", args: args{depth: 0}, want: "stack_test.go:37"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Caller(tt.args.depth); fmt.Sprintf("%v", got) != tt.want {
				t.Errorf("Caller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCallers(t *testing.T) {
	type args struct {
		depth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "depth_0", args: args{depth: 0}, want: "\n\tstack_test.go\n\ttesting.go\n\tasm_amd64.s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Callers(tt.args.depth)
			t.Logf("Callers() = %s", got)
			t.Logf("Callers(%%+s) = %+s", got)
			t.Logf("Callers(%%v) = %v", got)
			t.Logf("Callers(%%+v) = %+v", got)
			require.Equal(t, fmt.Sprintf("%+v", got.StackTrace()), fmt.Sprintf("%+v", got))
		})
	}
}
