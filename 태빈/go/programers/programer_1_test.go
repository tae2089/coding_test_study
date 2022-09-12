package coding

import "testing"

func TestFindNumber(t *testing.T) {
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
			name:"case1",
			args: args{
				n:3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumber(tt.args.n); got != tt.want {
				t.Errorf("FindNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
