//Package letter does Go routines! (soon)
package letter

type FreqMap map[rune]int

//Frequency calculates the number of rune occurences
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency calculates the number of rune occurences and capitalises on goroutines & channels
func ConcurrentFrequency(sts []string) FreqMap {
	ch := make(chan FreqMap, len(sts))
	totalFreq := FreqMap{}

	for _, s := range sts {
		go func(toMap string) {
			ch <- Frequency(toMap)
		}(s)
	}

	for range sts {
		// Statements that send or receive values from channels are blocking inside their own goroutine.
		for r, c := range <-ch {
			totalFreq[r] += c
		}
	}
	return totalFreq
}
