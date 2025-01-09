package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	nonce, hash, time_spent := minePow(4)
	fmt.Println(nonce, hash, time_spent)
}

func minePow(targetZeros int) (int, string, float64) {
	target := ""
	for i := 0; i < targetZeros; i++ {
		target += "0"
	}
	starttime1 := time.Now()
	nonce := 0
	nickname := "Hz_001"
	for {
		data := fmt.Sprintf("%s%d", nickname, nonce)
		sha256.Sum256([]byte(data))
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		if hash[0:targetZeros] == target {

			time_spent := time.Since(starttime1).Seconds()
			return nonce, hash, time_spent
		}
		nonce++
	}
}
