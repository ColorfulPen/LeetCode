package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		arg1 int
		arg2 int
	}
	tests := []struct {
		arg1 int
		arg2 int
		arg3 args
	}{
		{1, 2, args{1, 2}},
	}
	t.Log(tests[0].arg1)

	t.Log(add(1, 3))
}

func Test_add(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{1, 2}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		// TODO: Add test cases.
		{"t1",args{2},3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := Get(tt.args.index); gotRet != tt.wantRet {
				t.Errorf("Get() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
