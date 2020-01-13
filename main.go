package main

func main() {
}
func sale(sales[]int) int{
	sum := 0
	const purchaseLimit = 10_000
	const percentOfCashback = 5
	for _, value := range sales {
		if value >= purchaseLimit {
			value = value*percentOfCashback/100
			sum += value
		}
	}
	return 0
}


