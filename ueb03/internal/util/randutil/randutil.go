package randutil

import (
    "math/rand"
    "strconv"
    "time"
)

func Init(id string) {
    n, _ := strconv.Atoi(id)
    rand.Seed(time.Now().UnixNano() + int64(n))
}

// Generates a random number where min and max are inclusive.
func RandomInt(min, max int) int {
    return rand.Intn(max-min+1) + min
}
