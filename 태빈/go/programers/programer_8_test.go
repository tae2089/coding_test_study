package coding

import "testing"

func TestFindKim(t *testing.T) {
	type args struct {
		seoul []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{seoul: []string{"Jane", "Kim"}},
			want: "김서방은 1에 있다",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKim(tt.args.seoul); got != tt.want {
				t.Errorf("FindKim() = %v, want %v", got, tt.want)
			}
		})
	}
}
