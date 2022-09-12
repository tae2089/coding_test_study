package coding

import "testing"

func TestCalculateScarceMoney(t *testing.T) {
	type args struct {
		price int
		money int
		count int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name:"case1",
			args: args{
				price: 3,
				money: 20,
				count: 4,
			},
			want:10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateScarceMoney(tt.args.price, tt.args.money, tt.args.count); got != tt.want {
				t.Errorf("CalculateScarceMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
