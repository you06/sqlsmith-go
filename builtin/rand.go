package builtin

import "math/rand"

func rd (n int) int {
	return rand.Intn(n)
}
