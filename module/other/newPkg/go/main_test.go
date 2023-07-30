package newPkg

import "testing"

func TestLogger2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{s: "sample"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logger2(tt.args.s)
		})
	}
}
