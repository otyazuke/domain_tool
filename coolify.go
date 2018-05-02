package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicationVowel bool = true
	removeVowel      bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			v := -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						v = i
					}
				}
			}
			if v >= 0 {
				switch randBool() {
				case duplicationVowel:
					word = append(word[:v+1], word[v:]...)
				case removeVowel:
					word = append(word[:v], word[v+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}
