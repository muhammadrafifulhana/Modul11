package main

import "fmt"

func main() {

	validVotes := make(map[int]int)
	totalVotes := 0
	validVoteCount := 0

	for {
		var vote int
		_, err := fmt.Scan(&vote)
		if err != nil {
			fmt.Println("Input tidak valid. Abaikan.")
			continue
		}

		if vote == 0 {
			break
		}

		totalVotes++

		if vote >= 1 && vote <= 20 {
			validVotes[vote]++
			validVoteCount++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", validVoteCount)

	for candidate := 1; candidate <= 20; candidate++ {
		if count, exists := validVotes[candidate]; exists {
			fmt.Printf("%d: %d\n", candidate, count)
		}
	}
}
