package solution

func MaxBottlesDrunk(numBottles int, numExchange int) int {
	emptyBottle := numBottles
	total := numBottles

	for emptyBottle >= numExchange {
		emptyBottle = (emptyBottle - numExchange) + 1
		numExchange++
		total++
	}

	return total
}
