package coding

import "testing"

func TestTonurment(t *testing.T) {
	type args struct {
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name:"case1",
			args: args{
				n: 8,
				a: 4,
				b: 7,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tonurment(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Tonurment() = %v, want %v", got, tt.want)
			}
		})
	}
}
