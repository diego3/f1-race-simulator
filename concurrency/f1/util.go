package f1

import "math/rand"

func random(min, max int) int {
	return rand.Intn(max-min+1) + min
}
