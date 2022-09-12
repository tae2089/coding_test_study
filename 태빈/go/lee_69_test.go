package coding

import "testing"

func TestMysqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				x:1,
			}, 
			want:1,
		},
		{
			name: "case 2",
			args: args{
				x:4,
			}, 
			want:2,
		},
		{
			name: "case 3",
			args: args{
				x:8,
			}, 
			want:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mysqrt(tt.args.x); got != tt.want {
				t.Errorf("Mysqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
