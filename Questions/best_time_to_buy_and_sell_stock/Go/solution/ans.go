package solution

func MaxProfit(prices []int) int {
	bestPrice := 0
	point1 := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < point1 {
			point1 = prices[i]
		} else {
			bestPrice = max(bestPrice, (prices[i] - point1))
		}
	}
	return bestPrice
}
