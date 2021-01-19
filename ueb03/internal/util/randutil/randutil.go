package randutil

import (
    "math/rand"
    "time"
)

func Init(id uint64) {
    // add id to circumvent same seed by accident
    rand.Seed(time.Now().UnixNano() + int64(id))
}

// Generates a random number where min and max are inclusive.
func RandomInt(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func RoundedRandomInt(min, max, accuracy int) int {
    n := RandomInt(min, max)
    return n - (n % accuracy)
}
