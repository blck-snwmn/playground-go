package main

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/google/uuid"
)

func main() {
	check(uuid.New)
	check(func() uuid.UUID {
		ud, _ := uuid.NewV7()
		return ud
	})
}

func check(gen func() uuid.UUID) {
	ud := gen()
	fmt.Printf("=====%s=====\n", ud.Version())
	fmt.Printf("uuid=%s\n", ud.String())

	const genCount = 100
	uds := make([]uuid.UUID, 0, genCount)
	for i := 0; i < genCount; i++ {
		uds = append(uds, gen())
	}
	// fmt.Printf("uuids=%v\n", uds)
	sorted := slices.IsSortedFunc(uds, func(a, b uuid.UUID) int {
		return cmp.Compare(a.String(), b.String())
	})
	fmt.Printf("sorted=%t\n", sorted)
}
