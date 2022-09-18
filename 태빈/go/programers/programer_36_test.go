package coding

import "testing"

func TestFibonaci(t *testing.T) {
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
			name: "case1",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name:"case3", 
			args: args{
				n:2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonaci(tt.args.n); got != tt.want {
				t.Errorf("Fibonaci() = %v, want %v", got, tt.want)
			}
		})
	}
}