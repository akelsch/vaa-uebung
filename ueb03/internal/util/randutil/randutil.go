package randutil

import (
    "math/rand"
    "time"
)

func Init() {
    rand.Seed(time.Now().UnixNano())
}

// Generates a random number where min and max are inclusive.
func RandomInt(min, max int) int {
    return rand.Intn(max-min+1) + min
}
