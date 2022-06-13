package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	end := 0
	time := 0
	for i := range timeSeries {
		if timeSeries[i] < end {
			time += timeSeries[i] - (end - duration)
		} else {
			time += duration
		}
		end = timeSeries[i] + duration
	}
	return time
}

func main() {

}
