package coding

import "testing"

func TestWaterMellon(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				n: 3,
			},
			want: "수박수",
		},
		{
			name: "case1",
			args: args{
				n: 4,
			},
			want: "수박수박",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WaterMellon(tt.args.n); got != tt.want {
				t.Errorf("WaterMellon() = %v, want %v", got, tt.want)
			}
		})
	}
}