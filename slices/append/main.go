package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		day, value := costs[i].day, costs[i].value
		if day >= len(costsByDay) {
			diff := day - len(costsByDay)
			costsByDay = append(costsByDay, make([]float64, diff+1)...)
		}
		costsByDay[day] += value
	}
	return costsByDay
}
