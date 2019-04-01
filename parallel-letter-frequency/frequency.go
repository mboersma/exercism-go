// Package letter counts the frequency of letters in texts using parallel computation.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in given texts and returns
// this data as a FreqMap. Each text is analyzed concurrently.
func ConcurrentFrequency(list []string) FreqMap {
	ch := make(chan FreqMap, len(list))
	for _, s := range list {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	results := FreqMap{}
	for range list {
		for letter, freq := range <-ch {
			results[letter] += freq
		}
	}
	return results
}
