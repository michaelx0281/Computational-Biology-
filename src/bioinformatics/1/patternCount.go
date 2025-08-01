package main // work on trying to make packages better later

/*

Code Challenge: Implement PatternCount() (reproduced below).
     Input: Strings Text and Pattern.
     Output: Count(Text, Pattern).

PatternCount(Text, Pattern)
  count ← 0
  for i ← 0 to |Text| − |Pattern|
    if Text(i, |Pattern|) = Pattern
      count ← count + 1
  return count

*/

func PatternCount(Text string, Pattern string) int {
	count := 0
	for i := 0; i < len(Text)-len(Pattern)+1; i++ {
		if Text[i:i+len(Pattern)] == Pattern {
			count++
		}
	}
	return count
}
