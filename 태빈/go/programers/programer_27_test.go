package coding

import "testing"

func TestFindMbti(t *testing.T) {
	type args struct {
		survey  []string
		choices []int
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
				survey:  []string{"AN", "CF", "MJ", "RT", "NA"},
				choices: []int{5, 3, 2, 7, 5},
			},
			want: "TCMA",
		},
		{
			name: "case1",
			args: args{
				survey:  []string{"TR", "RT", "TR"},
				choices: []int{7, 1, 3},
			},
			want: "RCJA",
		},
		{
			name: "case2",
			args: args{
				survey:  []string{"TR"},
				choices: []int{4},
			},
			want: "RCJA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMbti(tt.args.survey, tt.args.choices); got != tt.want {
				t.Errorf("FindMbti() = %v, want %v", got, tt.want)
			}
		})
	}
}