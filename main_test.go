package main

import "testing"

func Test_sale(t *testing.T) {
	type args struct {
		sales []int
	}
	tests := []struct {
		name string
		sales [] int
		want int
	}{
		{name:"Purchase bonus", sales: []int{12_000, 8_000, 15_000,8_000}},
		{name: "No Bonus", sales: []int{5_000, 9_500, 8_000}},
		// TODO: Add test cases.
	}
	for _, test := range tests {
		got:=sale(test.sales)
		if got != test.want {
			t.Error("for bonus tests:", test.name, "got: ", got, "want",test.want)
		}
	}
}