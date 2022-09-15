package coding

import "testing"

func TestCreateSquare(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name:"case1",
			args: args{
				a: 5,
				b:3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateSquare(tt.args.a, tt.args.b)
		})
	}
}
