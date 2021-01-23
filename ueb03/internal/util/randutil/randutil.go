package randutil

import (
    "math/rand"
    "time"
)

func Init(id uint64) {
    // add id to circumvent same seed by accident
    rand.Seed(time.Now().UnixNano() + int64(id))
}

func RandomBool() bool {
    return RandomInt(0, 1) == 0
}

// Generates a random number where min and max are inclusive.
func RandomInt(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func RoundedRandomInt(min, max, accuracy int) int64 {
    n := RandomInt(min, max)
    n = n - (n % accuracy)
    return int64(n)
}

func RoundedRandomUint(min, max, accuracy int) uint64 {
    return uint64(RoundedRandomInt(min, max, accuracy))
}
