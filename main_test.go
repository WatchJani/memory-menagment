package main

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkTestSpeed(b *testing.B) {
	b.StopTimer()
	t2 := time.Now().Add(10 * time.Second)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		time.Now().Before(t2)
	}
}

func BenchmarkCompareSlice(b *testing.B) {
	b.StopTimer()

	key1 := []byte("jankokondic72621@jankokondic72621")
	key2 := []byte("jankokondic72621@jankokondic72621")

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		bytes.Compare(key1, key2)
	}
}

func BenchmarkMemoryBlock(b *testing.B) {
	b.StopTimer()

	tree := NewTree(25)

	keys := make([]int, 25)

	for index := range keys {
		keys[index] = rand.Intn(25)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for index := range 25 {
			tree.Add(index, keys[index])
		}

		tree.Clear()
	}
}
