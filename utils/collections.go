package utils

func FrequencyMap(values []int) map[int]int {
	freqMap := make(map[int]int)
	for _, value := range values {
		freqMap[value]++
	}
	return freqMap
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
