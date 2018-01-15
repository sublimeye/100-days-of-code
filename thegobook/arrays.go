package main

import (
	"crypto/sha256"
	"fmt"
)

const (
	ONE = iota
	TWO
	THR
)

type ByteSize int
type Ja float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 2 * iota
	MB
	GB
	TB
	PB Ja = 5 * iota
	EB
	ZB
	YB
)

// Indicies max appear in any order
// Indicides may be omitted
func arraysAsMaps() {
	// array := [...]string{0: "one1", 1: "2", 5: "3"}
	fmt.Printf("%d %d %d %f %f %f", KB, MB, GB, PB, EB, ZB)
}

func exercise41() {
	sha1 := sha256.Sum256([]byte("x"))
	// sha2 := sha256.Sum256([]byte("X"))
	// for _, val := range sha1 {
	// 	fmt.Print(val)
	// }
	x := "x"
	fmt.Printf("\n %v \n %x", sha1, x)
}

// func main() {
// 	exercise41()
// 	// arraysAsMaps()
// }
