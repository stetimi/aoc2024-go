package utils

func FrequencyMap(values []int) map[int]int {
	freqMap := make(map[int]int)
	for _, value := range values {
		freqMap[value]++
	}
	return freqMap
}
