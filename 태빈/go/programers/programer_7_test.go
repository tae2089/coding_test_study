package coding

import "testing"

func TestCollas(t *testing.T) {
	type args struct {
		num int
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
				num: 6,
			},
			want: 8,
		},
		{
			name: "case2",
			args: args{
				num: 16,
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				num: 626331,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collas(tt.args.num); got != tt.want {
				t.Errorf("Collas() = %v, want %v", got, tt.want)
			}
		})
	}
}