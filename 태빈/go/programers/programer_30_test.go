package coding

import "testing"

func TestRentCloth(t *testing.T) {
	type args struct {
		n       int
		lost    []int
		reserve []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RentCloth(tt.args.n, tt.args.lost, tt.args.reserve); got != tt.want {
				t.Errorf("RentCloth() = %v, want %v", got, tt.want)
			}
		})
	}
}
