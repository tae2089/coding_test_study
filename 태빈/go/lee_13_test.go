package coding_test

import (
	"coding"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type args struct {
			s string
	}
	tests := []struct {
			name string
			args args
			want int
	}{
			{
				name: "test01",
				args: args{
					s:"III",
				}, 
				want: 3,
			},
			{
				name: "test02",
				args: args{
					s:"LVIII",
				}, 
				want: 58,
			},
			{
				name: "test03",
				args: args{
					s:"IV",
				}, 
				want: 4,
			},
	}
	for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
					if got := coding.RomanToInt(tt.args.s); got != tt.want {
							t.Errorf("romanToInt() = %v, want %v", got, tt.want)
					}
			})
	}
}