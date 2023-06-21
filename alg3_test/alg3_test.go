package alg3_test

import (
	"fmt"
	"testing"
)

func TestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	k := 0
	j := 0
	max := 0
	curr := 0
	fmt.Println(s[0:0])
	for range s[k : j+1] {
		if k == j {
			if 1 >= max {
				max = 1
				curr = 1
			}
			continue
		}

		for i := k; i < j; i++ {
			if s[i] == s[j] {
				j += i + 1

				if curr >= max {
					max = curr
				}
				curr = curr - (i + 1)
				break
			}

			curr++

		}
		if curr >= max {
			max = curr
		}
	}

	return max
}
