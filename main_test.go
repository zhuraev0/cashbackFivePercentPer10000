package main

import "testing"

func Test_sale(t *testing.T) {
	tests := []struct {
		name string
		sales [] int
		want int
	}{
		{name: "Purchase bonus", sales: []int{25_000, 20_000, 15_000,20_000}, want:200},
		{name: "Purchase bonus", sales: []int{5_000, 8_000, 15_000,15_000}, want:75},
		{name: "Purchase bonus", sales: []int{5_000, 9_500, 8_000, 10_000}, want:25},
		{name: "Purchase bonus", sales: []int{10_000, 10_000, 10_000, 10_000}, want:100},
		{name: "NO bonus", sales: []int{1_000, 4_000, 6_000, 8_000}, want:0},
		// TODO: Add test cases.
	}
	for _, test := range tests {
		got:=sale(test.sales)
		if got != test.want {
			t.Error("for bonus tests:", test.name, "got: ", got, "want:", test.want)
		}
	}
}

