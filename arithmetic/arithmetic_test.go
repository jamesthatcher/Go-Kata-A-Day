package arithmetic

import "testing"

func TestArithmetic(t *testing.T) {
	type args struct {
		a        int
		b        int
		operator string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test add", args{4, 5, "add"}, 9},
		{"Test add", args{0, 5, "add"}, 5},
		{"Test add", args{900000, 500000, "add"}, 1400000},
		{"Test subtract", args{10, 5, "subtract"}, 5},
		{"Test subtract", args{0, 5, "subtract"}, -5},
		{"Test subtract", args{200, 5, "subtract"}, 195},
		{"Test mult", args{5, 5, "multiply"}, 25},
		{"Test mult", args{4, 5, "multiply"}, 20},
		{"Test mult", args{-4, 5, "multiply"}, -20},
		{"Test mult", args{5, 5, "divide"}, 1},
		{"Test mult", args{0, 5, "divide"}, 0},
		{"Test mult", args{-80, 40, "divide"}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Arithmetic(tt.args.a, tt.args.b, tt.args.operator); got != tt.want {
				t.Errorf("Arithmetic() = %v, want %v", got, tt.want)
			}
		})
	}
}
