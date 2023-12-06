package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func countMatches(winningNumbers, yourNumbers []string) int {
	matchCount := 0
	for _, wNum := range winningNumbers {
		for _, yNum := range yourNumbers {
			if wNum == yNum {
				matchCount++
				break
			}
		}
	}
	return matchCount
}

func main() {
	start := time.Now()

	file, err := os.Open("lotto.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lottoCards []string
	for scanner.Scan() {
		lottoCards = append(lottoCards, scanner.Text())
	}

	totalCards := 0
	cardCopies := make([]int, len(lottoCards))
	for i := range cardCopies {
		cardCopies[i] = 1
	}

	for i, card := range lottoCards {
		parts := strings.Split(card, "|")
		winningNumbers := strings.Fields(strings.TrimSpace(parts[0]))
		yourNumbers := strings.Fields(strings.TrimSpace(parts[1]))

		matchCount := countMatches(winningNumbers, yourNumbers)
		totalCards += cardCopies[i]

		for j := 1; j <= matchCount; j++ {
			nextIndex := i + j
			if nextIndex < len(lottoCards) {
				cardCopies[nextIndex] += cardCopies[i]
			}
		}
	}

	duration := time.Since(start)
	fmt.Printf("Total cards: %d\n", totalCards)
	fmt.Printf("Time elapsed: %v\n", duration)
}
