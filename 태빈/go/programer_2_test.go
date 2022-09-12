package coding

import "testing"

func TestIsOddOrEven(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"case1",
			args: args{
				num: 0,
			},
			want: "Even",
		}, 
		{
			name:"case2",
			args: args{
				num: 2,
			},
			want: "Even",
		},
		{
			name:"case3",
			args: args{
				num: 1,
			},
			want: "Odd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOddOrEven(tt.args.num); got != tt.want {
				t.Errorf("IsOddOrEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
