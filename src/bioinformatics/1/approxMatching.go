package main

func ApproxMatching(Text, Pattern string, d int) []int {
	list := make([]int, 0) //keep appending to list

	n := len(Text)
	k := len(Pattern)

	//range over all of Text
	for i := 0; i < n-k+1; i++ {
		screen := Text[i : i+k]
		//if approx match, append index
		if ApproxMatch(screen, Pattern, d) {
			list = append(list, i)
		}
	}

	return list
}

func ApproxMatch(screen, Pattern string, d int) bool {
	if HammingDist(screen, Pattern) <= d {
		return true
	}

	return false
}

func ApproxPatternCount(Text, Pattern string, d int) int {
	count := 0

	n := len(Text)
	k := len(Pattern)

	//range over all of Text
	for i := 0; i < n-k+1; i++ {
		screen := Text[i : i+k]
		//if approx match, append index
		if ApproxMatch(screen, Pattern, d) {
			count += 1
		}
	}

	return count
}
