package abbreviate_a_two_word_name

import "testing"

func TestAbbrevName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"should test that the solution returns the correct value", args{"Sam Harris"}, "S.H"},
		{"should test that the solution returns the correct value", args{"Patrick Feenan"}, "P.F"},
		{"should test that the solution returns the correct value", args{"Evan Cole"}, "E.C"},
		{"should test that the solution returns the correct value", args{"P Favuzzi"}, "P.F"},
		{"should test that the solution returns the correct value", args{"David Mendieta"}, "D.M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbbrevName(tt.args.name); got != tt.want {
				t.Errorf("AbbrevName() = %v, want %v", got, tt.want)
			}
		})
	}
}
