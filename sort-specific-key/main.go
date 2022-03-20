package main

import "fmt"

func main() {
	us := []user{
		{"s", "aaaa"},
		{"x", "bbbb"},
		{"y", "cccc"},
		{"z", "dddd"},
		{"1", "eeee"},
	}
	fmt.Println(us)
	sortSpecificKey(us, "y")
	fmt.Println(us)
}

type user struct {
	id   string
	name string
}

func (u user) Key() string {
	return u.id
}

type keyGetter[K comparable] interface {
	Key() K
}

func sortSpecificKey[K comparable, T keyGetter[K]](elems []T, key K) {
	for i, e := range elems {
		if e.Key() == key {
			copy(elems[1:i+1], elems[0:i])
			elems[0] = e
		}
	}
}
