package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {

	values := []float64{}

	for _, cost := range costs {
		if cost.day == day {
			values = append(values, cost.value)
		}
	}

	return values
}
