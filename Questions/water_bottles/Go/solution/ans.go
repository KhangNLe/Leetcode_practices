package solution

func NumWaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	emptyBottles := numBottles

	for emptyBottles/numExchange > 0 {
		drinkable := (emptyBottles / numExchange)
		total += drinkable
		emptyBottles = (emptyBottles % numExchange) + drinkable
	}
	return total
}
