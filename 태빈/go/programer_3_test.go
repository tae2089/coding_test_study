package coding

import "testing"

func TestAddStringNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"case1",
			args: args{
				n: 123,
			},
			want: 6,
		}, 
		{
			name:"case2",
			args: args{
				n: 987,
			},
			want: 24,
		}, 
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddStringNumbers(tt.args.n); got != tt.want {
				t.Errorf("AddStringNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
