package coding

import "testing"

func TestHidePhoneNumer(t *testing.T) {
	type args struct {
		phone_number string
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
				phone_number: "4444",
			},
			want: "4444",
		},
		{
			name: "case2",
			args: args{
				phone_number: "54444",
			},
			want: "*4444",
		},
		{
			name: "case3",
			args: args{
				phone_number: "01033334444",
			},
			want: "*******4444",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HidePhoneNumer(tt.args.phone_number); got != tt.want {
				t.Errorf("HidePhoneNumer() = %v, want %v", got, tt.want)
			}
		})
	}
}