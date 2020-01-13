package main

import "testing"

func Test_sale(t *testing.T) {
	tests := []struct {
		name string
		sales [] int
		want int
	}{
		{name: "More than 10_000", sales: []int{25_000, 20_000, 15_000,20_000}, want:4000},
		{name: "Less than 10_000", sales: []int{5_000, 9_500, 8_000, 1_000}, want:0},
		{name: "Equally 10_000", sales: []int{10_000, 10_000, 10_000, 10_000}, want:0},
		{name: "Mixed", sales: []int{25_000, 20_000, 6_000, 8_000}, want:2250},
		// TODO: Add test cases.
	}
	for _, test := range tests {
		got:=sale(test.sales)
		if got != test.want {
			t.Error("for bonus tests:", test.name, "got: ", got, "want:", test.want)
		}
	}
}

